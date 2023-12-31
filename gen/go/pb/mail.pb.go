// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: interfaces/mail.proto

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

type IMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`           // Email to send
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"` // Subject of mail
	Html    string `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`       // Html of mail
}

func (x *IMail) Reset() {
	*x = IMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMail) ProtoMessage() {}

func (x *IMail) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMail.ProtoReflect.Descriptor instead.
func (*IMail) Descriptor() ([]byte, []int) {
	return file_interfaces_mail_proto_rawDescGZIP(), []int{0}
}

func (x *IMail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IMail) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *IMail) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *IMail) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type SendMailDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To       string            `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject  string            `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Template *string           `protobuf:"bytes,4,opt,name=template,proto3,oneof" json:"template,omitempty"`
	Html     *string           `protobuf:"bytes,5,opt,name=html,proto3,oneof" json:"html,omitempty"`
	Context  map[string]string `protobuf:"bytes,6,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendMailDto) Reset() {
	*x = SendMailDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailDto) ProtoMessage() {}

func (x *SendMailDto) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailDto.ProtoReflect.Descriptor instead.
func (*SendMailDto) Descriptor() ([]byte, []int) {
	return file_interfaces_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailDto) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendMailDto) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailDto) GetTemplate() string {
	if x != nil && x.Template != nil {
		return *x.Template
	}
	return ""
}

func (x *SendMailDto) GetHtml() string {
	if x != nil && x.Html != nil {
		return *x.Html
	}
	return ""
}

func (x *SendMailDto) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_interfaces_mail_proto protoreflect.FileDescriptor

var file_interfaces_mail_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a,
	0x05, 0x49, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x74, 0x6d, 0x6c, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69,
	0x6c, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x68, 0x74, 0x6d, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x68, 0x74, 0x6d, 0x6c, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_interfaces_mail_proto_rawDescOnce sync.Once
	file_interfaces_mail_proto_rawDescData = file_interfaces_mail_proto_rawDesc
)

func file_interfaces_mail_proto_rawDescGZIP() []byte {
	file_interfaces_mail_proto_rawDescOnce.Do(func() {
		file_interfaces_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_interfaces_mail_proto_rawDescData)
	})
	return file_interfaces_mail_proto_rawDescData
}

var file_interfaces_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interfaces_mail_proto_goTypes = []interface{}{
	(*IMail)(nil),       // 0: mail.IMail
	(*SendMailDto)(nil), // 1: mail.SendMailDto
	nil,                 // 2: mail.SendMailDto.ContextEntry
}
var file_interfaces_mail_proto_depIdxs = []int32{
	2, // 0: mail.SendMailDto.context:type_name -> mail.SendMailDto.ContextEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interfaces_mail_proto_init() }
func file_interfaces_mail_proto_init() {
	if File_interfaces_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interfaces_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMail); i {
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
		file_interfaces_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailDto); i {
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
	file_interfaces_mail_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interfaces_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interfaces_mail_proto_goTypes,
		DependencyIndexes: file_interfaces_mail_proto_depIdxs,
		MessageInfos:      file_interfaces_mail_proto_msgTypes,
	}.Build()
	File_interfaces_mail_proto = out.File
	file_interfaces_mail_proto_rawDesc = nil
	file_interfaces_mail_proto_goTypes = nil
	file_interfaces_mail_proto_depIdxs = nil
}
