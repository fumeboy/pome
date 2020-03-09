package middleware

import (
	"context"
)

type MiddlewareFn func(ctx context.Context) (resp interface{}, err error)

type Middleware func(MiddlewareFn) MiddlewareFn

// Chain is a helper function for composing middlewares. Requests will
// traverse them in the order they're declared. That is, the first middleware
// is treated as the outermost middleware.
func Chain(outer Middleware, others ...Middleware) Middleware { // 中间件连成线
	return func(next MiddlewareFn) MiddlewareFn {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}
