package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"

	"github.com/afex/hystrix-go/hystrix"
)

// hystrix 熔断策略
var config = hystrix.CommandConfig{
	Timeout:               1000,
	MaxConcurrentRequests: 100,
	ErrorPercentThreshold: 25,
}

func init(){
	hystrix.ConfigureCommand("guestbook", config)
}

func HystrixMiddleware(next middleware.MiddlewareFn) middleware.MiddlewareFn {
	return func(ctx context.Context) (error) {
		fmt.Println("client's hystrixMiddleware")
		rpcMeta := getMeta(ctx)
		hystrixErr := hystrix.Do(rpcMeta.ServiceName, func() (err error) {
			return next(ctx)
		}, nil)
		return hystrixErr
	}
}
