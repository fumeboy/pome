package serverp

import (
	"fmt"
	"github.com/fumeboy/pome/util/logs"
)

func initLogger() (err error) {
	filename := fmt.Sprintf("%s/%s.log", conf.Log.Dir, conf.ServiceName)
	outputer, err := logs.NewFileOutputer(filename)
	if err != nil {
		return
	}

	level := logs.GetLogLevel(conf.Log.Level)
	logs.InitLogger(level, conf.Log.ChanSize, conf.ServiceName)
	logs.AddOutputer(outputer)

	if conf.Log.ConsoleLog {
		logs.AddOutputer(logs.NewConsoleOutputer())
	}
	return
}
