package main

import (
	"fmt"
	_ "github.com/fumeboy/pome/manager/registry/etcd"
	"github.com/fumeboy/pome/settings"
	"sync"
)
var wg *sync.WaitGroup
func main() {
	fmt.Println("start")
	wg = &sync.WaitGroup{}
	var err error
	settings.Init()
	if err = init_logger(); err != nil {
		panic("conf log")
	}
	if err = init_registry(); err != nil {
		panic("conf register")
	}
	if err = init_trace(); err != nil {
		panic("conf trace")
	}
	if err = init_server_proxy(); err != nil {
		panic("conf server")
	}
	if err = init_client_proxy(); err != nil {
		panic("conf client")
	}
	if err = init_kafka(); err != nil {
		panic("conf kafka")
	}
	wg.Wait()
}
