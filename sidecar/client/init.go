package client

import (
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/middleware"
)

func Init() {
	var c = conf.Client
	var mids = []middleware.Middleware{
		middleware.PrepareMiddleware,
		logMiddleware,
		traceMiddleware,
		prometheusMiddleware,
		middleware.NewRateLimitMiddleware(c.Limiter),
		HystrixMiddleware,
		newDiscoveryMiddleware(conf.Register),
		NewLoadBalanceMiddleware(c.Balance),
	}
	middlewareChain := middleware.Chain(mids[0], mids[1:]...)
	handle = middlewareChain(handle_fn)
}
