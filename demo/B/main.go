package main

import (
	"context"
	"fmt"
	"net"
	"pome/define"
	"pome/demo/proto"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct{}

var s proto.ServiceBbServer = &server{}

// B.MsgProxy -> A.Msg
func (s *server) MsgProxy(ctx context.Context, request *proto.ServiceBbMsgReq) (e *proto.EmptyB, e2 error) {
	fmt.Println("send::", request.Context)
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewServiceAaClient(conn)
	_, err = client.Msg(ctx, &proto.ServiceAaMsgReq{
		Context: request.Context,
	})
	if err != nil {
		panic(err)
	}
	return &proto.EmptyB{}, nil
}

func (s *server) Msg(ctx context.Context, request *proto.ServiceBbMsgReq) (e *proto.EmptyB, e2 error) {
	fmt.Println("recv::", request.Context)
	return &proto.EmptyB{}, nil
}

func (s *server) Do(ctx context.Context, request *proto.ServiceBbDoRequest) (*proto.ServiceBbDoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	reqnum, _ := strconv.Atoi(request.Str)
	if reqnum < 10 {
		// B.Do -> A.Do2
		conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client_inner := proto.NewServiceAaClient(conn)
		resp_inner, err := client_inner.Do2(ctx, &proto.ServiceAaDoRequest{
			Num: int32(reqnum + 1),
		})
		if err != nil {
			panic(err)
		}
		resp := &proto.ServiceBbDoResponse{
			NewStr: strconv.Itoa(int(resp_inner.NewNum)),
		}
		return resp, nil
	} else { // call C
		conn, err := net.Dial("tcp", fmt.Sprintf("10.0.0.3:%d", 8080))
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		conn.Write([]byte(strconv.Itoa(reqnum + 1)))

		// listen for reply
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			return nil, err
		}
		resp := &proto.ServiceBbDoResponse{
			NewStr: string(bs[:len]),
		}
		return resp, nil
	}
}
func (s *server) Do2(ctx context.Context, request *proto.ServiceBbDoRequest) (*proto.ServiceBbDoResponse, error) {
	reqnum, _ := strconv.Atoi(request.Str)
	resp := &proto.ServiceBbDoResponse{
		NewStr: strconv.Itoa(int(reqnum)),
	}
	return resp, nil
}

func main() {
	srv := grpc.NewServer()
	proto.RegisterServiceBbServer(srv, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", define.ServicePortGRPC))
	if err != nil {
		panic("failed launch server")
	}
	go srv.Serve(lis)

	lis2, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server B running")

	for {
		conn, err := lis2.Accept()
		if err != nil {
			panic("create conn failed")
		}
		go func(conn net.Conn) {
			fmt.Println("tcp conn accept")
			buf := make([]byte, 1024)
			for {
				// Read the incoming connection into the buffer.
				reqLen, err := conn.Read(buf)
				if err != nil {
					if err.Error() == "EOF" {
						fmt.Println("Disconned ")
						break
					} else {
						fmt.Println("Error reading:", err.Error())
						break
					}
				}
				str := string(buf[:reqLen])
				i, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}

				conn2, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", define.SidecarPortInnerGRPC), grpc.WithInsecure())
				if err != nil {
					panic(err)
				}
				defer conn2.Close()
				client_inner := proto.NewServiceAaClient(conn2)
				resp_inner, err := client_inner.Do2(context.TODO(), &proto.ServiceAaDoRequest{
					Num: int32(i + 1),
				})
				if err != nil {
					panic(err)
				}

				// Send a response back
				conn.Write([]byte(strconv.Itoa(int(resp_inner.NewNum))))
			}
			// Close the connection when you're done with it.
			conn.Close()
		}(conn)
	}
}
