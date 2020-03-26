// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v3/xray.proto

package envoy_config_trace_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type XRayConfig struct {
	DaemonEndpoint       *v3.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	SegmentName          string            `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	SamplingRuleManifest *v3.DataSource    `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d491b3510a2e630, []int{0}
}

func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRayConfig.Unmarshal(m, b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return xxx_messageInfo_XRayConfig.Size(m)
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() *v3.SocketAddress {
	if m != nil {
		return m.DaemonEndpoint
	}
	return nil
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *v3.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
}

func init() {
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v3.XRayConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v3/xray.proto", fileDescriptor_6d491b3510a2e630) }

var fileDescriptor_6d491b3510a2e630 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x5a, 0x5c, 0xbc, 0x2e, 0xad, 0x11, 0xfd, 0x63, 0x0c, 0x6d, 0x55, 0x9b, 0x52,
	0x13, 0xc2, 0x0a, 0xac, 0x9b, 0x6f, 0x51, 0x92, 0x5b, 0x12, 0x8c, 0x0c, 0xc1, 0x37, 0x31, 0x96,
	0xc6, 0xca, 0x12, 0x69, 0x57, 0xec, 0xae, 0x84, 0x75, 0xcb, 0x31, 0x87, 0x3c, 0x41, 0x1e, 0x25,
	0xf7, 0x40, 0xae, 0x79, 0x9d, 0x9c, 0x82, 0x25, 0x99, 0x60, 0xa2, 0xdb, 0xc2, 0xfc, 0x7e, 0xec,
	0xf7, 0xcd, 0x10, 0x1b, 0x79, 0x21, 0x4a, 0x27, 0x14, 0x7c, 0xcd, 0x62, 0x47, 0x4b, 0x08, 0xd1,
	0x29, 0x5c, 0x67, 0x23, 0xa1, 0xa4, 0x99, 0x14, 0x5a, 0x58, 0xdf, 0x2b, 0x82, 0xd6, 0x04, 0xad,
	0x08, 0x5a, 0xb8, 0xc3, 0xd1, 0x9e, 0x18, 0x0a, 0x59, 0x79, 0x10, 0x45, 0x12, 0x95, 0xaa, 0xd5,
	0xe1, 0x9f, 0x56, 0x66, 0x05, 0x0a, 0x1b, 0xe0, 0x57, 0x1e, 0x65, 0xe0, 0x00, 0xe7, 0x42, 0x83,
	0x66, 0x82, 0x2b, 0x47, 0x69, 0xd0, 0xf9, 0xce, 0xff, 0xfb, 0x6e, 0x5c, 0xa0, 0x54, 0x4c, 0x70,
	0xc6, 0xe3, 0x06, 0xf9, 0x59, 0x40, 0xc2, 0x22, 0xd0, 0xe8, 0xec, 0x1e, 0xf5, 0x60, 0x74, 0x67,
	0x12, 0xb2, 0xf4, 0xa1, 0x3c, 0xae, 0x3e, 0xb7, 0xce, 0xc8, 0xd7, 0x08, 0x30, 0x15, 0x3c, 0x40,
	0x1e, 0x65, 0x82, 0x71, 0x3d, 0x30, 0x6c, 0x63, 0xd2, 0x9b, 0x8e, 0xe9, 0x5e, 0xbf, 0x6d, 0x48,
	0x5a, 0xb8, 0x74, 0x21, 0xc2, 0x6b, 0xd4, 0x47, 0x75, 0x1d, 0xff, 0x4b, 0xed, 0x9e, 0x36, 0xaa,
	0x75, 0x40, 0x3e, 0x2b, 0x8c, 0x53, 0xe4, 0x3a, 0xe0, 0x90, 0xe2, 0xc0, 0xb4, 0x8d, 0x49, 0xd7,
	0xfb, 0xf4, 0xe2, 0x7d, 0x94, 0x66, 0xdf, 0xf0, 0x7b, 0xcd, 0xf0, 0x02, 0x52, 0xb4, 0x2e, 0xc9,
	0x0f, 0x05, 0x69, 0x96, 0x30, 0x1e, 0x07, 0x32, 0x4f, 0x30, 0x48, 0x81, 0xb3, 0x35, 0x2a, 0x3d,
	0xf8, 0x50, 0x05, 0xb0, 0xdb, 0x03, 0x9c, 0x80, 0x86, 0x85, 0xc8, 0x65, 0x88, 0xfe, 0xb7, 0x9d,
	0xef, 0xe7, 0x09, 0x9e, 0x37, 0xf6, 0xec, 0xf0, 0xfe, 0xf1, 0xf6, 0xf7, 0x7f, 0xf2, 0xaf, 0xed,
	0x3c, 0x53, 0x48, 0xb2, 0x2b, 0xa0, 0x6f, 0xfd, 0xbd, 0xd9, 0xc3, 0xcd, 0xd3, 0x73, 0xc7, 0xec,
	0x9b, 0x64, 0xcc, 0x44, 0xfd, 0x63, 0x26, 0xc5, 0xa6, 0xa4, 0xad, 0xd7, 0xf5, 0xba, 0x4b, 0x09,
	0xe5, 0x7c, 0xbb, 0xc8, 0xb9, 0xb1, 0xea, 0x54, 0x1b, 0x75, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x27, 0xc0, 0x0e, 0x2c, 0x02, 0x00, 0x00,
}
