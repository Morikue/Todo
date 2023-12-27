// Протокол использует синтаксис proto3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: todos.proto

// Определение пакета userservice внутри файла .proto

package todo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodoID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoID) Reset() {
	*x = TodoID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoID) ProtoMessage() {}

func (x *TodoID) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoID.ProtoReflect.Descriptor instead.
func (*TodoID) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{0}
}

func (x *TodoID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ShortTodoDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy   int32  `protobuf:"varint,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Assignee    int32  `protobuf:"varint,3,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ShortTodoDTO) Reset() {
	*x = ShortTodoDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortTodoDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortTodoDTO) ProtoMessage() {}

func (x *ShortTodoDTO) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortTodoDTO.ProtoReflect.Descriptor instead.
func (*ShortTodoDTO) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{1}
}

func (x *ShortTodoDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShortTodoDTO) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ShortTodoDTO) GetAssignee() int32 {
	if x != nil {
		return x.Assignee
	}
	return 0
}

func (x *ShortTodoDTO) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FullTodoDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy   int32                  `protobuf:"varint,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Assignee    int32                  `protobuf:"varint,3,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FullTodoDTO) Reset() {
	*x = FullTodoDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullTodoDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullTodoDTO) ProtoMessage() {}

func (x *FullTodoDTO) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullTodoDTO.ProtoReflect.Descriptor instead.
func (*FullTodoDTO) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{2}
}

func (x *FullTodoDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FullTodoDTO) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *FullTodoDTO) GetAssignee() int32 {
	if x != nil {
		return x.Assignee
	}
	return 0
}

func (x *FullTodoDTO) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FullTodoDTO) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FullTodoDTO) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedBy int32                  `protobuf:"varint,1,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Assignee  int32                  `protobuf:"varint,2,opt,name=assignee,proto3" json:"assignee,omitempty"`
	DateFrom  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *GetTodosRequest) Reset() {
	*x = GetTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosRequest) ProtoMessage() {}

func (x *GetTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosRequest.ProtoReflect.Descriptor instead.
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodosRequest) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *GetTodosRequest) GetAssignee() int32 {
	if x != nil {
		return x.Assignee
	}
	return 0
}

func (x *GetTodosRequest) GetDateFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.DateFrom
	}
	return nil
}

func (x *GetTodosRequest) GetDateTo() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTo
	}
	return nil
}

type GetTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FullTodoDTO `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetTodosResponse) Reset() {
	*x = GetTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosResponse) ProtoMessage() {}

func (x *GetTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosResponse.ProtoReflect.Descriptor instead.
func (*GetTodosResponse) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodosResponse) GetItems() []*FullTodoDTO {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_todos_proto protoreflect.FileDescriptor

var file_todos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x54, 0x6f, 0x64, 0x6f,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x7b, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x44,
	0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xf0, 0x01, 0x0a, 0x0b, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22,
	0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0xd5, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44,
	0x6f, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x1a, 0x18, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54,
	0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x44, 0x6f, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x1a,
	0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75,
	0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x1a, 0x18, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x6f, 0x64, 0x6f, 0x44, 0x54, 0x4f, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x44, 0x6f, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x13,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x72, 0x69, 0x6b, 0x75,
	0x65, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_proto_rawDescOnce sync.Once
	file_todos_proto_rawDescData = file_todos_proto_rawDesc
)

func file_todos_proto_rawDescGZIP() []byte {
	file_todos_proto_rawDescOnce.Do(func() {
		file_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_proto_rawDescData)
	})
	return file_todos_proto_rawDescData
}

var file_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_todos_proto_goTypes = []interface{}{
	(*TodoID)(nil),                // 0: todoservice.TodoID
	(*ShortTodoDTO)(nil),          // 1: todoservice.ShortTodoDTO
	(*FullTodoDTO)(nil),           // 2: todoservice.FullTodoDTO
	(*GetTodosRequest)(nil),       // 3: todoservice.GetTodosRequest
	(*GetTodosResponse)(nil),      // 4: todoservice.GetTodosResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_todos_proto_depIdxs = []int32{
	5,  // 0: todoservice.FullTodoDTO.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: todoservice.FullTodoDTO.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 2: todoservice.GetTodosRequest.date_from:type_name -> google.protobuf.Timestamp
	5,  // 3: todoservice.GetTodosRequest.date_to:type_name -> google.protobuf.Timestamp
	2,  // 4: todoservice.GetTodosResponse.items:type_name -> todoservice.FullTodoDTO
	1,  // 5: todoservice.TodoService.CreateToDo:input_type -> todoservice.ShortTodoDTO
	1,  // 6: todoservice.TodoService.UpdateToDo:input_type -> todoservice.ShortTodoDTO
	0,  // 7: todoservice.TodoService.GetTodoById:input_type -> todoservice.TodoID
	3,  // 8: todoservice.TodoService.GetToDos:input_type -> todoservice.GetTodosRequest
	0,  // 9: todoservice.TodoService.DeleteTodo:input_type -> todoservice.TodoID
	2,  // 10: todoservice.TodoService.CreateToDo:output_type -> todoservice.FullTodoDTO
	2,  // 11: todoservice.TodoService.UpdateToDo:output_type -> todoservice.FullTodoDTO
	2,  // 12: todoservice.TodoService.GetTodoById:output_type -> todoservice.FullTodoDTO
	4,  // 13: todoservice.TodoService.GetToDos:output_type -> todoservice.GetTodosResponse
	6,  // 14: todoservice.TodoService.DeleteTodo:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_todos_proto_init() }
func file_todos_proto_init() {
	if File_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoID); i {
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
		file_todos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortTodoDTO); i {
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
		file_todos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullTodoDTO); i {
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
		file_todos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosRequest); i {
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
		file_todos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosResponse); i {
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
			RawDescriptor: file_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_proto_goTypes,
		DependencyIndexes: file_todos_proto_depIdxs,
		MessageInfos:      file_todos_proto_msgTypes,
	}.Build()
	File_todos_proto = out.File
	file_todos_proto_rawDesc = nil
	file_todos_proto_goTypes = nil
	file_todos_proto_depIdxs = nil
}
