package tag
// from github.com/dovejb/quicktag
import (
	"fmt"
	. "reflect"
	"sort"
	. "unsafe"
)

var (
	styleConvert    = FieldName
	maxSelfRefLevel = 3
	tagNames        = []string{"yaml"}
	byteType        = TypeOf(uint8(0))
	dynTypeMap      = make(map[Type]Type)
)

func Q(x interface{}) interface{} {
	t := TypeOf(x)
	nt := dynTypeMap[t]
	if nt == nil {
		if t.Kind() == Ptr {
			t = t.Elem()
		}
		if t.Kind() != Struct {
			return x
		}
		nt = dynamicType(t)
	}
	return typeCast(x, nt)
}

func dynamicType(t Type) Type {
	if t.Kind() != Struct {
		panic("must struct type")
	}
	nt := dynTypeMap[t]
	if nt == nil {
		nt = genType(t, &genRecords{
			tmap: make(map[Type]int),
			max:  maxSelfRefLevel,
		})
		dynTypeMap[t] = nt
		dynTypeMap[PtrTo(t)] = nt
	}
	return nt
}

type genRecords struct {
	tmap map[Type]int
	max  int
}

func genType(t Type, r *genRecords) Type {
	switch t.Kind() {
	case Ptr:
		gt := genType(t.Elem(), r)
		if gt == nil {
			return nil
		}
		return PtrTo(gt)
	case Struct:
		break
	case Map:
		if t.Key().Kind() != String {
			panic("map key must be string")
		}
		gt := genType(t.Elem(), r)
		if gt == nil {
			return nil
		}
		return MapOf(t.Key(), gt)
	case Slice:
		gt := genType(t.Elem(), r)
		if gt == nil {
			return nil
		}
		return SliceOf(gt)
	case Array:
		gt := genType(t.Elem(), r)
		if gt == nil {
			return nil
		}
		return ArrayOf(t.Len(), gt)
	default:
		return t
	}
	if r.tmap[t] > r.max {
		return nil
	}
	r.tmap[t] += 1
	fs := make([]StructField, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tm := parseTag(string(f.Tag))
		custom := false
		if tm["yaml"] != "" {
			custom = true
		}
		autoName := styleConvert(f.Name)
		for _, tname := range tagNames {
			if tm[tname] == "" {
				tm[tname] = autoName
			}
		}
		var gt Type
		if tm["quicktag"] == "-" || custom {
			gt = f.Type
		} else {
			gt = genType(f.Type, r)
		}
		if gt == nil {
			blankTm := make(map[string]string)
			for _, tname := range tagNames {
				blankTm[tname] = "-"
			}
			fs = append(fs, StructField{
				Name: f.Name,
				Type: ArrayOf(int(f.Type.Size()), byteType),
				Tag:  makeTag(blankTm),
			})
		} else {
			if f.Anonymous && !custom {
				delete(tm, "yaml")
			}
			fs = append(fs, StructField{
				Name:      f.Name,
				Type:      gt,
				Tag:       makeTag(tm),
				Anonymous: f.Anonymous,
			})
		}
	}
	r.tmap[t] -= 1

	nt := StructOf(fs)

	return nt
}

func makeTag(m map[string]string) StructTag {
	if len(m) == 0 {
		return ""
	}
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	s := ""
	for _, k := range ks {
		s += fmt.Sprintf(`%s:"%s" `, k, m[k])
	}
	return StructTag(s)
}

type emptyInterface struct {
	pt Pointer
	pv Pointer
}

func pointerOfType(t Type) Pointer {
	p := *(*emptyInterface)(Pointer(&t))
	return p.pv
}

func typeCast(src interface{}, dstType Type) (dst interface{}) {
	srcType := TypeOf(src)
	eface := *(*emptyInterface)(Pointer(&src))
	if srcType.Kind() == Ptr {
		eface.pt = pointerOfType(PtrTo(dstType))
	} else {
		eface.pt = pointerOfType(dstType)
	}
	dst = *(*interface{})(Pointer(&eface))
	return
}
