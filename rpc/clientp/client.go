package clientp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware/loadbalance"
	"github.com/fumeboy/pome/rpc/middleware/trace"
	"sync"
	"time"

	"github.com/fumeboy/pome/registry"
	_ "github.com/fumeboy/pome/registry/etcd"
	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/rpc/middleware"
	"github.com/fumeboy/pome/util/logs"
	"golang.org/x/time/rate"
)

var initRegistryOnce sync.Once
var globalRegister registry.Registry

type client struct {
	opts     *confT
	register registry.Registry
	limiter  *rate.Limiter
	balance  loadbalance.LoadBalance
}

const (
	defaultConnTimeout  = 100 * time.Millisecond
	defaultReadTimeout  = time.Second
	defaultWriteTimeout = time.Second
)

func newClient(serviceName string) *client {
	conf.ServiceName = serviceName
	if err := initConfig(); err != nil{

	}
	fmt.Printf("init conf succ, conf:%#v\n", conf)
	client := &client{
		opts: conf,
		balance: loadbalance.NewRandomBalance(),
	}
	initRegistryOnce.Do(func() {
		ctx := context.TODO()
		var err error
		globalRegister, err = registry.InitRegistry(ctx,
			client.opts.RegisterName,
			registry.WithAddrs([]string{client.opts.RegisterAddr}),
			registry.WithTimeout(time.Second),
			registry.WithRegistryPath(client.opts.RegisterPath),
			registry.WithHeartBeat(10),
		)
		if err != nil {
			logs.Error(ctx, "init registry failed, err:%v", err)
			return
		}
	})

	if client.opts.MaxLimitQps > 0 {
		client.limiter = rate.NewLimiter(rate.Limit(client.opts.MaxLimitQps), client.opts.MaxLimitQps)
	}
	trace.InitTrace(client.opts.TraceClientServiceName, client.opts.TraceReportAddr, client.opts.TraceSampleType, client.opts.TraceSampleRate)
	client.register = globalRegister
	client.chainMiddleware()
	return client
}

func (this *client) getCaller(ctx context.Context) string {
	serverMeta := meta.GetServerMeta(ctx)
	if serverMeta == nil {
		return ""
	}
	return serverMeta.ServiceName
}

func (this *client) Call(ctx context.Context, method string, handle middleware.MiddlewareFn) (resp interface{}, err error) {
	caller := this.getCaller(ctx)
	ctx = meta.InitClientMeta(ctx, this.opts.ServiceName, method, caller)
	middlewareFunc := middlewareChain(handle)
	resp, err = middlewareFunc(ctx)
	if err != nil {
		return nil, err
	}
	return resp, err
}
