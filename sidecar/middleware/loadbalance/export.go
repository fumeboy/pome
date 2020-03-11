package loadbalance

import (
	"context"
	"github.com/fumeboy/pome/registry"
)

var (
	NewRandomBalance          = newRandomBalance
	NewRoundRobinBalance      = newRoundRobinBalance
	ErrNotHaveServiceInstance = errNotHaveServiceInstance
	WithBalanceContext        = withBalanceContext
	GetSelectedNodes          = getSelectedNodes
)

type (
	LoadBalance interface {
		Name() string
		Select(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error)
	}
)
