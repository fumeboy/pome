package utils

import "strings"

func FullName(serviceName, mthName string)string{
	return serviceName + "/" + mthName
}

func SplitFullName(fullname string) (result []string,ok bool) {
	strs := strings.Split(fullname, "/")
	if len(strs) != 2{
		return nil,false
	}
	return strs, true
}
