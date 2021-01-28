package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"pome/demo/A/guestbook"
)

func Get(ctx context.Context, r*guestbook.GetRequest)(resp*guestbook.GetResponse, err error){
	conn, err := grpc.Dial(proxyAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	req := r
	client := guestbook.NewGuestBookServiceClient(conn)
	ctx = metadata.NewOutgoingContext(ctx, map[string][]string{
	})
	resp, err = client.Get(ctx, req)
	return
}