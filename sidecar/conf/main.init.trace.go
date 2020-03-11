package conf

import (
	//"github.com/fumeboy/pome/sidecar/conf"
	//"github.com/fumeboy/pome/sidecar/middleware/trace"
	"github.com/fumeboy/pome/sidecar/middleware/trace"
)

func initTrace() (err error) {
	if !conf.Trace.SwitchOn {
		return
	}
	IfTrace = true
	return trace.InitTrace(service_name, conf.Trace.ReportAddr, conf.Trace.SampleType, conf.Trace.SampleRate)
}
