package serverp

import (
	"fmt"
	"log"
	"net"

	"github.com/fumeboy/pome/registry"
	_ "github.com/fumeboy/pome/registry/etcd"
	"github.com/fumeboy/pome/rpc/middleware"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

type serverT struct {
	*grpc.Server
	limiter          *rate.Limiter
	register         registry.Registry
	customMiddleware []middleware.Middleware
}

var server = &serverT{
	Server: grpc.NewServer(),
}

func run() {
	/*
		if conf.Prometheus.SwitchOn {
			go func() {
				http.Handle("/metrics", promhttp.Handler())
				addr := fmt.Sprintf("0.0.0.0:%d", conf.Prometheus.Port)
				log.Fatal(http.ListenAndServe(addr, nil))
			}()
		}*/

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	server.Serve(lis)
}

func grpc_server() *grpc.Server {
	return server.Server
}
