package proxy

import (
	"context"
	"github.com/fumeboy/pome/sidecar/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io"
)

var (
	clientStreamDescForProxying = &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}
)

func (s *handler)sync(srv interface{}, serverStream grpc.ServerStream, serviceName, mthName string) (err error){
	outgoingCtx, backendConn, err := s.sync_director(serverStream.Context(), serviceName, mthName)
	if err != nil {
		return err
	}
	clientCtx, clientCancel := context.WithCancel(outgoingCtx)
	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDescForProxying, backendConn, utils.FullName(serviceName,mthName))
	if err != nil {
		return err
	}
	s2cErrChan := forwardServerToClient(serverStream, clientStream)
	c2sErrChan := forwardClientToServer(clientStream, serverStream)
	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2cErrChan:
			if s2cErr == io.EOF {
				// this is the happy case where the sender has encountered io.EOF, and won't be sending anymore./
				clientStream.CloseSend()
				break
			} else {
				clientCancel()
				return grpc.Errorf(codes.Internal, "failed proxying s2c: %v", s2cErr)
			}
		case c2sErr := <-c2sErrChan:
			serverStream.SetTrailer(clientStream.Trailer())
			if c2sErr != io.EOF {
				return c2sErr
			}
			return nil
		}
	}
	return grpc.Errorf(codes.Internal, "gRPC proxying should never reach this stage.")
}

func forwardClientToServer(src grpc.ClientStream, dst grpc.ServerStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &Frame{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(f); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}
			if i == 0 {
				md, err := src.Header()
				if err != nil {
					ret <- err
					break
				}
				if err := dst.SendHeader(md); err != nil {
					ret <- err
					break
				}
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}

func forwardServerToClient(src grpc.ServerStream, dst grpc.ClientStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &Frame{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(f); err != nil {
				ret <- err // this can be io.EOF which is happy case
				break
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}