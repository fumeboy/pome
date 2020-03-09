package guestbook

import (
	"context"
	"fmt"
	"github.com/fumeboy/pome/rpc/clientp"
	"github.com/fumeboy/pome/rpc/meta"
	"github.com/fumeboy/pome/rpc/middleware"
)

const mth_add_name = "add"

func (this *clientT) Add(ctx context.Context, r *AddRequest) (resp *AddResponse, err error) {
	mkResp, err := this.Call(ctx, mth_add_name, add_mw(r))
	if err != nil {
		return nil, err
	}
	resp, ok := mkResp.(*AddResponse)
	if !ok {
		err = fmt.Errorf("invalid resp, not *guestbook.AddResponse")
		return nil, err
	}

	return resp, err
}

func add_mw(request interface{}) middleware.MiddlewareFn {
	return func(ctx context.Context) (resp interface{}, err error) {
		rpcMeta := meta.GetClientMeta(ctx)
		if rpcMeta.Conn == nil {
			return nil, clientp.ErrConnFailed
		}

		req := request.(*AddRequest)
		client := NewGuestBookServiceClient(rpcMeta.Conn)
		return client.Add(ctx, req)
	}
}
