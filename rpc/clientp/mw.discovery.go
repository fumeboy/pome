package clientp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"

	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/util/logs"
)

func newDiscoveryMiddleware(discovery registry.Registry) middleware.Middleware {
	return func(next middleware.MiddlewareFn) middleware.MiddlewareFn {
		return func(ctx context.Context) (resp interface{}, err error) {
			fmt.Println("client's DiscoveryMiddleware")
			//从ctx获取rpc的metadata
			rpcMeta := meta.GetClientMeta(ctx)
			if len(rpcMeta.AllNodes) > 0 {
				return next(ctx)
			}

			service, err := discovery.GetService(ctx, rpcMeta.ServiceName)
			if err != nil {
				logs.Error(ctx, "discovery service:%s failed, err:%v", rpcMeta.ServiceName, err)
				return
			}

			rpcMeta.AllNodes = service.Nodes
			resp, err = next(ctx)
			return
		}
	}
}
