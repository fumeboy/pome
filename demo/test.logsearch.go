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

// fmt.Println(logSearch("192.168.111.11", traceid, 0, 0))
func logSearch(ip string, trace_id int64, time_start int64, time_end int64) []string {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, define.SidecarPortCtrl), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientA := ctrl.NewSidecarClient(conn)
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	resp, err := clientA.SearchLogByTraceID(ctx, &ctrl.SearchLogByTraceIDRequest{
		TimeStart: time_start,
		TimeEnd:   time_end,
		TraceId:   trace_id,
	})
	if err != nil {
		panic(err)
	}
	return resp.LogRecords
}
