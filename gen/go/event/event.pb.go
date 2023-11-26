// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: interfaces/event.proto

package event

import (
	mail "./mail"
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

type Events int32

const (
	Events_A Events = 0
)

// Enum value maps for Events.
var (
	Events_name = map[int32]string{
		0: "A",
	}
	Events_value = map[string]int32{
		"A": 0,
	}
)

func (x Events) Enum() *Events {
	p := new(Events)
	*p = x
	return p
}

func (x Events) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Events) Descriptor() protoreflect.EnumDescriptor {
	return file_interfaces_event_proto_enumTypes[0].Descriptor()
}

func (Events) Type() protoreflect.EnumType {
	return &file_interfaces_event_proto_enumTypes[0]
}

func (x Events) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Events.Descriptor instead.
func (Events) EnumDescriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{0}
}

type IEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *IEvent) Reset() {
	*x = IEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IEvent) ProtoMessage() {}

func (x *IEvent) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IEvent.ProtoReflect.Descriptor instead.
func (*IEvent) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{0}
}

func (x *IEvent) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IEvent) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type CreateEventDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *CreateEventDto) Reset() {
	*x = CreateEventDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventDto) ProtoMessage() {}

func (x *CreateEventDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventDto.ProtoReflect.Descriptor instead.
func (*CreateEventDto) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEventDto) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type GetEventDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetEventDto) Reset() {
	*x = GetEventDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventDto) ProtoMessage() {}

func (x *GetEventDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventDto.ProtoReflect.Descriptor instead.
func (*GetEventDto) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EmitEventDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload *EmitEventDto_Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *EmitEventDto) Reset() {
	*x = EmitEventDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitEventDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitEventDto) ProtoMessage() {}

func (x *EmitEventDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitEventDto.ProtoReflect.Descriptor instead.
func (*EmitEventDto) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{3}
}

func (x *EmitEventDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmitEventDto) GetPayload() *EmitEventDto_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type EmitEventResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EmitEventResult) Reset() {
	*x = EmitEventResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitEventResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitEventResult) ProtoMessage() {}

func (x *EmitEventResult) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitEventResult.ProtoReflect.Descriptor instead.
func (*EmitEventResult) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{4}
}

func (x *EmitEventResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EmitEventResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmitEventDto_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*EmitEventDto_Payload_Mail
	Value isEmitEventDto_Payload_Value `protobuf_oneof:"value"`
}

func (x *EmitEventDto_Payload) Reset() {
	*x = EmitEventDto_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitEventDto_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitEventDto_Payload) ProtoMessage() {}

func (x *EmitEventDto_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitEventDto_Payload.ProtoReflect.Descriptor instead.
func (*EmitEventDto_Payload) Descriptor() ([]byte, []int) {
	return file_interfaces_event_proto_rawDescGZIP(), []int{3, 0}
}

func (m *EmitEventDto_Payload) GetValue() isEmitEventDto_Payload_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *EmitEventDto_Payload) GetMail() *mail.SendMailDto {
	if x, ok := x.GetValue().(*EmitEventDto_Payload_Mail); ok {
		return x.Mail
	}
	return nil
}

type isEmitEventDto_Payload_Value interface {
	isEmitEventDto_Payload_Value()
}

type EmitEventDto_Payload_Mail struct {
	Mail *mail.SendMailDto `protobuf:"bytes,2,opt,name=mail,json=payload,proto3,oneof"`
}

func (*EmitEventDto_Payload_Mail) isEmitEventDto_Payload_Value() {}

var File_interfaces_event_proto protoreflect.FileDescriptor

var file_interfaces_event_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x06, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0c,
	0x45, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x69, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x3e, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x44, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x45, 0x6d, 0x69, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x0f, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x32, 0xa1, 0x01, 0x0a, 0x10, 0x49, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a,
	0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x45, 0x6d, 0x69, 0x74,
	0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interfaces_event_proto_rawDescOnce sync.Once
	file_interfaces_event_proto_rawDescData = file_interfaces_event_proto_rawDesc
)

func file_interfaces_event_proto_rawDescGZIP() []byte {
	file_interfaces_event_proto_rawDescOnce.Do(func() {
		file_interfaces_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_interfaces_event_proto_rawDescData)
	})
	return file_interfaces_event_proto_rawDescData
}

var file_interfaces_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interfaces_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_interfaces_event_proto_goTypes = []interface{}{
	(Events)(0),                  // 0: event.Events
	(*IEvent)(nil),               // 1: event.IEvent
	(*CreateEventDto)(nil),       // 2: event.CreateEventDto
	(*GetEventDto)(nil),          // 3: event.GetEventDto
	(*EmitEventDto)(nil),         // 4: event.EmitEventDto
	(*EmitEventResult)(nil),      // 5: event.EmitEventResult
	(*EmitEventDto_Payload)(nil), // 6: event.EmitEventDto.Payload
	(*mail.SendMailDto)(nil),     // 7: mail.SendMailDto
}
var file_interfaces_event_proto_depIdxs = []int32{
	6, // 0: event.EmitEventDto.payload:type_name -> event.EmitEventDto.Payload
	7, // 1: event.EmitEventDto.Payload.mail:type_name -> mail.SendMailDto
	2, // 2: event.IEventController.Create:input_type -> event.CreateEventDto
	3, // 3: event.IEventController.Get:input_type -> event.GetEventDto
	4, // 4: event.IEventController.Emit:input_type -> event.EmitEventDto
	1, // 5: event.IEventController.Create:output_type -> event.IEvent
	1, // 6: event.IEventController.Get:output_type -> event.IEvent
	5, // 7: event.IEventController.Emit:output_type -> event.EmitEventResult
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interfaces_event_proto_init() }
func file_interfaces_event_proto_init() {
	if File_interfaces_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interfaces_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IEvent); i {
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
		file_interfaces_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventDto); i {
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
		file_interfaces_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventDto); i {
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
		file_interfaces_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmitEventDto); i {
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
		file_interfaces_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmitEventResult); i {
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
		file_interfaces_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmitEventDto_Payload); i {
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
	file_interfaces_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_interfaces_event_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_interfaces_event_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EmitEventDto_Payload_Mail)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interfaces_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interfaces_event_proto_goTypes,
		DependencyIndexes: file_interfaces_event_proto_depIdxs,
		EnumInfos:         file_interfaces_event_proto_enumTypes,
		MessageInfos:      file_interfaces_event_proto_msgTypes,
	}.Build()
	File_interfaces_event_proto = out.File
	file_interfaces_event_proto_rawDesc = nil
	file_interfaces_event_proto_goTypes = nil
	file_interfaces_event_proto_depIdxs = nil
}
