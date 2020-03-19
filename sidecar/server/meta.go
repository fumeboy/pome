package server

import (
	"context"
	"github.com/fumeboy/llog"
	"google.golang.org/grpc"
)

type metaT struct {
	RetCtx context.Context
	Log llog.CTX
	ServiceName string
	Method      string
	Cluster     string
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
		Log: llog.CTX{
			Fields: []*llog.KeyVal{},
		},
	}
	return context.WithValue(ctx, metaContextKey{}, meta)
}
