package producer

import (
	"context"
	"github.com/fumeboy/pome/courier/mq"
	"github.com/fumeboy/pome/courier/rpc"
)

func Director() rpc.AsyncHandlerT {
	return func(ctx context.Context, header *rpc.HeaderT, bytes []byte) (error) {
		var msg = &mq.MqMsg{
			Service: header.ServiceName,
			Method:  header.MthName,
			Body:    bytes,
		}
		mq.Send(header.NodeName, msg)
		return nil
	}
}
