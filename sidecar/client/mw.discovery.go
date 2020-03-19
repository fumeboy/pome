package client

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/sidecar/middleware"

	"github.com/fumeboy/pome/registry"
)

func newDiscoveryMiddleware(discovery registry.Registry) middleware.Middleware {
	return func(next middleware.MiddlewareFn) middleware.MiddlewareFn {
		return func(ctx context.Context) (err error) {
			fmt.Println("client's DiscoveryMiddleware")
			//从ctx获取rpc的metadata
			rpcMeta := getMeta(ctx)
			if len(rpcMeta.AllNodes) > 0 {
				return next(ctx)
			}

			service, err := discovery.GetService(ctx, rpcMeta.ServiceName)
			if err != nil {
				rpcMeta.Log.Error("discovery service:%s failed, err:%v", rpcMeta.ServiceName, err)
				return
			}

			rpcMeta.AllNodes = service.Nodes
			return next(ctx)
		}
	}
}
