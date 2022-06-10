// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ctrl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SidecarClient is the client API for Sidecar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SidecarClient interface {
	SearchLogByTraceID(ctx context.Context, in *SearchLogByTraceIDRequest, opts ...grpc.CallOption) (*SearchLogByTraceIDResponse, error)
	Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopResp, error)
	Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*StartResp, error)
	UpdateConfig(ctx context.Context, in *UCReq, opts ...grpc.CallOption) (*UCResp, error)
	ReadConfig(ctx context.Context, in *RCReq, opts ...grpc.CallOption) (*RCResp, error)
}

type sidecarClient struct {
	cc grpc.ClientConnInterface
}

func NewSidecarClient(cc grpc.ClientConnInterface) SidecarClient {
	return &sidecarClient{cc}
}

func (c *sidecarClient) SearchLogByTraceID(ctx context.Context, in *SearchLogByTraceIDRequest, opts ...grpc.CallOption) (*SearchLogByTraceIDResponse, error) {
	out := new(SearchLogByTraceIDResponse)
	err := c.cc.Invoke(ctx, "/ctrl.Sidecar/SearchLogByTraceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) Stop(ctx context.Context, in *StopReq, opts ...grpc.CallOption) (*StopResp, error) {
	out := new(StopResp)
	err := c.cc.Invoke(ctx, "/ctrl.Sidecar/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) Start(ctx context.Context, in *StartReq, opts ...grpc.CallOption) (*StartResp, error) {
	out := new(StartResp)
	err := c.cc.Invoke(ctx, "/ctrl.Sidecar/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) UpdateConfig(ctx context.Context, in *UCReq, opts ...grpc.CallOption) (*UCResp, error) {
	out := new(UCResp)
	err := c.cc.Invoke(ctx, "/ctrl.Sidecar/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) ReadConfig(ctx context.Context, in *RCReq, opts ...grpc.CallOption) (*RCResp, error) {
	out := new(RCResp)
	err := c.cc.Invoke(ctx, "/ctrl.Sidecar/ReadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SidecarServer is the server API for Sidecar service.
// All implementations should embed UnimplementedSidecarServer
// for forward compatibility
type SidecarServer interface {
	SearchLogByTraceID(context.Context, *SearchLogByTraceIDRequest) (*SearchLogByTraceIDResponse, error)
	Stop(context.Context, *StopReq) (*StopResp, error)
	Start(context.Context, *StartReq) (*StartResp, error)
	UpdateConfig(context.Context, *UCReq) (*UCResp, error)
	ReadConfig(context.Context, *RCReq) (*RCResp, error)
}

// UnimplementedSidecarServer should be embedded to have forward compatible implementations.
type UnimplementedSidecarServer struct {
}

func (UnimplementedSidecarServer) SearchLogByTraceID(context.Context, *SearchLogByTraceIDRequest) (*SearchLogByTraceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLogByTraceID not implemented")
}
func (UnimplementedSidecarServer) Stop(context.Context, *StopReq) (*StopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSidecarServer) Start(context.Context, *StartReq) (*StartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedSidecarServer) UpdateConfig(context.Context, *UCReq) (*UCResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedSidecarServer) ReadConfig(context.Context, *RCReq) (*RCResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadConfig not implemented")
}

// UnsafeSidecarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SidecarServer will
// result in compilation errors.
type UnsafeSidecarServer interface {
	mustEmbedUnimplementedSidecarServer()
}

func RegisterSidecarServer(s grpc.ServiceRegistrar, srv SidecarServer) {
	s.RegisterService(&Sidecar_ServiceDesc, srv)
}

func _Sidecar_SearchLogByTraceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLogByTraceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).SearchLogByTraceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrl.Sidecar/SearchLogByTraceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).SearchLogByTraceID(ctx, req.(*SearchLogByTraceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrl.Sidecar/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).Stop(ctx, req.(*StopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrl.Sidecar/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).Start(ctx, req.(*StartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrl.Sidecar/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).UpdateConfig(ctx, req.(*UCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_ReadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).ReadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctrl.Sidecar/ReadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).ReadConfig(ctx, req.(*RCReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sidecar_ServiceDesc is the grpc.ServiceDesc for Sidecar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sidecar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ctrl.Sidecar",
	HandlerType: (*SidecarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchLogByTraceID",
			Handler:    _Sidecar_SearchLogByTraceID_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Sidecar_Stop_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Sidecar_Start_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _Sidecar_UpdateConfig_Handler,
		},
		{
			MethodName: "ReadConfig",
			Handler:    _Sidecar_ReadConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ctrl.proto",
}