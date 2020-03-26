package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io"
)

func (s *handler)async(srv interface{}, serverStream grpc.ServerStream, serviceName, mthName string) (err error){
	f := &Frame{}
	bytes := []byte{}
	for i := 0; ; i++ {
		if err := serverStream.RecvMsg(f); err != nil {
			if err == io.EOF{
				break // success
			}else{
				return grpc.Errorf(codes.Internal, "bad")
			}
		}
		bytes = append(bytes, f.Payload...)
	}
	if err := s.async_director(serverStream.Context(), serviceName, mthName, bytes); err != nil{
		return err
	}
	return grpc.Errorf(codes.OK, "success")
}