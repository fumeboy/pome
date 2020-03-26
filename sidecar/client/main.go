package client

import (
	"fmt"
	"github.com/fumeboy/pome/sidecar/client/async"
	"github.com/fumeboy/pome/sidecar/client/sync"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"google.golang.org/grpc"
	"net"
)

func Init(){
	fmt.Println("client init")
	sidecar := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.ProxyHandler(sync.Director(),async.Director())))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Client.SidecarPort))
	if err != nil {
		panic("client launch failed")
	}
	if conf.Kafka.SwitchOn {
		go async.InitProducer()
	}
	sidecar.Serve(lis)
}