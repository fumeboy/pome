package conf

import (
	"github.com/fumeboy/pome/sidecar/client/loadbalance"
	"golang.org/x/time/rate"
)

//var errInitClientConf = &wrong.E{
//	Code:    "InitClientConf",
//	Message: "处理客户端配置失败",
//}

func (this *ClientT) init() (err error) {
	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	this.Balance = loadbalance.NewRandomBalance()
	return
}
