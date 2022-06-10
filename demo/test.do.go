package main

import (
	"context"
	"fmt"
	"pome/define"
	"pome/demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// fmt.Println(testDo(1))
// fmt.Println(testDo(10))
func testDo(in int32) int32 {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortOuterGRPC), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := proto.NewServiceAaClient(conn)
	resp, err := clientA.Do(context.Background(), &proto.ServiceAaDoRequest{
		Num: in,
	})
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return resp.NewNum
}
