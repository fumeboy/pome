package redirector

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/status"
)

func mw_log(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		fmt.Println("client's mw_log")
		startTime := time.Now()
		err = next(ctx,ctx2)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		ctx2.log.AddField("cost_us", cost)
		ctx2.log.AddField("methodName", ctx2.methodName)
		ctx2.log.AddField("trace_id", ctx2.trace_id)
		ctx2.log.AddField("service", ctx2.serviceName)
		ctx2.log.AddField("node", ctx2.nodeName)
		ctx2.log.Access("result=%v", errStatus.Code())

		return
	}
}
