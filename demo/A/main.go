package main

import (
	"context"
	"fmt"
	"net"
	"pome/define"
	"pome/demo/proto"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct{}

var s proto.ServiceAaServer = &server{}

func (s *server) Msg(ctx context.Context, request *proto.ServiceAaMsgReq) (e *proto.EmptyA, e2 error) {
	fmt.Println("recv::", request.Context)
	return &proto.EmptyA{}, nil
}

func (s *server) GraceStopTest(ctx context.Context, request *proto.ServiceAaMsgReq) (e *proto.ServiceAaMsgReq, e2 error) {
	fmt.Println("recv::", request.Context)
	time.Sleep(20 * time.Second)
	return request, nil
}

func (s *server) MsgProxy(ctx context.Context, request *proto.ServiceAaMsgReq) (e *proto.EmptyA, e2 error) {
	fmt.Println("send::", request.Context)
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewServiceBbClient(conn)
	_, err = client.Msg(ctx, &proto.ServiceBbMsgReq{
		Context: request.Context,
	})
	if err != nil {
		panic(err)
	}
	return &proto.EmptyA{}, nil
}

func (s *server) Do(ctx context.Context, request *proto.ServiceAaDoRequest) (*proto.ServiceAaDoResponse, error) {
	fmt.Println("Do")
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client_inner := proto.NewServiceBbClient(conn)
	// A.Do -> B.Do
	resp_inner, err := client_inner.Do(ctx, &proto.ServiceBbDoRequest{
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

func (s *server) Do2(ctx context.Context, request *proto.ServiceAaDoRequest) (*proto.ServiceAaDoResponse, error) {
	fmt.Println("Do2")
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client_inner := proto.NewServiceBbClient(conn)
	// A.Do2 -> B.Do2
	resp_inner, err := client_inner.Do2(ctx, &proto.ServiceBbDoRequest{
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", define.ServicePortGRPC))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server A running")
	srv.Serve(lis)
}
