package conf

import (
	"github.com/fumeboy/pome/sidecar/middleware/loadbalance"
	"golang.org/x/time/rate"
	"time"
)

var Client *ClientT

type ClientT struct {
	ConnTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	LimitQps     int
	SidecarPort  int
	// custom
	Limiter *rate.Limiter           `yaml:"-"`
	Balance loadbalance.LoadBalance `yaml:"-"`
}
