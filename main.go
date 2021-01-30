package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var P = App{}
type App struct { // 全局实例
	discoverer     *discoverer
	local cNode
}

type node struct { // service mesh 节点的 标识、地址、负载均衡、连接 信息
	id int64
	addr string
	NodePartLoadBalance
	NodePartConn
}

func (n *node) init(id int64, addr string) *node{
	n.id = id
	n.addr = addr
	n.Weight = 1
	return n
}

var rootContext, rootContextCancel = context.WithCancel(context.Background()) // 程序退出由该 ctx 控制
var nodeActiveContext, nodeActiveContextCancel = context.WithCancel(rootContext) // 在线、掉线由该 ctx 控制

func main(){
	fmt.Println("sidecar working...")
	go ExecUnitEtcd() // sidecar 由多个 ExecUnit 组成
	go ExecUnitProxy()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for _ = range c{
		rootContextCancel()
		fmt.Println("sidecar done")
		os.Exit(0)
	}
}