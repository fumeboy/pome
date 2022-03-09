package main

import (
	"context"
	"fmt"
	"net"
	"pome/define"
	"pome/demo/proto"
	"strconv"

	"google.golang.org/grpc"
)

type server struct{}

var s proto.ServiceAaServer = &server{}

func (s *server) Do(ctx context.Context, request *proto.ServiceAaDoRequest) (*proto.ServiceAaDoResponse, error) {
	fmt.Println("Do")

	resp := &proto.ServiceAaDoResponse{
		NewNum: int32(2),
	}
	return resp, nil
}

func (s *server) Do2(ctx context.Context, request *proto.ServiceAaDoRequest) (*proto.ServiceAaDoResponse, error) {
	fmt.Println("Do2")
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInner), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client_inner := proto.NewServiceBbClient(conn)
	// A.Do2 -> B.Do2
	resp_inner, err := client_inner.Do2(context.TODO(), &proto.ServiceBbDoRequest{
		Str: strconv.Itoa(int(request.Num + 1)),
	})
	if err != nil {
		panic(err)
	}
	respnum, _ := strconv.Atoi(resp_inner.NewStr)
	resp := &proto.ServiceAaDoResponse{
		NewNum: int32(respnum),
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
