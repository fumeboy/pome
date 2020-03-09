package clientp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"

	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc"
)

func finalMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		fmt.Println("client's finalMiddleware")
		//从ctx获取rpc的metadata
		rpcMeta := meta.GetClientMeta(ctx)
		address := fmt.Sprintf("%s:%d", rpcMeta.CurNode.IP, rpcMeta.CurNode.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			logs.Error(ctx, "connect %s failed, err:%v", address, err)
			return nil, errConnFailed
		}

		rpcMeta.Conn = conn
		defer conn.Close()
		resp, err = next(ctx)
		return
	}
}
