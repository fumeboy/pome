// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package mq

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MqMsg struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MqMsg) Reset()         { *m = MqMsg{} }
func (m *MqMsg) String() string { return proto.CompactTextString(m) }
func (*MqMsg) ProtoMessage()    {}
func (*MqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *MqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MqMsg.Unmarshal(m, b)
}
func (m *MqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MqMsg.Marshal(b, m, deterministic)
}
func (m *MqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MqMsg.Merge(m, src)
}
func (m *MqMsg) XXX_Size() int {
	return xxx_messageInfo_MqMsg.Size(m)
}
func (m *MqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MqMsg proto.InternalMessageInfo

func (m *MqMsg) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *MqMsg) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *MqMsg) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*MqMsg)(nil), "mq.MqMsg")
}

func init() {
	proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899)
}

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x2d, 0x54, 0xf2, 0xe5, 0x62, 0xf5, 0x2d, 0xf4,
	0x2d, 0x4e, 0x17, 0x92, 0xe0, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0x53, 0x4b, 0x32, 0xf2, 0x53,
	0x24, 0x98, 0xc0, 0x12, 0x50, 0x9e, 0x90, 0x10, 0x17, 0x4b, 0x52, 0x7e, 0x4a, 0xa5, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0x36, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x4e, 0x43, 0xa3, 0x66, 0x00, 0x00, 0x00,
}