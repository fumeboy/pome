package sync

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/status"
)

func mw_log(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		fmt.Println("server's logMiddleware")
		startTime := time.Now()
		err = next(ctx, ctx2)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		ctx2.log.AddField("cost_us", cost)
		ctx2.log.AddField("method", ctx2.method)

		ctx2.log.Access("result=%v", errStatus.Code())

		return
	}
}
