package utils

import "strings"

func FullName(nodeName, serviceName, mthName string) string {
	return "/" + nodeName + "." + serviceName + "/" + mthName
}

func SplitFullName(fullname string) (result []string, ok bool) {
	strs := strings.Split(fullname, "/")
	if len(strs) != 3 {
		return nil, false
	}
	strs2 := strings.Split(strs[1], ".")
	if len(strs2) != 2 {
		return nil, false
	}
	return []string{strs2[0], strs2[1], strs[2]}, true
}
