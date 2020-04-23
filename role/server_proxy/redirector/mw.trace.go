package redirector

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/manager/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/metadata"
)

func mw_trace(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		fmt.Println("server's traceMiddleware")
		//从ctx获取grpc的metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			//没有的话,	新建一个
			md = metadata.Pairs()
		} else{
			//md = md.Copy()
		}

		tracer := opentracing.GlobalTracer()
		parentSpanContext, err := tracer.Extract(opentracing.TextMap, trace.MDReaderWriter{md})
		if err != nil {
			ctx2.log.Warn("trace extract failed, parsing trace information: %v", err)
		}
		//开始追踪该方法
		serverSpan := tracer.StartSpan(
			ctx2.method,
			opentracing.ChildOf(parentSpanContext),
			ext.RPCServerOption(parentSpanContext),
			ext.SpanKindRPCServer,
		)
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		err = next(ctx,ctx2)
		//记录错误
		if err != nil {
			ext.Error.Set(serverSpan, true)
			serverSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
		}

		serverSpan.Finish()
		return
	}
}
