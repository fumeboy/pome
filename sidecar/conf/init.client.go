package conf

import (
	"github.com/fumeboy/pome/sidecar/utils/loadbalance"
	"golang.org/x/time/rate"
	"time"
)

type ClientT struct {
	SwitchOn     bool
	ConnTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	LimitQps     int
	SidecarPort  int
	// custom
	Limiter *rate.Limiter           `yaml:"-"`
	Balance loadbalance.LoadBalance `yaml:"-"`
}


//var errInitClientConf = &wrong.E{
//	Code:    "InitClientConf",
//	Message: "处理客户端配置失败",
//}

func (this *ClientT) init() (err error) {
	if !this.SwitchOn {
		return
	}
	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	this.Balance = loadbalance.NewRandomBalance()
	return
}
