// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: server/server.proto

package server

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Types of server events.
type ServerEvent_ServerEventType int32

const (
	ServerEvent_UNKNOWN         ServerEvent_ServerEventType = 0
	ServerEvent_REGISTERED      ServerEvent_ServerEventType = 1
	ServerEvent_SERVER_DOWN     ServerEvent_ServerEventType = 2
	ServerEvent_SERVER_UP       ServerEvent_ServerEventType = 3
	ServerEvent_PLAYERS_CHANGED ServerEvent_ServerEventType = 4
)

// Enum value maps for ServerEvent_ServerEventType.
var (
	ServerEvent_ServerEventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "REGISTERED",
		2: "SERVER_DOWN",
		3: "SERVER_UP",
		4: "PLAYERS_CHANGED",
	}
	ServerEvent_ServerEventType_value = map[string]int32{
		"UNKNOWN":         0,
		"REGISTERED":      1,
		"SERVER_DOWN":     2,
		"SERVER_UP":       3,
		"PLAYERS_CHANGED": 4,
	}
)

func (x ServerEvent_ServerEventType) Enum() *ServerEvent_ServerEventType {
	p := new(ServerEvent_ServerEventType)
	*p = x
	return p
}

func (x ServerEvent_ServerEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerEvent_ServerEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_server_proto_enumTypes[0].Descriptor()
}

func (ServerEvent_ServerEventType) Type() protoreflect.EnumType {
	return &file_server_server_proto_enumTypes[0]
}

func (x ServerEvent_ServerEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerEvent_ServerEventType.Descriptor instead.
func (ServerEvent_ServerEventType) EnumDescriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{5, 0}
}

// Request message for `MruVServerService.GetRegisteredServers`.
type GetRegisteredServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRegisteredServersRequest) Reset() {
	*x = GetRegisteredServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisteredServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisteredServersRequest) ProtoMessage() {}

func (x *GetRegisteredServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisteredServersRequest.ProtoReflect.Descriptor instead.
func (*GetRegisteredServersRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{0}
}

// Response message for `MruVServerService.GetRegisteredServers`.
type GetRegisteredServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*ServerInfo `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *GetRegisteredServersResponse) Reset() {
	*x = GetRegisteredServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisteredServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisteredServersResponse) ProtoMessage() {}

func (x *GetRegisteredServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisteredServersResponse.ProtoReflect.Descriptor instead.
func (*GetRegisteredServersResponse) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetRegisteredServersResponse) GetServers() []*ServerInfo {
	if x != nil {
		return x.Servers
	}
	return nil
}

// Request message for `MruVServerService.UpdateServerStatus`.
type UpdateServerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the server.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status of the server.
	Status ServerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=mruv.server.ServerStatus" json:"status,omitempty"`
	// How many players are registered on that server.
	Players uint32 `protobuf:"varint,3,opt,name=players,proto3" json:"players,omitempty"`
}

func (x *UpdateServerStatusRequest) Reset() {
	*x = UpdateServerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerStatusRequest) ProtoMessage() {}

func (x *UpdateServerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerStatusRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateServerStatusRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateServerStatusRequest) GetStatus() ServerStatus {
	if x != nil {
		return x.Status
	}
	return ServerStatus_UNKNOWN
}

func (x *UpdateServerStatusRequest) GetPlayers() uint32 {
	if x != nil {
		return x.Players
	}
	return 0
}

// Response message for `MruVServerService.UpdateServerStatus`.
type UpdateServerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateServerStatusResponse) Reset() {
	*x = UpdateServerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerStatusResponse) ProtoMessage() {}

func (x *UpdateServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateServerStatusResponse) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{3}
}

// Request message for `MruVServerService.ServerEventsStream`.
type ServerEventsStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the server from which we want to receive events.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServerEventsStreamRequest) Reset() {
	*x = ServerEventsStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerEventsStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEventsStreamRequest) ProtoMessage() {}

func (x *ServerEventsStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEventsStreamRequest.ProtoReflect.Descriptor instead.
func (*ServerEventsStreamRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{4}
}

func (x *ServerEventsStreamRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Server event.
// Response message for `MruVServerService.ServerEventsStream`.
type ServerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of a server event.
	Type ServerEvent_ServerEventType `protobuf:"varint,1,opt,name=type,proto3,enum=mruv.server.ServerEvent_ServerEventType" json:"type,omitempty"`
}

func (x *ServerEvent) Reset() {
	*x = ServerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEvent) ProtoMessage() {}

func (x *ServerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEvent.ProtoReflect.Descriptor instead.
func (*ServerEvent) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{5}
}

func (x *ServerEvent) GetType() ServerEvent_ServerEventType {
	if x != nil {
		return x.Type
	}
	return ServerEvent_UNKNOWN
}

var File_server_server_proto protoreflect.FileDescriptor

var file_server_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x72,
	0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x78, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x72, 0x75,
	0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x63, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x44, 0x10, 0x04, 0x32, 0xad, 0x04, 0x0a, 0x11, 0x4d, 0x72, 0x75, 0x56, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x6d, 0x72,
	0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x17, 0x2e,
	0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x86, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x32, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x12, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x26, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x72, 0x75, 0x56, 0x2d, 0x52, 0x50, 0x2f, 0x6d, 0x72, 0x75, 0x76,
	0x2d, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_server_proto_rawDescOnce sync.Once
	file_server_server_proto_rawDescData = file_server_server_proto_rawDesc
)

func file_server_server_proto_rawDescGZIP() []byte {
	file_server_server_proto_rawDescOnce.Do(func() {
		file_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_server_proto_rawDescData)
	})
	return file_server_server_proto_rawDescData
}

var file_server_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_server_proto_goTypes = []interface{}{
	(ServerEvent_ServerEventType)(0),     // 0: mruv.server.ServerEvent.ServerEventType
	(*GetRegisteredServersRequest)(nil),  // 1: mruv.server.GetRegisteredServersRequest
	(*GetRegisteredServersResponse)(nil), // 2: mruv.server.GetRegisteredServersResponse
	(*UpdateServerStatusRequest)(nil),    // 3: mruv.server.UpdateServerStatusRequest
	(*UpdateServerStatusResponse)(nil),   // 4: mruv.server.UpdateServerStatusResponse
	(*ServerEventsStreamRequest)(nil),    // 5: mruv.server.ServerEventsStreamRequest
	(*ServerEvent)(nil),                  // 6: mruv.server.ServerEvent
	(*ServerInfo)(nil),                   // 7: mruv.server.ServerInfo
	(ServerStatus)(0),                    // 8: mruv.server.ServerStatus
	(*ServerID)(nil),                     // 9: mruv.server.ServerID
}
var file_server_server_proto_depIdxs = []int32{
	7, // 0: mruv.server.GetRegisteredServersResponse.servers:type_name -> mruv.server.ServerInfo
	8, // 1: mruv.server.UpdateServerStatusRequest.status:type_name -> mruv.server.ServerStatus
	0, // 2: mruv.server.ServerEvent.type:type_name -> mruv.server.ServerEvent.ServerEventType
	7, // 3: mruv.server.MruVServerService.RegisterServer:input_type -> mruv.server.ServerInfo
	1, // 4: mruv.server.MruVServerService.GetRegisteredServers:input_type -> mruv.server.GetRegisteredServersRequest
	9, // 5: mruv.server.MruVServerService.GetServerInfo:input_type -> mruv.server.ServerID
	3, // 6: mruv.server.MruVServerService.UpdateServerStatus:input_type -> mruv.server.UpdateServerStatusRequest
	5, // 7: mruv.server.MruVServerService.ServerEventsStream:input_type -> mruv.server.ServerEventsStreamRequest
	9, // 8: mruv.server.MruVServerService.RegisterServer:output_type -> mruv.server.ServerID
	2, // 9: mruv.server.MruVServerService.GetRegisteredServers:output_type -> mruv.server.GetRegisteredServersResponse
	7, // 10: mruv.server.MruVServerService.GetServerInfo:output_type -> mruv.server.ServerInfo
	4, // 11: mruv.server.MruVServerService.UpdateServerStatus:output_type -> mruv.server.UpdateServerStatusResponse
	6, // 12: mruv.server.MruVServerService.ServerEventsStream:output_type -> mruv.server.ServerEvent
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_server_proto_init() }
func file_server_server_proto_init() {
	if File_server_server_proto != nil {
		return
	}
	file_server_server_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_server_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisteredServersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisteredServersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerEventsStreamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_server_proto_goTypes,
		DependencyIndexes: file_server_server_proto_depIdxs,
		EnumInfos:         file_server_server_proto_enumTypes,
		MessageInfos:      file_server_server_proto_msgTypes,
	}.Build()
	File_server_server_proto = out.File
	file_server_server_proto_rawDesc = nil
	file_server_server_proto_goTypes = nil
	file_server_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MruVServerServiceClient is the client API for MruVServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MruVServerServiceClient interface {
	// Register instance of server for further managing.
	RegisterServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerID, error)
	// Get all registered servers.
	GetRegisteredServers(ctx context.Context, in *GetRegisteredServersRequest, opts ...grpc.CallOption) (*GetRegisteredServersResponse, error)
	// Get game server status.
	GetServerInfo(ctx context.Context, in *ServerID, opts ...grpc.CallOption) (*ServerInfo, error)
	// Update game server status.
	UpdateServerStatus(ctx context.Context, in *UpdateServerStatusRequest, opts ...grpc.CallOption) (*UpdateServerStatusResponse, error)
	// Stream of server events. Events are streamed back in real-time for chosen server.
	//TODO: Change name to: SubscribeServerEvents
	ServerEventsStream(ctx context.Context, in *ServerEventsStreamRequest, opts ...grpc.CallOption) (MruVServerService_ServerEventsStreamClient, error)
}

type mruVServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMruVServerServiceClient(cc grpc.ClientConnInterface) MruVServerServiceClient {
	return &mruVServerServiceClient{cc}
}

func (c *mruVServerServiceClient) RegisterServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerID, error) {
	out := new(ServerID)
	err := c.cc.Invoke(ctx, "/mruv.server.MruVServerService/RegisterServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVServerServiceClient) GetRegisteredServers(ctx context.Context, in *GetRegisteredServersRequest, opts ...grpc.CallOption) (*GetRegisteredServersResponse, error) {
	out := new(GetRegisteredServersResponse)
	err := c.cc.Invoke(ctx, "/mruv.server.MruVServerService/GetRegisteredServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVServerServiceClient) GetServerInfo(ctx context.Context, in *ServerID, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := c.cc.Invoke(ctx, "/mruv.server.MruVServerService/GetServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVServerServiceClient) UpdateServerStatus(ctx context.Context, in *UpdateServerStatusRequest, opts ...grpc.CallOption) (*UpdateServerStatusResponse, error) {
	out := new(UpdateServerStatusResponse)
	err := c.cc.Invoke(ctx, "/mruv.server.MruVServerService/UpdateServerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mruVServerServiceClient) ServerEventsStream(ctx context.Context, in *ServerEventsStreamRequest, opts ...grpc.CallOption) (MruVServerService_ServerEventsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MruVServerService_serviceDesc.Streams[0], "/mruv.server.MruVServerService/ServerEventsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &mruVServerServiceServerEventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MruVServerService_ServerEventsStreamClient interface {
	Recv() (*ServerEvent, error)
	grpc.ClientStream
}

type mruVServerServiceServerEventsStreamClient struct {
	grpc.ClientStream
}

func (x *mruVServerServiceServerEventsStreamClient) Recv() (*ServerEvent, error) {
	m := new(ServerEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MruVServerServiceServer is the server API for MruVServerService service.
type MruVServerServiceServer interface {
	// Register instance of server for further managing.
	RegisterServer(context.Context, *ServerInfo) (*ServerID, error)
	// Get all registered servers.
	GetRegisteredServers(context.Context, *GetRegisteredServersRequest) (*GetRegisteredServersResponse, error)
	// Get game server status.
	GetServerInfo(context.Context, *ServerID) (*ServerInfo, error)
	// Update game server status.
	UpdateServerStatus(context.Context, *UpdateServerStatusRequest) (*UpdateServerStatusResponse, error)
	// Stream of server events. Events are streamed back in real-time for chosen server.
	//TODO: Change name to: SubscribeServerEvents
	ServerEventsStream(*ServerEventsStreamRequest, MruVServerService_ServerEventsStreamServer) error
}

// UnimplementedMruVServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMruVServerServiceServer struct {
}

func (*UnimplementedMruVServerServiceServer) RegisterServer(context.Context, *ServerInfo) (*ServerID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterServer not implemented")
}
func (*UnimplementedMruVServerServiceServer) GetRegisteredServers(context.Context, *GetRegisteredServersRequest) (*GetRegisteredServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredServers not implemented")
}
func (*UnimplementedMruVServerServiceServer) GetServerInfo(context.Context, *ServerID) (*ServerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerInfo not implemented")
}
func (*UnimplementedMruVServerServiceServer) UpdateServerStatus(context.Context, *UpdateServerStatusRequest) (*UpdateServerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServerStatus not implemented")
}
func (*UnimplementedMruVServerServiceServer) ServerEventsStream(*ServerEventsStreamRequest, MruVServerService_ServerEventsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerEventsStream not implemented")
}

func RegisterMruVServerServiceServer(s *grpc.Server, srv MruVServerServiceServer) {
	s.RegisterService(&_MruVServerService_serviceDesc, srv)
}

func _MruVServerService_RegisterServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVServerServiceServer).RegisterServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.server.MruVServerService/RegisterServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVServerServiceServer).RegisterServer(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVServerService_GetRegisteredServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVServerServiceServer).GetRegisteredServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.server.MruVServerService/GetRegisteredServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVServerServiceServer).GetRegisteredServers(ctx, req.(*GetRegisteredServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVServerService_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVServerServiceServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.server.MruVServerService/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVServerServiceServer).GetServerInfo(ctx, req.(*ServerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVServerService_UpdateServerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MruVServerServiceServer).UpdateServerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mruv.server.MruVServerService/UpdateServerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MruVServerServiceServer).UpdateServerStatus(ctx, req.(*UpdateServerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MruVServerService_ServerEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerEventsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MruVServerServiceServer).ServerEventsStream(m, &mruVServerServiceServerEventsStreamServer{stream})
}

type MruVServerService_ServerEventsStreamServer interface {
	Send(*ServerEvent) error
	grpc.ServerStream
}

type mruVServerServiceServerEventsStreamServer struct {
	grpc.ServerStream
}

func (x *mruVServerServiceServerEventsStreamServer) Send(m *ServerEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MruVServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mruv.server.MruVServerService",
	HandlerType: (*MruVServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterServer",
			Handler:    _MruVServerService_RegisterServer_Handler,
		},
		{
			MethodName: "GetRegisteredServers",
			Handler:    _MruVServerService_GetRegisteredServers_Handler,
		},
		{
			MethodName: "GetServerInfo",
			Handler:    _MruVServerService_GetServerInfo_Handler,
		},
		{
			MethodName: "UpdateServerStatus",
			Handler:    _MruVServerService_UpdateServerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerEventsStream",
			Handler:       _MruVServerService_ServerEventsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/server.proto",
}
