package serverp

import "github.com/fumeboy/pome/rpc/middleware/trace"

func initTrace() (err error) {
	if !conf.Trace.SwitchOn {
		return
	}
	return trace.InitTrace(conf.ServiceName, conf.Trace.ReportAddr, conf.Trace.SampleType, conf.Trace.SampleRate)
}
