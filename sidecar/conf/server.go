package conf

import "golang.org/x/time/rate"

var Server *ServerT

type ServerT struct {
	SidecarPort int
	Port        int
	Prometheus  prometheusConf
	LimitQps    int
	// custom
	Limiter *rate.Limiter `yaml:"-"`
}

type prometheusConf struct {
	SwitchOn bool
	Port     int
}
