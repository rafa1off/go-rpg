// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.1
// source: proto/raid.proto

package proto

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

type Raid_Boss_Race int32

const (
	Raid_Boss_HUMAN Raid_Boss_Race = 0
	Raid_Boss_ELF   Raid_Boss_Race = 1
	Raid_Boss_DWARF Raid_Boss_Race = 2
	Raid_Boss_ORC   Raid_Boss_Race = 3
)

// Enum value maps for Raid_Boss_Race.
var (
	Raid_Boss_Race_name = map[int32]string{
		0: "HUMAN",
		1: "ELF",
		2: "DWARF",
		3: "ORC",
	}
	Raid_Boss_Race_value = map[string]int32{
		"HUMAN": 0,
		"ELF":   1,
		"DWARF": 2,
		"ORC":   3,
	}
)

func (x Raid_Boss_Race) Enum() *Raid_Boss_Race {
	p := new(Raid_Boss_Race)
	*p = x
	return p
}

func (x Raid_Boss_Race) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Raid_Boss_Race) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raid_proto_enumTypes[0].Descriptor()
}

func (Raid_Boss_Race) Type() protoreflect.EnumType {
	return &file_proto_raid_proto_enumTypes[0]
}

func (x Raid_Boss_Race) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Raid_Boss_Race.Descriptor instead.
func (Raid_Boss_Race) EnumDescriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Raid_Boss_Class int32

const (
	Raid_Boss_HUNTER  Raid_Boss_Class = 0
	Raid_Boss_WARRIOR Raid_Boss_Class = 1
	Raid_Boss_MAGE    Raid_Boss_Class = 2
)

// Enum value maps for Raid_Boss_Class.
var (
	Raid_Boss_Class_name = map[int32]string{
		0: "HUNTER",
		1: "WARRIOR",
		2: "MAGE",
	}
	Raid_Boss_Class_value = map[string]int32{
		"HUNTER":  0,
		"WARRIOR": 1,
		"MAGE":    2,
	}
)

func (x Raid_Boss_Class) Enum() *Raid_Boss_Class {
	p := new(Raid_Boss_Class)
	*p = x
	return p
}

func (x Raid_Boss_Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Raid_Boss_Class) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raid_proto_enumTypes[1].Descriptor()
}

func (Raid_Boss_Class) Type() protoreflect.EnumType {
	return &file_proto_raid_proto_enumTypes[1]
}

func (x Raid_Boss_Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Raid_Boss_Class.Descriptor instead.
func (Raid_Boss_Class) EnumDescriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{0, 0, 1}
}

type Entry_Actions int32

const (
	Entry_HIT   Entry_Actions = 0
	Entry_SKILL Entry_Actions = 1
)

// Enum value maps for Entry_Actions.
var (
	Entry_Actions_name = map[int32]string{
		0: "HIT",
		1: "SKILL",
	}
	Entry_Actions_value = map[string]int32{
		"HIT":   0,
		"SKILL": 1,
	}
)

func (x Entry_Actions) Enum() *Entry_Actions {
	p := new(Entry_Actions)
	*p = x
	return p
}

func (x Entry_Actions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Entry_Actions) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raid_proto_enumTypes[2].Descriptor()
}

func (Entry_Actions) Type() protoreflect.EnumType {
	return &file_proto_raid_proto_enumTypes[2]
}

func (x Entry_Actions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Entry_Actions.Descriptor instead.
func (Entry_Actions) EnumDescriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{4, 0}
}

type Raid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool       `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Boss   *Raid_Boss `protobuf:"bytes,2,opt,name=boss,proto3" json:"boss,omitempty"`
}

func (x *Raid) Reset() {
	*x = Raid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Raid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Raid) ProtoMessage() {}

func (x *Raid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Raid.ProtoReflect.Descriptor instead.
func (*Raid) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{0}
}

func (x *Raid) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Raid) GetBoss() *Raid_Boss {
	if x != nil {
		return x.Boss
	}
	return nil
}

type RaidIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Raid *Raid `protobuf:"bytes,2,opt,name=raid,proto3,oneof" json:"raid,omitempty"`
}

func (x *RaidIn) Reset() {
	*x = RaidIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaidIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaidIn) ProtoMessage() {}

func (x *RaidIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaidIn.ProtoReflect.Descriptor instead.
func (*RaidIn) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{1}
}

func (x *RaidIn) GetId() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RaidIn) GetRaid() *Raid {
	if x != nil {
		return x.Raid
	}
	return nil
}

type RaidOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raid *Raid `protobuf:"bytes,1,opt,name=raid,proto3" json:"raid,omitempty"`
	ID   int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RaidOut) Reset() {
	*x = RaidOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaidOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaidOut) ProtoMessage() {}

func (x *RaidOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaidOut.ProtoReflect.Descriptor instead.
func (*RaidOut) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{2}
}

func (x *RaidOut) GetRaid() *Raid {
	if x != nil {
		return x.Raid
	}
	return nil
}

func (x *RaidOut) GetId() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details string `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	Raid    *Raid  `protobuf:"bytes,2,opt,name=raid,proto3,oneof" json:"raid,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{3}
}

func (x *Info) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Info) GetRaid() *Raid {
	if x != nil {
		return x.Raid
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RaidId      int32         `protobuf:"varint,1,opt,name=raidId,proto3" json:"raidId,omitempty"`
	CharacterId int32         `protobuf:"varint,2,opt,name=characterId,proto3" json:"characterId,omitempty"`
	Action      Entry_Actions `protobuf:"varint,3,opt,name=action,proto3,enum=raid.Entry_Actions" json:"action,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{4}
}

func (x *Entry) GetRaidId() int32 {
	if x != nil {
		return x.RaidId
	}
	return 0
}

func (x *Entry) GetCharacterId() int32 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *Entry) GetAction() Entry_Actions {
	if x != nil {
		return x.Action
	}
	return Entry_HIT
}

type Raid_Boss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Health int32           `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
	Race   Raid_Boss_Race  `protobuf:"varint,3,opt,name=race,proto3,enum=raid.Raid_Boss_Race" json:"race,omitempty"`
	Class  Raid_Boss_Class `protobuf:"varint,4,opt,name=class,proto3,enum=raid.Raid_Boss_Class" json:"class,omitempty"`
}

func (x *Raid_Boss) Reset() {
	*x = Raid_Boss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raid_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Raid_Boss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Raid_Boss) ProtoMessage() {}

func (x *Raid_Boss) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raid_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Raid_Boss.ProtoReflect.Descriptor instead.
func (*Raid_Boss) Descriptor() ([]byte, []int) {
	return file_proto_raid_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Raid_Boss) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Raid_Boss) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Raid_Boss) GetRace() Raid_Boss_Race {
	if x != nil {
		return x.Race
	}
	return Raid_Boss_HUMAN
}

func (x *Raid_Boss) GetClass() Raid_Boss_Class {
	if x != nil {
		return x.Class
	}
	return Raid_Boss_HUNTER
}

var File_proto_raid_proto protoreflect.FileDescriptor

var file_proto_raid_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x69, 0x64, 0x22, 0xab, 0x02, 0x0a, 0x04, 0x52, 0x61, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x52,
	0x61, 0x69, 0x64, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x73, 0x73, 0x1a, 0xe5,
	0x01, 0x0a, 0x04, 0x42, 0x6f, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x2e, 0x42, 0x6f,
	0x73, 0x73, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x04, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2b, 0x0a,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72,
	0x61, 0x69, 0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x52, 0x61,
	0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x4c, 0x46, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x57, 0x41, 0x52, 0x46, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x52, 0x43, 0x10, 0x03, 0x22, 0x2a, 0x0a, 0x05, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x52, 0x49, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x22, 0x46, 0x0a, 0x06, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x04, 0x72, 0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x48, 0x00, 0x52, 0x04, 0x72, 0x61,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x61, 0x69, 0x64, 0x22, 0x39,
	0x0a, 0x07, 0x52, 0x61, 0x69, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x61, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x52,
	0x61, 0x69, 0x64, 0x52, 0x04, 0x72, 0x61, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x72,
	0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x61, 0x69, 0x64,
	0x2e, 0x52, 0x61, 0x69, 0x64, 0x48, 0x00, 0x52, 0x04, 0x72, 0x61, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x61, 0x69, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x72, 0x61, 0x69, 0x64, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x07, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x49, 0x54, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x01, 0x32, 0x96, 0x01, 0x0a, 0x05, 0x52, 0x61,
	0x69, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e,
	0x72, 0x61, 0x69, 0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x1a, 0x0a, 0x2e, 0x72, 0x61, 0x69, 0x64,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x72,
	0x61, 0x69, 0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x1a, 0x0d, 0x2e, 0x72, 0x61, 0x69,
	0x64, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x0b, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a,
	0x0a, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x21, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x0c, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e,
	0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x72, 0x61, 0x69, 0x64, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raid_proto_rawDescOnce sync.Once
	file_proto_raid_proto_rawDescData = file_proto_raid_proto_rawDesc
)

func file_proto_raid_proto_rawDescGZIP() []byte {
	file_proto_raid_proto_rawDescOnce.Do(func() {
		file_proto_raid_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raid_proto_rawDescData)
	})
	return file_proto_raid_proto_rawDescData
}

var file_proto_raid_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_raid_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_raid_proto_goTypes = []interface{}{
	(Raid_Boss_Race)(0),  // 0: raid.Raid.Boss.Race
	(Raid_Boss_Class)(0), // 1: raid.Raid.Boss.Class
	(Entry_Actions)(0),   // 2: raid.Entry.Actions
	(*Raid)(nil),         // 3: raid.Raid
	(*RaidIn)(nil),       // 4: raid.RaidIn
	(*RaidOut)(nil),      // 5: raid.RaidOut
	(*Info)(nil),         // 6: raid.Info
	(*Entry)(nil),        // 7: raid.Entry
	(*Raid_Boss)(nil),    // 8: raid.Raid.Boss
}
var file_proto_raid_proto_depIdxs = []int32{
	8,  // 0: raid.Raid.boss:type_name -> raid.Raid.Boss
	3,  // 1: raid.RaidIn.raid:type_name -> raid.Raid
	3,  // 2: raid.RaidOut.raid:type_name -> raid.Raid
	3,  // 3: raid.Info.raid:type_name -> raid.Raid
	2,  // 4: raid.Entry.action:type_name -> raid.Entry.Actions
	0,  // 5: raid.Raid.Boss.race:type_name -> raid.Raid.Boss.Race
	1,  // 6: raid.Raid.Boss.class:type_name -> raid.Raid.Boss.Class
	3,  // 7: raid.Raids.Create:input_type -> raid.Raid
	4,  // 8: raid.Raids.Get:input_type -> raid.RaidIn
	7,  // 9: raid.Raids.Enter:input_type -> raid.Entry
	4,  // 10: raid.Raids.Leave:input_type -> raid.RaidIn
	6,  // 11: raid.Raids.Create:output_type -> raid.Info
	5,  // 12: raid.Raids.Get:output_type -> raid.RaidOut
	6,  // 13: raid.Raids.Enter:output_type -> raid.Info
	6,  // 14: raid.Raids.Leave:output_type -> raid.Info
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_raid_proto_init() }
func file_proto_raid_proto_init() {
	if File_proto_raid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Raid); i {
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
		file_proto_raid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaidIn); i {
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
		file_proto_raid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaidOut); i {
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
		file_proto_raid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_proto_raid_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_proto_raid_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Raid_Boss); i {
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
	file_proto_raid_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_proto_raid_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_raid_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raid_proto_goTypes,
		DependencyIndexes: file_proto_raid_proto_depIdxs,
		EnumInfos:         file_proto_raid_proto_enumTypes,
		MessageInfos:      file_proto_raid_proto_msgTypes,
	}.Build()
	File_proto_raid_proto = out.File
	file_proto_raid_proto_rawDesc = nil
	file_proto_raid_proto_goTypes = nil
	file_proto_raid_proto_depIdxs = nil
}
