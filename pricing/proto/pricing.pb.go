// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/pricing.proto

package pricing

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

type PricingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network   string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Contract  string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp int32  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Currency  string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *PricingRequest) Reset() {
	*x = PricingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pricing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingRequest) ProtoMessage() {}

func (x *PricingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pricing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingRequest.ProtoReflect.Descriptor instead.
func (*PricingRequest) Descriptor() ([]byte, []int) {
	return file_proto_pricing_proto_rawDescGZIP(), []int{0}
}

func (x *PricingRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *PricingRequest) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *PricingRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *PricingRequest) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PricingRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type PricingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalValue     float32 `protobuf:"fixed32,1,opt,name=total_value,json=totalValue,proto3" json:"total_value,omitempty"`
	UnitValue      float32 `protobuf:"fixed32,2,opt,name=unit_value,json=unitValue,proto3" json:"unit_value,omitempty"`
	AmountDecimals float32 `protobuf:"fixed32,3,opt,name=amount_decimals,json=amountDecimals,proto3" json:"amount_decimals,omitempty"`
}

func (x *PricingResponse) Reset() {
	*x = PricingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pricing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingResponse) ProtoMessage() {}

func (x *PricingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pricing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingResponse.ProtoReflect.Descriptor instead.
func (*PricingResponse) Descriptor() ([]byte, []int) {
	return file_proto_pricing_proto_rawDescGZIP(), []int{1}
}

func (x *PricingResponse) GetTotalValue() float32 {
	if x != nil {
		return x.TotalValue
	}
	return 0
}

func (x *PricingResponse) GetUnitValue() float32 {
	if x != nil {
		return x.UnitValue
	}
	return 0
}

func (x *PricingResponse) GetAmountDecimals() float32 {
	if x != nil {
		return x.AmountDecimals
	}
	return 0
}

var File_proto_pricing_proto protoreflect.FileDescriptor

var file_proto_pricing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x98,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x7a, 0x0a, 0x0f, 0x50, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x73, 0x32, 0x46, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x12, 0x3b, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pricing_proto_rawDescOnce sync.Once
	file_proto_pricing_proto_rawDescData = file_proto_pricing_proto_rawDesc
)

func file_proto_pricing_proto_rawDescGZIP() []byte {
	file_proto_pricing_proto_rawDescOnce.Do(func() {
		file_proto_pricing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pricing_proto_rawDescData)
	})
	return file_proto_pricing_proto_rawDescData
}

var file_proto_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pricing_proto_goTypes = []interface{}{
	(*PricingRequest)(nil),  // 0: pricing.PricingRequest
	(*PricingResponse)(nil), // 1: pricing.PricingResponse
}
var file_proto_pricing_proto_depIdxs = []int32{
	0, // 0: pricing.Pricing.Call:input_type -> pricing.PricingRequest
	1, // 1: pricing.Pricing.Call:output_type -> pricing.PricingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pricing_proto_init() }
func file_proto_pricing_proto_init() {
	if File_proto_pricing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pricing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingRequest); i {
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
		file_proto_pricing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingResponse); i {
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
			RawDescriptor: file_proto_pricing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pricing_proto_goTypes,
		DependencyIndexes: file_proto_pricing_proto_depIdxs,
		MessageInfos:      file_proto_pricing_proto_msgTypes,
	}.Build()
	File_proto_pricing_proto = out.File
	file_proto_pricing_proto_rawDesc = nil
	file_proto_pricing_proto_goTypes = nil
	file_proto_pricing_proto_depIdxs = nil
}
