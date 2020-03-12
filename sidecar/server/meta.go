package server

import (
	"context"
	"google.golang.org/grpc"
)

type metaT struct {
	RetCtx context.Context
	ServiceName string
	Method      string
	Cluster     string
	TraceID     string
	Env         string
	ServerIP    string
	ClientIP    string
	IDC         string
	Conn        *grpc.ClientConn
}

type metaContextKey struct{}

func getMeta(ctx context.Context) ( *metaT) {
	return ctx.Value(metaContextKey{}).(*metaT)
}

func initMeta(ctx context.Context, service, method string) context.Context {
	meta := &metaT{
		Method:      method,
		ServiceName: service,
	}
	return context.WithValue(ctx, metaContextKey{}, meta)
}
