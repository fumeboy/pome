package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"pome/define"
	"pome/demo/proto"
)

func main()  {
	fmt.Println("cli:")
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortOuter), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := proto.NewServiceAaClient(conn)
	respA, err := clientA.Do(context.TODO(), &proto.ServiceAaDoRequest{
		Num: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("A:", respA.NewNum)

	conn2, err := grpc.Dial(fmt.Sprintf("192.168.111.11:%d", define.SidecarPortOuter), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn2.Close()
	clientB := proto.NewServiceBbClient(conn2)
	respB, err := clientB.Do(context.TODO(), &proto.ServiceBbDoRequest{
		Str: "Hello",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("B:", respB.NewStr)
}
