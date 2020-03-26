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
		go client.Init()
	}
	if conf.Server.SwitchOn{
		wg.Add(1)
		go server.Init()
	}
	wg.Wait()
}
