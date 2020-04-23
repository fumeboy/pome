package server_proxy

import (
	"github.com/fumeboy/pome/courier/mq"
	"github.com/fumeboy/pome/courier/rpc"
	"github.com/fumeboy/pome/role/server_proxy/consumer"
	"github.com/fumeboy/pome/role/server_proxy/redirector"
	"github.com/fumeboy/pome/settings"
)

func RUN() {
	rpc.ServerReceiver(redirector.Director(), settings.ServerProxy.SidecarPort)
	mq.RegisteServerConsumerHandler(consumer.Handle)
}
