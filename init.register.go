package main

import (
	"context"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/manager/registry"
	"github.com/fumeboy/pome/settings"
)

func init_registry() (err error) {
	var this = settings.Register
	if !this.SwitchOn {
		return
	}
	ctx := context.TODO()
	registryInst, err := registry.InitRegistry(ctx,
		this.RegisterName,
		registry.WithAddrs([]string{this.RegisterAddr}),
		registry.WithTimeout(this.Timeout),
		registry.WithRegistryPath(this.RegisterPath),
		registry.WithHeartBeat(this.HeartBeat),
	)
	if err != nil {
		llog.Error("init registry failed, err:%v", err)
		return
	}
	this.Ins = registryInst
	return
}
