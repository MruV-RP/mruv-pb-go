// Code generated by protoc-gen-go. DO NOT EDIT.
// source: items/items_model.proto

package items

import (
	fmt "fmt"
	common "github.com/MruV-RP/mruv-pb-go/common"
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

type ItemType struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	BaseWeight           float32  `protobuf:"fixed32,4,opt,name=base_weight,json=baseWeight,proto3" json:"base_weight,omitempty"`
	BaseVolume           float32  `protobuf:"fixed32,5,opt,name=base_volume,json=baseVolume,proto3" json:"base_volume,omitempty"`
	ModelName            string   `protobuf:"bytes,6,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	ModelHash            int32    `protobuf:"varint,7,opt,name=model_hash,json=modelHash,proto3" json:"model_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemType) Reset()         { *m = ItemType{} }
func (m *ItemType) String() string { return proto.CompactTextString(m) }
func (*ItemType) ProtoMessage()    {}
func (*ItemType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{0}
}

func (m *ItemType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemType.Unmarshal(m, b)
}
func (m *ItemType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemType.Marshal(b, m, deterministic)
}
func (m *ItemType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemType.Merge(m, src)
}
func (m *ItemType) XXX_Size() int {
	return xxx_messageInfo_ItemType.Size(m)
}
func (m *ItemType) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemType.DiscardUnknown(m)
}

var xxx_messageInfo_ItemType proto.InternalMessageInfo

func (m *ItemType) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ItemType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ItemType) GetBaseWeight() float32 {
	if m != nil {
		return m.BaseWeight
	}
	return 0
}

func (m *ItemType) GetBaseVolume() float32 {
	if m != nil {
		return m.BaseVolume
	}
	return 0
}

func (m *ItemType) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ItemType) GetModelHash() int32 {
	if m != nil {
		return m.ModelHash
	}
	return 0
}

type ItemTypeID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemTypeID) Reset()         { *m = ItemTypeID{} }
func (m *ItemTypeID) String() string { return proto.CompactTextString(m) }
func (*ItemTypeID) ProtoMessage()    {}
func (*ItemTypeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{1}
}

func (m *ItemTypeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemTypeID.Unmarshal(m, b)
}
func (m *ItemTypeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemTypeID.Marshal(b, m, deterministic)
}
func (m *ItemTypeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemTypeID.Merge(m, src)
}
func (m *ItemTypeID) XXX_Size() int {
	return xxx_messageInfo_ItemTypeID.Size(m)
}
func (m *ItemTypeID) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemTypeID.DiscardUnknown(m)
}

var xxx_messageInfo_ItemTypeID proto.InternalMessageInfo

func (m *ItemTypeID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemTypeId           int32    `protobuf:"varint,2,opt,name=item_type_id,json=itemTypeId,proto3" json:"item_type_id,omitempty"`
	Weight               float32  `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Volume               float32  `protobuf:"fixed32,4,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{2}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetItemTypeId() int32 {
	if m != nil {
		return m.ItemTypeId
	}
	return 0
}

func (m *Item) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Item) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type ItemID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemID) Reset()         { *m = ItemID{} }
func (m *ItemID) String() string { return proto.CompactTextString(m) }
func (*ItemID) ProtoMessage()    {}
func (*ItemID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{3}
}

func (m *ItemID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemID.Unmarshal(m, b)
}
func (m *ItemID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemID.Marshal(b, m, deterministic)
}
func (m *ItemID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemID.Merge(m, src)
}
func (m *ItemID) XXX_Size() int {
	return xxx_messageInfo_ItemID.Size(m)
}
func (m *ItemID) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemID.DiscardUnknown(m)
}

var xxx_messageInfo_ItemID proto.InternalMessageInfo

func (m *ItemID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ContainerType struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ContainerItemTypeId  int32    `protobuf:"varint,2,opt,name=container_item_type_id,json=containerItemTypeId,proto3" json:"container_item_type_id,omitempty"`
	MaxNumber            int32    `protobuf:"varint,3,opt,name=max_number,json=maxNumber,proto3" json:"max_number,omitempty"`
	MaxVolume            float32  `protobuf:"fixed32,4,opt,name=max_volume,json=maxVolume,proto3" json:"max_volume,omitempty"`
	MaxWeight            float32  `protobuf:"fixed32,5,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	ValidItemTypes       []int32  `protobuf:"varint,8,rep,packed,name=valid_item_types,json=validItemTypes,proto3" json:"valid_item_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerType) Reset()         { *m = ContainerType{} }
func (m *ContainerType) String() string { return proto.CompactTextString(m) }
func (*ContainerType) ProtoMessage()    {}
func (*ContainerType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{4}
}

func (m *ContainerType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerType.Unmarshal(m, b)
}
func (m *ContainerType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerType.Marshal(b, m, deterministic)
}
func (m *ContainerType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerType.Merge(m, src)
}
func (m *ContainerType) XXX_Size() int {
	return xxx_messageInfo_ContainerType.Size(m)
}
func (m *ContainerType) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerType.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerType proto.InternalMessageInfo

func (m *ContainerType) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContainerType) GetContainerItemTypeId() int32 {
	if m != nil {
		return m.ContainerItemTypeId
	}
	return 0
}

func (m *ContainerType) GetMaxNumber() int32 {
	if m != nil {
		return m.MaxNumber
	}
	return 0
}

func (m *ContainerType) GetMaxVolume() float32 {
	if m != nil {
		return m.MaxVolume
	}
	return 0
}

func (m *ContainerType) GetMaxWeight() float32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *ContainerType) GetValidItemTypes() []int32 {
	if m != nil {
		return m.ValidItemTypes
	}
	return nil
}

type Container struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               int32         `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	ItemId               int64         `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemsInside          int32         `protobuf:"varint,4,opt,name=items_inside,json=itemsInside,proto3" json:"items_inside,omitempty"`
	Items                []*InsideItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{5}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Container) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Container) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Container) GetItemsInside() int32 {
	if m != nil {
		return m.ItemsInside
	}
	return 0
}

func (m *Container) GetItems() []*InsideItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ContainerID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerID) Reset()         { *m = ContainerID{} }
func (m *ContainerID) String() string { return proto.CompactTextString(m) }
func (*ContainerID) ProtoMessage()    {}
func (*ContainerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{6}
}

func (m *ContainerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerID.Unmarshal(m, b)
}
func (m *ContainerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerID.Marshal(b, m, deterministic)
}
func (m *ContainerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerID.Merge(m, src)
}
func (m *ContainerID) XXX_Size() int {
	return xxx_messageInfo_ContainerID.Size(m)
}
func (m *ContainerID) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerID.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerID proto.InternalMessageInfo

func (m *ContainerID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type InsideItem struct {
	ContainerId          int32            `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ItemId               int64            `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Position             *common.Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Rotation             *common.Rotation `protobuf:"bytes,4,opt,name=rotation,proto3" json:"rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InsideItem) Reset()         { *m = InsideItem{} }
func (m *InsideItem) String() string { return proto.CompactTextString(m) }
func (*InsideItem) ProtoMessage()    {}
func (*InsideItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_f923b2189f788928, []int{7}
}

func (m *InsideItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsideItem.Unmarshal(m, b)
}
func (m *InsideItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsideItem.Marshal(b, m, deterministic)
}
func (m *InsideItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsideItem.Merge(m, src)
}
func (m *InsideItem) XXX_Size() int {
	return xxx_messageInfo_InsideItem.Size(m)
}
func (m *InsideItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InsideItem.DiscardUnknown(m)
}

var xxx_messageInfo_InsideItem proto.InternalMessageInfo

func (m *InsideItem) GetContainerId() int32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *InsideItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *InsideItem) GetPosition() *common.Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *InsideItem) GetRotation() *common.Rotation {
	if m != nil {
		return m.Rotation
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemType)(nil), "mruv.ItemType")
	proto.RegisterType((*ItemTypeID)(nil), "mruv.ItemTypeID")
	proto.RegisterType((*Item)(nil), "mruv.Item")
	proto.RegisterType((*ItemID)(nil), "mruv.ItemID")
	proto.RegisterType((*ContainerType)(nil), "mruv.ContainerType")
	proto.RegisterType((*Container)(nil), "mruv.Container")
	proto.RegisterType((*ContainerID)(nil), "mruv.ContainerID")
	proto.RegisterType((*InsideItem)(nil), "mruv.InsideItem")
}

func init() { proto.RegisterFile("items/items_model.proto", fileDescriptor_f923b2189f788928) }

var fileDescriptor_f923b2189f788928 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0xd3, 0x64, 0xdb, 0x13, 0x2d, 0x4b, 0x94, 0x6d, 0x10, 0x17, 0x63, 0x44, 0x09,
	0x42, 0x5b, 0xd8, 0x7d, 0x03, 0xf5, 0xc2, 0x5e, 0xb8, 0x2c, 0x83, 0xac, 0xe0, 0x4d, 0x98, 0x76,
	0x86, 0x66, 0xa0, 0x93, 0x09, 0x99, 0x69, 0xed, 0xbe, 0x89, 0x97, 0x3e, 0x92, 0xf7, 0xbe, 0x8c,
	0xcc, 0xc9, 0x9f, 0xa6, 0xda, 0x9b, 0xd2, 0xf9, 0xbe, 0x93, 0x9c, 0xdf, 0x77, 0x4e, 0x06, 0xa6,
	0xc2, 0x70, 0xa9, 0x17, 0xf8, 0x9b, 0x49, 0xc5, 0xf8, 0x76, 0x5e, 0x56, 0xca, 0xa8, 0x70, 0x28,
	0xab, 0xdd, 0xfe, 0xc5, 0xf3, 0xb5, 0x92, 0x52, 0x15, 0x0b, 0x5d, 0x52, 0x23, 0x68, 0xe3, 0x25,
	0xbf, 0x1d, 0x18, 0x2d, 0x0d, 0x97, 0x5f, 0x1f, 0x4b, 0x1e, 0x4e, 0x60, 0x20, 0x58, 0xe4, 0xc4,
	0x4e, 0xea, 0x91, 0x81, 0x60, 0x61, 0x08, 0xc3, 0x82, 0x4a, 0x1e, 0x0d, 0x62, 0x27, 0x1d, 0x13,
	0xfc, 0x1f, 0xc6, 0x10, 0x30, 0xae, 0xd7, 0x95, 0x28, 0x8d, 0x50, 0x45, 0xe4, 0xa2, 0xd5, 0x97,
	0xc2, 0x57, 0x10, 0xac, 0xa8, 0xe6, 0xd9, 0x0f, 0x2e, 0x36, 0xb9, 0x89, 0x86, 0xb1, 0x93, 0x0e,
	0x08, 0x58, 0xe9, 0x1b, 0x2a, 0x5d, 0xc1, 0x5e, 0x6d, 0x77, 0x92, 0x47, 0xde, 0xb1, 0xe0, 0x01,
	0x95, 0xf0, 0x1a, 0x00, 0xf9, 0x33, 0xec, 0xee, 0x63, 0x8b, 0x31, 0x2a, 0x77, 0xb4, 0x6f, 0xe7,
	0x54, 0xe7, 0xd1, 0x05, 0xe2, 0xd6, 0xf6, 0x67, 0xaa, 0xf3, 0xe4, 0x25, 0x40, 0x9b, 0x68, 0xf9,
	0xe9, 0xdf, 0x4c, 0x49, 0x0e, 0x43, 0xeb, 0xf6, 0x74, 0x17, 0xb3, 0xc6, 0xf0, 0xc4, 0x4e, 0x2e,
	0x33, 0x8f, 0x25, 0xcf, 0x04, 0xc3, 0xcc, 0x1e, 0x01, 0xd1, 0xbe, 0x89, 0x85, 0x57, 0xe0, 0x37,
	0x91, 0x5c, 0x24, 0x6e, 0x4e, 0x56, 0x6f, 0x92, 0xd4, 0x51, 0x9b, 0x53, 0x12, 0x81, 0x6f, 0x3b,
	0x9d, 0x30, 0x60, 0xaf, 0xe4, 0x8f, 0x03, 0x4f, 0x3f, 0xaa, 0xc2, 0x50, 0x51, 0xf0, 0xea, 0xec,
	0xe4, 0x6f, 0xe1, 0x6a, 0xdd, 0x16, 0x64, 0x67, 0xb8, 0x9e, 0x75, 0xee, 0xf2, 0x08, 0x68, 0xe7,
	0x42, 0x0f, 0x59, 0xb1, 0x93, 0x2b, 0x5e, 0x21, 0xa4, 0x9d, 0x0b, 0x3d, 0xdc, 0xa1, 0xd0, 0xda,
	0x27, 0xac, 0xd6, 0xee, 0x0d, 0x9d, 0x1e, 0xda, 0xad, 0x79, 0x9d, 0xdd, 0x2c, 0x2d, 0x85, 0xcb,
	0x3d, 0xdd, 0x0a, 0x76, 0xa4, 0xd1, 0xd1, 0x28, 0x76, 0x53, 0x8f, 0x4c, 0x50, 0x6f, 0x39, 0x74,
	0xf2, 0xd3, 0x81, 0x71, 0x97, 0xee, 0xbf, 0x64, 0x53, 0xb8, 0x38, 0x8d, 0xe2, 0x9b, 0x9a, 0x7e,
	0x0a, 0x17, 0xf8, 0x6a, 0xc1, 0x10, 0xdd, 0x25, 0xbe, 0x3d, 0x2e, 0x59, 0xf8, 0xba, 0xde, 0x8c,
	0xce, 0x44, 0xa1, 0x05, 0xab, 0xc9, 0x3d, 0x12, 0xa0, 0xb6, 0x44, 0x29, 0x7c, 0x07, 0x1e, 0x1e,
	0x23, 0x2f, 0x76, 0xd3, 0xe0, 0xe6, 0x72, 0x6e, 0xbf, 0xf8, 0x79, 0x6d, 0x5a, 0x30, 0x52, 0xdb,
	0xc9, 0x35, 0x04, 0x1d, 0xd9, 0x99, 0x6f, 0xe3, 0x97, 0x03, 0x70, 0x7c, 0xc8, 0x36, 0xee, 0x2d,
	0xa1, 0x2d, 0x0c, 0x8e, 0xa3, 0x3f, 0x81, 0x1e, 0x9c, 0x40, 0xbf, 0x87, 0x51, 0xa9, 0xb4, 0xe8,
	0xee, 0x48, 0x70, 0x33, 0xa9, 0xa1, 0xee, 0x1b, 0x95, 0x74, 0xbe, 0xad, 0xad, 0x94, 0xa1, 0x58,
	0x3b, 0xec, 0xd7, 0x92, 0x46, 0x25, 0x9d, 0xff, 0xe1, 0xed, 0xf7, 0x37, 0x1b, 0x61, 0xf2, 0xdd,
	0x6a, 0xbe, 0x56, 0x72, 0xf1, 0xa5, 0xda, 0x3d, 0xcc, 0xc8, 0xfd, 0xc2, 0x56, 0xcf, 0xca, 0xd5,
	0x6c, 0xa3, 0xea, 0xeb, 0xbf, 0xf2, 0xf1, 0x76, 0xdf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0x31, 0x9b, 0x85, 0x14, 0x04, 0x00, 0x00,
}
