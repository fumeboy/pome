package main

import (
	"github.com/fumeboy/pome/demo/server/guestbook"
	"log"

	"github.com/fumeboy/pome/rpc/serverp"
)
func main() {
	err := serverp.Init()
	if err != nil {
		log.Fatal("init service failed, err:%v", err)
		return
	}
	guestbook.RegisterGuestBookServiceServer(serverp.GRPCServer(), serverIns)
	serverp.Run()
}
