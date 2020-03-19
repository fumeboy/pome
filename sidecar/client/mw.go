package client

import (
	"context"
	"github.com/fumeboy/llog"

	"github.com/fumeboy/pome/registry"
	"google.golang.org/grpc"
)

type mw_fn = func(ctx context.Context, ctx2 *ctxT) error

type ctxT struct {
	Log llog.CTX
	//服务提供方
	ServiceName string
	//调用的方法
	Method string
	TraceId string
	//当前节点
	CurNode *registry.Node
	//当前请求使用的连接
	Conn *grpc.ClientConn
	RetCtx context.Context
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
