package server

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

func tracerMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (err error) {
		fmt.Println("server's traceMiddleware")
		//从ctx获取grpc的metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			//没有的话,新建一个
			md = metadata.Pairs()
		}

		tracer := opentracing.GlobalTracer()
		parentSpanContext, err := tracer.Extract(opentracing.HTTPHeaders, trace.MetadataTextMap(md))
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			logs.Warn(ctx, "trace extract failed, parsing trace information: %v", err)
		}

		serverMeta := getMeta(ctx)
		//开始追踪该方法
		serverSpan := tracer.StartSpan(
			serverMeta.Method,
			ext.RPCServerOption(parentSpanContext),
			ext.SpanKindRPCServer,
		)

		serverSpan.SetTag(trace.TraceID, logs.GetTraceId(ctx))
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		err = next(ctx)
		//记录错误
		if err != nil {
			ext.Error.Set(serverSpan, true)
			serverSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		}

		serverSpan.Finish()
		return
	}
}
