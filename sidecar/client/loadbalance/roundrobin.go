package loadbalance

import (
	"github.com/fumeboy/pome/registry"
)

type roundRobinBalance struct {
	name  string
	index int
}

func newRoundRobinBalance() LoadBalance {
	return &roundRobinBalance{
		name: "roundrobin",
	}
}

func (r *roundRobinBalance) Name() string {
	return r.name
}

func (r *roundRobinBalance) Select(ctx *lbCtx, nodes []*registry.Node) (node *registry.Node, err error) {

	if len(nodes) == 0 {
		err = errNotHaveServiceInstance
		return
	}

	defer func() {
		if node != nil {
			setSelected(ctx, node)
		}
	}()

	var newNodes = filterNodes(ctx, nodes)
	if len(newNodes) == 0 {
		err = errAllNodesFailed
		return
	}

	r.index = (r.index + 1) % len(nodes)
	node = nodes[r.index]
	return
}
