package loadbalance

import (
	"fmt"

	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/registry"
)

type lbCtx struct {
	selectedNodeMap map[string]bool
}

func newBalanceContext() *lbCtx {
	return &lbCtx{
		selectedNodeMap: make(map[string]bool),
	}
}

func filterNodes(ctx *lbCtx, nodes []*registry.Node) []*registry.Node {
	var newNodes []*registry.Node
	if ctx == nil {
		return newNodes
	}

	for _, node := range nodes {
		addr := fmt.Sprintf("%s:%d", node.IP, node.Port)
		_, ok := ctx.selectedNodeMap[addr]
		if ok {
			llog.Debug("addr:%s ok", addr)
			continue
		}
		newNodes = append(newNodes, node)
	}

	return newNodes
}

func setSelected(sel *lbCtx, node *registry.Node) {
	addr := fmt.Sprintf("%s:%d", node.IP, node.Port)
	sel.selectedNodeMap[addr] = true
}
