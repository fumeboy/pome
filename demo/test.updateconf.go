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

func testUpdateConfig(ip string) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, define.SidecarPortCtrl), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := ctrl.NewSidecarClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	resp, err := clientA.UpdateConfig(ctx, &ctrl.UCReq{
		Context: `{"KeepAlive": "1000","KeepAliveTimeout": "800","AllowService":"Debug"}`,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	testReadConfig(ip)
	return
}
