package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/proxy"
	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var handle middleware.MiddlewareFn

func handle_fn (ctx context.Context) (err error) {
	rpcMeta := getMeta(ctx)
	address := fmt.Sprintf("%s:%d", rpcMeta.CurNode.IP, rpcMeta.CurNode.Port)
	conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
	if err != nil {
		logs.Error(ctx, "connect %s failed, err:%v", address, err)
		return errClientConnFailed
	}
	fmt.Println("handle done")
	rpcMeta.Conn = conn
	return
}

func Director(ctx context.Context, serviceAndMethodName string) (context.Context, *grpc.ClientConn, error) {
	serviceName, mthName, ok := proxy.ReadNames(serviceAndMethodName)
	fmt.Println("client:", serviceName, mthName)
	if ok {
		ctx = initMeta(ctx, serviceName, mthName)
		if err := handle(ctx); err == nil{
			metadata := getMeta(ctx)
			return ctx, metadata.Conn, nil
		}
	}
	return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
}
