package pome

import (
	"fmt"
	"github.com/fumeboy/pome/conf"
	async_client "github.com/fumeboy/pome/connector/async/client"
	"github.com/fumeboy/pome/connector/rpc"
	sync_client "github.com/fumeboy/pome/connector/sync/client"
	"sync"
)

func initClient(){
	fmt.Println("client init")
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		rpc.Receiver(sync_client.Director(), async_client.Director(), conf.Client.SidecarPort)
		wg.Done()
	}(wg)
	if conf.Kafka.SwitchOn {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			async_client.Init()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
