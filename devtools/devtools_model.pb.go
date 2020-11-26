// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devtools/devtools_model.proto

package devtools

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

// Position structure.
type Position struct {
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	// Short position name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the place which the position indicates.
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f0f820bd4c50c, []int{0}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Position) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// GTA V player outfit structure.
// You can find more information about RageMP outfits [here](https://wiki.rage.mp/index.php?title=Clothes "RageMP Clothes")
type Outfit struct {
	// Name of the outfit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Category of the outfit.
	Category             string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outfit) Reset()         { *m = Outfit{} }
func (m *Outfit) String() string { return proto.CompactTextString(m) }
func (*Outfit) ProtoMessage()    {}
func (*Outfit) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f0f820bd4c50c, []int{1}
}

func (m *Outfit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outfit.Unmarshal(m, b)
}
func (m *Outfit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outfit.Marshal(b, m, deterministic)
}
func (m *Outfit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outfit.Merge(m, src)
}
func (m *Outfit) XXX_Size() int {
	return xxx_messageInfo_Outfit.Size(m)
}
func (m *Outfit) XXX_DiscardUnknown() {
	xxx_messageInfo_Outfit.DiscardUnknown(m)
}

var xxx_messageInfo_Outfit proto.InternalMessageInfo

func (m *Outfit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Outfit) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// GTA V animation structure.
// You can find more information about RageMP outfits [here](https://wiki.rage.mp/index.php?title=Animations "RageMP Animations")
type Animation struct {
	// Animation dictionary name.
	Dict string `protobuf:"bytes,1,opt,name=dict,proto3" json:"dict,omitempty"`
	// Animation name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Animation category, specified by author.
	ReadableCategory     string   `protobuf:"bytes,3,opt,name=readable_category,json=readableCategory,proto3" json:"readable_category,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Animation) Reset()         { *m = Animation{} }
func (m *Animation) String() string { return proto.CompactTextString(m) }
func (*Animation) ProtoMessage()    {}
func (*Animation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f0f820bd4c50c, []int{2}
}

func (m *Animation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Animation.Unmarshal(m, b)
}
func (m *Animation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Animation.Marshal(b, m, deterministic)
}
func (m *Animation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animation.Merge(m, src)
}
func (m *Animation) XXX_Size() int {
	return xxx_messageInfo_Animation.Size(m)
}
func (m *Animation) XXX_DiscardUnknown() {
	xxx_messageInfo_Animation.DiscardUnknown(m)
}

var xxx_messageInfo_Animation proto.InternalMessageInfo

func (m *Animation) GetDict() string {
	if m != nil {
		return m.Dict
	}
	return ""
}

func (m *Animation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Animation) GetReadableCategory() string {
	if m != nil {
		return m.ReadableCategory
	}
	return ""
}

func (m *Animation) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// Player structure for storing authors of devtools saved records.
type Player struct {
	// Unique id.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Password hash.
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f0f820bd4c50c, []int{3}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Position)(nil), "mruv.devtools.Position")
	proto.RegisterType((*Outfit)(nil), "mruv.devtools.Outfit")
	proto.RegisterType((*Animation)(nil), "mruv.devtools.Animation")
	proto.RegisterType((*Player)(nil), "mruv.devtools.Player")
}

func init() { proto.RegisterFile("devtools/devtools_model.proto", fileDescriptor_b89f0f820bd4c50c) }

var fileDescriptor_b89f0f820bd4c50c = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x6b, 0x84, 0x30,
	0x10, 0xc5, 0x89, 0x6b, 0x65, 0x9d, 0x76, 0x4b, 0x9b, 0x93, 0x14, 0x0a, 0xe2, 0xa1, 0x08, 0x45,
	0x3d, 0xf4, 0xd2, 0x63, 0xff, 0x9c, 0x4b, 0xc5, 0x43, 0x0f, 0xbd, 0x2c, 0xd1, 0xa4, 0x9a, 0xa2,
	0x46, 0x62, 0x5c, 0xd6, 0xa5, 0x1f, 0xbe, 0x18, 0x57, 0xb1, 0x65, 0x6f, 0xef, 0xcd, 0x64, 0xde,
	0x2f, 0xf0, 0xe0, 0x96, 0xb2, 0x9d, 0x12, 0xa2, 0x6c, 0xa3, 0x49, 0x6c, 0x2b, 0x41, 0x59, 0x19,
	0x36, 0x52, 0x28, 0x81, 0x37, 0x95, 0xec, 0x76, 0xe1, 0xb4, 0xf2, 0xbe, 0x61, 0x1d, 0x8b, 0x96,
	0x2b, 0x2e, 0x6a, 0x7c, 0x01, 0x68, 0xef, 0x20, 0x17, 0xf9, 0x28, 0x41, 0xfb, 0xc1, 0xf5, 0x8e,
	0x31, 0xba, 0x7e, 0x70, 0x07, 0x67, 0x35, 0xba, 0x03, 0xc6, 0x60, 0xd6, 0xa4, 0x62, 0x8e, 0xe9,
	0x22, 0xdf, 0x4e, 0xb4, 0xc6, 0x2e, 0x9c, 0x53, 0xd6, 0x66, 0x92, 0x37, 0x43, 0x98, 0x73, 0xa6,
	0x57, 0xcb, 0x91, 0xf7, 0x08, 0xd6, 0x7b, 0xa7, 0xbe, 0xb8, 0x9a, 0xef, 0xd1, 0xe2, 0xfe, 0x06,
	0xd6, 0x19, 0x51, 0x2c, 0x17, 0x72, 0xc4, 0xda, 0xc9, 0xec, 0xbd, 0x1f, 0xb0, 0x9f, 0x6b, 0x5e,
	0x11, 0xfd, 0x4d, 0x0c, 0x26, 0xe5, 0x99, 0x9a, 0x8e, 0x07, 0x3d, 0x07, 0x1a, 0x8b, 0xc0, 0x7b,
	0xb8, 0x96, 0x8c, 0x50, 0x92, 0x96, 0x6c, 0x3b, 0x27, 0xaf, 0xf4, 0x83, 0xab, 0x69, 0xf1, 0x7a,
	0x9c, 0xff, 0xa1, 0x9b, 0xff, 0xe8, 0x4f, 0x60, 0xc5, 0x25, 0xe9, 0x99, 0xc4, 0x97, 0x60, 0x70,
	0xaa, 0xc1, 0x9b, 0xc4, 0xe0, 0xf4, 0x24, 0x16, 0x83, 0x59, 0x90, 0xb6, 0x38, 0x92, 0xb4, 0x7e,
	0xf1, 0x3f, 0xef, 0x72, 0xae, 0x8a, 0x2e, 0x0d, 0x33, 0x51, 0x45, 0x6f, 0xb2, 0xfb, 0x08, 0x92,
	0x38, 0x1a, 0x9a, 0x08, 0x9a, 0x34, 0xc8, 0xc5, 0x5c, 0x55, 0x6a, 0xe9, 0x96, 0x1e, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x6f, 0x33, 0x26, 0xc6, 0x01, 0x00, 0x00,
}
