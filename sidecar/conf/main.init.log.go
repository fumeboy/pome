package conf

import (
	"fmt"
	"github.com/fumeboy/llog"
)

func initLogger() (err error) {
	filename := fmt.Sprintf("%s/%s.log", conf.Log.Path, conf.NodeName)
	outputer, err := llog.NewFileOutputer(filename)
	if err != nil {
		return
	}

	level := llog.GetLogLevel(conf.Log.Level)
	llog.InitLogger(level, conf.Log.ChanSize, &llog.KeyVal{"serviceName", conf.NodeName})
	llog.AddOutputer(outputer)

	if conf.Log.ConsoleLog {
		llog.AddOutputer(llog.NewConsoleOutputer())
	}
	return
}
