package sync

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func sync_handle_inner(ctx context.Context, ctx2 *ctxT) (err error) {
	address := fmt.Sprintf("%s:%d", ctx2.targetNode.IP, ctx2.targetNode.Port)
	conn, err := grpc.DialContext(ctx, address, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
	if err != nil {
		return errClientConnFailed
	}
	ctx2.ret_conn = conn
	ctx2.ret_ctx = ctx
	return err
}

func Director() proxy.SyncDirectorT {
	var sync_handle = prepare(
		mw_log(
			mw_trace(
				mw_prom(
					mw_rate_limit(
						conf.Client.Limiter,
						mw_hystrix(
							mw_finder(
								conf.Client.Balance,
								sync_handle_inner)))))))
	return func(ctx context.Context, serviceName, MethodName string) (context.Context, *grpc.ClientConn, error) {
		ctx2 := init_mw_ctx(serviceName, MethodName)
		if err := sync_handle(ctx, ctx2); err == nil {
			return ctx2.ret_ctx, ctx2.ret_conn, nil
		}
		return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}
}