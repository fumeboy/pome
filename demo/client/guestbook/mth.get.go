package guestbook

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/clientp"
	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/rpc/middleware"
)

const mth_get_name = "get"

func (this *clientT) Get(ctx context.Context, r*GetRequest)(resp*GetResponse, err error){
	mkResp, err := this.Call(ctx, mth_get_name, get_mw(r))
	if err != nil {
		return nil, err
	}
	resp, ok := mkResp.(*GetResponse)
	if !ok {
		err = fmt.Errorf("invalid resp, not *guestbook.GetResponse")
		return nil, err
	}

	return resp, err
}


func get_mw(request interface{}) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		rpcMeta := meta.GetClientMeta(ctx)
		if rpcMeta.Conn == nil {
			return nil, clientp.ErrConnFailed
		}

		req := request.(*GetRequest)
		client := NewGuestBookServiceClient(rpcMeta.Conn)
		return client.Get(ctx, req)
	}
}
