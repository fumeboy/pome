package clientp

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
	ConnTimeout:            defaultConnTimeout,
	WriteTimeout:           defaultWriteTimeout,
	ReadTimeout:            defaultReadTimeout,
}

type confT struct {
	ConnTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration

	ServiceName  string
	//注册中心名字
	RegisterName string
	//注册中心地址
	RegisterAddr string
	//注册中心路径
	RegisterPath string
	//限流的qps
	MaxLimitQps int
	//trace report address
	TraceReportAddr string
	//trace sample type
	TraceSampleType string
	//trace sample rate
	TraceSampleRate float64
	//clientServiceName
	TraceClientServiceName string
}

func initConfig() (err error) {
	exeFilePath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return
	}
	if runtime.GOOS == "windows" {
		exeFilePath = strings.Replace(exeFilePath, "\\", "/", -1)
	}
	var ConfigFile = path.Join(strings.ToLower(exeFilePath),"..", conf.ServiceName, fmt.Sprintf("%s.yaml", util.GetEnv()))

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
