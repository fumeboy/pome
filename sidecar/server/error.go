package server

import "github.com/fumeboy/pome/util/wrong"

var errServerConnFailed = &wrong.E{
	Code:    "ServerConnFailed",
	Message: "服务端发起连接失败",
}