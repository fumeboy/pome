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

func testReadConfig(ip string) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, define.SidecarPortCtrl), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := ctrl.NewSidecarClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	resp, err := clientA.ReadConfig(ctx, &ctrl.RCReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return
}
