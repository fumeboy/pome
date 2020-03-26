package async

import (
	"github.com/fumeboy/pome/sidecar/proxy"
	"github.com/fumeboy/pome/sidecar/utils/mq"
)

func Director() proxy.AsyncDirectorT {
	return func(serviceName, MethodName string,bytes []byte)  {
		var msg = &mq.MqMsg{
			Method:               MethodName,
			Body:                 bytes,
		}
		send(serviceName, msg)
	}
}
