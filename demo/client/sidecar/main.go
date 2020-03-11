package main

import (
	"fmt"
	"github.com/fumeboy/pome/sidecar/client"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"github.com/fumeboy/pome/sidecar/server"
	"google.golang.org/grpc"
	"net"
)

func main(){
	var bool_s = conf.Server != nil
	var bool_c = conf.Client != nil
	if bool_s {
		server.Init()
		sidecar_s := grpc.NewServer(
			grpc.CustomCodec(proxy.Codec()),
			grpc.UnknownServiceHandler(proxy.TransparentHandler(server.Director)))
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Server.SidecarPort))
		if err != nil {
			panic("client launch failed")
		}
		if bool_c {
			go sidecar_s.Serve(lis)
		}else{
			sidecar_s.Serve(lis)
		}
	}
	if bool_c {
		client.Init()
		sidecar_c := grpc.NewServer(
			grpc.CustomCodec(proxy.Codec()),
			grpc.UnknownServiceHandler(proxy.TransparentHandler(client.Director)))
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Client.SidecarPort))
		if err != nil {
			panic("client launch failed")
		}
		sidecar_c.Serve(lis)
	}
}