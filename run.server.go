package pome

import (
	"github.com/fumeboy/pome/conf"
	async_server "github.com/fumeboy/pome/connector/async/server"
	"github.com/fumeboy/pome/connector/rpc"
	sync_server "github.com/fumeboy/pome/connector/sync/server"
	"sync"
)

func initServer() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		rpc.Receiver(sync_server.Director(), nil, conf.Server.SidecarPort)
		wg.Done()
	}(wg)
	if conf.Kafka.SwitchOn {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			async_server.Init()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
