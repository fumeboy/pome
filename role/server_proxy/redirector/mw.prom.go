package redirector

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/settings"
	"github.com/fumeboy/pome/manager/prometheus"
	"time"
)

var (
	defaultMetrics = prometheus.NewServerMetrics()
)

func mw_prom(next mw_fn) mw_fn {
	return func(ctx context.Context,ctx2 *ctxT) (err error) {
		fmt.Println("server's promMiddleware")
		defaultMetrics.IncrRequest(ctx, settings.NodeName(), ctx2.method)

		startTime := time.Now()
		err = next(ctx,ctx2)

		defaultMetrics.IncrCode(ctx, settings.NodeName(), ctx2.method, err)
		defaultMetrics.Latency(ctx, settings.NodeName(),
			ctx2.method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}
