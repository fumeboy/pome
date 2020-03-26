package server

import (
	"context"
	"github.com/fumeboy/pome/conf"
	"github.com/fumeboy/pome/connector/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Director() rpc.SyncDirectorT {
	var address = conf.Server.MainoneAddr

	var handle = prepare(
		mw_log(
			mw_prom(
				mw_rate_limit(
					conf.Server.Limiter,
					mw_trace(
						func(ctx context.Context, ctx2 *ctxT) (err error) {
							ctx2.ret_ctx = ctx
							return
						})))))

	return func(ctx context.Context, _, MethodName string) (context.Context,string, error) {
		ctx2 := init_mw_ctx(MethodName)
		if err := handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, address, nil
		}
		return nil, "", grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
}