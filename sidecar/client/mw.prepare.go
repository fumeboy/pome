package client

import (
	"context"
	"github.com/fumeboy/pome/sidecar/middleware/trace"

	"google.golang.org/grpc/metadata"
)

func prepare(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		//处理traceId
		var traceId string
		//从ctx获取grpc的metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			vals, ok := md[trace.TraceID]
			if ok && len(vals) > 0 {
				traceId = vals[0]
			}
		}

		if len(traceId) == 0 {
			traceId = trace.GenTraceId()
		}
		ctx2.TraceId = traceId
		return next(ctx,ctx2)
	}
}
