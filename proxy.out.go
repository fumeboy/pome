package main

import (
	"google.golang.org/grpc/v2"
)

/*
	流量流出 (OUT)。我和其他节点的 sidecar 通信

	可以维持多个长连接， 多个 grpc 请求复用这些连接
	参考 etcd 租约设计，有 grpc 请求就延长连接的持有，超出固定时间空闲就关闭
*/

func handlerOut(_ interface{}, s grpc.ServerStream) (err error) {
	endpoint := P.discoverer.direct(serviceNameFrom(s))
	if endpoint == nil {
		panic(serviceNameFrom(s))
	}
	conn, err := (*cNode)(endpoint).Conn(nodeActiveContext)
	if err != nil {
		return err
	}
	err = s.(grpc.ServerStreamRedirect).Redirect(conn)
	if err != nil {
		panic(err)
	}
	return nil
}
