package conf

import (
	//"github.com/fumeboy/pome/sidecar/config"
	//"github.com/fumeboy/pome/sidecar/utils/trace"
	"github.com/fumeboy/pome/sidecar/utils/trace"
)

type traceConf struct {
	SwitchOn   bool
	ReportAddr string
	SampleType string
	SampleRate float64
}

func (this *traceConf)init() (err error) {
	if !this.SwitchOn {
		return
	}
	return trace.InitTrace(config.NodeName, this.ReportAddr, this.SampleType, this.SampleRate)
}
