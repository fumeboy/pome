package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"pome/ctrl"
	"pome/define"
	"sync"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var logger *log.Entry

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	path := "./log.txt"
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	 WithMaxAge 和 WithRotationCount二者只能设置一个
	`WithMaxAge` 设置文件清理前的最长保存时间
	`WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	log.SetOutput(io.MultiWriter(os.Stdout, writer))
	logger = log.WithFields(log.Fields{})
}

var P = App{}

type App struct { // 全局实例
	discoverer *discoverer
	local      nodeGRPC
	ctx        context.Context
	cancel     func()
	stop       bool
	startCh    chan bool
}

type node struct { // service mesh 节点的 标识、地址、负载均衡、连接 信息
	id   int64
	addr string
	NodePartLoadBalance
	NodePartConn
	lock sync.Mutex
}

func (n *node) init(id int64, addr string) *node {
	n.id = id
	n.addr = addr
	return n
}

func main() {
	fmt.Println("sidecar working...")
	defer fmt.Println("sidecar exit!!!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	lnCtrl, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortCtrl))
	if err != nil {
		panic(err)
	}
	srvCtrl := grpc.NewServer()
	ctrl.RegisterSidecarServer(srvCtrl, SidecarSrv)

	go srvCtrl.Serve(lnCtrl)

	for {
		if P.stop {
			<-P.startCh
		}

		var nodeActiveContext, nodeActiveContextCancel = context.WithCancel(context.Background()) // 在线、掉线由该 ctx 控制
		P = App{
			ctx:     nodeActiveContext,
			cancel:  nodeActiveContextCancel,
			startCh: make(chan bool, 1),
		}

		var exitFinishChannel = make(chan bool, 2)
		go ExecUnitEtcd(nodeActiveContext, nodeActiveContextCancel, exitFinishChannel) // sidecar 由多个 ExecUnit 组成
		go ExecUnitProxy(nodeActiveContext, nodeActiveContextCancel, exitFinishChannel)
		// ...
		select {
		case <-c:
			nodeActiveContextCancel()
			srvCtrl.GracefulStop()
			lnCtrl.Close()
			os.Exit(0)
		case <-nodeActiveContext.Done():
			fmt.Println("reconnect...")
			for i := 0; i < 2; i++ {
				<-exitFinishChannel
			}
		}
	}
}
