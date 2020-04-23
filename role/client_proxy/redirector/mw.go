package redirector

import (
	"context"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/courier/rpc"

	"github.com/fumeboy/pome/manager/registry"
)

type mw_fn = func(ctx context.Context, ctx2 *ctxT) error

type ctxT struct {
	log      llog.CTX
	nodeName string
	//服务提供方
	serviceName string
	//调用的方法
	methodName string
	trace_id   string

	targetNode *registry.Node
	//当前请求使用的连接
	ret_address string
	ret_ctx     context.Context
}

func init_mw_ctx(header *rpc.HeaderT) (*ctxT) {
	return &ctxT{
		nodeName:    header.NodeName,
		methodName:  header.MthName,
		serviceName: header.ServiceName,
		log: llog.CTX{
			Fields: []*llog.KeyVal{},
		},
	}
}
