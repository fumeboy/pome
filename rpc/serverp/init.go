package serverp

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/middleware"
	"github.com/fumeboy/pome/util/logs"
	"golang.org/x/time/rate"
)

func init_serv(m ...middleware.Middleware) (err error) {
	if len(m) > 0 {
		server.customMiddleware = m
	}
	if err = initConfig(); err != nil {
		return
	}
	fmt.Printf("init conf succ, conf:%#v\n", conf)
	//初始化限流器
	if conf.Limit.SwitchOn {
		server.limiter = rate.NewLimiter(rate.Limit(conf.Limit.QPSLimit), conf.Limit.QPSLimit)
	}
	if err = initLogger(); err != nil {
		logs.Error(context.TODO(), "init logger failed, err:%v", err)
		return
	}
	//初始化注册中心
	if err = initRegister(); err != nil {
		logs.Error(context.TODO(), "init register failed, err:%v", err)
		return
	}
	if err = initTrace(); err != nil {
		logs.Error(context.TODO(), "init tracing failed, err:%v", err)
	}
	initMiddlewareFn()
	return
}
