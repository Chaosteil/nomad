// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package hashicorp_nomad_plugins_base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hclspec "hashicorp/nomad/plugins/shared/hclspec"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PluginType enumerates the type of plugins Nomad supports
type PluginType int32

const (
	PluginType_UNKNOWN PluginType = 0
	PluginType_DRIVER  PluginType = 1
	PluginType_DEVICE  PluginType = 2
)

var PluginType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DRIVER",
	2: "DEVICE",
}
var PluginType_value = map[string]int32{
	"UNKNOWN": 0,
	"DRIVER":  1,
	"DEVICE":  2,
}

func (x PluginType) String() string {
	return proto.EnumName(PluginType_name, int32(x))
}
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{0}
}

// PluginInfoRequest is used to request the plugins basic information.
type PluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoRequest) Reset()         { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()    {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{0}
}
func (m *PluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoRequest.Unmarshal(m, b)
}
func (m *PluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoRequest.Marshal(b, m, deterministic)
}
func (dst *PluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoRequest.Merge(dst, src)
}
func (m *PluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInfoRequest.Size(m)
}
func (m *PluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoRequest proto.InternalMessageInfo

// PluginInfoResponse returns basic information about the plugin such
// that Nomad can decide whether to load the plugin or not.
type PluginInfoResponse struct {
	// type indicates what type of plugin this is.
	Type PluginType `protobuf:"varint,1,opt,name=type,proto3,enum=hashicorp.nomad.plugins.base.PluginType" json:"type,omitempty"`
	// plugin_api_version indicates the version of the Nomad Plugin API
	// this plugin is built against.
	PluginApiVersion string `protobuf:"bytes,2,opt,name=plugin_api_version,json=pluginApiVersion,proto3" json:"plugin_api_version,omitempty"`
	// plugin_version is the semver version of this individual plugin.
	// This is divorce from Nomad’s development and versioning.
	PluginVersion string `protobuf:"bytes,3,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`
	// name is the name of the plugin
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoResponse) Reset()         { *m = PluginInfoResponse{} }
func (m *PluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*PluginInfoResponse) ProtoMessage()    {}
func (*PluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{1}
}
func (m *PluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoResponse.Unmarshal(m, b)
}
func (m *PluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoResponse.Marshal(b, m, deterministic)
}
func (dst *PluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoResponse.Merge(dst, src)
}
func (m *PluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_PluginInfoResponse.Size(m)
}
func (m *PluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoResponse proto.InternalMessageInfo

func (m *PluginInfoResponse) GetType() PluginType {
	if m != nil {
		return m.Type
	}
	return PluginType_UNKNOWN
}

func (m *PluginInfoResponse) GetPluginApiVersion() string {
	if m != nil {
		return m.PluginApiVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetPluginVersion() string {
	if m != nil {
		return m.PluginVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigSchemaRequest is used to request the configurations schema.
type ConfigSchemaRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigSchemaRequest) Reset()         { *m = ConfigSchemaRequest{} }
func (m *ConfigSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaRequest) ProtoMessage()    {}
func (*ConfigSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{2}
}
func (m *ConfigSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaRequest.Unmarshal(m, b)
}
func (m *ConfigSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaRequest.Merge(dst, src)
}
func (m *ConfigSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaRequest.Size(m)
}
func (m *ConfigSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaRequest proto.InternalMessageInfo

// ConfigSchemaResponse returns the plugins configuration schema.
type ConfigSchemaResponse struct {
	// spec is the plugins configuration schema
	Spec                 *hclspec.Spec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConfigSchemaResponse) Reset()         { *m = ConfigSchemaResponse{} }
func (m *ConfigSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaResponse) ProtoMessage()    {}
func (*ConfigSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{3}
}
func (m *ConfigSchemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaResponse.Unmarshal(m, b)
}
func (m *ConfigSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaResponse.Merge(dst, src)
}
func (m *ConfigSchemaResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaResponse.Size(m)
}
func (m *ConfigSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaResponse proto.InternalMessageInfo

func (m *ConfigSchemaResponse) GetSpec() *hclspec.Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// SetConfigRequest is used to set the configuration
type SetConfigRequest struct {
	// msgpack_config is the configuration encoded as MessagePack.
	MsgpackConfig        []byte   `protobuf:"bytes,1,opt,name=msgpack_config,json=msgpackConfig,proto3" json:"msgpack_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{4}
}
func (m *SetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigRequest.Unmarshal(m, b)
}
func (m *SetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *SetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigRequest.Merge(dst, src)
}
func (m *SetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigRequest.Size(m)
}
func (m *SetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigRequest proto.InternalMessageInfo

func (m *SetConfigRequest) GetMsgpackConfig() []byte {
	if m != nil {
		return m.MsgpackConfig
	}
	return nil
}

// SetConfigResponse is used to respond to setting the configuration
type SetConfigResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigResponse) Reset()         { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()    {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_6491f5f52ef6eb79, []int{5}
}
func (m *SetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigResponse.Unmarshal(m, b)
}
func (m *SetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigResponse.Marshal(b, m, deterministic)
}
func (dst *SetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigResponse.Merge(dst, src)
}
func (m *SetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigResponse.Size(m)
}
func (m *SetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PluginInfoRequest)(nil), "hashicorp.nomad.plugins.base.PluginInfoRequest")
	proto.RegisterType((*PluginInfoResponse)(nil), "hashicorp.nomad.plugins.base.PluginInfoResponse")
	proto.RegisterType((*ConfigSchemaRequest)(nil), "hashicorp.nomad.plugins.base.ConfigSchemaRequest")
	proto.RegisterType((*ConfigSchemaResponse)(nil), "hashicorp.nomad.plugins.base.ConfigSchemaResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "hashicorp.nomad.plugins.base.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "hashicorp.nomad.plugins.base.SetConfigResponse")
	proto.RegisterEnum("hashicorp.nomad.plugins.base.PluginType", PluginType_name, PluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BasePluginClient is the client API for BasePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasePluginClient interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
}

type basePluginClient struct {
	cc *grpc.ClientConn
}

func NewBasePluginClient(cc *grpc.ClientConn) BasePluginClient {
	return &basePluginClient{cc}
}

func (c *basePluginClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error) {
	out := new(ConfigSchemaResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/ConfigSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasePluginServer is the server API for BasePlugin service.
type BasePluginServer interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(context.Context, *ConfigSchemaRequest) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
}

func RegisterBasePluginServer(s *grpc.Server, srv BasePluginServer) {
	s.RegisterService(&_BasePlugin_serviceDesc, srv)
}

func _BasePlugin_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_ConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).ConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/ConfigSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).ConfigSchema(ctx, req.(*ConfigSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.base.BasePlugin",
	HandlerType: (*BasePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _BasePlugin_PluginInfo_Handler,
		},
		{
			MethodName: "ConfigSchema",
			Handler:    _BasePlugin_ConfigSchema_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _BasePlugin_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base.proto",
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_base_6491f5f52ef6eb79) }

var fileDescriptor_base_6491f5f52ef6eb79 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0x6a, 0xa8, 0xf4, 0xf4, 0x87, 0x38, 0x55, 0x28, 0xc1, 0x8b, 0x12, 0x10, 0x8a, 0x94,
	0x89, 0x8d, 0x78, 0x21, 0x78, 0xa1, 0xd6, 0x5e, 0x14, 0xa1, 0x4a, 0xaa, 0xd5, 0xbb, 0x30, 0x4d,
	0xa7, 0x4d, 0xb0, 0x99, 0x99, 0x66, 0x52, 0xa5, 0xcf, 0xe6, 0x0b, 0xed, 0x63, 0x2c, 0x99, 0x49,
	0xba, 0xd9, 0x65, 0xb7, 0x74, 0xaf, 0x32, 0x9c, 0xef, 0xe7, 0x9c, 0xf3, 0x1d, 0x02, 0xb0, 0x22,
	0x92, 0x62, 0x91, 0xf2, 0x8c, 0xa3, 0x97, 0x11, 0x91, 0x51, 0x1c, 0xf2, 0x54, 0x60, 0xc6, 0x13,
	0xb2, 0xc6, 0x62, 0x77, 0xd8, 0xc6, 0x4c, 0xe2, 0x9c, 0x63, 0xbf, 0x3b, 0xa1, 0xae, 0x42, 0xdd,
	0x02, 0x75, 0x65, 0x44, 0x52, 0xba, 0x76, 0xa3, 0x70, 0x27, 0x05, 0x0d, 0xf3, 0x6f, 0x90, 0x3f,
	0xb4, 0xa9, 0xd3, 0x83, 0x67, 0xdf, 0x15, 0x71, 0xc6, 0x36, 0xdc, 0xa7, 0xfb, 0x03, 0x95, 0x99,
	0xf3, 0xdf, 0x00, 0x54, 0xad, 0x4a, 0xc1, 0x99, 0xa4, 0xe8, 0x03, 0x98, 0xd9, 0x51, 0xd0, 0xbe,
	0x31, 0x30, 0x86, 0x5d, 0x6f, 0x88, 0xcf, 0xcd, 0x83, 0xb5, 0xfe, 0xc7, 0x51, 0x50, 0x5f, 0xa9,
	0xd0, 0x08, 0x90, 0x26, 0x04, 0x44, 0xc4, 0xc1, 0x5f, 0x9a, 0xca, 0x98, 0xb3, 0x7e, 0x7d, 0x60,
	0x0c, 0x9b, 0xbe, 0xa5, 0x91, 0x4f, 0x22, 0x5e, 0xea, 0x3a, 0x7a, 0x05, 0xdd, 0x82, 0x5d, 0x32,
	0x9f, 0x28, 0x66, 0x47, 0x57, 0x4b, 0x1a, 0x02, 0x93, 0x91, 0x84, 0xf6, 0x4d, 0x05, 0xaa, 0xb7,
	0xf3, 0x02, 0x7a, 0x13, 0xce, 0x36, 0xf1, 0x76, 0x11, 0x46, 0x34, 0x21, 0xe5, 0x52, 0xbf, 0xe1,
	0xf9, 0xed, 0x72, 0xb1, 0xd5, 0x47, 0x30, 0xf3, 0x3c, 0xd4, 0x56, 0x2d, 0x6f, 0xf4, 0xe0, 0x56,
	0x3a, 0x47, 0x5c, 0xe4, 0x88, 0x17, 0x82, 0x86, 0xbe, 0x52, 0x3a, 0xef, 0xc1, 0x5a, 0xd0, 0x4c,
	0x9b, 0x17, 0xdd, 0xf2, 0xf9, 0x13, 0xb9, 0x15, 0x24, 0xfc, 0x13, 0x84, 0x0a, 0x50, 0xfe, 0x6d,
	0xbf, 0x53, 0x54, 0x35, 0x3b, 0x8f, 0xbf, 0x22, 0xd5, 0x13, 0xbd, 0x1e, 0x03, 0xdc, 0xa4, 0x87,
	0x5a, 0xf0, 0xf4, 0xe7, 0xfc, 0xeb, 0xfc, 0xdb, 0xaf, 0xb9, 0x55, 0x43, 0x00, 0x8d, 0x2f, 0xfe,
	0x6c, 0x39, 0xf5, 0x2d, 0x43, 0xbd, 0xa7, 0xcb, 0xd9, 0x64, 0x6a, 0xd5, 0xbd, 0xab, 0x3a, 0xc0,
	0x67, 0x22, 0xa9, 0xd6, 0xa1, 0x7d, 0xe9, 0x90, 0xdf, 0x0f, 0xb9, 0x97, 0x5c, 0xaa, 0x72, 0x7f,
	0xfb, 0xcd, 0xe5, 0x02, 0x3d, 0xb2, 0x53, 0x43, 0xff, 0xa0, 0x5d, 0x8d, 0x17, 0x8d, 0xcf, 0x7b,
	0xdc, 0x73, 0x21, 0xdb, 0x7b, 0x8c, 0xe4, 0xd4, 0x98, 0x41, 0xf3, 0x14, 0x21, 0xc2, 0xe7, 0x2d,
	0xee, 0x9e, 0xc9, 0x76, 0x2f, 0xe6, 0x97, 0xfd, 0x56, 0x0d, 0xf5, 0xe3, 0xbc, 0xbd, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x9e, 0x46, 0x0c, 0x9b, 0x03, 0x00, 0x00,
}