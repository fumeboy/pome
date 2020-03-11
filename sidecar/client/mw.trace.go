package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/middleware/trace"
	"github.com/fumeboy/pome/util/logs"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/metadata"
)

func traceMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("client's traceMiddleware")
		tracer := opentracing.GlobalTracer()
		var parentSpanCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentSpanCtx = parent.Context()
		}

		opts := []opentracing.StartSpanOption{
			opentracing.ChildOf(parentSpanCtx),
			ext.SpanKindRPCClient,
			opentracing.Tag{Key: string(ext.Component), Value: "pome"},
			opentracing.Tag{Key: trace.TraceID, Value: logs.GetTraceId(ctx)},
		}

		rpcMeta := getMeta(ctx)
		clientSpan := tracer.StartSpan(rpcMeta.ServiceName, opts...)

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		if err := tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, trace.MetadataTextMap(md)); err != nil {
			logs.Debug(ctx, "grpc_opentracing: failed serializing trace information: %v", err)
		}

		ctx = metadata.NewOutgoingContext(ctx, md)
		ctx = metadata.AppendToOutgoingContext(ctx, trace.TraceID, logs.GetTraceId(ctx))
		ctx = opentracing.ContextWithSpan(ctx, clientSpan)

		err = next(ctx)
		//记录错误
		if err != nil {
			ext.Error.Set(clientSpan, true)
			clientSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		}

		clientSpan.Finish()
		return
	}
}
