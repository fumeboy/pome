package conf

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/util"
	"golang.org/x/time/rate"
)

type ServerT struct {
	SwitchOn    bool
	SidecarPort int
	MainonePort int
	Prometheus  prometheusConf
	LimitQps    int
	// custom
	Limiter     *rate.Limiter `yaml:"-"`
	MainoneAddr string        `yaml:"-"`
}

type prometheusConf struct {
	SwitchOn bool
	Port     int
}

func (this *ServerT) init() (err error) {
	if !this.SwitchOn {
		return
	}

	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	service := &registry.Service{
		Name: config.NodeName,
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
	this.MainoneAddr = fmt.Sprintf("%s:%d", "127.0.0.1", this.MainonePort)
	return
}
