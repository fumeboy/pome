package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func Receiver(fn SyncDirectorT, fn2 AsyncDirectorT, port int){
	sidecar := grpc.NewServer(
		grpc.CustomCodec(Codec()),
		grpc.UnknownServiceHandler(ProxyHandler(fn,fn2)))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("client launch failed")
	}
	sidecar.Serve(lis)
}
