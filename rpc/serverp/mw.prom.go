package serverp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"
	"github.com/fumeboy/pome/rpc/middleware/prometheus"
	"time"

	"github.com/fumeboy/pome/rpc/meta"
)

var (
	defaultMetrics = prometheus.NewServerMetrics()
)

func prometheusMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		fmt.Println("server's promMiddleware")
		serverMeta := meta.GetServerMeta(ctx)
		defaultMetrics.IncrRequest(ctx, serverMeta.ServiceName, serverMeta.Method)

		startTime := time.Now()
		resp, err = next(ctx)

		defaultMetrics.IncrCode(ctx, serverMeta.ServiceName, serverMeta.Method, err)
		defaultMetrics.Latency(ctx, serverMeta.ServiceName,
			serverMeta.Method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}
