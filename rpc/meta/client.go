package meta

import (
	"context"

	"github.com/fumeboy/pome/registry"
	"google.golang.org/grpc"
)

type ClientMeta struct {
	//调用方名字
	Caller string
	//服务提供方
	ServiceName string
	//调用的方法
	Method string
	//调用方集群
	CallerCluster string
	//服务提供方集群
	ServiceCluster string
	//TraceID
	TraceID string
	//环境
	Env string
	//调用方IDC
	CallerIDC string
	//服务提供方IDC
	ServiceIDC string
	//当前节点
	CurNode *registry.Node
	//历史选择节点
	HistoryNodes []*registry.Node
	//服务提供方的节点列表
	AllNodes []*registry.Node
	//当前请求使用的连接
	Conn *grpc.ClientConn
}

type clientMetaContextKey struct{}

func GetClientMeta(ctx context.Context) *ClientMeta {
	meta, ok := ctx.Value(clientMetaContextKey{}).(*ClientMeta)
	if !ok {
		meta = &ClientMeta{}
	}

	return meta
}

func InitClientMeta(ctx context.Context, service, method, caller string) context.Context {
	meta := &ClientMeta{
		Method:      method,
		ServiceName: service,
		Caller:      caller,
	}
	return context.WithValue(ctx, clientMetaContextKey{}, meta)
}
