package redirector

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/manager/prometheus"
	"time"
)

var (
	defaultMetrics = prometheus.NewClientMetrics()
)

func mw_prom(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		fmt.Println("client's promMiddleware")
		// TODO
		defaultMetrics.IncrRequest(ctx, ctx2.serviceName, ctx2.methodName)

		startTime := time.Now()
		err = next(ctx, ctx2)

		defaultMetrics.IncrCode(ctx, ctx2.serviceName, ctx2.methodName, err)
		defaultMetrics.Latency(ctx, ctx2.serviceName,
			ctx2.methodName, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}

