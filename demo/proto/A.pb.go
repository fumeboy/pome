// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/A.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServiceAaDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ServiceAaDoRequest) Reset() {
	*x = ServiceAaDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_A_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAaDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAaDoRequest) ProtoMessage() {}

func (x *ServiceAaDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_A_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAaDoRequest.ProtoReflect.Descriptor instead.
func (*ServiceAaDoRequest) Descriptor() ([]byte, []int) {
	return file_proto_A_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceAaDoRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ServiceAaDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewNum int32 `protobuf:"varint,1,opt,name=new_num,json=newNum,proto3" json:"new_num,omitempty"`
}

func (x *ServiceAaDoResponse) Reset() {
	*x = ServiceAaDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_A_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAaDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAaDoResponse) ProtoMessage() {}

func (x *ServiceAaDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_A_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAaDoResponse.ProtoReflect.Descriptor instead.
func (*ServiceAaDoResponse) Descriptor() ([]byte, []int) {
	return file_proto_A_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceAaDoResponse) GetNewNum() int32 {
	if x != nil {
		return x.NewNum
	}
	return 0
}

var File_proto_A_proto protoreflect.FileDescriptor

var file_proto_A_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x61, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2e,
	0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x44, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x32, 0x4a,
	0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x12, 0x3d, 0x0a, 0x02, 0x44,
	0x6f, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x61, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x44, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_A_proto_rawDescOnce sync.Once
	file_proto_A_proto_rawDescData = file_proto_A_proto_rawDesc
)

func file_proto_A_proto_rawDescGZIP() []byte {
	file_proto_A_proto_rawDescOnce.Do(func() {
		file_proto_A_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_A_proto_rawDescData)
	})
	return file_proto_A_proto_rawDescData
}

var file_proto_A_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_A_proto_goTypes = []interface{}{
	(*ServiceAaDoRequest)(nil),  // 0: proto.ServiceAaDoRequest
	(*ServiceAaDoResponse)(nil), // 1: proto.ServiceAaDoResponse
}
var file_proto_A_proto_depIdxs = []int32{
	0, // 0: proto.ServiceAa.Do:input_type -> proto.ServiceAaDoRequest
	1, // 1: proto.ServiceAa.Do:output_type -> proto.ServiceAaDoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_A_proto_init() }
func file_proto_A_proto_init() {
	if File_proto_A_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_A_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAaDoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_A_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAaDoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_A_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_A_proto_goTypes,
		DependencyIndexes: file_proto_A_proto_depIdxs,
		MessageInfos:      file_proto_A_proto_msgTypes,
	}.Build()
	File_proto_A_proto = out.File
	file_proto_A_proto_rawDesc = nil
	file_proto_A_proto_goTypes = nil
	file_proto_A_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceAaClient is the client API for ServiceAa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAaClient interface {
	Do(ctx context.Context, in *ServiceAaDoRequest, opts ...grpc.CallOption) (*ServiceAaDoResponse, error)
}

type serviceAaClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAaClient(cc grpc.ClientConnInterface) ServiceAaClient {
	return &serviceAaClient{cc}
}

func (c *serviceAaClient) Do(ctx context.Context, in *ServiceAaDoRequest, opts ...grpc.CallOption) (*ServiceAaDoResponse, error) {
	out := new(ServiceAaDoResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAa/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAaServer is the server API for ServiceAa service.
type ServiceAaServer interface {
	Do(context.Context, *ServiceAaDoRequest) (*ServiceAaDoResponse, error)
}

// UnimplementedServiceAaServer can be embedded to have forward compatible implementations.
type UnimplementedServiceAaServer struct {
}

func (*UnimplementedServiceAaServer) Do(context.Context, *ServiceAaDoRequest) (*ServiceAaDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}

func RegisterServiceAaServer(s *grpc.Server, srv ServiceAaServer) {
	s.RegisterService(&_ServiceAa_serviceDesc, srv)
}

func _ServiceAa_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAaDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAaServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAa/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAaServer).Do(ctx, req.(*ServiceAaDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceAa",
	HandlerType: (*ServiceAaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _ServiceAa_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/A.proto",
}
