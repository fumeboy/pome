package main

import (
	"fmt"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/settings"
)

func init_logger() (err error) {
	var this = settings.Log
	filename := fmt.Sprintf("%s/%s.log", this.Path, settings.NodeName())
	outputer, err := llog.NewFileOutputer(filename)
	if err != nil {
		return
	}

	level := llog.GetLogLevel(this.Level)
	llog.InitLogger(level, this.ChanSize,
		//&llog.KeyVal{"serviceName", config.NodeName}
	)
	llog.AddOutputer(outputer)

	if this.ConsoleLog {
		llog.AddOutputer(llog.NewConsoleOutputer())
	}
	return
}
