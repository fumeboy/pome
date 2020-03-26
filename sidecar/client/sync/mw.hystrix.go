package sync

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
)

// hystrix 熔断策略
var config = hystrix.CommandConfig{
	Timeout:               1000,
	MaxConcurrentRequests: 100,
	ErrorPercentThreshold: 25,
}

func init() {
	hystrix.ConfigureCommand("guestbook", config)
}

func mw_hystrix(next mw_fn) mw_fn {
	return func(ctx context.Context, ctx2 *ctxT) error {
		return hystrix.Do(ctx2.targetServiceName, func() error {
			return next(ctx, ctx2)
		}, nil)
	}
}
