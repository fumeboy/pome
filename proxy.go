package main

import (
	"context"
	"fmt"
	"net"
	"pome/define"

	"google.golang.org/grpc"
)

func ExecUnitProxy(nodeActiveContext context.Context, nodeActiveContextCancel func(), exitChannel chan bool) {
	P.local = nodeGRPC{
		addr: "127.0.0.1",
	}

	lnINgrpc, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortOuterGRPC))
	if err != nil {
		panic(err)
	}
	srvIN := grpc.NewServer(grpc.UnknownServiceHandler(handlerInGRPC(nodeActiveContext)))
	go srvIN.Serve(lnINgrpc)

	lnINtcp, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortOuterTCP))
	if err != nil {
		panic(err)
	}
	go handlerInTCP(nodeActiveContext, lnINtcp)

	lnOUTgrpc, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortInnerGRPC))
	if err != nil {
		panic(err)
	}
	srvOUT := grpc.NewServer(grpc.UnknownServiceHandler(handlerOutGRPC(nodeActiveContext)))
	go srvOUT.Serve(lnOUTgrpc)

	lnOUTtcp, err := net.Listen("tcp", fmt.Sprintf(":%d", define.SidecarPortInnerTCP))
	if err != nil {
		panic(err)
	}
	go handlerOutTCP(nodeActiveContext, lnOUTtcp)

	fmt.Println("sidecar proxy listening")

	<-nodeActiveContext.Done()
	srvIN.GracefulStop()
	srvOUT.GracefulStop()
	lnINgrpc.Close()
	lnOUTgrpc.Close()
	lnINtcp.Close()
	lnOUTtcp.Close()
	exitChannel <- true
	fmt.Println("proxy down")
}
