package conf

import (
	"fmt"
	"github.com/fumeboy/llog"
)

type logConf struct {
	Level      string
	Path       string
	ChanSize   int
	ConsoleLog bool
}

func (this *logConf) init() (err error) {
	filename := fmt.Sprintf("%s/%s.log", this.Path, config.NodeName)
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
