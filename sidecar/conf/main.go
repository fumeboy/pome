package conf

import (
	"fmt"
	"github.com/fumeboy/pome/registry"
	"github.com/fumeboy/pome/util"
	"github.com/fumeboy/llog"
	"github.com/fumeboy/pome/util/tag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	_ "github.com/fumeboy/pome/registry/etcd"
)

var conf = &confT{}

var (
	Register registry.Registry
	IfTrace bool
)

type confT struct {
	NodeName  string
	Server     *ServerT
	Client     *ClientT
	Register   registerConf
	Trace      traceConf
	Log        logConf
}

type traceConf struct {
	SwitchOn   bool
	ReportAddr string
	SampleType string
	SampleRate float64
}

type registerConf struct {
	SwitchOn     bool
	RegisterPath string
	Timeout      time.Duration
	HeartBeat    int64
	RegisterName string
	RegisterAddr string
}

type logConf struct {
	Level      string
	Path       string
	ChanSize   int
	ConsoleLog bool
}


// TODO 从配置中心获取配置
func init() {
	fmt.Println("start")
	if exeFilePath, err := filepath.Abs(os.Args[0]); err != nil {
		panic("init config failed")
	} else {
		if runtime.GOOS == "windows" {
			exeFilePath = strings.Replace(exeFilePath, "\\", "/", -1)
		}
		var ConfigFile = path.Join(strings.ToLower(exeFilePath), "..", fmt.Sprintf("%s.yaml", util.GetEnv()))
		if data, err := ioutil.ReadFile(ConfigFile); err != nil {
			panic("init config failed2")
		} else {
			if err = yaml.Unmarshal(data, tag.Q(conf)); err != nil {
				panic("init config failed3")
			}else{
				Server = conf.Server
				Client = conf.Client
				if err = initLogger(); err != nil {
					llog.Error("init logger failed, err:%v", err)
					return
				}
				if err = initRegister(); err != nil {
					llog.Error("init register failed, err:%v", err)
					return
				}
				if err = initTrace(); err != nil {
					llog.Error("init tracing failed, err:%v", err)
				}
				if Server != nil {
					err = Server.init()
					if err != nil{
						panic("init config Server init")
					}
				}
				if Client != nil {
					err = Client.init()
					if err != nil{
						panic("init config Client init")
					}
				}
			}
		}
	}
}
