package conf

import (
	"fmt"
	"github.com/fumeboy/pome/util/logs"
)

func initLogger() (err error) {
	filename := fmt.Sprintf("%s/%s.log", conf.Log.Path, service_name)
	outputer, err := logs.NewFileOutputer(filename)
	if err != nil {
		return
	}

	level := logs.GetLogLevel(conf.Log.Level)
	logs.InitLogger(level, conf.Log.ChanSize, service_name)
	logs.AddOutputer(outputer)

	if conf.Log.ConsoleLog {
		logs.AddOutputer(logs.NewConsoleOutputer())
	}
	return
}
