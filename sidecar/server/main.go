package server

import (
	"fmt"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"github.com/fumeboy/pome/sidecar/server/async"
	syncH "github.com/fumeboy/pome/sidecar/server/sync"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Init(){
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		sidecar := grpc.NewServer(
			grpc.CustomCodec(proxy.Codec()),
			grpc.UnknownServiceHandler(proxy.ProxyHandler(syncH.Director(),nil)))
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Server.SidecarPort))
		if err != nil {
			panic("client launch failed")
		}
		sidecar.Serve(lis)
		wg.Done()
	}(wg)
	if conf.Kafka.SwitchOn {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			async.InitConsumer()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
