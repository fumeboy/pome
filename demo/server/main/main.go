package main

import (
	"fmt"
	"github.com/fumeboy/pome/demo/server/main/guestbook"
	"google.golang.org/grpc"
	"net"
)
func main() {
	srv := grpc.NewServer()
	guestbook.RegisterGuestBookServiceServer(srv, serverIns)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9100))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server running")
	srv.Serve(lis)
}
