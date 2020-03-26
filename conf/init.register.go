package conf

import (
	"context"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/manager/registry"
	"time"
)


type registerConf struct {
	SwitchOn     bool
	RegisterPath string
	Timeout      time.Duration
	HeartBeat    int64
	RegisterName string
	RegisterAddr string
}

func (this *registerConf) init() (err error) {
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

	Register = registryInst
	return
}
