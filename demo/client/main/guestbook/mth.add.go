package guestbook

import (
	"context"
	"github.com/fumeboy/llog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (this *clientT) Add(ctx context.Context, r *AddRequest) (resp *AddResponse, err error) {
	address := "127.0.0.1:"+ sidecar_port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		llog.Error("connect %s failed, err:%v", address, err)
		return nil, err
	}

	defer conn.Close()

	req := r
	client := NewGuestBookServiceClient(conn)
	ctx = metadata.NewOutgoingContext(ctx, map[string][]string{
		"pome-ifasync":{"1"},
	})
	resp, err = client.Add(ctx, req)
	return
}
