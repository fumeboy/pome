package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/middleware/trace"
	"time"

	"google.golang.org/grpc/status"
)

func logMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("client's logMiddleware")
		startTime := time.Now()
		err = next(ctx)

		rpcMeta := getMeta(ctx)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		rpcMeta.Log.AddField("cost_us", cost)
		rpcMeta.Log.AddField("method", rpcMeta.Method)
		rpcMeta.Log.AddField("trace_id", trace.GetTraceId(ctx))
		rpcMeta.Log.AddField("server", rpcMeta.ServiceName)

		rpcMeta.Log.AddField("caller_cluster", rpcMeta.CallerCluster)
		rpcMeta.Log.AddField("upstream_cluster", rpcMeta.ServiceCluster)
		rpcMeta.Log.AddField("rpc", 1)
		rpcMeta.Log.AddField("env", rpcMeta.Env)

		var upstreamInfo string
		for _, node := range rpcMeta.HistoryNodes {
			upstreamInfo += fmt.Sprintf("%s:%d,", node.IP, node.Port)
		}

		rpcMeta.Log.AddField("upstream", upstreamInfo)
		rpcMeta.Log.AddField("caller_idc", rpcMeta.CallerIDC)
		rpcMeta.Log.AddField("upstream_idc", rpcMeta.ServiceIDC)
		rpcMeta.Log.Access("result=%v", errStatus.Code())

		return
	}
}
