package settings

import (
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
	//custom
	Ins registry.Registry `yaml:"-"`
}
