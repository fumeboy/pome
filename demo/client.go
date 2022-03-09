package main

import (
	"context"
	"fmt"
	"pome/define"
	"pome/demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// fmt.Println("cli:")
	// conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortOuter), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(conn.GetState())
	// defer conn.Close()
	// clientA := proto.NewServiceAaClient(conn)
	// respA, err := clientA.Do(context.TODO(), &proto.ServiceAaDoRequest{
	// 	Num: 1,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("A:", respA.NewNum)

	conn2, err := grpc.Dial(fmt.Sprintf("192.168.111.11:%d", define.SidecarPortOuter), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
