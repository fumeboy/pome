package serverp

import (
	"fmt"
	"github.com/fumeboy/pome/util"
	"github.com/fumeboy/pome/util/tag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var conf = &confT{
	Port: 8080,
	Prometheus: prometheusConf{
		SwitchOn: true,
		Port:     8081,
	},
	ServiceName: "server",
	Log: logConf{
		Level:      "debug",
		Dir:        "./util/logs/",
		ChanSize:   10000,
		ConsoleLog: true,
	},
	Register: registerConf{
	},
	Limit: limitConf{
		SwitchOn: true,
		QPSLimit: 50000,
	},
}

type confT struct {
	Port        int            `yaml:"port"`
	Prometheus  prometheusConf `yaml:"prometheus"`
	ServiceName string         `yaml:"service_name"`
	Register    registerConf   `yaml:"register"`
	Log         logConf        `yaml:"log"`
	Limit       limitConf      `yaml:"limit"`
	Trace       traceConf      `yaml:"trace"`

	//内部的配置项
	ConfigFile string `yaml:"-"`
}

type traceConf struct {
	SwitchOn   bool    `yaml:"switch_on"`
	ReportAddr string  `yaml:"report_addr"`
	SampleType string  `yaml:"sample_type"`
	SampleRate float64 `yaml:"sample_rate"`
}

type limitConf struct {
	QPSLimit int  `yaml:"qps"`
	SwitchOn bool `yaml:"switch_on"`
}

type prometheusConf struct {
	SwitchOn bool `yaml:"switch_on"`
	Port     int  `yaml:"port"`
}

type registerConf struct {
	SwitchOn     bool          `yaml:"switch_on"`
	RegisterPath string        `yaml:"register_path"`
	Timeout      time.Duration `yaml:"timeout"`
	HeartBeat    int64         `yaml:"heart_beat"`
	RegisterName string        `yaml:"register_name"`
	RegisterAddr string        `yaml:"register_addr"`
}

type logConf struct {
	Level      string `yaml:"level"`
	Dir        string `yaml:"path"`
	ChanSize   int    `yaml:"chan_size"`
	ConsoleLog bool   `yaml:"console_log"`
}


func initConfig() (err error) {
	exeFilePath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return
	}
	if runtime.GOOS == "windows" {
		exeFilePath = strings.Replace(exeFilePath, "\\", "/", -1)
	}
	var ConfigFile = path.Join(strings.ToLower(exeFilePath),"..", fmt.Sprintf("%s.yaml", util.GetEnv()))

	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, tag.Q(conf))
	if err != nil {
		return
	}
	return
}