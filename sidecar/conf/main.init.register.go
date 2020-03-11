package conf

import (
	"context"
	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/util/logs"
)

func initRegister() (err error) {
	if !conf.Register.SwitchOn {
		return
	}

	ctx := context.TODO()
	registryInst, err := registry.InitRegistry(ctx,
		conf.Register.RegisterName,
		registry.WithAddrs([]string{conf.Register.RegisterAddr}),
		registry.WithTimeout(conf.Register.Timeout),
		registry.WithRegistryPath(conf.Register.RegisterPath),
		registry.WithHeartBeat(conf.Register.HeartBeat),
	)
	if err != nil {
		logs.Error(ctx, "init registry failed, err:%v", err)
		return
	}

	Register = registryInst
	return
}
