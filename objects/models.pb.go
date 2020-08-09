// Code generated by protoc-gen-go. DO NOT EDIT.
// source: objects/models.proto

package objects

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// SA-MP Object type data structure.
type ObjectModel struct {
	Model     uint32  `protobuf:"varint,1,opt,name=model,proto3" json:"model,omitempty"`
	ModelName string  `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Name      string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Category  string  `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Length    float64 `protobuf:"fixed64,5,opt,name=length,proto3" json:"length,omitempty"`
	Width     float64 `protobuf:"fixed64,6,opt,name=width,proto3" json:"width,omitempty"`
	Height    float64 `protobuf:"fixed64,7,opt,name=height,proto3" json:"height,omitempty"`
	// if no size specified, it will be calculated based on length, width and height
	Size                 float64  `protobuf:"fixed64,8,opt,name=size,proto3" json:"size,omitempty"`
	Tags                 []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	HasCollision         bool     `protobuf:"varint,10,opt,name=has_collision,json=hasCollision,proto3" json:"has_collision,omitempty"`
	BreaksOnHit          bool     `protobuf:"varint,11,opt,name=breaks_on_hit,json=breaksOnHit,proto3" json:"breaks_on_hit,omitempty"`
	HasAnimation         bool     `protobuf:"varint,12,opt,name=has_animation,json=hasAnimation,proto3" json:"has_animation,omitempty"`
	VisibleByTime        bool     `protobuf:"varint,16,opt,name=visible_by_time,json=visibleByTime,proto3" json:"visible_by_time,omitempty"`
	VisibleFrom          uint32   `protobuf:"varint,17,opt,name=visible_from,json=visibleFrom,proto3" json:"visible_from,omitempty"`
	VisibleTo            uint32   `protobuf:"varint,18,opt,name=visible_to,json=visibleTo,proto3" json:"visible_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectModel) Reset()         { *m = ObjectModel{} }
func (m *ObjectModel) String() string { return proto.CompactTextString(m) }
func (*ObjectModel) ProtoMessage()    {}
func (*ObjectModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{0}
}

func (m *ObjectModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectModel.Unmarshal(m, b)
}
func (m *ObjectModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectModel.Marshal(b, m, deterministic)
}
func (m *ObjectModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectModel.Merge(m, src)
}
func (m *ObjectModel) XXX_Size() int {
	return xxx_messageInfo_ObjectModel.Size(m)
}
func (m *ObjectModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectModel.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectModel proto.InternalMessageInfo

func (m *ObjectModel) GetModel() uint32 {
	if m != nil {
		return m.Model
	}
	return 0
}

func (m *ObjectModel) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ObjectModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectModel) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ObjectModel) GetLength() float64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ObjectModel) GetWidth() float64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ObjectModel) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ObjectModel) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ObjectModel) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ObjectModel) GetHasCollision() bool {
	if m != nil {
		return m.HasCollision
	}
	return false
}

func (m *ObjectModel) GetBreaksOnHit() bool {
	if m != nil {
		return m.BreaksOnHit
	}
	return false
}

func (m *ObjectModel) GetHasAnimation() bool {
	if m != nil {
		return m.HasAnimation
	}
	return false
}

func (m *ObjectModel) GetVisibleByTime() bool {
	if m != nil {
		return m.VisibleByTime
	}
	return false
}

func (m *ObjectModel) GetVisibleFrom() uint32 {
	if m != nil {
		return m.VisibleFrom
	}
	return 0
}

func (m *ObjectModel) GetVisibleTo() uint32 {
	if m != nil {
		return m.VisibleTo
	}
	return 0
}

// Request message for rpc `CreateObjectModel`.
type CreateObjectModelRequest struct {
	ObjectType           *ObjectModel `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateObjectModelRequest) Reset()         { *m = CreateObjectModelRequest{} }
func (m *CreateObjectModelRequest) String() string { return proto.CompactTextString(m) }
func (*CreateObjectModelRequest) ProtoMessage()    {}
func (*CreateObjectModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{1}
}

func (m *CreateObjectModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectModelRequest.Unmarshal(m, b)
}
func (m *CreateObjectModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectModelRequest.Marshal(b, m, deterministic)
}
func (m *CreateObjectModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectModelRequest.Merge(m, src)
}
func (m *CreateObjectModelRequest) XXX_Size() int {
	return xxx_messageInfo_CreateObjectModelRequest.Size(m)
}
func (m *CreateObjectModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectModelRequest proto.InternalMessageInfo

func (m *CreateObjectModelRequest) GetObjectType() *ObjectModel {
	if m != nil {
		return m.ObjectType
	}
	return nil
}

// Response message for rpc `CreateObjectModel`.
type CreateObjectModelResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateObjectModelResponse) Reset()         { *m = CreateObjectModelResponse{} }
func (m *CreateObjectModelResponse) String() string { return proto.CompactTextString(m) }
func (*CreateObjectModelResponse) ProtoMessage()    {}
func (*CreateObjectModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{2}
}

func (m *CreateObjectModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectModelResponse.Unmarshal(m, b)
}
func (m *CreateObjectModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectModelResponse.Marshal(b, m, deterministic)
}
func (m *CreateObjectModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectModelResponse.Merge(m, src)
}
func (m *CreateObjectModelResponse) XXX_Size() int {
	return xxx_messageInfo_CreateObjectModelResponse.Size(m)
}
func (m *CreateObjectModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectModelResponse proto.InternalMessageInfo

func (m *CreateObjectModelResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request message for rpc `GetObjectModel`.
type GetObjectModelRequest struct {
	Model                uint32   `protobuf:"varint,1,opt,name=model,proto3" json:"model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectModelRequest) Reset()         { *m = GetObjectModelRequest{} }
func (m *GetObjectModelRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectModelRequest) ProtoMessage()    {}
func (*GetObjectModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{3}
}

func (m *GetObjectModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectModelRequest.Unmarshal(m, b)
}
func (m *GetObjectModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectModelRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectModelRequest.Merge(m, src)
}
func (m *GetObjectModelRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectModelRequest.Size(m)
}
func (m *GetObjectModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectModelRequest proto.InternalMessageInfo

func (m *GetObjectModelRequest) GetModel() uint32 {
	if m != nil {
		return m.Model
	}
	return 0
}

// Response message for rpc `GetObjectModel`.
type GetObjectModelResponse struct {
	ObjectType           *ObjectModel `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetObjectModelResponse) Reset()         { *m = GetObjectModelResponse{} }
func (m *GetObjectModelResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectModelResponse) ProtoMessage()    {}
func (*GetObjectModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{4}
}

func (m *GetObjectModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectModelResponse.Unmarshal(m, b)
}
func (m *GetObjectModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectModelResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectModelResponse.Merge(m, src)
}
func (m *GetObjectModelResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectModelResponse.Size(m)
}
func (m *GetObjectModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectModelResponse proto.InternalMessageInfo

func (m *GetObjectModelResponse) GetObjectType() *ObjectModel {
	if m != nil {
		return m.ObjectType
	}
	return nil
}

// Request message for rpc `UpdateObjectModel`.
type UpdateObjectModelRequest struct {
	ObjectType           *ObjectModel `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateObjectModelRequest) Reset()         { *m = UpdateObjectModelRequest{} }
func (m *UpdateObjectModelRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateObjectModelRequest) ProtoMessage()    {}
func (*UpdateObjectModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{5}
}

func (m *UpdateObjectModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateObjectModelRequest.Unmarshal(m, b)
}
func (m *UpdateObjectModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateObjectModelRequest.Marshal(b, m, deterministic)
}
func (m *UpdateObjectModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateObjectModelRequest.Merge(m, src)
}
func (m *UpdateObjectModelRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateObjectModelRequest.Size(m)
}
func (m *UpdateObjectModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateObjectModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateObjectModelRequest proto.InternalMessageInfo

func (m *UpdateObjectModelRequest) GetObjectType() *ObjectModel {
	if m != nil {
		return m.ObjectType
	}
	return nil
}

// Response message for rpc `UpdateObjectModel`.
type UpdateObjectModelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateObjectModelResponse) Reset()         { *m = UpdateObjectModelResponse{} }
func (m *UpdateObjectModelResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateObjectModelResponse) ProtoMessage()    {}
func (*UpdateObjectModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{6}
}

func (m *UpdateObjectModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateObjectModelResponse.Unmarshal(m, b)
}
func (m *UpdateObjectModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateObjectModelResponse.Marshal(b, m, deterministic)
}
func (m *UpdateObjectModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateObjectModelResponse.Merge(m, src)
}
func (m *UpdateObjectModelResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateObjectModelResponse.Size(m)
}
func (m *UpdateObjectModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateObjectModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateObjectModelResponse proto.InternalMessageInfo

// Request message for rpc `DeleteObjectModel`.
type DeleteObjectModelRequest struct {
	Model                uint32   `protobuf:"varint,1,opt,name=model,proto3" json:"model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectModelRequest) Reset()         { *m = DeleteObjectModelRequest{} }
func (m *DeleteObjectModelRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectModelRequest) ProtoMessage()    {}
func (*DeleteObjectModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{7}
}

func (m *DeleteObjectModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectModelRequest.Unmarshal(m, b)
}
func (m *DeleteObjectModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectModelRequest.Marshal(b, m, deterministic)
}
func (m *DeleteObjectModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectModelRequest.Merge(m, src)
}
func (m *DeleteObjectModelRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectModelRequest.Size(m)
}
func (m *DeleteObjectModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectModelRequest proto.InternalMessageInfo

func (m *DeleteObjectModelRequest) GetModel() uint32 {
	if m != nil {
		return m.Model
	}
	return 0
}

// Response message for rpc `DeleteObjectModel`.
type DeleteObjectModelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectModelResponse) Reset()         { *m = DeleteObjectModelResponse{} }
func (m *DeleteObjectModelResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectModelResponse) ProtoMessage()    {}
func (*DeleteObjectModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eec9892f5f88c85, []int{8}
}

func (m *DeleteObjectModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectModelResponse.Unmarshal(m, b)
}
func (m *DeleteObjectModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectModelResponse.Marshal(b, m, deterministic)
}
func (m *DeleteObjectModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectModelResponse.Merge(m, src)
}
func (m *DeleteObjectModelResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectModelResponse.Size(m)
}
func (m *DeleteObjectModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectModelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ObjectModel)(nil), "mruv.objects.ObjectModel")
	proto.RegisterType((*CreateObjectModelRequest)(nil), "mruv.objects.CreateObjectModelRequest")
	proto.RegisterType((*CreateObjectModelResponse)(nil), "mruv.objects.CreateObjectModelResponse")
	proto.RegisterType((*GetObjectModelRequest)(nil), "mruv.objects.GetObjectModelRequest")
	proto.RegisterType((*GetObjectModelResponse)(nil), "mruv.objects.GetObjectModelResponse")
	proto.RegisterType((*UpdateObjectModelRequest)(nil), "mruv.objects.UpdateObjectModelRequest")
	proto.RegisterType((*UpdateObjectModelResponse)(nil), "mruv.objects.UpdateObjectModelResponse")
	proto.RegisterType((*DeleteObjectModelRequest)(nil), "mruv.objects.DeleteObjectModelRequest")
	proto.RegisterType((*DeleteObjectModelResponse)(nil), "mruv.objects.DeleteObjectModelResponse")
}

func init() { proto.RegisterFile("objects/models.proto", fileDescriptor_4eec9892f5f88c85) }

var fileDescriptor_4eec9892f5f88c85 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0x95, 0xd3, 0xcb, 0xdf, 0x7c, 0x49, 0xda, 0x3f, 0xa3, 0xd2, 0x4e, 0x52, 0x10, 0xc1, 0xa5,
	0x6d, 0x44, 0x95, 0x18, 0xca, 0x8e, 0x1d, 0x2d, 0x02, 0x36, 0xa5, 0x28, 0x84, 0x2e, 0xd8, 0x58,
	0x76, 0xf2, 0x61, 0x0f, 0xd8, 0x1e, 0xe3, 0x99, 0x04, 0x05, 0xe8, 0x06, 0x89, 0x07, 0x40, 0xbc,
	0x09, 0x2b, 0xde, 0x83, 0x57, 0xe0, 0x41, 0x90, 0xc7, 0x93, 0x92, 0x8b, 0x03, 0x0b, 0xd8, 0xcd,
	0x9c, 0x39, 0x3e, 0xdf, 0xe5, 0x1c, 0x19, 0x36, 0xb9, 0xfb, 0x0a, 0x7b, 0x52, 0x58, 0x21, 0xef,
	0x63, 0x20, 0xda, 0x71, 0xc2, 0x25, 0x27, 0xe5, 0x30, 0x19, 0x0c, 0xdb, 0xfa, 0xa9, 0x7e, 0xd5,
	0xe3, 0xdc, 0x0b, 0xd0, 0x72, 0x62, 0x66, 0x39, 0x51, 0xc4, 0xa5, 0x23, 0x19, 0x8f, 0x34, 0xd7,
	0xfc, 0xb6, 0x04, 0xa5, 0x33, 0xc5, 0x3c, 0x4d, 0x25, 0xc8, 0x26, 0xac, 0x28, 0x2d, 0x6a, 0x34,
	0x8c, 0x66, 0xa5, 0x93, 0x5d, 0xc8, 0x35, 0x00, 0x75, 0xb0, 0x23, 0x27, 0x44, 0x5a, 0x68, 0x18,
	0xcd, 0x62, 0xa7, 0xa8, 0x90, 0x27, 0x4e, 0x88, 0x84, 0xc0, 0xb2, 0x7a, 0x58, 0x52, 0x0f, 0xea,
	0x4c, 0xea, 0xb0, 0xd6, 0x73, 0x24, 0x7a, 0x3c, 0x19, 0xd1, 0x65, 0x85, 0x5f, 0xde, 0xc9, 0x16,
	0xac, 0x06, 0x18, 0x79, 0xd2, 0xa7, 0x2b, 0x0d, 0xa3, 0x69, 0x74, 0xf4, 0x2d, 0x2d, 0xfe, 0x96,
	0xf5, 0xa5, 0x4f, 0x57, 0x15, 0x9c, 0x5d, 0x52, 0xb6, 0x8f, 0xcc, 0xf3, 0x25, 0xfd, 0x2f, 0x63,
	0x67, 0xb7, 0xb4, 0xaa, 0x60, 0xef, 0x90, 0xae, 0x29, 0x54, 0x9d, 0x53, 0x4c, 0x3a, 0x9e, 0xa0,
	0xc5, 0xc6, 0x52, 0xda, 0x49, 0x7a, 0x26, 0xbb, 0x50, 0xf1, 0x1d, 0x61, 0xf7, 0x78, 0x10, 0x30,
	0xc1, 0x78, 0x44, 0xa1, 0x61, 0x34, 0xd7, 0x3a, 0x65, 0xdf, 0x11, 0x27, 0x63, 0x8c, 0x98, 0x50,
	0x71, 0x13, 0x74, 0x5e, 0x0b, 0x9b, 0x47, 0xb6, 0xcf, 0x24, 0x2d, 0x29, 0x52, 0x29, 0x03, 0xcf,
	0xa2, 0xc7, 0x4c, 0x8e, 0x85, 0x9c, 0x88, 0x85, 0x6a, 0x87, 0xb4, 0x7c, 0x29, 0x74, 0x7f, 0x8c,
	0x91, 0x7d, 0xd8, 0x18, 0x32, 0xc1, 0xdc, 0x00, 0x6d, 0x77, 0x64, 0x4b, 0x16, 0x22, 0xfd, 0x5f,
	0xd1, 0x2a, 0x1a, 0x3e, 0x1e, 0x75, 0x59, 0x88, 0xe4, 0x06, 0x94, 0xc7, 0xbc, 0x97, 0x09, 0x0f,
	0x69, 0x55, 0xed, 0xbb, 0xa4, 0xb1, 0x87, 0x09, 0x0f, 0xd3, 0xad, 0x8f, 0x29, 0x92, 0x53, 0xa2,
	0x08, 0x45, 0x8d, 0x74, 0xb9, 0x79, 0x0e, 0xf4, 0x24, 0x41, 0x47, 0xe2, 0x84, 0x7f, 0x1d, 0x7c,
	0x33, 0x40, 0x21, 0xc9, 0x3d, 0x28, 0x65, 0xfe, 0xdb, 0x72, 0x14, 0xa3, 0x32, 0xb3, 0x74, 0x54,
	0x6b, 0x4f, 0x06, 0xa3, 0x3d, 0xf9, 0x19, 0x64, 0x60, 0x77, 0x14, 0xa3, 0x79, 0x08, 0xb5, 0x1c,
	0x5d, 0x11, 0xf3, 0x48, 0x20, 0x59, 0x87, 0x02, 0xeb, 0xeb, 0x70, 0x14, 0x58, 0xdf, 0x6c, 0xc1,
	0x95, 0x47, 0x28, 0x73, 0x3a, 0xc8, 0x0d, 0x92, 0xd9, 0x85, 0xad, 0x59, 0xba, 0x16, 0xfe, 0x9b,
	0x8e, 0xcf, 0x81, 0x3e, 0x8f, 0xfb, 0xff, 0x7e, 0x13, 0x3b, 0x50, 0xcb, 0xd1, 0xcd, 0x1a, 0x36,
	0x6f, 0x03, 0x7d, 0x80, 0x01, 0xe6, 0x16, 0xcd, 0x1f, 0x7e, 0x07, 0x6a, 0x39, 0x5f, 0x64, 0x72,
	0x47, 0x5f, 0x97, 0x61, 0xfb, 0x34, 0x19, 0x9c, 0x4f, 0xbc, 0x89, 0x67, 0x98, 0x0c, 0x59, 0x0f,
	0xc9, 0x05, 0x54, 0xe7, 0x1c, 0x21, 0xfb, 0xd3, 0x33, 0x2c, 0x8a, 0x42, 0xfd, 0xe0, 0x8f, 0x3c,
	0x3d, 0xd0, 0xf6, 0xc7, 0xef, 0x3f, 0xbe, 0x14, 0xaa, 0xe6, 0x86, 0x35, 0xbc, 0x63, 0xfd, 0xda,
	0x82, 0x20, 0x1f, 0x60, 0x7d, 0xda, 0x34, 0xb2, 0x3b, 0xad, 0x99, 0x9b, 0x80, 0xfa, 0xcd, 0xdf,
	0x93, 0x74, 0xd5, 0xeb, 0xaa, 0x6a, 0x8d, 0x6c, 0xcf, 0x54, 0xb5, 0xde, 0xab, 0xa5, 0x5d, 0x90,
	0xcf, 0x06, 0x54, 0xe7, 0x5c, 0x98, 0x9d, 0x7e, 0x91, 0xfd, 0xb3, 0xd3, 0x2f, 0xb6, 0xf3, 0x50,
	0xf5, 0xb1, 0x77, 0xb4, 0x3b, 0xd7, 0xc7, 0x44, 0x7c, 0xda, 0xba, 0xa7, 0x4f, 0x06, 0x54, 0xe7,
	0xac, 0x9c, 0xed, 0x69, 0x51, 0x3a, 0x66, 0x7b, 0x5a, 0x98, 0x89, 0xf1, 0x6e, 0x6e, 0x2d, 0xda,
	0xcd, 0xf1, 0xc1, 0x8b, 0x3d, 0x8f, 0x49, 0x7f, 0xe0, 0xb6, 0x7b, 0x3c, 0xb4, 0xd2, 0xf8, 0xb4,
	0x3a, 0x4f, 0xad, 0x54, 0xbd, 0x15, 0xbb, 0x2d, 0x8f, 0xeb, 0x8f, 0x84, 0xbb, 0xaa, 0xfe, 0xf6,
	0x77, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x19, 0xd4, 0x8a, 0x31, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MruVObjectModelsServiceClient is the client API for MruVObjectModelsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MruVObjectModelsServiceClient interface {
	// Create an object model.
	CreateObjectModel(ctx context.Context, in *CreateObjectModelRequest, opts ...grpc.CallOption) (*CreateObjectModelResponse, error)
	// Get an object model.
	GetObjectModel(ctx context.Context, in *GetObjectModelRequest, opts ...grpc.CallOption) (*GetObjectModelResponse, error)
	// Update an object model.
	UpdateObjectModel(ctx context.Context, in *UpdateObjectModelRequest, opts ...grpc.CallOption) (*UpdateObjectModelResponse, error)
	// Delete an object model.
	DeleteObjectModel(ctx context.Context, in *DeleteObjectModelRequest, opts ...grpc.CallOption) (*DeleteObjectModelResponse, error)
}

type mruVObjectModelsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMruVObjectModelsServiceClient(cc *grpc.ClientConn) MruVObjectModelsServiceClient {
	return &mruVObjectModelsServiceClient{cc}
}

func (c *mruVObjectModelsServiceClient) CreateObjectModel(ctx context.Context, in *CreateObjectModelRequest, opts ...grpc.CallOption) (*CreateObjectModelResponse, error) {
	out := new(CreateObjectModelResponse)
	err := c.cc.Invoke(ctx, "/mruv.objects.MruVObjectModelsService/CreateObjectModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVObjectModelsServiceClient) GetObjectModel(ctx context.Context, in *GetObjectModelRequest, opts ...grpc.CallOption) (*GetObjectModelResponse, error) {
	out := new(GetObjectModelResponse)
	err := c.cc.Invoke(ctx, "/mruv.objects.MruVObjectModelsService/GetObjectModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVObjectModelsServiceClient) UpdateObjectModel(ctx context.Context, in *UpdateObjectModelRequest, opts ...grpc.CallOption) (*UpdateObjectModelResponse, error) {
	out := new(UpdateObjectModelResponse)
	err := c.cc.Invoke(ctx, "/mruv.objects.MruVObjectModelsService/UpdateObjectModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVObjectModelsServiceClient) DeleteObjectModel(ctx context.Context, in *DeleteObjectModelRequest, opts ...grpc.CallOption) (*DeleteObjectModelResponse, error) {
	out := new(DeleteObjectModelResponse)
	err := c.cc.Invoke(ctx, "/mruv.objects.MruVObjectModelsService/DeleteObjectModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MruVObjectModelsServiceServer is the server API for MruVObjectModelsService service.
type MruVObjectModelsServiceServer interface {
	// Create an object model.
	CreateObjectModel(context.Context, *CreateObjectModelRequest) (*CreateObjectModelResponse, error)
	// Get an object model.
	GetObjectModel(context.Context, *GetObjectModelRequest) (*GetObjectModelResponse, error)
	// Update an object model.
	UpdateObjectModel(context.Context, *UpdateObjectModelRequest) (*UpdateObjectModelResponse, error)
	// Delete an object model.
	DeleteObjectModel(context.Context, *DeleteObjectModelRequest) (*DeleteObjectModelResponse, error)
}

// UnimplementedMruVObjectModelsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMruVObjectModelsServiceServer struct {
}

func (*UnimplementedMruVObjectModelsServiceServer) CreateObjectModel(ctx context.Context, req *CreateObjectModelRequest) (*CreateObjectModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectModel not implemented")
}
func (*UnimplementedMruVObjectModelsServiceServer) GetObjectModel(ctx context.Context, req *GetObjectModelRequest) (*GetObjectModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectModel not implemented")
}
func (*UnimplementedMruVObjectModelsServiceServer) UpdateObjectModel(ctx context.Context, req *UpdateObjectModelRequest) (*UpdateObjectModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectModel not implemented")
}
func (*UnimplementedMruVObjectModelsServiceServer) DeleteObjectModel(ctx context.Context, req *DeleteObjectModelRequest) (*DeleteObjectModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectModel not implemented")
}

func RegisterMruVObjectModelsServiceServer(s *grpc.Server, srv MruVObjectModelsServiceServer) {
	s.RegisterService(&_MruVObjectModelsService_serviceDesc, srv)
}

func _MruVObjectModelsService_CreateObjectModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVObjectModelsServiceServer).CreateObjectModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.objects.MruVObjectModelsService/CreateObjectModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVObjectModelsServiceServer).CreateObjectModel(ctx, req.(*CreateObjectModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVObjectModelsService_GetObjectModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVObjectModelsServiceServer).GetObjectModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.objects.MruVObjectModelsService/GetObjectModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVObjectModelsServiceServer).GetObjectModel(ctx, req.(*GetObjectModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVObjectModelsService_UpdateObjectModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVObjectModelsServiceServer).UpdateObjectModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.objects.MruVObjectModelsService/UpdateObjectModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVObjectModelsServiceServer).UpdateObjectModel(ctx, req.(*UpdateObjectModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVObjectModelsService_DeleteObjectModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVObjectModelsServiceServer).DeleteObjectModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.objects.MruVObjectModelsService/DeleteObjectModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVObjectModelsServiceServer).DeleteObjectModel(ctx, req.(*DeleteObjectModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MruVObjectModelsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mruv.objects.MruVObjectModelsService",
	HandlerType: (*MruVObjectModelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectModel",
			Handler:    _MruVObjectModelsService_CreateObjectModel_Handler,
		},
		{
			MethodName: "GetObjectModel",
			Handler:    _MruVObjectModelsService_GetObjectModel_Handler,
		},
		{
			MethodName: "UpdateObjectModel",
			Handler:    _MruVObjectModelsService_UpdateObjectModel_Handler,
		},
		{
			MethodName: "DeleteObjectModel",
			Handler:    _MruVObjectModelsService_DeleteObjectModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objects/models.proto",
}
