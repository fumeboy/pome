package loadbalance

import (
	"math/rand"

	"github.com/fumeboy/pome/manager/registry"
)

type randomBalance struct {
	name string
}

func newRandomBalance() LoadBalance {
	return &randomBalance{
		name: "random",
	}
}

func (r *randomBalance) Name() string {
	return r.name
}

func (r *randomBalance) Select(ctx *lbCtx, nodes []*registry.Node) (node *registry.Node, err error) {

	if len(nodes) == 0 {
		err = errNotHaveServiceInstance
		return
	}

	defer func() {
		if node != nil {
			setSelected(ctx, node)
		}
	}()

	var newNodes  = filterNodes(ctx, nodes)
	if len(newNodes) == 0 {
		err = errAllNodesFailed
		return
	}

	var totalWeight int
	for _, val := range newNodes {
		if val.Weight == 0 {// 0 是 int 的初始值，表示该 node 第一次进行分配，对其进行权重初始化
			val.Weight = defaultNodeWeight
		}
		totalWeight += val.Weight
	}

	curWeight := rand.Intn(totalWeight)
	curIndex := -1
	for index, node := range nodes {
		curWeight -= node.Weight
		if curWeight < 0 {
			curIndex = index
			break
		}
	}

	if curIndex == -1 {
		err = errAllNodesFailed
		return
	}

	node = nodes[curIndex]
	return
}
