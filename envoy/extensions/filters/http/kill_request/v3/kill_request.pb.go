// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/extensions/filters/http/kill_request/v3/kill_request.proto

package kill_requestv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/Nordix/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// On which direction should the filter check for the `kill_request_header`.
// Default to `REQUEST` if unspecified.
type KillRequest_Direction int32

const (
	KillRequest_REQUEST  KillRequest_Direction = 0
	KillRequest_RESPONSE KillRequest_Direction = 1
)

// Enum value maps for KillRequest_Direction.
var (
	KillRequest_Direction_name = map[int32]string{
		0: "REQUEST",
		1: "RESPONSE",
	}
	KillRequest_Direction_value = map[string]int32{
		"REQUEST":  0,
		"RESPONSE": 1,
	}
)

func (x KillRequest_Direction) Enum() *KillRequest_Direction {
	p := new(KillRequest_Direction)
	*p = x
	return p
}

func (x KillRequest_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KillRequest_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_enumTypes[0].Descriptor()
}

func (KillRequest_Direction) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_enumTypes[0]
}

func (x KillRequest_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KillRequest_Direction.Descriptor instead.
func (KillRequest_Direction) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescGZIP(), []int{0, 0}
}

// Configuration for KillRequest filter.
type KillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The probability that a Kill request will be triggered.
	Probability *v3.FractionalPercent `protobuf:"bytes,1,opt,name=probability,proto3" json:"probability,omitempty"`
	// The name of the kill request header. If this field is not empty, it will override the :ref:`default header <config_http_filters_kill_request_http_header>` name. Otherwise the default header name will be used.
	KillRequestHeader string                `protobuf:"bytes,2,opt,name=kill_request_header,json=killRequestHeader,proto3" json:"kill_request_header,omitempty"`
	Direction         KillRequest_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=envoy.extensions.filters.http.kill_request.v3.KillRequest_Direction" json:"direction,omitempty"`
}

func (x *KillRequest) Reset() {
	*x = KillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillRequest) ProtoMessage() {}

func (x *KillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillRequest.ProtoReflect.Descriptor instead.
func (*KillRequest) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescGZIP(), []int{0}
}

func (x *KillRequest) GetProbability() *v3.FractionalPercent {
	if x != nil {
		return x.Probability
	}
	return nil
}

func (x *KillRequest) GetKillRequestHeader() string {
	if x != nil {
		return x.KillRequestHeader
	}
	return ""
}

func (x *KillRequest) GetDirection() KillRequest_Direction {
	if x != nil {
		return x.Direction
	}
	return KillRequest_REQUEST
}

var File_envoy_extensions_filters_http_kill_request_v3_kill_request_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x2f,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x33, 0x1a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x0b, 0x4b, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x13, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01,
	0x02, 0xc8, 0x01, 0x00, 0x52, 0x11, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x42, 0xbe, 0x01,
	0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x4b,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescData = file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDesc
)

func file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDescData
}

var file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_goTypes = []interface{}{
	(KillRequest_Direction)(0),   // 0: envoy.extensions.filters.http.kill_request.v3.KillRequest.Direction
	(*KillRequest)(nil),          // 1: envoy.extensions.filters.http.kill_request.v3.KillRequest
	(*v3.FractionalPercent)(nil), // 2: envoy.type.v3.FractionalPercent
}
var file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.kill_request.v3.KillRequest.probability:type_name -> envoy.type.v3.FractionalPercent
	0, // 1: envoy.extensions.filters.http.kill_request.v3.KillRequest.direction:type_name -> envoy.extensions.filters.http.kill_request.v3.KillRequest.Direction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_init() }
func file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_init() {
	if File_envoy_extensions_filters_http_kill_request_v3_kill_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillRequest); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_kill_request_v3_kill_request_proto = out.File
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_rawDesc = nil
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_goTypes = nil
	file_envoy_extensions_filters_http_kill_request_v3_kill_request_proto_depIdxs = nil
}
