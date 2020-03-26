package async

import (
	"context"
	"github.com/fumeboy/pome/sidecar/conf"
	"github.com/fumeboy/pome/sidecar/proxy"
	"github.com/fumeboy/pome/sidecar/utils"
	"github.com/fumeboy/pome/sidecar/utils/mq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)


var (
	clientStreamDescForProxying = &grpc.StreamDesc{
		ServerStreams: false,
		ClientStreams: true,
	}
)

func handle(v *mq.MqMsg) (err error){
	var address = conf.Server.MainoneAddr
	ctx := metadata.NewOutgoingContext(context.TODO(), map[string][]string{})
	conn, _ := grpc.DialContext(ctx, address,grpc.WithInsecure(), grpc.WithCodec(proxy.Codec()))
	clientStream, _ := conn.NewStream(ctx, clientStreamDescForProxying,utils.FullName(conf.NodeName(), v.Method))
	err = clientStream.SendMsg(&proxy.Frame{v.Body})
	defer clientStream.CloseSend()
	return //TODO
}
