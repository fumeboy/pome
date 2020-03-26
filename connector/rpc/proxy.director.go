package rpc

import (
	"github.com/fumeboy/pome/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type (
	SyncDirectorT  = func(ctx context.Context, serviceName, mthName string) (context.Context, string, error)
	AsyncDirectorT = func(ctx context.Context, serviceName, MethodName string, value []byte) (error)
	handler        struct {
		sync_director  SyncDirectorT
		async_director AsyncDirectorT
	}
)

func ReadMetadata(srv grpc.ServerStream) (target_service_name string, target_method_name string, ifasync bool, ok bool) {
	fullname, ok := grpc.MethodFromServerStream(srv)
	if ok {
		names, ok := utils.SplitFullName(fullname)
		if ok {
			if md, ok := metadata.FromIncomingContext(srv.Context()); ok {
				if _, ok := md["pome-ifasync"]; ok {
					ifasync = true
				}
			}
			return names[0], names[1], ifasync, true
		}
	}
	return "", "", false, false
}

func ProxyHandler(fn SyncDirectorT, fn2 AsyncDirectorT) grpc.StreamHandler {
	streamer := &handler{fn, fn2}
	return streamer.handler
}

func (s *handler) handler(srv interface{}, serverStream grpc.ServerStream) error {
	serviceName, mthName, ifasync, ok := ReadMetadata(serverStream)
	if !ok {
		return grpc.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}
	if s.async_director != nil {
		if ifasync {
			return s.async(srv, serverStream, serviceName, mthName)
		} else {
			return s.sync(srv, serverStream, serviceName, mthName)
		}
	} else {
		return s.sync(srv, serverStream, serviceName, mthName)
	}
}
