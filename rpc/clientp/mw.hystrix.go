package clientp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/fumeboy/pome/rpc/meta"
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
	return func(ctx context.Context) (interface{}, error) {
		fmt.Println("client's hystrixMiddleware")
		rpcMeta := meta.GetClientMeta(ctx)
		var resp interface{}
		hystrixErr := hystrix.Do(rpcMeta.ServiceName, func() (err error) {
			resp, err = next(ctx)
			return err
		}, nil)

		if hystrixErr != nil {
			return nil, hystrixErr
		}

		return resp, hystrixErr
	}
}
