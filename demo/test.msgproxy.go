package main

import (
	"context"
	"fmt"
	"pome/define"
	"pome/demo/proto"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testmsgp() {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortOuterGRPC), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := proto.NewServiceAaClient(conn)

	// conn2, err := grpc.Dial(fmt.Sprintf("192.168.111.11:%d", define.SidecarPortOuterGRPC), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn2.Close()
	// clientB := proto.NewServiceBbClient(conn2)

	i := 0
	for i < 1 {
		i++
		// time.Sleep(1 * time.Second)
		_, err := clientA.MsgProxy(context.TODO(), &proto.ServiceAaMsgReq{
			Context: strconv.Itoa(i),
		})
		if err != nil {
			panic(err)
		}
		// _, err = clientB.MsgProxy(context.TODO(), &proto.ServiceBbMsgReq{
		// 	Context: strconv.Itoa(i),
		// })
		if err != nil {
			panic(err)
		}
	}
}
