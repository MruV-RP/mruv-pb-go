// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plots/plots.proto

package plots

import (
	context "context"
	fmt "fmt"
	common "github.com/MruV-RP/mruv-pb-go/common"
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

type Plot struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Points               []*common.Position `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Plot) Reset()         { *m = Plot{} }
func (m *Plot) String() string { return proto.CompactTextString(m) }
func (*Plot) ProtoMessage()    {}
func (*Plot) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{0}
}

func (m *Plot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plot.Unmarshal(m, b)
}
func (m *Plot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plot.Marshal(b, m, deterministic)
}
func (m *Plot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plot.Merge(m, src)
}
func (m *Plot) XXX_Size() int {
	return xxx_messageInfo_Plot.Size(m)
}
func (m *Plot) XXX_DiscardUnknown() {
	xxx_messageInfo_Plot.DiscardUnknown(m)
}

var xxx_messageInfo_Plot proto.InternalMessageInfo

func (m *Plot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plot) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plot) GetPoints() []*common.Position {
	if m != nil {
		return m.Points
	}
	return nil
}

// Request message for rpc `CreatePlot`.
type CreatePlotRequest struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Points               []*common.Position `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreatePlotRequest) Reset()         { *m = CreatePlotRequest{} }
func (m *CreatePlotRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlotRequest) ProtoMessage()    {}
func (*CreatePlotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{1}
}

func (m *CreatePlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlotRequest.Unmarshal(m, b)
}
func (m *CreatePlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlotRequest.Marshal(b, m, deterministic)
}
func (m *CreatePlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlotRequest.Merge(m, src)
}
func (m *CreatePlotRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlotRequest.Size(m)
}
func (m *CreatePlotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlotRequest proto.InternalMessageInfo

func (m *CreatePlotRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePlotRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreatePlotRequest) GetPoints() []*common.Position {
	if m != nil {
		return m.Points
	}
	return nil
}

// Response message for rpc `CreatePlot`.
type CreatePlotResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlotResponse) Reset()         { *m = CreatePlotResponse{} }
func (m *CreatePlotResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePlotResponse) ProtoMessage()    {}
func (*CreatePlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{2}
}

func (m *CreatePlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlotResponse.Unmarshal(m, b)
}
func (m *CreatePlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlotResponse.Marshal(b, m, deterministic)
}
func (m *CreatePlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlotResponse.Merge(m, src)
}
func (m *CreatePlotResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePlotResponse.Size(m)
}
func (m *CreatePlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlotResponse proto.InternalMessageInfo

func (m *CreatePlotResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request message for rpc `GetPlot`.
type GetPlotRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlotRequest) Reset()         { *m = GetPlotRequest{} }
func (m *GetPlotRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlotRequest) ProtoMessage()    {}
func (*GetPlotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{3}
}

func (m *GetPlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlotRequest.Unmarshal(m, b)
}
func (m *GetPlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlotRequest.Marshal(b, m, deterministic)
}
func (m *GetPlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlotRequest.Merge(m, src)
}
func (m *GetPlotRequest) XXX_Size() int {
	return xxx_messageInfo_GetPlotRequest.Size(m)
}
func (m *GetPlotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlotRequest proto.InternalMessageInfo

func (m *GetPlotRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `GetPlot`.
type GetPlotResponse struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Points               []*common.Position `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	Area                 float64            `protobuf:"fixed64,4,opt,name=area,proto3" json:"area,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPlotResponse) Reset()         { *m = GetPlotResponse{} }
func (m *GetPlotResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlotResponse) ProtoMessage()    {}
func (*GetPlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{4}
}

func (m *GetPlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlotResponse.Unmarshal(m, b)
}
func (m *GetPlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlotResponse.Marshal(b, m, deterministic)
}
func (m *GetPlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlotResponse.Merge(m, src)
}
func (m *GetPlotResponse) XXX_Size() int {
	return xxx_messageInfo_GetPlotResponse.Size(m)
}
func (m *GetPlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlotResponse proto.InternalMessageInfo

func (m *GetPlotResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPlotResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetPlotResponse) GetPoints() []*common.Position {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *GetPlotResponse) GetArea() float64 {
	if m != nil {
		return m.Area
	}
	return 0
}

// Request message for rpc `UpdatePlot`.
type UpdatePlotRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlotRequest) Reset()         { *m = UpdatePlotRequest{} }
func (m *UpdatePlotRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePlotRequest) ProtoMessage()    {}
func (*UpdatePlotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{5}
}

func (m *UpdatePlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlotRequest.Unmarshal(m, b)
}
func (m *UpdatePlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlotRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlotRequest.Merge(m, src)
}
func (m *UpdatePlotRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePlotRequest.Size(m)
}
func (m *UpdatePlotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlotRequest proto.InternalMessageInfo

func (m *UpdatePlotRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePlotRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePlotRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Response message for rpc `UpdatePlot`.
type UpdatePlotResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlotResponse) Reset()         { *m = UpdatePlotResponse{} }
func (m *UpdatePlotResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePlotResponse) ProtoMessage()    {}
func (*UpdatePlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{6}
}

func (m *UpdatePlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlotResponse.Unmarshal(m, b)
}
func (m *UpdatePlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlotResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlotResponse.Merge(m, src)
}
func (m *UpdatePlotResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePlotResponse.Size(m)
}
func (m *UpdatePlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlotResponse proto.InternalMessageInfo

// Request message for rpc `DeletePlot`.
type DeletePlotRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlotRequest) Reset()         { *m = DeletePlotRequest{} }
func (m *DeletePlotRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePlotRequest) ProtoMessage()    {}
func (*DeletePlotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{7}
}

func (m *DeletePlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlotRequest.Unmarshal(m, b)
}
func (m *DeletePlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlotRequest.Marshal(b, m, deterministic)
}
func (m *DeletePlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlotRequest.Merge(m, src)
}
func (m *DeletePlotRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePlotRequest.Size(m)
}
func (m *DeletePlotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlotRequest proto.InternalMessageInfo

func (m *DeletePlotRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response message for rpc `DeletePlot`.
type DeletePlotResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlotResponse) Reset()         { *m = DeletePlotResponse{} }
func (m *DeletePlotResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePlotResponse) ProtoMessage()    {}
func (*DeletePlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32035c16824dd6e6, []int{8}
}

func (m *DeletePlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlotResponse.Unmarshal(m, b)
}
func (m *DeletePlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlotResponse.Marshal(b, m, deterministic)
}
func (m *DeletePlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlotResponse.Merge(m, src)
}
func (m *DeletePlotResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePlotResponse.Size(m)
}
func (m *DeletePlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlotResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Plot)(nil), "mruv.plots.Plot")
	proto.RegisterType((*CreatePlotRequest)(nil), "mruv.plots.CreatePlotRequest")
	proto.RegisterType((*CreatePlotResponse)(nil), "mruv.plots.CreatePlotResponse")
	proto.RegisterType((*GetPlotRequest)(nil), "mruv.plots.GetPlotRequest")
	proto.RegisterType((*GetPlotResponse)(nil), "mruv.plots.GetPlotResponse")
	proto.RegisterType((*UpdatePlotRequest)(nil), "mruv.plots.UpdatePlotRequest")
	proto.RegisterType((*UpdatePlotResponse)(nil), "mruv.plots.UpdatePlotResponse")
	proto.RegisterType((*DeletePlotRequest)(nil), "mruv.plots.DeletePlotRequest")
	proto.RegisterType((*DeletePlotResponse)(nil), "mruv.plots.DeletePlotResponse")
}

func init() { proto.RegisterFile("plots/plots.proto", fileDescriptor_32035c16824dd6e6) }

var fileDescriptor_32035c16824dd6e6 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x69, 0x3b, 0xac, 0xec, 0x3b, 0x38, 0x6e, 0xc3, 0x2a, 0xa5, 0xfe, 0xa1, 0x74, 0x15,
	0x06, 0xa1, 0x0d, 0x8e, 0xdf, 0x40, 0x05, 0x4f, 0xc2, 0x50, 0x51, 0x70, 0x0f, 0x42, 0xa6, 0x0d,
	0x35, 0xd8, 0x36, 0xb1, 0x49, 0x07, 0x41, 0xbc, 0x78, 0xf4, 0xea, 0xc5, 0xef, 0xe5, 0x57, 0xf0,
	0x83, 0x48, 0xd2, 0x3a, 0xfd, 0x37, 0xbb, 0xc7, 0xbd, 0x0c, 0xe1, 0xcd, 0x9b, 0xe7, 0xf7, 0xe4,
	0x7d, 0x32, 0x05, 0x57, 0x14, 0x5c, 0x49, 0x6c, 0x7e, 0x63, 0x51, 0x73, 0xc5, 0x11, 0x94, 0x75,
	0xb3, 0x8f, 0x4d, 0xc5, 0x7f, 0x90, 0x73, 0x9e, 0x17, 0x14, 0x13, 0xc1, 0x30, 0xa9, 0x2a, 0xae,
	0x88, 0x62, 0xbc, 0xea, 0x3a, 0xfd, 0xf3, 0x94, 0x97, 0x25, 0xaf, 0xb0, 0x14, 0x44, 0x31, 0x52,
	0xb4, 0xd5, 0xf0, 0x33, 0x2c, 0xb6, 0x05, 0x57, 0x08, 0xc1, 0xa2, 0x22, 0x25, 0xf5, 0xac, 0xc0,
	0x5a, 0x9f, 0x26, 0x66, 0x8d, 0x02, 0x58, 0x66, 0x54, 0xa6, 0x35, 0x13, 0x5a, 0xc7, 0xb3, 0xcd,
	0xd6, 0xb0, 0x84, 0x22, 0x38, 0x11, 0x9c, 0x55, 0x4a, 0x7a, 0x4e, 0xe0, 0xac, 0x97, 0x9b, 0xbb,
	0xb1, 0xb1, 0xd3, 0x92, 0xe2, 0x2d, 0x97, 0x4c, 0xb7, 0x25, 0x5d, 0x53, 0xf8, 0x15, 0xdc, 0x97,
	0x35, 0x25, 0x8a, 0x6a, 0x64, 0x42, 0xbf, 0x34, 0x54, 0xde, 0x10, 0xf9, 0x31, 0xa0, 0x21, 0x59,
	0x0a, 0x5e, 0x49, 0x8a, 0x56, 0x60, 0xb3, 0xcc, 0x80, 0x6f, 0x27, 0x36, 0xcb, 0xc2, 0x00, 0x56,
	0xaf, 0xa9, 0x1a, 0x9a, 0x9b, 0x76, 0xfc, 0xb4, 0xe0, 0xce, 0xa1, 0xa5, 0x53, 0xb9, 0x89, 0x0b,
	0x68, 0x08, 0xa9, 0x29, 0xf1, 0x16, 0x81, 0xb5, 0xb6, 0x12, 0xb3, 0x0e, 0x3f, 0x80, 0xfb, 0x4e,
	0x64, 0x93, 0x71, 0x4e, 0x1c, 0x1f, 0xdc, 0xd9, 0x57, 0xbb, 0x73, 0x66, 0xee, 0xc2, 0x73, 0x40,
	0x43, 0xe9, 0xf6, 0xa6, 0xe1, 0x05, 0xb8, 0xaf, 0x68, 0x41, 0xaf, 0x05, 0xea, 0xa3, 0xc3, 0xa6,
	0xf6, 0xe8, 0xe6, 0xb7, 0x03, 0x67, 0x6f, 0xea, 0xe6, 0xbd, 0x2e, 0xca, 0xb7, 0xb4, 0xde, 0xb3,
	0x94, 0xa2, 0x8f, 0x00, 0x7d, 0x2a, 0xe8, 0x61, 0xdc, 0xbf, 0xe5, 0x78, 0xf6, 0x4e, 0xfc, 0x47,
	0x57, 0x6d, 0x77, 0xe6, 0xdc, 0x1f, 0x7f, 0xfe, 0xfe, 0xb2, 0x97, 0xe1, 0x29, 0xde, 0x3f, 0x6b,
	0xff, 0x22, 0xe8, 0x12, 0x6e, 0x75, 0x61, 0x21, 0x7f, 0x78, 0x7a, 0x1c, 0xb2, 0x7f, 0xff, 0xe8,
	0x5e, 0x27, 0x7b, 0xcf, 0xc8, 0x9e, 0xa1, 0xd5, 0x41, 0x16, 0x7f, 0x63, 0xd9, 0x77, 0x94, 0x02,
	0xf4, 0x13, 0x1a, 0x7b, 0x9f, 0x85, 0x32, 0xf6, 0x7e, 0x64, 0xb0, 0x1d, 0x64, 0x73, 0x04, 0xd2,
	0xcf, 0x72, 0x0c, 0x99, 0x05, 0x31, 0x86, 0xcc, 0x23, 0xf8, 0x0f, 0x79, 0x3a, 0x81, 0xbc, 0x78,
	0x72, 0x79, 0x91, 0x33, 0xf5, 0xa9, 0xd9, 0xe9, 0xc7, 0x87, 0x75, 0x48, 0x51, 0xb2, 0xc5, 0x5a,
	0x2b, 0x12, 0xbb, 0x28, 0xe7, 0x6d, 0xef, 0xee, 0xc4, 0x7c, 0x30, 0x9e, 0xff, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x7b, 0xe3, 0x9f, 0x85, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MruVPlotsServiceClient is the client API for MruVPlotsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MruVPlotsServiceClient interface {
	// Create a plot.
	CreatePlot(ctx context.Context, in *CreatePlotRequest, opts ...grpc.CallOption) (*CreatePlotResponse, error)
	// Get a plot.
	GetPlot(ctx context.Context, in *GetPlotRequest, opts ...grpc.CallOption) (*GetPlotResponse, error)
	// Update a plot.
	UpdatePlot(ctx context.Context, in *UpdatePlotRequest, opts ...grpc.CallOption) (*UpdatePlotResponse, error)
	// Delete a plot.
	DeletePlot(ctx context.Context, in *DeletePlotRequest, opts ...grpc.CallOption) (*DeletePlotResponse, error)
}

type mruVPlotsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMruVPlotsServiceClient(cc *grpc.ClientConn) MruVPlotsServiceClient {
	return &mruVPlotsServiceClient{cc}
}

func (c *mruVPlotsServiceClient) CreatePlot(ctx context.Context, in *CreatePlotRequest, opts ...grpc.CallOption) (*CreatePlotResponse, error) {
	out := new(CreatePlotResponse)
	err := c.cc.Invoke(ctx, "/mruv.plots.MruVPlotsService/CreatePlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVPlotsServiceClient) GetPlot(ctx context.Context, in *GetPlotRequest, opts ...grpc.CallOption) (*GetPlotResponse, error) {
	out := new(GetPlotResponse)
	err := c.cc.Invoke(ctx, "/mruv.plots.MruVPlotsService/GetPlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVPlotsServiceClient) UpdatePlot(ctx context.Context, in *UpdatePlotRequest, opts ...grpc.CallOption) (*UpdatePlotResponse, error) {
	out := new(UpdatePlotResponse)
	err := c.cc.Invoke(ctx, "/mruv.plots.MruVPlotsService/UpdatePlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVPlotsServiceClient) DeletePlot(ctx context.Context, in *DeletePlotRequest, opts ...grpc.CallOption) (*DeletePlotResponse, error) {
	out := new(DeletePlotResponse)
	err := c.cc.Invoke(ctx, "/mruv.plots.MruVPlotsService/DeletePlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MruVPlotsServiceServer is the server API for MruVPlotsService service.
type MruVPlotsServiceServer interface {
	// Create a plot.
	CreatePlot(context.Context, *CreatePlotRequest) (*CreatePlotResponse, error)
	// Get a plot.
	GetPlot(context.Context, *GetPlotRequest) (*GetPlotResponse, error)
	// Update a plot.
	UpdatePlot(context.Context, *UpdatePlotRequest) (*UpdatePlotResponse, error)
	// Delete a plot.
	DeletePlot(context.Context, *DeletePlotRequest) (*DeletePlotResponse, error)
}

// UnimplementedMruVPlotsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMruVPlotsServiceServer struct {
}

func (*UnimplementedMruVPlotsServiceServer) CreatePlot(ctx context.Context, req *CreatePlotRequest) (*CreatePlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlot not implemented")
}
func (*UnimplementedMruVPlotsServiceServer) GetPlot(ctx context.Context, req *GetPlotRequest) (*GetPlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlot not implemented")
}
func (*UnimplementedMruVPlotsServiceServer) UpdatePlot(ctx context.Context, req *UpdatePlotRequest) (*UpdatePlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlot not implemented")
}
func (*UnimplementedMruVPlotsServiceServer) DeletePlot(ctx context.Context, req *DeletePlotRequest) (*DeletePlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlot not implemented")
}

func RegisterMruVPlotsServiceServer(s *grpc.Server, srv MruVPlotsServiceServer) {
	s.RegisterService(&_MruVPlotsService_serviceDesc, srv)
}

func _MruVPlotsService_CreatePlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVPlotsServiceServer).CreatePlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.plots.MruVPlotsService/CreatePlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVPlotsServiceServer).CreatePlot(ctx, req.(*CreatePlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVPlotsService_GetPlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVPlotsServiceServer).GetPlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.plots.MruVPlotsService/GetPlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVPlotsServiceServer).GetPlot(ctx, req.(*GetPlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVPlotsService_UpdatePlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVPlotsServiceServer).UpdatePlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.plots.MruVPlotsService/UpdatePlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVPlotsServiceServer).UpdatePlot(ctx, req.(*UpdatePlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVPlotsService_DeletePlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVPlotsServiceServer).DeletePlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.plots.MruVPlotsService/DeletePlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVPlotsServiceServer).DeletePlot(ctx, req.(*DeletePlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MruVPlotsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mruv.plots.MruVPlotsService",
	HandlerType: (*MruVPlotsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlot",
			Handler:    _MruVPlotsService_CreatePlot_Handler,
		},
		{
			MethodName: "GetPlot",
			Handler:    _MruVPlotsService_GetPlot_Handler,
		},
		{
			MethodName: "UpdatePlot",
			Handler:    _MruVPlotsService_UpdatePlot_Handler,
		},
		{
			MethodName: "DeletePlot",
			Handler:    _MruVPlotsService_DeletePlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plots/plots.proto",
}
