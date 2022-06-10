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

func testStop(node_id int64) {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortCtrl), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := ctrl.NewSidecarClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	resp, err := clientA.Stop(ctx, &ctrl.StopReq{NodeId: node_id})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return
}
