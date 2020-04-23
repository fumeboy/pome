package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io"
)

func (s *handler)async(srv interface{}, serverStream grpc.ServerStream, header *HeaderT) (err error){
	f := &Frame{}
	for i := 0; ; i++ {
		if err := serverStream.RecvMsg(f); err != nil {
			if err == io.EOF{
				break // success
			}else{
				return grpc.Errorf(codes.Internal, "bad")
			}
		}
		if err := s.async_handler(serverStream.Context(), header, f.Payload); err != nil{
			//TODO
		}
	}
	return grpc.Errorf(codes.OK, "success")
}