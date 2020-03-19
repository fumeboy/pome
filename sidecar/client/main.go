package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var handle = prepare(
	mw_log(
		mw_trace(
			mw_prom(
				mw_rate_limit(
					conf.Client.Limiter,
					mw_hystrix(
						mw_loadbalance(
							conf.Client.Balance,
							inner)))))))

func inner(ctx context.Context, ctx2 *ctxT) (err error) {
	address := fmt.Sprintf("%s:%d", ctx2.CurNode.IP, ctx2.CurNode.Port)
	conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
	if err != nil {
		llog.Error("connect %s failed, err:%v", address, err)
		return errClientConnFailed
	}
	ctx2.Conn = conn
	ctx2.RetCtx = ctx
	return err
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
