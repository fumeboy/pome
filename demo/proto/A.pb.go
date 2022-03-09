// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/A.proto

package proto

import (
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
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x32, 0x8a,
	0x01, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x12, 0x3d, 0x0a, 0x02,
	0x44, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x61, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x44,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03, 0x44,
	0x6f, 0x32, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x61, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x61, 0x44,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	0, // 1: proto.ServiceAa.Do2:input_type -> proto.ServiceAaDoRequest
	1, // 2: proto.ServiceAa.Do:output_type -> proto.ServiceAaDoResponse
	1, // 3: proto.ServiceAa.Do2:output_type -> proto.ServiceAaDoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
