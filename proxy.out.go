package main

import (
	"google.golang.org/grpc/v2"
	"google.golang.org/grpc/v2/codes"
	"google.golang.org/grpc/v2/status"
)

/*
	我和其他节点的 sidecar 通信

	可以维持多个长连接， 多个 grpc 请求复用这些连接

	参考 etcd 租约设计，有 grpc 请求就延长连接的持有，超出固定时间空闲就关闭
*/
type configProxyOut struct {
	port int
}

func handlerOut(_ interface{}, s grpc.ServerStream) (err error) {
	// 获取请求流的目的 Method 名称
	fullMethodName, ok := grpc.MethodFromServerStream(s)
	if !ok {
		return status.Errorf(codes.Internal, "failed to get method from server stream")
	}

	// 根据请求流头部信息，判断出正确的对应的目的方
	endpoint := P.discoverer.direct(fullMethodName)
	if endpoint == nil {
		panic("")
	}
	conn, err := (*cNode)(endpoint).Conn(nodeContext)
	if err != nil {
		return err
	}

	err = s.(grpc.ServerStreamRedirect).Redirect(conn, fullMethodName)
	if err != nil {
		panic(err)
	}
	return nil
}
