package guestbook

import (
	"context"
	"fmt"
)

const sidecar_port = "9110"

type clientT struct {
	serviceName string
}

func new_client(serviceName string) *clientT {
	c :=  &clientT{
		serviceName: serviceName,
	}
	return c
}

func Run() {
	clientIns := new_client("guestbook")
	r := &AddRequest{
		Msg: &Msg{
			Email:   "test@qq.com",
			Content: "dkfdkfdkfd",
		},
	}
	_, err := clientIns.Add(context.TODO(), r)
	fmt.Println("add msg result:", err)

	getReq := &GetRequest{
		Offset: 0,
		Limit:  10,
	}
	result, err := clientIns.Get(context.TODO(), getReq)
	if err != nil {
		fmt.Println("get msg failed,", err)
		return
	}
	for _, msg := range result.Msgs {
		fmt.Println("email:", msg.Email, "content:", msg.Content)
	}
}
