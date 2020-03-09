package serverp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"
	"time"

	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc/status"
)

func logMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		fmt.Println("server's logMiddleware")
		startTime := time.Now()
		resp, err = next(ctx)

		serverMeta := meta.GetServerMeta(ctx)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		logs.AddField(ctx, "cost_us", cost)
		logs.AddField(ctx, "method", serverMeta.Method)

		logs.AddField(ctx, "cluster", serverMeta.Cluster)
		logs.AddField(ctx, "env", serverMeta.Env)
		logs.AddField(ctx, "server_ip", serverMeta.ServerIP)
		logs.AddField(ctx, "client_ip", serverMeta.ClientIP)
		logs.AddField(ctx, "idc", serverMeta.IDC)
		logs.Access(ctx, "result=%v", errStatus.Code())

		return
	}
}

