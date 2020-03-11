package server

import (
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/middleware"
)

func Init() {
	var c = conf.Server
	var mids []middleware.Middleware
	mids = append(mids, logMiddleware)
	if c.Prometheus.SwitchOn {
		mids = append(mids, prometheusMiddleware)
	}
	if c.LimitQps != 0 {
		mids = append(mids, middleware.NewRateLimitMiddleware(c.Limiter))
	}
	if conf.IfTrace {
		mids = append(mids, tracerMiddleware)
	}
	middlewareChain := middleware.Chain(middleware.PrepareMiddleware, mids...)
	handle = middlewareChain(handle_fn)
}
