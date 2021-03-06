// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: common/spatial.proto

package common

import (
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

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_spatial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_common_spatial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_common_spatial_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Rotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rx float64 `protobuf:"fixed64,1,opt,name=rx,proto3" json:"rx,omitempty"`
	Ry float64 `protobuf:"fixed64,2,opt,name=ry,proto3" json:"ry,omitempty"`
	Rz float64 `protobuf:"fixed64,3,opt,name=rz,proto3" json:"rz,omitempty"`
}

func (x *Rotation) Reset() {
	*x = Rotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_spatial_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rotation) ProtoMessage() {}

func (x *Rotation) ProtoReflect() protoreflect.Message {
	mi := &file_common_spatial_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rotation.ProtoReflect.Descriptor instead.
func (*Rotation) Descriptor() ([]byte, []int) {
	return file_common_spatial_proto_rawDescGZIP(), []int{1}
}

func (x *Rotation) GetRx() float64 {
	if x != nil {
		return x.Rx
	}
	return 0
}

func (x *Rotation) GetRy() float64 {
	if x != nil {
		return x.Ry
	}
	return 0
}

func (x *Rotation) GetRz() float64 {
	if x != nil {
		return x.Rz
	}
	return 0
}

var File_common_spatial_proto protoreflect.FileDescriptor

var file_common_spatial_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x72, 0x75, 0x76, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x22, 0x3a, 0x0a, 0x08, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x72, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x72, 0x7a, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x72, 0x75, 0x56, 0x2d, 0x52, 0x50, 0x2f, 0x6d, 0x72, 0x75, 0x76,
	0x2d, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_spatial_proto_rawDescOnce sync.Once
	file_common_spatial_proto_rawDescData = file_common_spatial_proto_rawDesc
)

func file_common_spatial_proto_rawDescGZIP() []byte {
	file_common_spatial_proto_rawDescOnce.Do(func() {
		file_common_spatial_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_spatial_proto_rawDescData)
	})
	return file_common_spatial_proto_rawDescData
}

var file_common_spatial_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_spatial_proto_goTypes = []interface{}{
	(*Position)(nil), // 0: mruv.common.Position
	(*Rotation)(nil), // 1: mruv.common.Rotation
}
var file_common_spatial_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_spatial_proto_init() }
func file_common_spatial_proto_init() {
	if File_common_spatial_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_spatial_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_common_spatial_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rotation); i {
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
			RawDescriptor: file_common_spatial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_spatial_proto_goTypes,
		DependencyIndexes: file_common_spatial_proto_depIdxs,
		MessageInfos:      file_common_spatial_proto_msgTypes,
	}.Build()
	File_common_spatial_proto = out.File
	file_common_spatial_proto_rawDesc = nil
	file_common_spatial_proto_goTypes = nil
	file_common_spatial_proto_depIdxs = nil
}
