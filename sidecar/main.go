package main

import (
	"github.com/fumeboy/pome/sidecar/client"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/server"
	"sync"
)

func main(){
	conf.Init()
	var wg = &sync.WaitGroup{}
	if conf.Client.SwitchOn{
		wg.Add(1)
		go func() {
			client.Init()
			wg.Done()
		}()
	}
	if conf.Server.SwitchOn{
		wg.Add(1)
		go func() {
			server.Init()
			wg.Done()
		}()
	}
	wg.Wait()
}
