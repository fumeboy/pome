package main

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/manager/registry"
	"github.com/fumeboy/pome/role/server_proxy"
	"github.com/fumeboy/pome/settings"
	"github.com/fumeboy/pome/utils"
	"golang.org/x/time/rate"
)

func init_server_proxy() (err error) {
	var this = settings.ServerProxy
	if !this.SwitchOn {
		return
	}

	this.Limiter = rate.NewLimiter(rate.Limit(this.LimitQps), this.LimitQps)
	service := &registry.Service{
		Name: settings.NodeName(),
	}

	ip, err := utils.GetLocalIP()
	if err != nil {
		return
	}
	service.Nodes = append(service.Nodes, &registry.Node{
		IP:   ip,
		Port: this.SidecarPort,
	}, )

	settings.Register.Ins.Register(context.TODO(), service)
	this.MainoneAddr = fmt.Sprintf("%s:%d", "127.0.0.1", this.MainonePort)

	wg.Add(1)
	go server_proxy.RUN()
	return
}
