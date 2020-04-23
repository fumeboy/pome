package client_proxy

import (
	"fmt"
	"github.com/fumeboy/pome/courier/rpc"
	"github.com/fumeboy/pome/role/client_proxy/producer"
	"github.com/fumeboy/pome/role/client_proxy/redirector"
	"github.com/fumeboy/pome/settings"
)

func RUN(){
	fmt.Println("client init")
	rpc.ClientReceiver(redirector.Director(), producer.Director(), settings.ClientProxy.SidecarPort)
}
