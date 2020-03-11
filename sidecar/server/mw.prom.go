package server

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/middleware/prometheus"
	"time"
)

var (
	defaultMetrics = prometheus.NewServerMetrics()
)

func prometheusMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("server's promMiddleware")
		serverMeta := getMeta(ctx)
		defaultMetrics.IncrRequest(ctx, serverMeta.ServiceName, serverMeta.Method)

		startTime := time.Now()
		err = next(ctx)

		defaultMetrics.IncrCode(ctx, serverMeta.ServiceName, serverMeta.Method, err)
		defaultMetrics.Latency(ctx, serverMeta.ServiceName,
			serverMeta.Method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}
