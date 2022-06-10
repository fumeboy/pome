package main

import (
	"fmt"
	"pome/define"
	"pome/demo/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testRaw(in int32) {
	conn, err := grpc.Dial(fmt.Sprintf("192.168.111.10:%d", define.SidecarPortOuterGRPC), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	proto.NewServiceAaClient(conn)
}

func BenchmarkA(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testRaw(1)
		}
	})
}
