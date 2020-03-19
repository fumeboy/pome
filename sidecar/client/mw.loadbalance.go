package client

import (
	"context"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/sidecar/client/loadbalance"
	"github.com/fumeboy/pome/sidecar/conf"
)
func mw_loadbalance(balancer loadbalance.LoadBalance,next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		service, err := conf.Register.GetService(ctx, ctx2.ServiceName)
		if err != nil {
			llog.Error("discovery service:%s failed, err:%v", ctx2.ServiceName, err)
			return err
		}
		nodes := service.Nodes
		if len(nodes) == 0 {
			err = loadbalance.ErrNotHaveServiceInstance
			llog.Error( "not have instance")
			return
		}
		//生成loadbalance的上下文,用来过滤已经选择的节点
		lbctx := loadbalance.NewBalanceContext()
		for {
			currentNode, err := balancer.Select(lbctx, nodes)
			if err != nil {
				return
			}
			llog.Debug("select node:%#v", currentNode)
			if err = next(ctx,ctx2); err != nil {
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