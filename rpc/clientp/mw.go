package clientp

import "github.com/fumeboy/pome/rpc/middleware"

var middlewareChain middleware.Middleware

func (this *client) chainMiddleware() {
	var mids []middleware.Middleware
	mids = append(
		mids,
		middleware.PrepareMiddleware,
		logMiddleware,
		traceMiddleware,
		prometheusMiddleware,
	)
	if this.limiter != nil {
		mids = append(mids, middleware.NewRateLimitMiddleware(this.limiter))
	}
	mids = append(
		mids,
		HystrixMiddleware,
		newDiscoveryMiddleware(this.register),
		NewLoadBalanceMiddleware(this.balance),
		finalMiddleware,
	)
	middlewareChain = middleware.Chain(mids[0], mids[1:]...)
}
