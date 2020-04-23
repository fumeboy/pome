package rpc

import (
	"github.com/fumeboy/pome/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)
const (
	async_flag = "pome-ifasync"
)
type (
	SyncHandlerT  = func(ctx context.Context, header *HeaderT) (context.Context, string, error)
	AsyncHandlerT = func(ctx context.Context, header *HeaderT, value []byte) (error)
	handler       struct {
		sync_handler  SyncHandlerT
		async_handler AsyncHandlerT
	}
	HeaderT struct {
		fullname    string
		NodeName    string
		ServiceName string
		MthName     string
		ifAsync     bool
	}
)

func ReadHeader(srv grpc.ServerStream) (header *HeaderT, ok bool) {
	header = &HeaderT{}
	fullname, ok := grpc.MethodFromServerStream(srv)
	if ok {
		header.fullname = fullname
		names, ok := utils.SplitFullName(fullname)
		if ok {
			header.NodeName = names[0]
			header.ServiceName = names[1]
			header.MthName = names[2]

			if md, ok := metadata.FromIncomingContext(srv.Context()); ok {
				if _, ok := md[async_flag]; ok {
					header.ifAsync = true
				}
			}
			return header, true
		}
	}
	return nil, false
}

func ProxyHandler(fn SyncHandlerT, fn2 AsyncHandlerT) grpc.StreamHandler {
	streamer := &handler{fn, fn2}
	return streamer.director
}

func (s *handler) director(srv interface{}, serverStream grpc.ServerStream) error {
	header, ok := ReadHeader(serverStream)
	if !ok {
		return grpc.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}
	if s.async_handler != nil {
		if header.ifAsync {
			return s.async(srv, serverStream, header)
		} else {
			return s.sync(srv, serverStream, header)
		}
	} else {
		return s.sync(srv, serverStream, header)
	}
}
