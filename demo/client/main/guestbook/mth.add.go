package guestbook

import (
	"context"
	"github.com/fumeboy/pome/util/logs"
	"google.golang.org/grpc"
)

func (this *clientT) Add(ctx context.Context, r *AddRequest) (resp *AddResponse, err error) {
	address := "127.0.0.1:"+ sidecar_port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logs.Error(ctx, "connect %s failed, err:%v", address, err)
		return nil, err
	}

	defer conn.Close()

	req := r
	client := NewGuestBookServiceClient(conn)
	resp, err = client.Add(ctx, req)
	return
}
