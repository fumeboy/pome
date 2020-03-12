package conf

import (
	"context"
	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/util"
	"golang.org/x/time/rate"
)

func (this *ServerT) init() (err error) {
	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	service := &registry.Service{
		Name: conf.NodeName,
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
