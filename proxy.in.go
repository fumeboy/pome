package main

import (
	"fmt"
	"google.golang.org/grpc/v2"
	"net"
	"pome/define"
)

func ExecUnitProxy() {
	P.local = cNode{
		addr: fmt.Sprintf("127.0.0.1:%d", define.ServicePort),
	}
	sIn := grpc.NewServer(grpc.UnknownServiceHandler(handlerIn))
	sOut := grpc.NewServer(grpc.UnknownServiceHandler(handlerOut))
	lIN, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortOuter))
	if err != nil {
		panic(err)
	}
	lOUT, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortInner))
	if err != nil {
		panic(err)
	}
	go sIn.Serve(lIN)
	go sOut.Serve(lOUT)
	// grpc server 貌似没法用 ctx 控制
	for {
		select {
		case <-rootContext.Done():
			// 正常退出
			sIn.Stop()
			sOut.Stop()
			lIN.Close()
			lOUT.Close()
			return
		case <-nodeActiveContext.Done():
			// 进入这个 case 说明是因为其他 ExecUnit 中止了
			panic("TODO")
		}
	}
}

/*
	流量流入 (IN)。其他节点的 sidecar 和 我 通信

	与service程序维持一个长连接， 多个 grpc 请求复用这个连接

	参考 etcd 租约设计，有 grpc 请求就延长连接的持有，超出固定时间空闲就关闭
*/

func handlerIn(_ interface{}, s grpc.ServerStream) (err error) {
	conn, err := P.local.Conn(nodeActiveContext)
	fmt.Println("IN:", serviceNameFrom(s))
	if err != nil {
		return err
	}
	err = s.(grpc.ServerStreamRedirect).Redirect(conn)
	if err != nil {
		panic(err)
	}
	return nil
}
