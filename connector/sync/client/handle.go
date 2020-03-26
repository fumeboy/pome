package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/conf"
	"github.com/fumeboy/pome/connector/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Director() rpc.SyncDirectorT {
	var sync_handle = prepare(
		mw_log(
			mw_trace(
				mw_prom(
					mw_rate_limit(
						conf.Client.Limiter,
						mw_hystrix(
							mw_finder(
								conf.Client.Balance,
								func(ctx context.Context, ctx2 *ctxT) error {
									address := fmt.Sprintf("%s:%d", ctx2.targetNode.IP, ctx2.targetNode.Port)
									ctx2.ret_address = address
									ctx2.ret_ctx = ctx
									return nil
								})))))))

	return func(ctx context.Context, serviceName, MethodName string) (context.Context, string, error) {
		ctx2 := init_mw_ctx(serviceName, MethodName)
		if err := sync_handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, ctx2.ret_address, nil
		}
		return nil, "", grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
}
