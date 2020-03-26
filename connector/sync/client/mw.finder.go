package client

import (
	"context"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/conf"
	"github.com/fumeboy/pome/connector/sync/client/loadbalance"
)
func mw_finder(balancer loadbalance.LoadBalance,next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		service, err := conf.Register.GetService(ctx, ctx2.targetServiceName)
		if err != nil {
			llog.Error("discovery service:%s failed, err:%v", ctx2.targetServiceName, err)
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
				return err
			}
			llog.Debug("select node:%#v", currentNode)
			ctx2.targetNode = currentNode
			if err = next(ctx,ctx2); err != nil {
				//连接错误的话，进行重试
				if isConnFailed(err) {
					continue
				}
				return err
			}
			break
		}
		return
	}
}