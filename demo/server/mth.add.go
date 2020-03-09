package main

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/demo/server/guestbook"
	"github.com/fumeboy/pome/demo/server/model"
	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/rpc/middleware"
	"github.com/fumeboy/pome/rpc/serverp"
	"github.com/fumeboy/pome/util/logs"
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
		logs.Error(ctx, "add msg failed, err:%v", err)
		return
	}
	return
}

func mwAdd(request interface{})middleware.MiddlewareFn{
	r := request.(*guestbook.AddRequest)
	return func(ctx context.Context) (resp interface{}, err error){
		fmt.Println("method: Add")
		ctrl := &add{}
		err = ctrl.checkParams(r)
		if err != nil {
			return
		}
		resp, err = ctrl.run(ctx, r)
		return
	}
}

func (this *serverT) Add(ctx context.Context, r*guestbook.AddRequest)(resp*guestbook.AddResponse, err error){
	ctx = meta.InitServerMeta(ctx,"guestbook", "add")
	mwResp, err := serverp.LoadMethod(mwAdd(r))(ctx)
	if err != nil {
		return
	}
	resp = mwResp.(*guestbook.AddResponse)
	return
}
