package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/v2"
	"google.golang.org/grpc/v2/codes"
	"google.golang.org/grpc/v2/status"
	"net"
)

func ExecUnitProxy(ctx context.Context) {
	sIn := grpc.NewServer(grpc.UnknownServiceHandler(handlerIn))
	sOut := grpc.NewServer(grpc.UnknownServiceHandler(handlerOut))
	lIN, err := net.Listen("tcp", fmt.Sprintf(":%d", CONFIG.configProxyIn.port))
	if err != nil {
		panic("")
	}
	lOUT, err := net.Listen("tcp", fmt.Sprintf(":%d", CONFIG.configProxyOut.port))
	if err != nil {
		panic("")
	}
	go sIn.Serve(lIN)
	go sOut.Serve(lOUT)
	for {
		select {
		case <-ctx.Done():
			sIn.Stop()
			sOut.Stop()
			lIN.Close()
			lOUT.Close()
			return
		}
	}
}

/*
	其他节点的 sidecar 和 我 通信

	与service程序维持一个长连接， 多个 grpc 请求复用这个连接

	参考 etcd 租约设计，有 grpc 请求就延长连接的持有，超出固定时间空闲就关闭
*/

type configProxyIn struct {
	port int
}

func handlerIn(_ interface{}, s grpc.ServerStream) (err error) {
	fullMethodName, ok := grpc.MethodFromServerStream(s)
	if !ok {
		return status.Errorf(codes.Internal, "failed to get method from server stream")
	}
	conn, err := P.local.Conn(nodeContext)
	if err != nil {
		return err
	}
	err = s.(grpc.ServerStreamRedirect).Redirect(conn, fullMethodName)
	if err != nil {
		panic(err)
	}
	return nil
}
