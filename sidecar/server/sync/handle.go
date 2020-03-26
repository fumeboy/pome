package sync

import (
	"context"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Director() proxy.SyncDirectorT {
	var address = conf.Server.MainoneAddr
	var inner = func(ctx context.Context, ctx2 *ctxT) (err error) {
		conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
		ctx2.ret_conn = conn
		ctx2.ret_ctx = ctx
		return
	}

	var handle = prepare(
		mw_log(
			mw_prom(
				mw_rate_limit(
					conf.Server.Limiter,
					mw_trace(
						inner)))))

	return func(ctx context.Context, _, MethodName string) (context.Context, *grpc.ClientConn, error) {
		ctx2 := init_mw_ctx(MethodName)
		if err := handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, ctx2.ret_conn, nil
		}
		return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
}