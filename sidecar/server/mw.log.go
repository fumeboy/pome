package server

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"time"

	"google.golang.org/grpc/status"
)

func logMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("server's logMiddleware")
		startTime := time.Now()
		err = next(ctx)
		meta := getMeta(ctx)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		meta.Log.AddField("cost_us", cost)
		meta.Log.AddField("method", meta.Method)

		meta.Log.AddField("cluster", meta.Cluster)
		meta.Log.AddField("env", meta.Env)
		meta.Log.AddField("server_ip", meta.ServerIP)
		meta.Log.AddField("client_ip", meta.ClientIP)
		meta.Log.AddField("idc", meta.IDC)
		meta.Log.Access("result=%v", errStatus.Code())

		return
	}
}
