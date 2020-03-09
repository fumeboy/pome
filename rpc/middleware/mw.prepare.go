package middleware

import (
	"context"
	"github.com/fumeboy/pome/rpc/middleware/trace"

	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc/metadata"
)

func PrepareMiddleware(next MiddlewareFn) MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
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
			traceId = logs.GenTraceId()
		}

		ctx = logs.WithFieldContext(ctx)
		ctx = logs.WithTraceId(ctx, traceId)
		resp, err = next(ctx)
		return
	}
}
