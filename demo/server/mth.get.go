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
)

type get struct {
	offset uint32
	limit  uint32
	methodT
}

func (this *get) msgRead(r *guestbook.GetRequest) {
	this.offset = r.GetOffset()
	this.limit = r.GetLimit()
}

func (this *get) msgExec(ctx context.Context) (result []*model.Msg, err error) {
	return model.GetMsg(ctx, this.offset, this.limit)
}

//检查请求参数，如果该函数返回错误，则Run函数不会执行
func (this *get) checkParams(ctx context.Context, r *guestbook.GetRequest) (err error) {
	if r.GetOffset() < 0 || r.GetLimit() <= 0 {
		err = status.Errorf(codes.InvalidArgument, "add msg failed")
		return
	}
	return
}

//SayHello函数的实现
func (this *get) run(ctx context.Context, r *guestbook.GetRequest) (resp *guestbook.GetResponse, err error) {
	resp = &guestbook.GetResponse{}
	this.msgRead(r)
	result, err := this.msgExec(ctx)
	if err != nil {
		logs.Error(ctx, "get msg failed, err:%v", err)
		return
	}
	for _, one := range result {
		msg := &guestbook.Msg{
			Email:   one.Email,
			Content: one.Content,
		}
		resp.Msgs = append(resp.Msgs, msg)
	}
	return
}

func mwGet(request interface{}) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		fmt.Println("method: Get")
		r := request.(*guestbook.GetRequest)
		ctrl := &get{}
		err = ctrl.checkParams(ctx, r)
		if err != nil {
			return
		}
		resp, err = ctrl.run(ctx, r)
		return
	}
}

func (this *serverT) Get(ctx context.Context, r *guestbook.GetRequest) (resp *guestbook.GetResponse, err error) {
	ctx = meta.InitServerMeta(ctx, "guestbook", "get")
	mwResp, err := serverp.LoadMethod(mwGet(r))(ctx)
	if err != nil {
		return
	}
	resp = mwResp.(*guestbook.GetResponse)
	return
}
