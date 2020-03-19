package server

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var handle = prepare(
	mw_log(
		mw_prom(
			mw_rate_limit(
				conf.Server.Limiter,
				mw_trace(
					inner)))))

func inner(ctx context.Context, ctx2 *ctxT) (err error) {
	address := fmt.Sprintf("%s:%d", "127.0.0.1", conf.Server.Port)
	conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
	if err != nil {
		ctx2.Log.Error("connect %s failed, err:%v", address, err)
		return errServerConnFailed
	}
	ctx2.Conn = conn
	ctx2.RetCtx = ctx
	return
}

func Director(ctx context.Context, serviceAndMethodName string) (context.Context, *grpc.ClientConn, error) {
	serviceName, mthName, ok := proxy.ReadNames(serviceAndMethodName)
	if ok {
		ctx2 := init_mw_ctx(serviceName, mthName)
		if err := handle(ctx, ctx2); err == nil {
			return ctx2.RetCtx, ctx2.Conn, nil
		}
	}
	return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
}
