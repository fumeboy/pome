package main

import (
	//"github.com/fumeboy/pome/sidecar/config"
	//"github.com/fumeboy/pome/sidecar/utils/trace"
	"github.com/fumeboy/pome/manager/trace"
	"github.com/fumeboy/pome/settings"
)

func init_trace() (err error) {
	var this = settings.Trace
	if !this.SwitchOn {
		return
	}
	return trace.InitTrace(settings.NodeName(), this.ReportAddr, this.SampleType, this.SampleRate)
}
