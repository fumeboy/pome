package settings

import (
	"github.com/fumeboy/pome/role/client_proxy/redirector/loadbalance"
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
