// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estates/elevators.proto

package elevators

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

// Request message for rpc `CreateElevator`.
type CreateElevatorRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateElevatorRequest) Reset()         { *m = CreateElevatorRequest{} }
func (m *CreateElevatorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateElevatorRequest) ProtoMessage()    {}
func (*CreateElevatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{0}
}

func (m *CreateElevatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateElevatorRequest.Unmarshal(m, b)
}
func (m *CreateElevatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateElevatorRequest.Marshal(b, m, deterministic)
}
func (m *CreateElevatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateElevatorRequest.Merge(m, src)
}
func (m *CreateElevatorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateElevatorRequest.Size(m)
}
func (m *CreateElevatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateElevatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateElevatorRequest proto.InternalMessageInfo

// Response message for rpc `CreateElevator`.
type CreateElevatorResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateElevatorResponse) Reset()         { *m = CreateElevatorResponse{} }
func (m *CreateElevatorResponse) String() string { return proto.CompactTextString(m) }
func (*CreateElevatorResponse) ProtoMessage()    {}
func (*CreateElevatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{1}
}

func (m *CreateElevatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateElevatorResponse.Unmarshal(m, b)
}
func (m *CreateElevatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateElevatorResponse.Marshal(b, m, deterministic)
}
func (m *CreateElevatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateElevatorResponse.Merge(m, src)
}
func (m *CreateElevatorResponse) XXX_Size() int {
	return xxx_messageInfo_CreateElevatorResponse.Size(m)
}
func (m *CreateElevatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateElevatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateElevatorResponse proto.InternalMessageInfo

func (m *CreateElevatorResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request message for rpc `GetElevator`.
type GetElevatorRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorRequest) Reset()         { *m = GetElevatorRequest{} }
func (m *GetElevatorRequest) String() string { return proto.CompactTextString(m) }
func (*GetElevatorRequest) ProtoMessage()    {}
func (*GetElevatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{2}
}

func (m *GetElevatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorRequest.Unmarshal(m, b)
}
func (m *GetElevatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorRequest.Marshal(b, m, deterministic)
}
func (m *GetElevatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorRequest.Merge(m, src)
}
func (m *GetElevatorRequest) XXX_Size() int {
	return xxx_messageInfo_GetElevatorRequest.Size(m)
}
func (m *GetElevatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorRequest proto.InternalMessageInfo

func (m *GetElevatorRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `GetElevator`.
type GetElevatorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorResponse) Reset()         { *m = GetElevatorResponse{} }
func (m *GetElevatorResponse) String() string { return proto.CompactTextString(m) }
func (*GetElevatorResponse) ProtoMessage()    {}
func (*GetElevatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{3}
}

func (m *GetElevatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorResponse.Unmarshal(m, b)
}
func (m *GetElevatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorResponse.Marshal(b, m, deterministic)
}
func (m *GetElevatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorResponse.Merge(m, src)
}
func (m *GetElevatorResponse) XXX_Size() int {
	return xxx_messageInfo_GetElevatorResponse.Size(m)
}
func (m *GetElevatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorResponse proto.InternalMessageInfo

// Request message for rpc `UpdateElevator`.
type UpdateElevatorRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateElevatorRequest) Reset()         { *m = UpdateElevatorRequest{} }
func (m *UpdateElevatorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateElevatorRequest) ProtoMessage()    {}
func (*UpdateElevatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{4}
}

func (m *UpdateElevatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateElevatorRequest.Unmarshal(m, b)
}
func (m *UpdateElevatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateElevatorRequest.Marshal(b, m, deterministic)
}
func (m *UpdateElevatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateElevatorRequest.Merge(m, src)
}
func (m *UpdateElevatorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateElevatorRequest.Size(m)
}
func (m *UpdateElevatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateElevatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateElevatorRequest proto.InternalMessageInfo

func (m *UpdateElevatorRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `UpdateElevator`.
type UpdateElevatorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateElevatorResponse) Reset()         { *m = UpdateElevatorResponse{} }
func (m *UpdateElevatorResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateElevatorResponse) ProtoMessage()    {}
func (*UpdateElevatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{5}
}

func (m *UpdateElevatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateElevatorResponse.Unmarshal(m, b)
}
func (m *UpdateElevatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateElevatorResponse.Marshal(b, m, deterministic)
}
func (m *UpdateElevatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateElevatorResponse.Merge(m, src)
}
func (m *UpdateElevatorResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateElevatorResponse.Size(m)
}
func (m *UpdateElevatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateElevatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateElevatorResponse proto.InternalMessageInfo

// Request message for rpc `DeleteElevator`.
type DeleteElevatorRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteElevatorRequest) Reset()         { *m = DeleteElevatorRequest{} }
func (m *DeleteElevatorRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteElevatorRequest) ProtoMessage()    {}
func (*DeleteElevatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{6}
}

func (m *DeleteElevatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteElevatorRequest.Unmarshal(m, b)
}
func (m *DeleteElevatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteElevatorRequest.Marshal(b, m, deterministic)
}
func (m *DeleteElevatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteElevatorRequest.Merge(m, src)
}
func (m *DeleteElevatorRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteElevatorRequest.Size(m)
}
func (m *DeleteElevatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteElevatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteElevatorRequest proto.InternalMessageInfo

func (m *DeleteElevatorRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `DeleteElevator`.
type DeleteElevatorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteElevatorResponse) Reset()         { *m = DeleteElevatorResponse{} }
func (m *DeleteElevatorResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteElevatorResponse) ProtoMessage()    {}
func (*DeleteElevatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{7}
}

func (m *DeleteElevatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteElevatorResponse.Unmarshal(m, b)
}
func (m *DeleteElevatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteElevatorResponse.Marshal(b, m, deterministic)
}
func (m *DeleteElevatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteElevatorResponse.Merge(m, src)
}
func (m *DeleteElevatorResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteElevatorResponse.Size(m)
}
func (m *DeleteElevatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteElevatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteElevatorResponse proto.InternalMessageInfo

// Request message for rpc `GetElevatorFloors`.
type GetElevatorFloorsRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorFloorsRequest) Reset()         { *m = GetElevatorFloorsRequest{} }
func (m *GetElevatorFloorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetElevatorFloorsRequest) ProtoMessage()    {}
func (*GetElevatorFloorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{8}
}

func (m *GetElevatorFloorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorFloorsRequest.Unmarshal(m, b)
}
func (m *GetElevatorFloorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorFloorsRequest.Marshal(b, m, deterministic)
}
func (m *GetElevatorFloorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorFloorsRequest.Merge(m, src)
}
func (m *GetElevatorFloorsRequest) XXX_Size() int {
	return xxx_messageInfo_GetElevatorFloorsRequest.Size(m)
}
func (m *GetElevatorFloorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorFloorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorFloorsRequest proto.InternalMessageInfo

func (m *GetElevatorFloorsRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `GetElevatorFloors`.
type GetElevatorFloorsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorFloorsResponse) Reset()         { *m = GetElevatorFloorsResponse{} }
func (m *GetElevatorFloorsResponse) String() string { return proto.CompactTextString(m) }
func (*GetElevatorFloorsResponse) ProtoMessage()    {}
func (*GetElevatorFloorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd9d47894856e98, []int{9}
}

func (m *GetElevatorFloorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorFloorsResponse.Unmarshal(m, b)
}
func (m *GetElevatorFloorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorFloorsResponse.Marshal(b, m, deterministic)
}
func (m *GetElevatorFloorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorFloorsResponse.Merge(m, src)
}
func (m *GetElevatorFloorsResponse) XXX_Size() int {
	return xxx_messageInfo_GetElevatorFloorsResponse.Size(m)
}
func (m *GetElevatorFloorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorFloorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorFloorsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateElevatorRequest)(nil), "mruv.elevators.CreateElevatorRequest")
	proto.RegisterType((*CreateElevatorResponse)(nil), "mruv.elevators.CreateElevatorResponse")
	proto.RegisterType((*GetElevatorRequest)(nil), "mruv.elevators.GetElevatorRequest")
	proto.RegisterType((*GetElevatorResponse)(nil), "mruv.elevators.GetElevatorResponse")
	proto.RegisterType((*UpdateElevatorRequest)(nil), "mruv.elevators.UpdateElevatorRequest")
	proto.RegisterType((*UpdateElevatorResponse)(nil), "mruv.elevators.UpdateElevatorResponse")
	proto.RegisterType((*DeleteElevatorRequest)(nil), "mruv.elevators.DeleteElevatorRequest")
	proto.RegisterType((*DeleteElevatorResponse)(nil), "mruv.elevators.DeleteElevatorResponse")
	proto.RegisterType((*GetElevatorFloorsRequest)(nil), "mruv.elevators.GetElevatorFloorsRequest")
	proto.RegisterType((*GetElevatorFloorsResponse)(nil), "mruv.elevators.GetElevatorFloorsResponse")
}

func init() { proto.RegisterFile("estates/elevators.proto", fileDescriptor_afd9d47894856e98) }

var fileDescriptor_afd9d47894856e98 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xbb, 0x4e, 0xfb, 0x30,
	0x18, 0xc5, 0xd5, 0x4a, 0xff, 0xff, 0x60, 0xd4, 0x20, 0x4c, 0xd3, 0x4b, 0xca, 0x00, 0xe1, 0xd2,
	0x8b, 0xd4, 0x58, 0x94, 0x37, 0xe0, 0x3a, 0x21, 0xa1, 0x22, 0x18, 0xd8, 0xd2, 0xe6, 0x23, 0x44,
	0x4a, 0xe3, 0x60, 0x3b, 0x59, 0x2a, 0x16, 0x46, 0x18, 0x79, 0x34, 0x5e, 0x81, 0x07, 0x41, 0x4d,
	0x0d, 0x6d, 0xe2, 0xa4, 0x74, 0x8d, 0xcf, 0x77, 0x7e, 0x9f, 0x7d, 0x8e, 0x82, 0xea, 0xc0, 0x85,
	0x2d, 0x80, 0x13, 0xf0, 0x21, 0xb6, 0x05, 0x65, 0xdc, 0x0a, 0x19, 0x15, 0x14, 0x6b, 0x13, 0x16,
	0xc5, 0xd6, 0xef, 0x57, 0x63, 0xc7, 0xa5, 0xd4, 0xf5, 0x81, 0xd8, 0xa1, 0x47, 0xec, 0x20, 0xa0,
	0xc2, 0x16, 0x1e, 0x0d, 0xa4, 0xda, 0xac, 0x23, 0xfd, 0x8c, 0x81, 0x2d, 0xe0, 0x42, 0x0e, 0x0c,
	0xe1, 0x39, 0x02, 0x2e, 0xcc, 0x0e, 0xaa, 0x65, 0x0f, 0x78, 0x48, 0x03, 0x0e, 0x58, 0x43, 0x65,
	0xcf, 0x69, 0x94, 0x76, 0x4b, 0x9d, 0xca, 0xb0, 0xec, 0x39, 0xe6, 0x01, 0xc2, 0x57, 0x20, 0x32,
	0xf3, 0x8a, 0x4a, 0x47, 0xdb, 0x29, 0xd5, 0xdc, 0xcc, 0x6c, 0x23, 0xfd, 0x2e, 0x74, 0x54, 0xbe,
	0x32, 0xdf, 0x40, 0xb5, 0xac, 0x70, 0x61, 0x71, 0x0e, 0x3e, 0xac, 0x65, 0x91, 0x15, 0x4a, 0x8b,
	0x1e, 0x6a, 0x2c, 0x2d, 0x77, 0xe9, 0x53, 0xca, 0x78, 0x91, 0x4b, 0x0b, 0x35, 0x73, 0xb4, 0x73,
	0xa3, 0xc1, 0xdb, 0x3f, 0x54, 0xbd, 0x66, 0xd1, 0xfd, 0xcf, 0x31, 0xbf, 0x05, 0x16, 0x7b, 0x63,
	0xc0, 0x31, 0xd2, 0xd2, 0xcf, 0x89, 0x0f, 0xad, 0x74, 0x50, 0x56, 0x6e, 0x0e, 0xc6, 0xd1, 0x5f,
	0x32, 0x79, 0x05, 0xfd, 0xf5, 0xf3, 0xeb, 0xa3, 0xbc, 0x69, 0x56, 0x48, 0x7c, 0xbc, 0xe8, 0x04,
	0x66, 0x68, 0x63, 0x69, 0x5b, 0x6c, 0x66, 0xdd, 0xd4, 0xe4, 0x8c, 0xfd, 0x95, 0x1a, 0x89, 0x33,
	0x12, 0x5c, 0x15, 0xe3, 0x14, 0x8e, 0x4c, 0x3d, 0xe7, 0x05, 0x4f, 0x91, 0x96, 0x8e, 0x4a, 0xbd,
	0x6b, 0x6e, 0xe6, 0xea, 0x5d, 0x0b, 0x12, 0x97, 0xf0, 0x41, 0x01, 0x3c, 0x1d, 0xb2, 0x0a, 0xcf,
	0x6d, 0x8b, 0x0a, 0x2f, 0xe8, 0x8a, 0x84, 0xf7, 0xf2, 0xe0, 0xef, 0x25, 0xb4, 0xa5, 0x94, 0x03,
	0x77, 0x56, 0x3c, 0x68, 0xaa, 0x6b, 0x46, 0x77, 0x0d, 0xa5, 0x5c, 0x63, 0x2f, 0x59, 0xa3, 0x85,
	0x9b, 0xea, 0x1a, 0xe4, 0x31, 0x91, 0x9e, 0x76, 0x1f, 0xda, 0xae, 0x27, 0x9e, 0xa2, 0x91, 0x35,
	0xa6, 0x13, 0x32, 0xab, 0x65, 0x7f, 0x78, 0x43, 0x66, 0x84, 0x7e, 0x38, 0xea, 0xbb, 0x74, 0x31,
	0x36, 0xfa, 0x9f, 0xfc, 0x0d, 0x4e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x29, 0x6c, 0x28,
	0x56, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MruVElevatorsServiceClient is the client API for MruVElevatorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MruVElevatorsServiceClient interface {
	// Create an elevator.
	CreateElevator(ctx context.Context, in *CreateElevatorRequest, opts ...grpc.CallOption) (*CreateElevatorResponse, error)
	// Get an elevator.
	GetElevator(ctx context.Context, in *GetElevatorRequest, opts ...grpc.CallOption) (*GetElevatorResponse, error)
	// Update an elevator.
	UpdateElevator(ctx context.Context, in *UpdateElevatorRequest, opts ...grpc.CallOption) (*UpdateElevatorResponse, error)
	// Delete an elevator.
	DeleteElevator(ctx context.Context, in *DeleteElevatorRequest, opts ...grpc.CallOption) (*DeleteElevatorResponse, error)
	// Get available elevator floors.
	GetElevatorFloors(ctx context.Context, in *GetElevatorFloorsRequest, opts ...grpc.CallOption) (*GetElevatorFloorsResponse, error)
}

type mruVElevatorsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMruVElevatorsServiceClient(cc *grpc.ClientConn) MruVElevatorsServiceClient {
	return &mruVElevatorsServiceClient{cc}
}

func (c *mruVElevatorsServiceClient) CreateElevator(ctx context.Context, in *CreateElevatorRequest, opts ...grpc.CallOption) (*CreateElevatorResponse, error) {
	out := new(CreateElevatorResponse)
	err := c.cc.Invoke(ctx, "/mruv.elevators.MruVElevatorsService/CreateElevator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVElevatorsServiceClient) GetElevator(ctx context.Context, in *GetElevatorRequest, opts ...grpc.CallOption) (*GetElevatorResponse, error) {
	out := new(GetElevatorResponse)
	err := c.cc.Invoke(ctx, "/mruv.elevators.MruVElevatorsService/GetElevator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVElevatorsServiceClient) UpdateElevator(ctx context.Context, in *UpdateElevatorRequest, opts ...grpc.CallOption) (*UpdateElevatorResponse, error) {
	out := new(UpdateElevatorResponse)
	err := c.cc.Invoke(ctx, "/mruv.elevators.MruVElevatorsService/UpdateElevator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVElevatorsServiceClient) DeleteElevator(ctx context.Context, in *DeleteElevatorRequest, opts ...grpc.CallOption) (*DeleteElevatorResponse, error) {
	out := new(DeleteElevatorResponse)
	err := c.cc.Invoke(ctx, "/mruv.elevators.MruVElevatorsService/DeleteElevator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVElevatorsServiceClient) GetElevatorFloors(ctx context.Context, in *GetElevatorFloorsRequest, opts ...grpc.CallOption) (*GetElevatorFloorsResponse, error) {
	out := new(GetElevatorFloorsResponse)
	err := c.cc.Invoke(ctx, "/mruv.elevators.MruVElevatorsService/GetElevatorFloors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MruVElevatorsServiceServer is the server API for MruVElevatorsService service.
type MruVElevatorsServiceServer interface {
	// Create an elevator.
	CreateElevator(context.Context, *CreateElevatorRequest) (*CreateElevatorResponse, error)
	// Get an elevator.
	GetElevator(context.Context, *GetElevatorRequest) (*GetElevatorResponse, error)
	// Update an elevator.
	UpdateElevator(context.Context, *UpdateElevatorRequest) (*UpdateElevatorResponse, error)
	// Delete an elevator.
	DeleteElevator(context.Context, *DeleteElevatorRequest) (*DeleteElevatorResponse, error)
	// Get available elevator floors.
	GetElevatorFloors(context.Context, *GetElevatorFloorsRequest) (*GetElevatorFloorsResponse, error)
}

// UnimplementedMruVElevatorsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMruVElevatorsServiceServer struct {
}

func (*UnimplementedMruVElevatorsServiceServer) CreateElevator(ctx context.Context, req *CreateElevatorRequest) (*CreateElevatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateElevator not implemented")
}
func (*UnimplementedMruVElevatorsServiceServer) GetElevator(ctx context.Context, req *GetElevatorRequest) (*GetElevatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElevator not implemented")
}
func (*UnimplementedMruVElevatorsServiceServer) UpdateElevator(ctx context.Context, req *UpdateElevatorRequest) (*UpdateElevatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateElevator not implemented")
}
func (*UnimplementedMruVElevatorsServiceServer) DeleteElevator(ctx context.Context, req *DeleteElevatorRequest) (*DeleteElevatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteElevator not implemented")
}
func (*UnimplementedMruVElevatorsServiceServer) GetElevatorFloors(ctx context.Context, req *GetElevatorFloorsRequest) (*GetElevatorFloorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElevatorFloors not implemented")
}

func RegisterMruVElevatorsServiceServer(s *grpc.Server, srv MruVElevatorsServiceServer) {
	s.RegisterService(&_MruVElevatorsService_serviceDesc, srv)
}

func _MruVElevatorsService_CreateElevator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateElevatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVElevatorsServiceServer).CreateElevator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.elevators.MruVElevatorsService/CreateElevator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVElevatorsServiceServer).CreateElevator(ctx, req.(*CreateElevatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVElevatorsService_GetElevator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetElevatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVElevatorsServiceServer).GetElevator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.elevators.MruVElevatorsService/GetElevator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVElevatorsServiceServer).GetElevator(ctx, req.(*GetElevatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVElevatorsService_UpdateElevator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateElevatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVElevatorsServiceServer).UpdateElevator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.elevators.MruVElevatorsService/UpdateElevator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVElevatorsServiceServer).UpdateElevator(ctx, req.(*UpdateElevatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVElevatorsService_DeleteElevator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteElevatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVElevatorsServiceServer).DeleteElevator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.elevators.MruVElevatorsService/DeleteElevator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVElevatorsServiceServer).DeleteElevator(ctx, req.(*DeleteElevatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVElevatorsService_GetElevatorFloors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetElevatorFloorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVElevatorsServiceServer).GetElevatorFloors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.elevators.MruVElevatorsService/GetElevatorFloors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVElevatorsServiceServer).GetElevatorFloors(ctx, req.(*GetElevatorFloorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MruVElevatorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mruv.elevators.MruVElevatorsService",
	HandlerType: (*MruVElevatorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateElevator",
			Handler:    _MruVElevatorsService_CreateElevator_Handler,
		},
		{
			MethodName: "GetElevator",
			Handler:    _MruVElevatorsService_GetElevator_Handler,
		},
		{
			MethodName: "UpdateElevator",
			Handler:    _MruVElevatorsService_UpdateElevator_Handler,
		},
		{
			MethodName: "DeleteElevator",
			Handler:    _MruVElevatorsService_DeleteElevator_Handler,
		},
		{
			MethodName: "GetElevatorFloors",
			Handler:    _MruVElevatorsService_GetElevatorFloors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "estates/elevators.proto",
}