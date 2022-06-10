package main

import (
	"context"
	"fmt"
	// "pome/define"
	"pome/demo/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testmsg() {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", 20002), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := proto.NewServiceAaClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	_, err = clientA.Msg(ctx, &proto.ServiceAaMsgReq{
		Context: "1",
	})
	if err != nil {
		panic(err)
	}
}
