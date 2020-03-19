package client

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 限流器
type Limiter interface {
	Allow() bool
}

func mw_rate_limit(l Limiter, next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) (err error) {
		allow := l.Allow()
		if !allow {
			err = status.Error(codes.ResourceExhausted, "rate limited")
			return
		}
		return next(ctx, ctx2)
	}
}
