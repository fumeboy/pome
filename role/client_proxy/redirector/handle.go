package redirector

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/courier/rpc"
	"github.com/fumeboy/pome/settings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Director() rpc.SyncHandlerT {
	var sync_handle = prepare(
		mw_log(
			mw_trace(
				mw_prom(
					mw_rate_limit(
						settings.ClientProxy.Limiter,
						mw_hystrix(
							mw_finder(
								settings.ClientProxy.Balance,
								func(ctx context.Context, ctx2 *ctxT) error {
									address := fmt.Sprintf("%s:%d", ctx2.targetNode.IP, ctx2.targetNode.Port)
									ctx2.ret_address = address
									ctx2.ret_ctx = ctx
									return nil
								})))))))

	return func(ctx context.Context, header *rpc.HeaderT) (context.Context, string, error) {
		ctx2 := init_mw_ctx(header)
		if err := sync_handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, ctx2.ret_address, nil
		}
		return nil, "", grpc.Errorf(codes.Unimplemented, "Unknown methodName")
	}
}
