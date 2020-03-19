package loadbalance

import (
	"context"
	"fmt"

	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/llog"
)

type selectedNodes struct {
	selectedNodeMap map[string]bool
}

type loadbalanceFilterNodes struct{}

func withBalanceContext(ctx context.Context) context.Context {
	sel := &selectedNodes{
		selectedNodeMap: make(map[string]bool),
	}
	return context.WithValue(ctx, loadbalanceFilterNodes{}, sel)
}

func getSelectedNodes(ctx context.Context) *selectedNodes {
	sel, ok := ctx.Value(loadbalanceFilterNodes{}).(*selectedNodes)
	if !ok {
		return nil
	}
	return sel
}

func filterNodes(ctx context.Context, nodes []*registry.Node) []*registry.Node {
	var newNodes []*registry.Node
	sel := GetSelectedNodes(ctx)
	if sel == nil {
		return newNodes
	}

	for _, node := range nodes {
		addr := fmt.Sprintf("%s:%d", node.IP, node.Port)
		_, ok := sel.selectedNodeMap[addr]
		if ok {
			llog.Debug("addr:%s ok", addr)
			continue
		}
		newNodes = append(newNodes, node)
	}

	return newNodes
}

func setSelected(ctx context.Context, node *registry.Node) {
	sel := GetSelectedNodes(ctx)
	if sel == nil {
		return
	}
	addr := fmt.Sprintf("%s:%d", node.IP, node.Port)
	sel.selectedNodeMap[addr] = true
}
