package clientp


import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"
	"github.com/fumeboy/pome/rpc/middleware/prometheus"
	"time"

	"github.com/fumeboy/pome/rpc/meta"
)

var (
	defaultMetrics = prometheus.NewClientMetrics()
)

func prometheusMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		fmt.Println("client's promMiddleware")
		rpcMeta := meta.GetClientMeta(ctx)
		defaultMetrics.IncrRequest(ctx, rpcMeta.ServiceName, rpcMeta.Method)

		startTime := time.Now()
		resp, err = next(ctx)

		defaultMetrics.IncrCode(ctx, rpcMeta.ServiceName, rpcMeta.Method, err)
		defaultMetrics.Latency(ctx, rpcMeta.ServiceName,
			rpcMeta.Method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}

