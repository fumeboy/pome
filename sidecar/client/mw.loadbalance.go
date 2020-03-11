package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"
	"github.com/fumeboy/pome/sidecar/middleware/loadbalance"
	"github.com/fumeboy/pome/util/logs"
)

func NewLoadBalanceMiddleware(balancer loadbalance.LoadBalance) middleware.Middleware {
	return func(next middleware.MiddlewareFn) middleware.MiddlewareFn {
		return func(ctx context.Context) (err error) {
			fmt.Println("client's loadBalanceMiddleware")
			//从ctx获取rpc的metadata
			rpcMeta := getMeta(ctx)
			if len(rpcMeta.AllNodes) == 0 {
				err = loadbalance.ErrNotHaveServiceInstance
				logs.Error(ctx, "not have instance")
				return
			}
			//生成loadbalance的上下文,用来过滤已经选择的节点
			ctx = loadbalance.WithBalanceContext(ctx)
			for {
				rpcMeta.CurNode, err = balancer.Select(ctx, rpcMeta.AllNodes)
				if err != nil {
					return
				}
				logs.Debug(ctx, "select node:%#v", rpcMeta.CurNode)
				rpcMeta.HistoryNodes = append(rpcMeta.HistoryNodes, rpcMeta.CurNode)
				if err = next(ctx); err != nil {
					//连接错误的话，进行重试
					if isConnFailed(err) {
						continue
					}
					return
				}
				break
			}
			return
		}
	}
}
