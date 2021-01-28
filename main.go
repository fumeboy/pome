package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	discoverer     *discoverer
	local cNode
}

type node struct {
	id int64
	addr string
	NodePartLoadBalance
	NodePartConn
}

func IPaddress() string {
	return ""
}

func (n *node) init(id int64, addr string) *node{
	n.id = id
	n.addr = addr
	return n
}

var P = App{}
var rootContext, rootContextCancel = context.WithCancel(context.Background())
var nodeContext, nodeContextCancel = context.WithCancel(rootContext)

func main(){
	go ExecUnitEtcd(rootContext)
	go ExecUnitProxy(rootContext)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for _ = range c{
		rootContextCancel()
		os.Exit(0)
	}
}