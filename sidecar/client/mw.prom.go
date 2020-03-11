package client


import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/middleware/prometheus"
	"time"
)

var (
	defaultMetrics = prometheus.NewClientMetrics()
)

func prometheusMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("client's promMiddleware")
		rpcMeta := getMeta(ctx)
		defaultMetrics.IncrRequest(ctx, rpcMeta.ServiceName, rpcMeta.Method)

		startTime := time.Now()
		err = next(ctx)

		defaultMetrics.IncrCode(ctx, rpcMeta.ServiceName, rpcMeta.Method, err)
		defaultMetrics.Latency(ctx, rpcMeta.ServiceName,
			rpcMeta.Method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}

