package main

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/demo/server/main/guestbook"
	"github.com/fumeboy/pome/demo/server/main/model"
	"github.com/fumeboy/llog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type add struct {
	email   string
	content string
	methodT
}

func (this *add) msgRead(r *guestbook.AddRequest) {
	this.email = r.Msg.GetEmail()
	this.content = r.Msg.GetContent()
}

func (this *add) msgExec(ctx context.Context) (err error) {
	msg := &model.Msg{
		Email:   this.email,
		Content: this.content,
		Timestamp: time.Now().Unix(),
	}
	return model.AddMsg(ctx, msg)
}

//检查请求参数，如果该函数返回错误，则Run函数不会执行
func (this *add) checkParams(r *guestbook.AddRequest) (err error) {
	if len(r.Msg.Email) == 0 || len(r.Msg.Content) == 0 {
		err = status.Errorf(codes.InvalidArgument, "add msg failed")
		return
	}
	return
}

//SayHello函数的实现
func (this *add) run(ctx context.Context, r *guestbook.AddRequest) (
	resp *guestbook.AddResponse, err error) {
	resp = &guestbook.AddResponse{}
	this.msgRead(r)
	err = this.msgExec(ctx)
	if err != nil {
		llog.Error("add msg failed, err:%v", err)
		return
	}
	return
}


func (this *serverT) Add(ctx context.Context, r*guestbook.AddRequest)(resp*guestbook.AddResponse, err error){
	fmt.Println("method: Add")
	ctrl := &add{}
	err = ctrl.checkParams(r)
	if err != nil {
		return
	}
	resp, err = ctrl.run(ctx, r)
	if err != nil {
		return
	}
	resp.Code = "abc"
	return
}
