package sync

import (
	"context"
	"github.com/fumeboy/llog"

	"google.golang.org/grpc"
)

type mw_fn = func(ctx context.Context, ctx2 *ctxT) error

type ctxT struct {
	method   string
	trace_id string
	ret_conn *grpc.ClientConn
	ret_ctx  context.Context
	log      llog.CTX
}

func init_mw_ctx(method string) (*ctxT) {
	return &ctxT{
		method: method,
		log: llog.CTX{
			Fields: []*llog.KeyVal{},
		},
	}
}
