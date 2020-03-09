package middleware

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 限流器
type Limiter interface {
	Allow() bool
}

func NewRateLimitMiddleware(l Limiter) Middleware {
	return func(next MiddlewareFn) MiddlewareFn {
		return func(ctx context.Context) (resp interface{}, err error) {
			allow := l.Allow()
			if !allow {
				err = status.Error(codes.ResourceExhausted, "rate limited")
				return
			}

			return next(ctx)
		}
	}
}
