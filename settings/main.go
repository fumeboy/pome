package settings

import "github.com/fumeboy/conf"

var config = &confT{}

var (
	NodeName    = func() string { return config.NodeName }
	ClientProxy = &config.Client
	ServerProxy = &config.Server
	Kafka       = &config.Kafka
	Register    = &config.Register
	Log         = &config.Log
	Trace       = &config.Trace
)

type confT struct {
	NodeName string

	Register registerConf
	Trace    traceConf
	Log      logConf
	Kafka    kafkaConf

	Server ServerT
	Client ClientT
}

func Init(){conf.ReadConfig(config)}
