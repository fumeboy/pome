package conf

import (
	"fmt"
	"github.com/fumeboy/conf"
)

// TODO 从配置中心获取配置
func Init() {
	fmt.Println("start")
	var err error
	conf.ReadConfig(config)
	if err = config.Log.init(); err != nil {
		panic("conf log")
	}
	if err = config.Register.init(); err != nil {
		panic("conf register")
	}
	if err = config.Trace.init(); err != nil {
		panic("conf trace")
	}
	if err = config.Kafka.init(); err != nil {
		panic("conf kafka")
	}
	if err = config.Server.init(); err != nil {
		panic("conf server")
	}
	if err = config.Client.init(); err != nil {
		panic("conf client")
	}
}
