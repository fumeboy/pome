package main

import (
	"context"
	"fmt"
	"pome/ctrl"
	"pome/define"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testStart() {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortCtrl), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := ctrl.NewSidecarClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	resp, err := clientA.Start(ctx, &ctrl.StartReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return
}
