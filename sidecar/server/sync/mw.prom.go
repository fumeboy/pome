package sync

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/utils/prometheus"
	"time"
)

var (
	defaultMetrics = prometheus.NewServerMetrics()
)

func mw_prom(next mw_fn) mw_fn {
	return func(ctx context.Context,ctx2 *ctxT) (err error) {
		fmt.Println("server's promMiddleware")
		defaultMetrics.IncrRequest(ctx, conf.NodeName(), ctx2.method)

		startTime := time.Now()
		err = next(ctx,ctx2)

		defaultMetrics.IncrCode(ctx, conf.NodeName(), ctx2.method, err)
		defaultMetrics.Latency(ctx, conf.NodeName(),
			ctx2.method, time.Since(startTime).Nanoseconds()/1000)
		return
	}
}
