// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: interfaces/listener.proto

package pb

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

type IListener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Event       *IEvent `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *IListener) Reset() {
	*x = IListener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IListener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IListener) ProtoMessage() {}

func (x *IListener) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IListener.ProtoReflect.Descriptor instead.
func (*IListener) Descriptor() ([]byte, []int) {
	return file_interfaces_listener_proto_rawDescGZIP(), []int{0}
}

func (x *IListener) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IListener) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IListener) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IListener) GetEvent() *IEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type AddListenerDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	EventName   string `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
}

func (x *AddListenerDto) Reset() {
	*x = AddListenerDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_listener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListenerDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListenerDto) ProtoMessage() {}

func (x *AddListenerDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_listener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListenerDto.ProtoReflect.Descriptor instead.
func (*AddListenerDto) Descriptor() ([]byte, []int) {
	return file_interfaces_listener_proto_rawDescGZIP(), []int{1}
}

func (x *AddListenerDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddListenerDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddListenerDto) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

type GetListenersByEventDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventName string `protobuf:"bytes,1,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
}

func (x *GetListenersByEventDto) Reset() {
	*x = GetListenersByEventDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_listener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListenersByEventDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListenersByEventDto) ProtoMessage() {}

func (x *GetListenersByEventDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_listener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListenersByEventDto.ProtoReflect.Descriptor instead.
func (*GetListenersByEventDto) Descriptor() ([]byte, []int) {
	return file_interfaces_listener_proto_rawDescGZIP(), []int{2}
}

func (x *GetListenersByEventDto) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

type GetListenersByEventResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*IListener `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetListenersByEventResult) Reset() {
	*x = GetListenersByEventResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_listener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListenersByEventResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListenersByEventResult) ProtoMessage() {}

func (x *GetListenersByEventResult) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_listener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListenersByEventResult.ProtoReflect.Descriptor instead.
func (*GetListenersByEventResult) Descriptor() ([]byte, []int) {
	return file_interfaces_listener_proto_rawDescGZIP(), []int{3}
}

func (x *GetListenersByEventResult) GetResult() []*IListener {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_interfaces_listener_proto protoreflect.FileDescriptor

var file_interfaces_listener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a,
	0x09, 0x49, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x99, 0x01, 0x0a, 0x13, 0x49, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x18,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x49, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x4c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x23, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x05, 0x5a, 0x03, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interfaces_listener_proto_rawDescOnce sync.Once
	file_interfaces_listener_proto_rawDescData = file_interfaces_listener_proto_rawDesc
)

func file_interfaces_listener_proto_rawDescGZIP() []byte {
	file_interfaces_listener_proto_rawDescOnce.Do(func() {
		file_interfaces_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_interfaces_listener_proto_rawDescData)
	})
	return file_interfaces_listener_proto_rawDescData
}

var file_interfaces_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_interfaces_listener_proto_goTypes = []interface{}{
	(*IListener)(nil),                 // 0: listener.IListener
	(*AddListenerDto)(nil),            // 1: listener.AddListenerDto
	(*GetListenersByEventDto)(nil),    // 2: listener.GetListenersByEventDto
	(*GetListenersByEventResult)(nil), // 3: listener.GetListenersByEventResult
	(*IEvent)(nil),                    // 4: event.IEvent
}
var file_interfaces_listener_proto_depIdxs = []int32{
	4, // 0: listener.IListener.event:type_name -> event.IEvent
	0, // 1: listener.GetListenersByEventResult.result:type_name -> listener.IListener
	1, // 2: listener.IListenerController.Add:input_type -> listener.AddListenerDto
	2, // 3: listener.IListenerController.Get:input_type -> listener.GetListenersByEventDto
	0, // 4: listener.IListenerController.Add:output_type -> listener.IListener
	3, // 5: listener.IListenerController.Get:output_type -> listener.GetListenersByEventResult
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interfaces_listener_proto_init() }
func file_interfaces_listener_proto_init() {
	if File_interfaces_listener_proto != nil {
		return
	}
	file_interfaces_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interfaces_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IListener); i {
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
		file_interfaces_listener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListenerDto); i {
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
		file_interfaces_listener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListenersByEventDto); i {
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
		file_interfaces_listener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListenersByEventResult); i {
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
			RawDescriptor: file_interfaces_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interfaces_listener_proto_goTypes,
		DependencyIndexes: file_interfaces_listener_proto_depIdxs,
		MessageInfos:      file_interfaces_listener_proto_msgTypes,
	}.Build()
	File_interfaces_listener_proto = out.File
	file_interfaces_listener_proto_rawDesc = nil
	file_interfaces_listener_proto_goTypes = nil
	file_interfaces_listener_proto_depIdxs = nil
}