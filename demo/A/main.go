package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"pome/define"
	"pome/demo/proto"
)

type server struct{}

var s proto.ServiceAaServer = &server{}

func (s *server) Do(ctx context.Context, request *proto.ServiceAaDoRequest) (*proto.ServiceAaDoResponse, error) {
	resp := &proto.ServiceAaDoResponse{
		NewNum: request.Num + 1,
	}
	return resp, nil
}

func main() {
	srv := grpc.NewServer()
	proto.RegisterServiceAaServer(srv, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", define.ServicePort))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server A running")
	srv.Serve(lis)
	fmt.Println("server A done")
}

