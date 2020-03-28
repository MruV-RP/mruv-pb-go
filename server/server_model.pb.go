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

// Data that describe server
type ServerInfo struct {
	// Short name of the server
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Host (ip) of the server
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// Port of the server
	Port string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	// Platform, on which server is running
	Platform             string   `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
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

// State of the server
type ServerStatus struct {
	// Is server active and working
	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	// How many players are registered on that server
	Players              int32    `protobuf:"varint,2,opt,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStatus) Reset()         { *m = ServerStatus{} }
func (m *ServerStatus) String() string { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()    {}
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d74db2732d33f2e, []int{2}
}

func (m *ServerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatus.Unmarshal(m, b)
}
func (m *ServerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatus.Marshal(b, m, deterministic)
}
func (m *ServerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatus.Merge(m, src)
}
func (m *ServerStatus) XXX_Size() int {
	return xxx_messageInfo_ServerStatus.Size(m)
}
func (m *ServerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatus proto.InternalMessageInfo

func (m *ServerStatus) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ServerStatus) GetPlayers() int32 {
	if m != nil {
		return m.Players
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerID)(nil), "mruv.ServerID")
	proto.RegisterType((*ServerInfo)(nil), "mruv.ServerInfo")
	proto.RegisterType((*ServerStatus)(nil), "mruv.ServerStatus")
}

func init() { proto.RegisterFile("server/server_model.proto", fileDescriptor_3d74db2732d33f2e) }

var fileDescriptor_3d74db2732d33f2e = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x69, 0x1a, 0x6b, 0x7c, 0x88, 0x87, 0x3d, 0xc8, 0xda, 0x93, 0x04, 0x11, 0x2f, 0x69,
	0x0e, 0xfe, 0x01, 0x11, 0x2f, 0x1e, 0x04, 0xd9, 0x82, 0x07, 0x2f, 0xb2, 0x69, 0xb6, 0x6d, 0x20,
	0xdb, 0xb7, 0xbc, 0x7d, 0x1b, 0xf0, 0xdf, 0xcb, 0xee, 0xc6, 0x9e, 0x66, 0xe6, 0x9b, 0xc3, 0x30,
	0x70, 0xe7, 0x0d, 0x4d, 0x86, 0xda, 0x2c, 0x3f, 0x16, 0x7b, 0x33, 0x6e, 0x1c, 0x21, 0xa3, 0x28,
	0x2d, 0x85, 0xa9, 0x5e, 0x43, 0xb5, 0x4d, 0xdd, 0xfb, 0x9b, 0xb8, 0x81, 0x62, 0xe8, 0xe5, 0xe2,
	0x7e, 0xf1, 0xb4, 0x54, 0xc5, 0xd0, 0xd7, 0x3d, 0xc0, 0xdc, 0x9d, 0xf6, 0x28, 0x04, 0x94, 0x27,
	0x6d, 0x4d, 0xea, 0xaf, 0x54, 0xf2, 0x91, 0x1d, 0xd1, 0xb3, 0x2c, 0x32, 0x8b, 0x3e, 0x32, 0x87,
	0xc4, 0x72, 0x99, 0x59, 0xf4, 0x62, 0x0d, 0x95, 0x1b, 0x35, 0xef, 0x91, 0xac, 0x2c, 0x13, 0x3f,
	0xe7, 0xfa, 0x05, 0xae, 0xf3, 0xca, 0x96, 0x35, 0x07, 0x2f, 0x6e, 0x61, 0xa5, 0x77, 0x3c, 0x4c,
	0x79, 0xa9, 0x52, 0x73, 0x12, 0x12, 0x2e, 0xdd, 0xa8, 0x7f, 0x0d, 0xf9, 0x34, 0x77, 0xa1, 0xfe,
	0xe3, 0xeb, 0xe3, 0xf7, 0xc3, 0x61, 0xe0, 0x63, 0xe8, 0x36, 0x3b, 0xb4, 0xed, 0x07, 0x85, 0xaf,
	0x46, 0x7d, 0xb6, 0xf1, 0x5e, 0xe3, 0xba, 0xe6, 0x80, 0xf3, 0xfb, 0x6e, 0x95, 0x8e, 0x3f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x28, 0xcf, 0xaf, 0x69, 0x15, 0x01, 0x00, 0x00,
}
