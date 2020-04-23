package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func receiver(fn SyncHandlerT, fn2 AsyncHandlerT, port int){
	sidecar := grpc.NewServer(
		grpc.CustomCodec(Codec()),
		grpc.UnknownServiceHandler(ProxyHandler(fn,fn2)))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("client launch failed")
	}
	sidecar.Serve(lis)
}

var ClientReceiver = receiver
var ServerReceiver = func(fn SyncHandlerT, port int){receiver(fn,nil,port)}
