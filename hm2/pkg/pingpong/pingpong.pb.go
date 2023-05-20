// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: pkg/pingpong/pingpong.proto

package pingpong

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pingpong_pingpong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pingpong_pingpong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pingpong_pingpong_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pingpong_pingpong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pingpong_pingpong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pingpong_pingpong_proto_rawDescGZIP(), []int{1}
}

func (x *PongResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_pingpong_pingpong_proto protoreflect.FileDescriptor

var file_pkg_pingpong_pingpong_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x43, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67,
	0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pingpong_pingpong_proto_rawDescOnce sync.Once
	file_pkg_pingpong_pingpong_proto_rawDescData = file_pkg_pingpong_pingpong_proto_rawDesc
)

func file_pkg_pingpong_pingpong_proto_rawDescGZIP() []byte {
	file_pkg_pingpong_pingpong_proto_rawDescOnce.Do(func() {
		file_pkg_pingpong_pingpong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pingpong_pingpong_proto_rawDescData)
	})
	return file_pkg_pingpong_pingpong_proto_rawDescData
}

var file_pkg_pingpong_pingpong_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pingpong_pingpong_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: pingpong.PingRequest
	(*PongResponse)(nil), // 1: pingpong.PongResponse
}
var file_pkg_pingpong_pingpong_proto_depIdxs = []int32{
	0, // 0: pingpong.PingPong.Ping:input_type -> pingpong.PingRequest
	1, // 1: pingpong.PingPong.Ping:output_type -> pingpong.PongResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pingpong_pingpong_proto_init() }
func file_pkg_pingpong_pingpong_proto_init() {
	if File_pkg_pingpong_pingpong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pingpong_pingpong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_pkg_pingpong_pingpong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
			RawDescriptor: file_pkg_pingpong_pingpong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pingpong_pingpong_proto_goTypes,
		DependencyIndexes: file_pkg_pingpong_pingpong_proto_depIdxs,
		MessageInfos:      file_pkg_pingpong_pingpong_proto_msgTypes,
	}.Build()
	File_pkg_pingpong_pingpong_proto = out.File
	file_pkg_pingpong_pingpong_proto_rawDesc = nil
	file_pkg_pingpong_pingpong_proto_goTypes = nil
	file_pkg_pingpong_pingpong_proto_depIdxs = nil
}
