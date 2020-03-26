package client

import (
	"context"
	"github.com/fumeboy/llog"

	"github.com/fumeboy/pome/manager/registry"
)

type mw_fn = func(ctx context.Context, ctx2 *ctxT) error

type ctxT struct {
	log llog.CTX
	//服务提供方
	targetServiceName string
	//调用的方法
	method   string
	trace_id string

	targetNode *registry.Node
	//当前请求使用的连接
	ret_address string
	ret_ctx     context.Context
}

func init_mw_ctx(service, method string) (*ctxT) {
	return &ctxT{
		method:            method,
		targetServiceName: service,
		log: llog.CTX{
			Fields: []*llog.KeyVal{},
		},
	}
}
