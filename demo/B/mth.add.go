package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pome/demo/B/guestbook"
	"pome/demo/B/model"
	"time"
)

func (this *server) Add(ctx context.Context, r*guestbook.AddRequest)(resp*guestbook.AddResponse, err error){
	email, content := r.Msg.GetEmail(), r.Msg.GetContent()
	if len(email) == 0 || len(content) == 0 {
		err = status.Errorf(codes.InvalidArgument, "add msg failed")
		return
	}

	msg := &model.Msg{
		Email:   email,
		Content: content,
		Timestamp: time.Now().Unix(),
	}
	if err = model.AddMsg(msg); err != nil{
		return
	}

	resp = &guestbook.AddResponse{Code: ""}
	return
}
