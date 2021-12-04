// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: todoproto/todogrpc.proto

package todoproto

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

type ToDoUpdateType int32

const (
	ToDoUpdateType_Added   ToDoUpdateType = 0
	ToDoUpdateType_Updated ToDoUpdateType = 1
	ToDoUpdateType_Deleted ToDoUpdateType = 2
	ToDoUpdateType_Noop    ToDoUpdateType = 3
)

// Enum value maps for ToDoUpdateType.
var (
	ToDoUpdateType_name = map[int32]string{
		0: "Added",
		1: "Updated",
		2: "Deleted",
		3: "Noop",
	}
	ToDoUpdateType_value = map[string]int32{
		"Added":   0,
		"Updated": 1,
		"Deleted": 2,
		"Noop":    3,
	}
)

func (x ToDoUpdateType) Enum() *ToDoUpdateType {
	p := new(ToDoUpdateType)
	*p = x
	return p
}

func (x ToDoUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToDoUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_todoproto_todogrpc_proto_enumTypes[0].Descriptor()
}

func (ToDoUpdateType) Type() protoreflect.EnumType {
	return &file_todoproto_todogrpc_proto_enumTypes[0]
}

func (x ToDoUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToDoUpdateType.Descriptor instead.
func (ToDoUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{0}
}

type NewToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task   string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	UserID int32  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *NewToDo) Reset() {
	*x = NewToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewToDo) ProtoMessage() {}

func (x *NewToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewToDo.ProtoReflect.Descriptor instead.
func (*NewToDo) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{0}
}

func (x *NewToDo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *NewToDo) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task   string `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	UserID int32  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Done   bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	TodoID string `protobuf:"bytes,4,opt,name=todoID,proto3" json:"todoID,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ToDo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *ToDo) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ToDo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *ToDo) GetTodoID() string {
	if x != nil {
		return x.TodoID
	}
	return ""
}

type NoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParams) Reset() {
	*x = NoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParams) ProtoMessage() {}

func (x *NoParams) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParams.ProtoReflect.Descriptor instead.
func (*NoParams) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{2}
}

type ToDoUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType ToDoUpdateType `protobuf:"varint,1,opt,name=updateType,proto3,enum=todoproto.ToDoUpdateType" json:"updateType,omitempty"`
	Todo       *ToDo          `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *ToDoUpdate) Reset() {
	*x = ToDoUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoUpdate) ProtoMessage() {}

func (x *ToDoUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoUpdate.ProtoReflect.Descriptor instead.
func (*ToDoUpdate) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{3}
}

func (x *ToDoUpdate) GetUpdateType() ToDoUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return ToDoUpdateType_Added
}

func (x *ToDoUpdate) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type AddToDoOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *NewToDo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *AddToDoOp) Reset() {
	*x = AddToDoOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToDoOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToDoOp) ProtoMessage() {}

func (x *AddToDoOp) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToDoOp.ProtoReflect.Descriptor instead.
func (*AddToDoOp) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{4}
}

func (x *AddToDoOp) GetTodo() *NewToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateToDoOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoID string `protobuf:"bytes,1,opt,name=todoID,proto3" json:"todoID,omitempty"`
	Done   bool   `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *UpdateToDoOp) Reset() {
	*x = UpdateToDoOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateToDoOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateToDoOp) ProtoMessage() {}

func (x *UpdateToDoOp) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateToDoOp.ProtoReflect.Descriptor instead.
func (*UpdateToDoOp) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateToDoOp) GetTodoID() string {
	if x != nil {
		return x.TodoID
	}
	return ""
}

func (x *UpdateToDoOp) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type DeleteToDoOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoID string `protobuf:"bytes,1,opt,name=todoID,proto3" json:"todoID,omitempty"`
}

func (x *DeleteToDoOp) Reset() {
	*x = DeleteToDoOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteToDoOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteToDoOp) ProtoMessage() {}

func (x *DeleteToDoOp) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteToDoOp.ProtoReflect.Descriptor instead.
func (*DeleteToDoOp) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteToDoOp) GetTodoID() string {
	if x != nil {
		return x.TodoID
	}
	return ""
}

type ToDoOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Op:
	//	*ToDoOp_Add
	//	*ToDoOp_Upd
	//	*ToDoOp_Del
	Op isToDoOp_Op `protobuf_oneof:"op"`
}

func (x *ToDoOp) Reset() {
	*x = ToDoOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoOp) ProtoMessage() {}

func (x *ToDoOp) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoOp.ProtoReflect.Descriptor instead.
func (*ToDoOp) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{7}
}

func (m *ToDoOp) GetOp() isToDoOp_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *ToDoOp) GetAdd() *AddToDoOp {
	if x, ok := x.GetOp().(*ToDoOp_Add); ok {
		return x.Add
	}
	return nil
}

func (x *ToDoOp) GetUpd() *UpdateToDoOp {
	if x, ok := x.GetOp().(*ToDoOp_Upd); ok {
		return x.Upd
	}
	return nil
}

func (x *ToDoOp) GetDel() *DeleteToDoOp {
	if x, ok := x.GetOp().(*ToDoOp_Del); ok {
		return x.Del
	}
	return nil
}

type isToDoOp_Op interface {
	isToDoOp_Op()
}

type ToDoOp_Add struct {
	Add *AddToDoOp `protobuf:"bytes,1,opt,name=add,proto3,oneof"`
}

type ToDoOp_Upd struct {
	Upd *UpdateToDoOp `protobuf:"bytes,2,opt,name=upd,proto3,oneof"`
}

type ToDoOp_Del struct {
	Del *DeleteToDoOp `protobuf:"bytes,3,opt,name=del,proto3,oneof"`
}

func (*ToDoOp_Add) isToDoOp_Op() {}

func (*ToDoOp_Upd) isToDoOp_Op() {}

func (*ToDoOp_Del) isToDoOp_Op() {}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgID int64 `protobuf:"varint,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{8}
}

func (x *Ping) GetMsgID() int64 {
	if x != nil {
		return x.MsgID
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgID int64 `protobuf:"varint,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoproto_todogrpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_todoproto_todogrpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_todoproto_todogrpc_proto_rawDescGZIP(), []int{9}
}

func (x *Pong) GetMsgID() int64 {
	if x != nil {
		return x.MsgID
	}
	return 0
}

var File_todoproto_todogrpc_proto protoreflect.FileDescriptor

var file_todoproto_todogrpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x6f, 0x64, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x44, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x04,
	0x54, 0x6f, 0x44, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x22, 0x0a, 0x0a, 0x08,
	0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x6c, 0x0a, 0x0a, 0x54, 0x6f, 0x44, 0x6f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x33, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x44,
	0x6f, 0x4f, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x3a, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x64, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64,
	0x6f, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x22,
	0x92, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x12, 0x28, 0x0a, 0x03, 0x61, 0x64,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x64, 0x64, 0x12, 0x2b, 0x0a, 0x03, 0x75, 0x70, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x48, 0x00, 0x52, 0x03, 0x75, 0x70,
	0x64, 0x12, 0x2b, 0x0a, 0x03, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x48, 0x00, 0x52, 0x03, 0x64, 0x65, 0x6c, 0x42, 0x04,
	0x0a, 0x02, 0x6f, 0x70, 0x22, 0x1c, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x44, 0x22, 0x1c, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73,
	0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x44,
	0x2a, 0x3f, 0x0a, 0x0e, 0x54, 0x6f, 0x44, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6f, 0x70, 0x10,
	0x03, 0x32, 0xe0, 0x02, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x72, 0x70, 0x63, 0x12, 0x36,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x44, 0x6f, 0x12,
	0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x6f, 0x44, 0x6f, 0x1a, 0x0f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x44, 0x6f, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x44, 0x6f, 0x73, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0f, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x3e, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x44, 0x6f, 0x73, 0x12,
	0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x3a, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4f, 0x70, 0x73, 0x12, 0x11, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x4f, 0x70, 0x1a, 0x15,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x0f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x32, 0x0a, 0x04, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x13, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x6f, 0x6f, 0x66, 0x70, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todoproto_todogrpc_proto_rawDescOnce sync.Once
	file_todoproto_todogrpc_proto_rawDescData = file_todoproto_todogrpc_proto_rawDesc
)

func file_todoproto_todogrpc_proto_rawDescGZIP() []byte {
	file_todoproto_todogrpc_proto_rawDescOnce.Do(func() {
		file_todoproto_todogrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_todoproto_todogrpc_proto_rawDescData)
	})
	return file_todoproto_todogrpc_proto_rawDescData
}

var file_todoproto_todogrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todoproto_todogrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_todoproto_todogrpc_proto_goTypes = []interface{}{
	(ToDoUpdateType)(0),  // 0: todoproto.ToDoUpdateType
	(*NewToDo)(nil),      // 1: todoproto.NewToDo
	(*ToDo)(nil),         // 2: todoproto.ToDo
	(*NoParams)(nil),     // 3: todoproto.NoParams
	(*ToDoUpdate)(nil),   // 4: todoproto.ToDoUpdate
	(*AddToDoOp)(nil),    // 5: todoproto.AddToDoOp
	(*UpdateToDoOp)(nil), // 6: todoproto.UpdateToDoOp
	(*DeleteToDoOp)(nil), // 7: todoproto.DeleteToDoOp
	(*ToDoOp)(nil),       // 8: todoproto.ToDoOp
	(*Ping)(nil),         // 9: todoproto.Ping
	(*Pong)(nil),         // 10: todoproto.Pong
}
var file_todoproto_todogrpc_proto_depIdxs = []int32{
	0,  // 0: todoproto.ToDoUpdate.updateType:type_name -> todoproto.ToDoUpdateType
	2,  // 1: todoproto.ToDoUpdate.todo:type_name -> todoproto.ToDo
	1,  // 2: todoproto.AddToDoOp.todo:type_name -> todoproto.NewToDo
	5,  // 3: todoproto.ToDoOp.add:type_name -> todoproto.AddToDoOp
	6,  // 4: todoproto.ToDoOp.upd:type_name -> todoproto.UpdateToDoOp
	7,  // 5: todoproto.ToDoOp.del:type_name -> todoproto.DeleteToDoOp
	1,  // 6: todoproto.TodoGrpc.CreateNewToDo:input_type -> todoproto.NewToDo
	3,  // 7: todoproto.TodoGrpc.GetAllToDos:input_type -> todoproto.NoParams
	3,  // 8: todoproto.TodoGrpc.MonitorToDos:input_type -> todoproto.NoParams
	8,  // 9: todoproto.TodoGrpc.ApplyOps:input_type -> todoproto.ToDoOp
	9,  // 10: todoproto.TodoGrpc.KeepAlive:input_type -> todoproto.Ping
	3,  // 11: todoproto.TodoGrpc.Kill:input_type -> todoproto.NoParams
	2,  // 12: todoproto.TodoGrpc.CreateNewToDo:output_type -> todoproto.ToDo
	2,  // 13: todoproto.TodoGrpc.GetAllToDos:output_type -> todoproto.ToDo
	4,  // 14: todoproto.TodoGrpc.MonitorToDos:output_type -> todoproto.ToDoUpdate
	4,  // 15: todoproto.TodoGrpc.ApplyOps:output_type -> todoproto.ToDoUpdate
	10, // 16: todoproto.TodoGrpc.KeepAlive:output_type -> todoproto.Pong
	3,  // 17: todoproto.TodoGrpc.Kill:output_type -> todoproto.NoParams
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_todoproto_todogrpc_proto_init() }
func file_todoproto_todogrpc_proto_init() {
	if File_todoproto_todogrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todoproto_todogrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewToDo); i {
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
		file_todoproto_todogrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_todoproto_todogrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParams); i {
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
		file_todoproto_todogrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoUpdate); i {
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
		file_todoproto_todogrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToDoOp); i {
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
		file_todoproto_todogrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateToDoOp); i {
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
		file_todoproto_todogrpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteToDoOp); i {
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
		file_todoproto_todogrpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoOp); i {
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
		file_todoproto_todogrpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_todoproto_todogrpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
	file_todoproto_todogrpc_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ToDoOp_Add)(nil),
		(*ToDoOp_Upd)(nil),
		(*ToDoOp_Del)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todoproto_todogrpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todoproto_todogrpc_proto_goTypes,
		DependencyIndexes: file_todoproto_todogrpc_proto_depIdxs,
		EnumInfos:         file_todoproto_todogrpc_proto_enumTypes,
		MessageInfos:      file_todoproto_todogrpc_proto_msgTypes,
	}.Build()
	File_todoproto_todogrpc_proto = out.File
	file_todoproto_todogrpc_proto_rawDesc = nil
	file_todoproto_todogrpc_proto_goTypes = nil
	file_todoproto_todogrpc_proto_depIdxs = nil
}
