package guestbook

import (
	"context"
	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc"
)

const mth_get_name = "get"

func (this *clientT) Get(ctx context.Context, r*GetRequest)(resp*GetResponse, err error){
	address := "127.0.0.1:"+ sidecar_port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logs.Error(ctx, "connect %s failed, err:%v", address, err)
		return nil, err
	}

	defer conn.Close()

	req := r
	client := NewGuestBookServiceClient(conn)
	resp, err = client.Get(ctx, req)
	return
}