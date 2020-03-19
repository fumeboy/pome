package loadbalance

import (
	"github.com/fumeboy/pome/registry"
)

var (
	NewRandomBalance          = newRandomBalance
	NewRoundRobinBalance      = newRoundRobinBalance
	ErrNotHaveServiceInstance = errNotHaveServiceInstance
	NewBalanceContext        = newBalanceContext
)

type (
	LoadBalance interface {
		Name() string
		Select(ctx *lbCtx, nodes []*registry.Node) (node *registry.Node, err error)
	}
)
