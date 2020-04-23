package redirector

import (
	"context"
	"github.com/fumeboy/pome/settings"
	"github.com/fumeboy/pome/courier/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Director() rpc.SyncHandlerT {
	var address = settings.ServerProxy.MainoneAddr

	var handle = prepare(
		mw_log(
			mw_prom(
				mw_rate_limit(
					settings.ServerProxy.Limiter,
					mw_trace(
						func(ctx context.Context, ctx2 *ctxT) (err error) {
							ctx2.ret_ctx = ctx
							return
						})))))

	return func(ctx context.Context, header *rpc.HeaderT) (context.Context,string, error) {
		ctx2 := init_mw_ctx(header.MthName)
		if err := handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, address, nil
		}
		return nil, "", grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
}