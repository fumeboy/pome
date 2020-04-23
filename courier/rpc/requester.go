package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var clientStreamDescForProxying = &grpc.StreamDesc{
	ServerStreams: false,
	ClientStreams: true,
}

type RequestMsg struct {
	Header map[string][]string
	Body   []byte
}

func AsyncRequester(address string, target_name string, msg *RequestMsg) {
	ctx := metadata.NewOutgoingContext(context.TODO(), msg.Header)
	conn, _ := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithCodec(Codec()))
	clientStream, _ := conn.NewStream(ctx, clientStreamDescForProxying, target_name)
	_ = clientStream.SendMsg(&Frame{msg.Body})
	_ = clientStream.CloseSend()
	return //TODO
}
