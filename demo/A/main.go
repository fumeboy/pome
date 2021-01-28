package main

import (
	"context"
	"fmt"
	"pome/demo/A/guestbook"
)

const proxyAddress = "127.0.0.1:"+ "21000"
// 真实的 B 端地址是 127.0.0.1: 9100

func main() {
	r := &guestbook.AddRequest{
		Msg: &guestbook.Msg{
			Email:   "test@qq.com",
			Content: "dkfdkfdkfd",
		},
	}
	_, err := Add(context.TODO(), r)
	fmt.Println("add msg result:", err)

	getReq := &guestbook.GetRequest{
		Offset: 0,
		Limit:  10,
	}
	result, err := Get(context.TODO(), getReq)
	if err != nil {
		fmt.Println("get msg failed,", err)
		return
	}
	for _, msg := range result.Msgs {
		fmt.Println("email:", msg.Email, "content:", msg.Content)
	}
}
