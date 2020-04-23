package settings

import "golang.org/x/time/rate"

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
