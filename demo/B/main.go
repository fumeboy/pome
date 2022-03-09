package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"pome/define"
	"pome/demo/proto"
	"strconv"
)

type server struct{}

var s proto.ServiceBbServer = &server{}

func (s *server) Do(ctx context.Context, request *proto.ServiceBbDoRequest) (*proto.ServiceBbDoResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInner), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := proto.NewServiceAaClient(conn)
	respA, err := clientA.Do2(context.TODO(), &proto.ServiceAaDoRequest{
		Num: 2020,
	})
	if err != nil {
		panic(err)
	}
	resp := &proto.ServiceBbDoResponse{
		NewStr: request.Str + " World! " + strconv.Itoa(int(respA.NewNum)),
	}
	return resp, nil
}
func (s *server) Do2(ctx context.Context, request *proto.ServiceBbDoRequest) (*proto.ServiceBbDoResponse, error) {
	resp := &proto.ServiceBbDoResponse{
		NewStr: "2000",
	}
	return resp, nil
}

func main() {
	srv := grpc.NewServer()
	proto.RegisterServiceBbServer(srv, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", define.ServicePort))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server B running")
	srv.Serve(lis)
	fmt.Println("server B done")
}
