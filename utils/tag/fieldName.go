package tag

import "strings"

func FieldName(s string) string {
	if len(s) <= 1 {
		return strings.ToLower(s)
	}
	isUpper := func(c byte) bool {
		return c >= 'A' && c <= 'Z'
	}
	toLower := func(c byte) byte {
		if isUpper(c) {
			return c - 'A' + 'a'
		}
		return c
	}
	r := strings.Builder{}
	r.WriteByte(toLower(s[0]))
	i, l := 1, len(s)
	for i < l {
		if isUpper(s[i]) {
			if !isUpper(s[i-1]) {
				r.WriteByte('_')
			} else if i < l-1 && !isUpper(s[i+1]) {
				r.WriteByte('_')
			}
			r.WriteByte(toLower(s[i]))
		} else {
			r.WriteByte(s[i])
		}
		i += 1
	}
	return r.String()
}
