package conf

import (
	"github.com/fumeboy/pome/registry"
	_ "github.com/fumeboy/pome/registry/etcd"
)

var config = &confT{}

var (
	Register registry.Registry
	NodeName = func() string{return config.NodeName}
    Client *ClientT = &config.Client
	Server *ServerT = &config.Server
	Kafka = &config.Kafka
)

type confT struct {
	NodeName  string

	Register   registerConf
	Trace      traceConf
	Log        logConf
	Kafka      kafkaConf

	Server     ServerT
	Client     ClientT
}