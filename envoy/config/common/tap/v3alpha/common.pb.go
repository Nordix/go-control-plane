// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/tap/v3alpha/common.proto

package envoy_config_common_tap_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/service/tap/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc79ce0810a312a, []int{0}
}

func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig.Size(m)
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v3alpha.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v3alpha.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

type CommonExtensionConfig_TapDSConfig struct {
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc79ce0810a312a, []int{0, 0}
}

func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Size(m)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AdminConfig struct {
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc79ce0810a312a, []int{1}
}

func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminConfig.Unmarshal(m, b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return xxx_messageInfo_AdminConfig.Size(m)
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.config.common.tap.v3alpha.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.config.common.tap.v3alpha.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.config.common.tap.v3alpha.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/tap/v3alpha/common.proto", fileDescriptor_9bc79ce0810a312a)
}

var fileDescriptor_9bc79ce0810a312a = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0x6e, 0xd5, 0xdd, 0x49, 0x16, 0x64, 0x40, 0x94, 0x08, 0xba, 0x2e, 0x8b, 0x88,
	0xee, 0x26, 0xd0, 0x45, 0x84, 0x22, 0xc2, 0x4e, 0x15, 0x56, 0xbc, 0xa9, 0xad, 0xf7, 0xe5, 0x98,
	0x8c, 0x71, 0xa0, 0x99, 0x19, 0x92, 0x69, 0x68, 0xdf, 0x40, 0x7c, 0x04, 0x1f, 0xc6, 0x3b, 0x2f,
	0x7c, 0x23, 0xe9, 0x95, 0xcc, 0x9f, 0xb4, 0x29, 0x58, 0x22, 0xde, 0x25, 0x73, 0xbe, 0xf3, 0x9b,
	0xef, 0x3b, 0x67, 0xf0, 0x39, 0x13, 0xb5, 0x5c, 0x25, 0xa9, 0x14, 0x9f, 0x79, 0x9e, 0xa4, 0xb2,
	0x28, 0xa4, 0x48, 0x34, 0xa8, 0xa4, 0xbe, 0x84, 0xb9, 0xfa, 0x02, 0xfe, 0x28, 0x56, 0xa5, 0xd4,
	0x92, 0x3c, 0xb2, 0xea, 0xd8, 0xa9, 0x63, 0x5f, 0xd2, 0xa0, 0x62, 0xaf, 0x8e, 0x9e, 0x39, 0x1c,
	0x28, 0xde, 0x02, 0x94, 0xcc, 0xf3, 0x67, 0x95, 0x5c, 0x94, 0x29, 0x73, 0xb0, 0xe8, 0x89, 0xd3,
	0x56, 0xac, 0xac, 0x79, 0xca, 0xf6, 0x5e, 0x1a, 0x3d, 0x5e, 0x64, 0x0a, 0x12, 0x10, 0x42, 0x6a,
	0xd0, 0x5c, 0x8a, 0x2a, 0xa9, 0x59, 0x59, 0x71, 0x29, 0xb8, 0xc8, 0xbd, 0xe4, 0x5e, 0x0d, 0x73,
	0x9e, 0x81, 0x66, 0x49, 0xf3, 0xe1, 0x0a, 0xa7, 0x3f, 0xfa, 0xf8, 0xee, 0xc8, 0xc2, 0xde, 0x2e,
	0x35, 0x13, 0xa6, 0x6b, 0x64, 0xad, 0x90, 0x0f, 0x38, 0x84, 0xac, 0xe0, 0x62, 0xe6, 0xac, 0xdd,
	0x47, 0x27, 0xe8, 0x69, 0x30, 0x38, 0x8f, 0x3b, 0x12, 0xc6, 0x57, 0xa6, 0xc9, 0x31, 0xae, 0x6f,
	0x4c, 0x02, 0xd8, 0xfe, 0x92, 0xf7, 0xf8, 0xb8, 0x32, 0x16, 0xd3, 0x86, 0xd9, 0xb3, 0xcc, 0x33,
	0xcf, 0xf4, 0x41, 0x77, 0x68, 0x1f, 0x41, 0x6d, 0x58, 0xa1, 0x6b, 0xf6, 0xb0, 0x1c, 0x87, 0x1a,
	0x54, 0x56, 0x35, 0xac, 0x03, 0xcb, 0xa2, 0x9d, 0xfe, 0xfe, 0x9a, 0xd6, 0xdc, 0xf3, 0x66, 0xba,
	0x75, 0x6d, 0xc9, 0xee, 0x37, 0xfa, 0x85, 0x70, 0xd0, 0x2a, 0x93, 0x29, 0x3e, 0xde, 0xd9, 0x96,
	0x9f, 0x4c, 0x93, 0x02, 0x14, 0xdf, 0xdc, 0x65, 0x56, 0x1b, 0xbb, 0xb6, 0xa9, 0xd5, 0xd2, 0xc3,
	0x35, 0xbd, 0xf9, 0x0d, 0xf5, 0xee, 0xa0, 0x49, 0x98, 0xb6, 0xce, 0xc9, 0x03, 0xdc, 0x17, 0x50,
	0x30, 0x3b, 0x91, 0x23, 0x7a, 0x7b, 0x4d, 0xfb, 0x65, 0xef, 0x04, 0x4d, 0xec, 0xe1, 0xf0, 0xfa,
	0xfb, 0xcf, 0xaf, 0x0f, 0x47, 0xf8, 0x6a, 0x6f, 0xb4, 0xc1, 0x3f, 0x46, 0x1b, 0xbe, 0x32, 0xa4,
	0x97, 0xf8, 0xc5, 0x7f, 0x91, 0x28, 0xc1, 0x81, 0x4f, 0xae, 0x57, 0x8a, 0x91, 0x83, 0xdf, 0x14,
	0x9d, 0xe6, 0x38, 0x68, 0x6d, 0x9c, 0x9c, 0xe1, 0x23, 0x2f, 0xe1, 0x99, 0x1d, 0x4c, 0x2b, 0xcc,
	0xa1, 0xab, 0xbc, 0xcb, 0x86, 0x03, 0x63, 0xe3, 0x02, 0x3f, 0xef, 0xb2, 0xd1, 0x22, 0xd3, 0xd7,
	0xf8, 0x82, 0x4b, 0x37, 0x63, 0x55, 0xca, 0xe5, 0xaa, 0x6b, 0xd1, 0x34, 0x70, 0x21, 0xc6, 0xe6,
	0x9d, 0x8f, 0xd1, 0xa7, 0x5b, 0xf6, 0xc1, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x50,
	0xbe, 0xbc, 0xd1, 0x03, 0x00, 0x00,
}
