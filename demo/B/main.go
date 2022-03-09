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
	client_inner := proto.NewServiceAaClient(conn)
	// B.Do -> A.Do2
	reqnum, _ := strconv.Atoi(request.Str)
	resp_inner, err := client_inner.Do2(context.TODO(), &proto.ServiceAaDoRequest{
		Num: int32(reqnum + 1),
	})
	if err != nil {
		panic(err)
	}
	resp := &proto.ServiceBbDoResponse{
		NewStr: strconv.Itoa(int(resp_inner.NewNum)),
	}
	return resp, nil
}
func (s *server) Do2(ctx context.Context, request *proto.ServiceBbDoRequest) (*proto.ServiceBbDoResponse, error) {
	reqnum, _ := strconv.Atoi(request.Str)
	resp := &proto.ServiceBbDoResponse{
		NewStr: strconv.Itoa(int(reqnum + 1)),
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
