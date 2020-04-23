package redirector

import "github.com/fumeboy/pome/utils/wrong"

var errClientConnFailed = &wrong.E{
	Code:    "ConnFailed",
	Message: "客户端发起连接失败",
}

func isConnFailed(err error) bool {
	e, ok := err.(*wrong.E)
	if !ok {
		return false
	}
	var result bool
	if e == errClientConnFailed {
		result = true
	}
	return result
}
