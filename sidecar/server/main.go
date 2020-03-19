package server

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var handle middleware.MiddlewareFn

func handle_fn(ctx context.Context) (err error) {
	rpcMeta := getMeta(ctx)
	address := fmt.Sprintf("%s:%d", "127.0.0.1", conf.Server.Port)
	conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()),grpc.WithInsecure())
	if err != nil {
		rpcMeta.Log.Error("connect %s failed, err:%v", address, err)
		return errServerConnFailed
	}
	rpcMeta.Conn = conn
	rpcMeta.RetCtx = ctx
	return
}

func Director(ctx context.Context, serviceAndMethodName string) (context.Context, *grpc.ClientConn, error) {
	serviceName, mthName, ok := proxy.ReadNames(serviceAndMethodName)
	if ok {
		ctx = initMeta(ctx, serviceName, mthName)
		if err := handle(ctx); err == nil{
			metadata_ := getMeta(ctx)
			return metadata_.RetCtx, metadata_.Conn, nil
		}
	}
	return nil,nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
}
