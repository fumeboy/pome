package middleware

import (
	"context"
	"github.com/fumeboy/pome/sidecar/middleware/trace"

	"google.golang.org/grpc/metadata"
)

func PrepareMiddleware(next MiddlewareFn) MiddlewareFn {
	return func(ctx context.Context) (err error) {
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
		ctx = trace.WithTraceId(ctx, traceId)
		return next(ctx)
	}
}
