// Code generated by protoc-gen-go. DO NOT EDIT.
// source: guestbook.proto

package guestbook

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRequest struct {
	Msg                  *Msg     `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

type AddResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Msg struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{2}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Msg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type GetRequest struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetResponse struct {
	Msgs                 []*Msg   `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{4}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetMsgs() []*Msg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "guestbook.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "guestbook.AddResponse")
	proto.RegisterType((*Msg)(nil), "guestbook.Msg")
	proto.RegisterType((*GetRequest)(nil), "guestbook.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "guestbook.GetResponse")
}

func init() { proto.RegisterFile("guestbook.proto", fileDescriptor_fbda7dd58e0f267b) }

var fileDescriptor_fbda7dd58e0f267b = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x1b, 0x52, 0x8a, 0x7a, 0x51, 0x01, 0x59, 0x50, 0x45, 0x4c, 0xc1, 0x13, 0x53, 0x24,
	0x8a, 0x60, 0x60, 0x2b, 0x4b, 0xa6, 0x2e, 0xe6, 0x17, 0xd0, 0xf8, 0x6a, 0x45, 0x6d, 0x72, 0xa5,
	0x67, 0xf8, 0x01, 0xfc, 0x72, 0xe4, 0x4b, 0x4a, 0x5a, 0xd4, 0xcd, 0xef, 0x3d, 0x7f, 0xf6, 0x3b,
	0x1b, 0xae, 0xdc, 0x17, 0xb2, 0x5f, 0x12, 0xad, 0xf3, 0xed, 0x8e, 0x3c, 0xa9, 0xf1, 0x9f, 0xa1,
	0x73, 0x80, 0xb9, 0xb5, 0x06, 0x3f, 0x83, 0xa3, 0x32, 0x88, 0x6b, 0x76, 0x69, 0x94, 0x45, 0x0f,
	0xc9, 0xec, 0x32, 0xef, 0xb9, 0x05, 0x3b, 0x13, 0x22, 0x7d, 0x0f, 0x89, 0xec, 0xe7, 0x2d, 0x35,
	0x8c, 0x4a, 0xc1, 0xb0, 0x24, 0x8b, 0x42, 0x8c, 0x8d, 0xac, 0xf5, 0x33, 0xc4, 0x0b, 0x76, 0xea,
	0x06, 0xce, 0xb1, 0xfe, 0xa8, 0x36, 0x5d, 0xd6, 0x0a, 0x95, 0xc2, 0x45, 0x49, 0x8d, 0xc7, 0xc6,
	0xa7, 0x67, 0xe2, 0xef, 0xa5, 0x7e, 0x05, 0x28, 0xd0, 0xef, 0x9b, 0x4c, 0x61, 0x44, 0xab, 0x15,
	0xa3, 0x17, 0x7c, 0x62, 0x3a, 0x15, 0x4e, 0xdd, 0x54, 0x75, 0xd5, 0xd2, 0x13, 0xd3, 0x0a, 0xfd,
	0x08, 0x89, 0xb0, 0x5d, 0x2b, 0x0d, 0xc3, 0x9a, 0x1d, 0xa7, 0x51, 0x16, 0x9f, 0x98, 0x43, 0xb2,
	0xd9, 0x4f, 0x04, 0xd7, 0x45, 0xf0, 0xdf, 0x88, 0xd6, 0xef, 0xb8, 0xfb, 0xae, 0x4a, 0x54, 0x2f,
	0x10, 0xcf, 0xad, 0x55, 0xb7, 0x07, 0x44, 0xff, 0x3a, 0x77, 0xd3, 0xff, 0x76, 0x7b, 0x9d, 0x1e,
	0x04, 0xae, 0x40, 0x7f, 0xc4, 0xf5, 0xb3, 0x1c, 0x71, 0x07, 0x35, 0xf5, 0x60, 0x39, 0x92, 0xff,
	0x78, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xda, 0xfe, 0xe3, 0x82, 0xa2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GuestBookServiceClient is the client API for GuestBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuestBookServiceClient interface {
	//添加留言
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	//查看留言
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type guestBookServiceClient struct {
	cc *grpc.ClientConn
}

func NewGuestBookServiceClient(cc *grpc.ClientConn) GuestBookServiceClient {
	return &guestBookServiceClient{cc}
}

func (c *guestBookServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/guestbook.GuestBookService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestBookServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/guestbook.GuestBookService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestBookServiceServer is the server API for GuestBookService service.
type GuestBookServiceServer interface {
	//添加留言
	Add(context.Context, *AddRequest) (*AddResponse, error)
	//查看留言
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedGuestBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuestBookServiceServer struct {
}

func (*UnimplementedGuestBookServiceServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedGuestBookServiceServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterGuestBookServiceServer(s *grpc.Server, srv GuestBookServiceServer) {
	s.RegisterService(&_GuestBookService_serviceDesc, srv)
}

func _GuestBookService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestBookServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestbook.GuestBookService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestBookServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestBookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestBookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestbook.GuestBookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestBookServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuestBookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guestbook.GuestBookService",
	HandlerType: (*GuestBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _GuestBookService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GuestBookService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestbook.proto",
}