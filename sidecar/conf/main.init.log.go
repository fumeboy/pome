package conf

import (
	"fmt"
	"github.com/fumeboy/pome/util/logs"
)

func initLogger() (err error) {
	filename := fmt.Sprintf("%s/%s.log", conf.Log.Path, conf.NodeName)
	outputer, err := logs.NewFileOutputer(filename)
	if err != nil {
		return
	}

	level := logs.GetLogLevel(conf.Log.Level)
	logs.InitLogger(level, conf.Log.ChanSize, conf.NodeName)
	logs.AddOutputer(outputer)

	if conf.Log.ConsoleLog {
		logs.AddOutputer(logs.NewConsoleOutputer())
	}
	return
}
