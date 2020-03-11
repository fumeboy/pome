package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"time"

	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc/status"
)

func logMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("client's logMiddleware")
		ctx = logs.WithFieldContext(ctx)
		startTime := time.Now()
		err = next(ctx)

		rpcMeta := getMeta(ctx)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		logs.AddField(ctx, "cost_us", cost)
		logs.AddField(ctx, "method", rpcMeta.Method)
		logs.AddField(ctx, "server", rpcMeta.ServiceName)

		logs.AddField(ctx, "caller_cluster", rpcMeta.CallerCluster)
		logs.AddField(ctx, "upstream_cluster", rpcMeta.ServiceCluster)
		logs.AddField(ctx, "rpc", 1)
		logs.AddField(ctx, "env", rpcMeta.Env)

		var upstreamInfo string
		for _, node := range rpcMeta.HistoryNodes {
			upstreamInfo += fmt.Sprintf("%s:%d,", node.IP, node.Port)
		}

		logs.AddField(ctx, "upstream", upstreamInfo)
		logs.AddField(ctx, "caller_idc", rpcMeta.CallerIDC)
		logs.AddField(ctx, "upstream_idc", rpcMeta.ServiceIDC)
		logs.Access(ctx, "result=%v", errStatus.Code())

		return
	}
}
