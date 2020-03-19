package server

import (
	"context"
	"github.com/fumeboy/llog"

	"google.golang.org/grpc"
)

type mw_fn = func(ctx context.Context, ctx2 *ctxT) error

type ctxT struct {
	ServiceName string
	Method string
	TraceId string
	Conn *grpc.ClientConn
	RetCtx context.Context
	Log llog.CTX
}

func init_mw_ctx(service, method string) (*ctxT) {
	return &ctxT{
		Method:      method,
		ServiceName: service,
		Log: llog.CTX{
			Fields: []*llog.KeyVal{},
		},
	}
}
