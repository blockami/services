// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/abi.proto

package abi

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

type AbiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Block   int32  `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *AbiRequest) Reset() {
	*x = AbiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_abi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbiRequest) ProtoMessage() {}

func (x *AbiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_abi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbiRequest.ProtoReflect.Descriptor instead.
func (*AbiRequest) Descriptor() ([]byte, []int) {
	return file_proto_abi_proto_rawDescGZIP(), []int{0}
}

func (x *AbiRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *AbiRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AbiRequest) GetBlock() int32 {
	if x != nil {
		return x.Block
	}
	return 0
}

type AbiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abi string `protobuf:"bytes,1,opt,name=abi,proto3" json:"abi,omitempty"`
}

func (x *AbiResponse) Reset() {
	*x = AbiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_abi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbiResponse) ProtoMessage() {}

func (x *AbiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_abi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbiResponse.ProtoReflect.Descriptor instead.
func (*AbiResponse) Descriptor() ([]byte, []int) {
	return file_proto_abi_proto_rawDescGZIP(), []int{1}
}

func (x *AbiResponse) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

var File_proto_abi_proto protoreflect.FileDescriptor

var file_proto_abi_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x62, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x62, 0x69, 0x22, 0x56, 0x0a, 0x0a, 0x41, 0x62, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1f,
	0x0a, 0x0b, 0x41, 0x62, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x62, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x62, 0x69, 0x32,
	0x39, 0x0a, 0x03, 0x41, 0x62, 0x69, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x62, 0x69, 0x12, 0x0f, 0x2e, 0x61, 0x62, 0x69, 0x2e, 0x41, 0x62, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x62, 0x69, 0x2e, 0x41, 0x62, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x62, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_abi_proto_rawDescOnce sync.Once
	file_proto_abi_proto_rawDescData = file_proto_abi_proto_rawDesc
)

func file_proto_abi_proto_rawDescGZIP() []byte {
	file_proto_abi_proto_rawDescOnce.Do(func() {
		file_proto_abi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_abi_proto_rawDescData)
	})
	return file_proto_abi_proto_rawDescData
}

var file_proto_abi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_abi_proto_goTypes = []interface{}{
	(*AbiRequest)(nil),  // 0: abi.AbiRequest
	(*AbiResponse)(nil), // 1: abi.AbiResponse
}
var file_proto_abi_proto_depIdxs = []int32{
	0, // 0: abi.Abi.ContractAbi:input_type -> abi.AbiRequest
	1, // 1: abi.Abi.ContractAbi:output_type -> abi.AbiResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_abi_proto_init() }
func file_proto_abi_proto_init() {
	if File_proto_abi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_abi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbiRequest); i {
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
		file_proto_abi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbiResponse); i {
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
			RawDescriptor: file_proto_abi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_abi_proto_goTypes,
		DependencyIndexes: file_proto_abi_proto_depIdxs,
		MessageInfos:      file_proto_abi_proto_msgTypes,
	}.Build()
	File_proto_abi_proto = out.File
	file_proto_abi_proto_rawDesc = nil
	file_proto_abi_proto_goTypes = nil
	file_proto_abi_proto_depIdxs = nil
}