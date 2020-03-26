package pome

import (
	"github.com/fumeboy/pome/conf"
	"sync"
)

func RUN(){
	conf.Init()
	var wg = &sync.WaitGroup{}
	if conf.Client.SwitchOn{
		wg.Add(1)
		go func() {
			initClient()
			wg.Done()
		}()
	}
	if conf.Server.SwitchOn{
		wg.Add(1)
		go func() {
			initServer()
			wg.Done()
		}()
	}
	wg.Wait()
}
