// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mafia/welcome.proto

package welcome

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

type MessageType int32

const (
	MessageType_WaitList           MessageType = 0
	MessageType_SessionConnectAddr MessageType = 1
	MessageType_InvalidName        MessageType = 2
	MessageType_RepeatedName       MessageType = 3
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "WaitList",
		1: "SessionConnectAddr",
		2: "InvalidName",
		3: "RepeatedName",
	}
	MessageType_value = map[string]int32{
		"WaitList":           0,
		"SessionConnectAddr": 1,
		"InvalidName":        2,
		"RepeatedName":       3,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_mafia_welcome_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_mafia_welcome_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_mafia_welcome_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_welcome_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_welcome_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_mafia_welcome_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type WelcomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType  MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=welcome.MessageType" json:"message_type,omitempty"`
	ResultString string      `protobuf:"bytes,2,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"`
}

func (x *WelcomeResponse) Reset() {
	*x = WelcomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_welcome_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeResponse) ProtoMessage() {}

func (x *WelcomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_welcome_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeResponse.ProtoReflect.Descriptor instead.
func (*WelcomeResponse) Descriptor() ([]byte, []int) {
	return file_mafia_welcome_proto_rawDescGZIP(), []int{1}
}

func (x *WelcomeResponse) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_WaitList
}

func (x *WelcomeResponse) GetResultString() string {
	if x != nil {
		return x.ResultString
	}
	return ""
}

var File_mafia_welcome_proto protoreflect.FileDescriptor

var file_mafia_welcome_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x2c,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x0f,
	0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2a, 0x56, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x57, 0x61, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x10, 0x03, 0x32, 0x4b, 0x0a, 0x07, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x77, 0x65,
	0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x57,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x6c, 0x63,
	0x6f, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mafia_welcome_proto_rawDescOnce sync.Once
	file_mafia_welcome_proto_rawDescData = file_mafia_welcome_proto_rawDesc
)

func file_mafia_welcome_proto_rawDescGZIP() []byte {
	file_mafia_welcome_proto_rawDescOnce.Do(func() {
		file_mafia_welcome_proto_rawDescData = protoimpl.X.CompressGZIP(file_mafia_welcome_proto_rawDescData)
	})
	return file_mafia_welcome_proto_rawDescData
}

var file_mafia_welcome_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mafia_welcome_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mafia_welcome_proto_goTypes = []interface{}{
	(MessageType)(0),        // 0: welcome.MessageType
	(*ConnectRequest)(nil),  // 1: welcome.ConnectRequest
	(*WelcomeResponse)(nil), // 2: welcome.WelcomeResponse
}
var file_mafia_welcome_proto_depIdxs = []int32{
	0, // 0: welcome.WelcomeResponse.message_type:type_name -> welcome.MessageType
	1, // 1: welcome.Welcome.Connect:input_type -> welcome.ConnectRequest
	2, // 2: welcome.Welcome.Connect:output_type -> welcome.WelcomeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mafia_welcome_proto_init() }
func file_mafia_welcome_proto_init() {
	if File_mafia_welcome_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mafia_welcome_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_mafia_welcome_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeResponse); i {
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
			RawDescriptor: file_mafia_welcome_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mafia_welcome_proto_goTypes,
		DependencyIndexes: file_mafia_welcome_proto_depIdxs,
		EnumInfos:         file_mafia_welcome_proto_enumTypes,
		MessageInfos:      file_mafia_welcome_proto_msgTypes,
	}.Build()
	File_mafia_welcome_proto = out.File
	file_mafia_welcome_proto_rawDesc = nil
	file_mafia_welcome_proto_goTypes = nil
	file_mafia_welcome_proto_depIdxs = nil
}
