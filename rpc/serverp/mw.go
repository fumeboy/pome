package serverp

import "github.com/fumeboy/pome/rpc/middleware"

var middlewareChain middleware.Middleware

func initMiddlewareFn()  {
	var mids []middleware.Middleware
	mids = append(mids, logMiddleware)
	if conf.Prometheus.SwitchOn {
		mids = append(mids, prometheusMiddleware)
	}
	if conf.Limit.SwitchOn {
		mids = append(mids, middleware.NewRateLimitMiddleware(server.limiter))
	}
	if conf.Trace.SwitchOn {
		mids = append(mids, tracerMiddleware)
	}
	if len(server.customMiddleware) != 0 {
		mids = append(mids, server.customMiddleware...)
	}
	middlewareChain = middleware.Chain(middleware.PrepareMiddleware, mids...)
}

func loadMethod(handle middleware.MiddlewareFn)middleware.MiddlewareFn{
	return middlewareChain(handle)
}