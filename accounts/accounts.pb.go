// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts/accounts.proto

package accounts

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

// Request message for rpc `RegisterAccount`.
type RegisterAccountRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountRequest) Reset()         { *m = RegisterAccountRequest{} }
func (m *RegisterAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountRequest) ProtoMessage()    {}
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{0}
}

func (m *RegisterAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountRequest.Unmarshal(m, b)
}
func (m *RegisterAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountRequest.Marshal(b, m, deterministic)
}
func (m *RegisterAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountRequest.Merge(m, src)
}
func (m *RegisterAccountRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountRequest.Size(m)
}
func (m *RegisterAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountRequest proto.InternalMessageInfo

func (m *RegisterAccountRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *RegisterAccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterAccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// Response message for rpc `RegisterAccount`.
type RegisterAccountResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	AccountId            uint32   `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterAccountResponse) Reset()         { *m = RegisterAccountResponse{} }
func (m *RegisterAccountResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterAccountResponse) ProtoMessage()    {}
func (*RegisterAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{1}
}

func (m *RegisterAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAccountResponse.Unmarshal(m, b)
}
func (m *RegisterAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAccountResponse.Marshal(b, m, deterministic)
}
func (m *RegisterAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAccountResponse.Merge(m, src)
}
func (m *RegisterAccountResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterAccountResponse.Size(m)
}
func (m *RegisterAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAccountResponse proto.InternalMessageInfo

func (m *RegisterAccountResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterAccountResponse) GetAccountId() uint32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// Request message for rpc `LogIn`.
type LogInRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{2}
}

func (m *LogInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInRequest.Unmarshal(m, b)
}
func (m *LogInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInRequest.Marshal(b, m, deterministic)
}
func (m *LogInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInRequest.Merge(m, src)
}
func (m *LogInRequest) XXX_Size() int {
	return xxx_messageInfo_LogInRequest.Size(m)
}
func (m *LogInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInRequest proto.InternalMessageInfo

func (m *LogInRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Response message for rpc `LogIn`.
type LogInResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	AccountId            uint32   `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{3}
}

func (m *LogInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInResponse.Unmarshal(m, b)
}
func (m *LogInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInResponse.Marshal(b, m, deterministic)
}
func (m *LogInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInResponse.Merge(m, src)
}
func (m *LogInResponse) XXX_Size() int {
	return xxx_messageInfo_LogInResponse.Size(m)
}
func (m *LogInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogInResponse proto.InternalMessageInfo

func (m *LogInResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogInResponse) GetAccountId() uint32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// Request message for rpc `GetAccount`.
type GetAccountRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{4}
}

func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(m, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

// Response message for rpc `GetAccount`.
type GetAccountResponse struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountResponse) Reset()         { *m = GetAccountResponse{} }
func (m *GetAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountResponse) ProtoMessage()    {}
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{5}
}

func (m *GetAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountResponse.Unmarshal(m, b)
}
func (m *GetAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountResponse.Merge(m, src)
}
func (m *GetAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountResponse.Size(m)
}
func (m *GetAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountResponse proto.InternalMessageInfo

func (m *GetAccountResponse) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *GetAccountResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// Request message for rpc `GetAccountCharacters`.
type GetAccountCharactersRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountCharactersRequest) Reset()         { *m = GetAccountCharactersRequest{} }
func (m *GetAccountCharactersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountCharactersRequest) ProtoMessage()    {}
func (*GetAccountCharactersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{6}
}

func (m *GetAccountCharactersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountCharactersRequest.Unmarshal(m, b)
}
func (m *GetAccountCharactersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountCharactersRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountCharactersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountCharactersRequest.Merge(m, src)
}
func (m *GetAccountCharactersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountCharactersRequest.Size(m)
}
func (m *GetAccountCharactersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountCharactersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountCharactersRequest proto.InternalMessageInfo

func (m *GetAccountCharactersRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

// Response message for rpc `GetAccountCharacters`.
type GetAccountCharactersResponse struct {
	CharacterIds         []uint32 `protobuf:"varint,1,rep,packed,name=character_ids,json=characterIds,proto3" json:"character_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountCharactersResponse) Reset()         { *m = GetAccountCharactersResponse{} }
func (m *GetAccountCharactersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountCharactersResponse) ProtoMessage()    {}
func (*GetAccountCharactersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{7}
}

func (m *GetAccountCharactersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountCharactersResponse.Unmarshal(m, b)
}
func (m *GetAccountCharactersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountCharactersResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountCharactersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountCharactersResponse.Merge(m, src)
}
func (m *GetAccountCharactersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountCharactersResponse.Size(m)
}
func (m *GetAccountCharactersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountCharactersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountCharactersResponse proto.InternalMessageInfo

func (m *GetAccountCharactersResponse) GetCharacterIds() []uint32 {
	if m != nil {
		return m.CharacterIds
	}
	return nil
}

// Request message for rpc `IsAccountExist`.
type IsAccountExistRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAccountExistRequest) Reset()         { *m = IsAccountExistRequest{} }
func (m *IsAccountExistRequest) String() string { return proto.CompactTextString(m) }
func (*IsAccountExistRequest) ProtoMessage()    {}
func (*IsAccountExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{8}
}

func (m *IsAccountExistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAccountExistRequest.Unmarshal(m, b)
}
func (m *IsAccountExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAccountExistRequest.Marshal(b, m, deterministic)
}
func (m *IsAccountExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountExistRequest.Merge(m, src)
}
func (m *IsAccountExistRequest) XXX_Size() int {
	return xxx_messageInfo_IsAccountExistRequest.Size(m)
}
func (m *IsAccountExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountExistRequest proto.InternalMessageInfo

func (m *IsAccountExistRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

// Response message for rpc `IsAccountExist`.
type IsAccountExistResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAccountExistResponse) Reset()         { *m = IsAccountExistResponse{} }
func (m *IsAccountExistResponse) String() string { return proto.CompactTextString(m) }
func (*IsAccountExistResponse) ProtoMessage()    {}
func (*IsAccountExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da251b36fd5e693, []int{9}
}

func (m *IsAccountExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAccountExistResponse.Unmarshal(m, b)
}
func (m *IsAccountExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAccountExistResponse.Marshal(b, m, deterministic)
}
func (m *IsAccountExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountExistResponse.Merge(m, src)
}
func (m *IsAccountExistResponse) XXX_Size() int {
	return xxx_messageInfo_IsAccountExistResponse.Size(m)
}
func (m *IsAccountExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountExistResponse proto.InternalMessageInfo

func (m *IsAccountExistResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *IsAccountExistResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterAccountRequest)(nil), "mruv.accounts.RegisterAccountRequest")
	proto.RegisterType((*RegisterAccountResponse)(nil), "mruv.accounts.RegisterAccountResponse")
	proto.RegisterType((*LogInRequest)(nil), "mruv.accounts.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "mruv.accounts.LogInResponse")
	proto.RegisterType((*GetAccountRequest)(nil), "mruv.accounts.GetAccountRequest")
	proto.RegisterType((*GetAccountResponse)(nil), "mruv.accounts.GetAccountResponse")
	proto.RegisterType((*GetAccountCharactersRequest)(nil), "mruv.accounts.GetAccountCharactersRequest")
	proto.RegisterType((*GetAccountCharactersResponse)(nil), "mruv.accounts.GetAccountCharactersResponse")
	proto.RegisterType((*IsAccountExistRequest)(nil), "mruv.accounts.IsAccountExistRequest")
	proto.RegisterType((*IsAccountExistResponse)(nil), "mruv.accounts.IsAccountExistResponse")
}

func init() { proto.RegisterFile("accounts/accounts.proto", fileDescriptor_5da251b36fd5e693) }

var fileDescriptor_5da251b36fd5e693 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x15, 0x57, 0x29, 0xed, 0xa8, 0x2e, 0x62, 0x48, 0x93, 0xc8, 0x4d, 0x45, 0x6b, 0x68,
	0x09, 0xa0, 0xd8, 0x82, 0xbe, 0x40, 0xa1, 0x42, 0x10, 0x09, 0x24, 0x64, 0x24, 0x0e, 0x1c, 0x28,
	0x8e, 0xbd, 0x72, 0x57, 0x4a, 0xbc, 0x66, 0x77, 0x1d, 0x90, 0x10, 0x42, 0xe2, 0xc0, 0x85, 0x23,
	0x17, 0xde, 0x8b, 0x57, 0xe0, 0x41, 0x90, 0xd7, 0x6b, 0x37, 0x71, 0x4c, 0x13, 0xd1, 0x5b, 0x66,
	0x3c, 0x33, 0xdf, 0xec, 0xbf, 0xff, 0x06, 0x3a, 0x7e, 0x10, 0xb0, 0x34, 0x96, 0xc2, 0x2d, 0x7e,
	0x38, 0x09, 0x67, 0x92, 0xa1, 0x39, 0xe1, 0xe9, 0xd4, 0x29, 0x92, 0x56, 0x2f, 0x62, 0x2c, 0x1a,
	0x13, 0xd7, 0x4f, 0xa8, 0xeb, 0xc7, 0x31, 0x93, 0xbe, 0xa4, 0x2c, 0xd6, 0xc5, 0xf6, 0x7b, 0x68,
	0x7b, 0x24, 0xa2, 0x42, 0x12, 0xfe, 0x38, 0xef, 0xf0, 0xc8, 0x87, 0x94, 0x08, 0x89, 0x2d, 0x68,
	0x8e, 0x59, 0x44, 0xe3, 0x6e, 0x63, 0xbf, 0xd1, 0xdf, 0xf4, 0xf2, 0x00, 0x2d, 0xd8, 0x48, 0x7c,
	0x21, 0x3e, 0x32, 0x1e, 0x76, 0x0d, 0xf5, 0xa1, 0x8c, 0xb3, 0x0e, 0x32, 0xf1, 0xe9, 0xb8, 0xbb,
	0x96, 0x77, 0xa8, 0xc0, 0xf6, 0xa0, 0xb3, 0x40, 0x10, 0x09, 0x8b, 0x05, 0xc1, 0x2e, 0x5c, 0x13,
	0x69, 0x10, 0x10, 0x21, 0x14, 0x64, 0xc3, 0x2b, 0x42, 0xdc, 0x03, 0xd0, 0x07, 0x38, 0xa3, 0x39,
	0xc8, 0xf4, 0x36, 0x75, 0x66, 0x18, 0xda, 0x27, 0xb0, 0xf5, 0x82, 0x45, 0xc3, 0xf8, 0xbf, 0x77,
	0xb5, 0x9f, 0x83, 0xa9, 0x27, 0x5c, 0x75, 0x97, 0x7b, 0x70, 0xe3, 0x19, 0x91, 0xab, 0x88, 0x67,
	0x9f, 0x00, 0xce, 0x96, 0x6a, 0x72, 0xfd, 0xf2, 0xa5, 0x98, 0xc6, 0xac, 0x98, 0xc7, 0xb0, 0x7b,
	0x31, 0xe1, 0xf4, 0xdc, 0xe7, 0x7e, 0x20, 0x09, 0x17, 0x97, 0x63, 0x4f, 0xa1, 0x57, 0xdf, 0xa4,
	0x17, 0xb8, 0x0d, 0x66, 0x50, 0x64, 0xcf, 0x68, 0x98, 0x09, 0xb0, 0xd6, 0x37, 0xbd, 0xad, 0x32,
	0x39, 0x0c, 0x85, 0x3d, 0x80, 0x9d, 0xa1, 0xd0, 0x33, 0x9e, 0x7e, 0xa2, 0x62, 0xe9, 0x51, 0xdb,
	0xd5, 0x72, 0x4d, 0x6b, 0xc3, 0x3a, 0xc9, 0x12, 0x85, 0xce, 0x3a, 0xc2, 0x6d, 0x30, 0x4a, 0x79,
	0x0d, 0x1a, 0x3e, 0xfa, 0xd1, 0x84, 0x9b, 0x2f, 0x79, 0xfa, 0x46, 0x0f, 0x11, 0xaf, 0x09, 0x9f,
	0xd2, 0x80, 0xe0, 0x57, 0xb8, 0x5e, 0xf1, 0x13, 0x1e, 0x3a, 0x73, 0x96, 0x77, 0xea, 0x1d, 0x6d,
	0x1d, 0x2d, 0x2b, 0xcb, 0x37, 0xb4, 0xf7, 0xbe, 0xfd, 0xfe, 0xf3, 0xd3, 0xe8, 0xd8, 0x3b, 0xee,
	0xf4, 0x61, 0xf9, 0xb8, 0x5c, 0xae, 0xab, 0xf1, 0x1d, 0x34, 0x95, 0x75, 0x70, 0xb7, 0x32, 0x6f,
	0xd6, 0x92, 0x56, 0xaf, 0xfe, 0xa3, 0x46, 0x58, 0x0a, 0xd1, 0xb2, 0x71, 0x0e, 0x91, 0xdf, 0xfc,
	0xf7, 0x06, 0x6c, 0xcf, 0x6b, 0x87, 0x77, 0x2a, 0xc3, 0x6a, 0x6f, 0xc2, 0x3a, 0x5c, 0x52, 0xa5,
	0xd9, 0x77, 0x15, 0xfb, 0x00, 0x6f, 0xcd, 0xb1, 0x3f, 0x2b, 0xf8, 0x97, 0xf2, 0x98, 0x24, 0x44,
	0x06, 0x70, 0xe1, 0x1b, 0xdc, 0xaf, 0x4c, 0x5f, 0x30, 0xbd, 0x75, 0x70, 0x49, 0x85, 0x66, 0xf7,
	0x14, 0xbb, 0x8d, 0xad, 0x3a, 0x36, 0xfe, 0x6a, 0x40, 0xab, 0xce, 0xa9, 0x78, 0xff, 0x9f, 0x93,
	0x17, 0xde, 0x80, 0xf5, 0x60, 0xa5, 0xda, 0x95, 0xb4, 0x28, 0x1f, 0x80, 0x78, 0xd2, 0x7f, 0x7b,
	0x14, 0x51, 0x79, 0x9e, 0x8e, 0x9c, 0x80, 0x4d, 0xdc, 0xcc, 0x97, 0x03, 0xef, 0x95, 0x9b, 0x91,
	0x06, 0xc9, 0x68, 0x10, 0xb1, 0xb2, 0x79, 0xb4, 0xae, 0xfe, 0x58, 0x8f, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x13, 0x9a, 0x58, 0xa0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MruVAccountsServiceClient is the client API for MruVAccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MruVAccountsServiceClient interface {
	// Register a new account.
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error)
	// Sign into an existing account.
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	// Check, is account with specified login exist. If yes, it returns account id.
	IsAccountExist(ctx context.Context, in *IsAccountExistRequest, opts ...grpc.CallOption) (*IsAccountExistResponse, error)
	// Get an account.
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// Get an account characters.
	GetAccountCharacters(ctx context.Context, in *GetAccountCharactersRequest, opts ...grpc.CallOption) (*GetAccountCharactersResponse, error)
}

type mruVAccountsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMruVAccountsServiceClient(cc *grpc.ClientConn) MruVAccountsServiceClient {
	return &mruVAccountsServiceClient{cc}
}

func (c *mruVAccountsServiceClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error) {
	out := new(RegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/mruv.accounts.MruVAccountsService/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVAccountsServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/mruv.accounts.MruVAccountsService/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVAccountsServiceClient) IsAccountExist(ctx context.Context, in *IsAccountExistRequest, opts ...grpc.CallOption) (*IsAccountExistResponse, error) {
	out := new(IsAccountExistResponse)
	err := c.cc.Invoke(ctx, "/mruv.accounts.MruVAccountsService/IsAccountExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVAccountsServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/mruv.accounts.MruVAccountsService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVAccountsServiceClient) GetAccountCharacters(ctx context.Context, in *GetAccountCharactersRequest, opts ...grpc.CallOption) (*GetAccountCharactersResponse, error) {
	out := new(GetAccountCharactersResponse)
	err := c.cc.Invoke(ctx, "/mruv.accounts.MruVAccountsService/GetAccountCharacters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MruVAccountsServiceServer is the server API for MruVAccountsService service.
type MruVAccountsServiceServer interface {
	// Register a new account.
	RegisterAccount(context.Context, *RegisterAccountRequest) (*RegisterAccountResponse, error)
	// Sign into an existing account.
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	// Check, is account with specified login exist. If yes, it returns account id.
	IsAccountExist(context.Context, *IsAccountExistRequest) (*IsAccountExistResponse, error)
	// Get an account.
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// Get an account characters.
	GetAccountCharacters(context.Context, *GetAccountCharactersRequest) (*GetAccountCharactersResponse, error)
}

// UnimplementedMruVAccountsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMruVAccountsServiceServer struct {
}

func (*UnimplementedMruVAccountsServiceServer) RegisterAccount(ctx context.Context, req *RegisterAccountRequest) (*RegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (*UnimplementedMruVAccountsServiceServer) LogIn(ctx context.Context, req *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (*UnimplementedMruVAccountsServiceServer) IsAccountExist(ctx context.Context, req *IsAccountExistRequest) (*IsAccountExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAccountExist not implemented")
}
func (*UnimplementedMruVAccountsServiceServer) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedMruVAccountsServiceServer) GetAccountCharacters(ctx context.Context, req *GetAccountCharactersRequest) (*GetAccountCharactersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountCharacters not implemented")
}

func RegisterMruVAccountsServiceServer(s *grpc.Server, srv MruVAccountsServiceServer) {
	s.RegisterService(&_MruVAccountsService_serviceDesc, srv)
}

func _MruVAccountsService_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVAccountsServiceServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.accounts.MruVAccountsService/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVAccountsServiceServer).RegisterAccount(ctx, req.(*RegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVAccountsService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVAccountsServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.accounts.MruVAccountsService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVAccountsServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVAccountsService_IsAccountExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAccountExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVAccountsServiceServer).IsAccountExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.accounts.MruVAccountsService/IsAccountExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVAccountsServiceServer).IsAccountExist(ctx, req.(*IsAccountExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVAccountsService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVAccountsServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.accounts.MruVAccountsService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVAccountsServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVAccountsService_GetAccountCharacters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountCharactersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVAccountsServiceServer).GetAccountCharacters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.accounts.MruVAccountsService/GetAccountCharacters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVAccountsServiceServer).GetAccountCharacters(ctx, req.(*GetAccountCharactersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MruVAccountsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mruv.accounts.MruVAccountsService",
	HandlerType: (*MruVAccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _MruVAccountsService_RegisterAccount_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _MruVAccountsService_LogIn_Handler,
		},
		{
			MethodName: "IsAccountExist",
			Handler:    _MruVAccountsService_IsAccountExist_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _MruVAccountsService_GetAccount_Handler,
		},
		{
			MethodName: "GetAccountCharacters",
			Handler:    _MruVAccountsService_GetAccountCharacters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/accounts.proto",
}
