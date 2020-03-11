package conf

import (
	"context"
	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/util"
	"github.com/fumeboy/pome/util/wrong"
	"golang.org/x/time/rate"
)

var errInitServerConf = &wrong.E{
	Code:    "InitServerConf",
	Message: "处理服务端配置失败",
}

func (this *ServerT) init() (err error) {
	if this.ServiceName == "" {
		return errInitServerConf
	}
	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	service := &registry.Service{
		Name: this.ServiceName,
	}

	ip, err := util.GetLocalIP()
	if err != nil {
		return
	}
	service.Nodes = append(service.Nodes, &registry.Node{
		IP:   ip,
		Port: this.SidecarPort,
	}, )

	Register.Register(context.TODO(), service)

	return
}
