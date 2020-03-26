package client

import (
	"context"
	"github.com/fumeboy/pome/connector/async/mq"
	"github.com/fumeboy/pome/connector/rpc"
)

func Director() rpc.AsyncDirectorT {
	return func(ctx context.Context, serviceName, MethodName string,bytes []byte) (error) {
		var msg = &mq.MqMsg{
			Method:               MethodName,
			Body:                 bytes,
		}
		mq.Send(serviceName, msg)
		return nil
	}
}
