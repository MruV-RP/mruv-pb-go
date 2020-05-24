// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/server_model.proto

package server

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

type ServerStatus int32

const (
	ServerStatus_UNKNOWN ServerStatus = 0
	ServerStatus_ON      ServerStatus = 1
	ServerStatus_OFF     ServerStatus = 2
)

var ServerStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "ON",
	2: "OFF",
}

var ServerStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"ON":      1,
	"OFF":     2,
}

func (x ServerStatus) String() string {
	return proto.EnumName(ServerStatus_name, int32(x))
}

func (ServerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d74db2732d33f2e, []int{0}
}

type ServerID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerID) Reset()         { *m = ServerID{} }
func (m *ServerID) String() string { return proto.CompactTextString(m) }
func (*ServerID) ProtoMessage()    {}
func (*ServerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d74db2732d33f2e, []int{0}
}

func (m *ServerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerID.Unmarshal(m, b)
}
func (m *ServerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerID.Marshal(b, m, deterministic)
}
func (m *ServerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerID.Merge(m, src)
}
func (m *ServerID) XXX_Size() int {
	return xxx_messageInfo_ServerID.Size(m)
}
func (m *ServerID) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerID.DiscardUnknown(m)
}

var xxx_messageInfo_ServerID proto.InternalMessageInfo

func (m *ServerID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Data that describe server.
type ServerInfo struct {
	// Id of the server.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Short name of the server.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Host (ip) of the server.
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// Port of the server.
	Port string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	// Platform of the server.
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
	// Status of the server.
	Status ServerStatus `protobuf:"varint,6,opt,name=status,proto3,enum=mruv.server.ServerStatus" json:"status,omitempty"`
	// How many players are playing on the server.
	Players              int32    `protobuf:"varint,7,opt,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d74db2732d33f2e, []int{1}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServerInfo) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *ServerInfo) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ServerInfo) GetStatus() ServerStatus {
	if m != nil {
		return m.Status
	}
	return ServerStatus_UNKNOWN
}

func (m *ServerInfo) GetPlayers() int32 {
	if m != nil {
		return m.Players
	}
	return 0
}

func init() {
	proto.RegisterEnum("mruv.server.ServerStatus", ServerStatus_name, ServerStatus_value)
	proto.RegisterType((*ServerID)(nil), "mruv.server.ServerID")
	proto.RegisterType((*ServerInfo)(nil), "mruv.server.ServerInfo")
}

func init() { proto.RegisterFile("server/server_model.proto", fileDescriptor_3d74db2732d33f2e) }

var fileDescriptor_3d74db2732d33f2e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x4d, 0xbb, 0xb5, 0xf3, 0x26, 0xa3, 0xdc, 0x53, 0xb6, 0xa7, 0x32, 0x44, 0x8a, 0xd8,
	0x16, 0xf5, 0x1b, 0x88, 0x0c, 0x44, 0x6c, 0xa5, 0x43, 0x05, 0x5f, 0xa4, 0xb5, 0xdd, 0x56, 0x68,
	0x97, 0x92, 0xa4, 0x03, 0x3f, 0x9e, 0xdf, 0x4c, 0x92, 0x54, 0x19, 0x3e, 0xe5, 0x77, 0xbf, 0xff,
	0x85, 0x3b, 0x0e, 0xe6, 0xa2, 0xe2, 0x87, 0x8a, 0xc7, 0xe6, 0xf9, 0x68, 0x59, 0x59, 0x35, 0x51,
	0xc7, 0x99, 0x64, 0x38, 0x6d, 0x79, 0x7f, 0x88, 0x4c, 0xb0, 0x5c, 0xc0, 0x64, 0xad, 0xe9, 0xe1,
	0x1e, 0x67, 0x60, 0xd5, 0x25, 0x25, 0x3e, 0x09, 0xec, 0xcc, 0xaa, 0xcb, 0xe5, 0x37, 0x01, 0x18,
	0xc2, 0xfd, 0x86, 0xfd, 0x8f, 0x11, 0x61, 0xb4, 0xcf, 0xdb, 0x8a, 0x5a, 0x3e, 0x09, 0x4e, 0x33,
	0xcd, 0xca, 0xed, 0x98, 0x90, 0xd4, 0x36, 0x4e, 0xb1, 0x72, 0x1d, 0xe3, 0x92, 0x8e, 0x8c, 0x53,
	0x8c, 0x0b, 0x98, 0x74, 0x4d, 0x2e, 0x37, 0x8c, 0xb7, 0x74, 0xac, 0xfd, 0x5f, 0x8d, 0xd7, 0xe0,
	0x08, 0x99, 0xcb, 0x5e, 0x50, 0xc7, 0x27, 0xc1, 0xec, 0x66, 0x1e, 0x1d, 0x2d, 0x1c, 0x99, 0x85,
	0xd6, 0xba, 0x21, 0x1b, 0x1a, 0x91, 0x82, 0xdb, 0x35, 0xf9, 0x57, 0xc5, 0x05, 0x75, 0x7d, 0x12,
	0x8c, 0xb3, 0xdf, 0xf2, 0xf2, 0x0a, 0xce, 0x8e, 0x7f, 0xe0, 0x14, 0xdc, 0x97, 0xe4, 0x31, 0x49,
	0xdf, 0x12, 0xef, 0x04, 0x1d, 0xb0, 0xd2, 0xc4, 0x23, 0xe8, 0x82, 0x9d, 0xae, 0x56, 0x9e, 0x75,
	0x77, 0xf1, 0x7e, 0xbe, 0xad, 0xe5, 0xae, 0x2f, 0xa2, 0x4f, 0xd6, 0xc6, 0x4f, 0xbc, 0x7f, 0x0d,
	0xb3, 0xe7, 0x58, 0x8d, 0x0f, 0xbb, 0x22, 0xdc, 0xb2, 0xe1, 0x9c, 0x85, 0xa3, 0x2f, 0x79, 0xfb,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xfc, 0x86, 0xb9, 0x66, 0x01, 0x00, 0x00,
}
