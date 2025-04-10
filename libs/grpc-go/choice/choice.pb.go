// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/choice/choice.proto

package choice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Choice struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Information   *string                `protobuf:"bytes,4,opt,name=information,proto3,oneof" json:"information,omitempty"`
	Image         *string                `protobuf:"bytes,5,opt,name=image,proto3,oneof" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Choice) Reset() {
	*x = Choice{}
	mi := &file_proto_choice_choice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Choice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice) ProtoMessage() {}

func (x *Choice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_choice_choice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice.ProtoReflect.Descriptor instead.
func (*Choice) Descriptor() ([]byte, []int) {
	return file_proto_choice_choice_proto_rawDescGZIP(), []int{0}
}

func (x *Choice) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Choice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Choice) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Choice) GetInformation() string {
	if x != nil && x.Information != nil {
		return *x.Information
	}
	return ""
}

func (x *Choice) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

var File_proto_choice_choice_proto protoreflect.FileDescriptor

const file_proto_choice_choice_proto_rawDesc = "" +
	"\n" +
	"\x19proto/choice/choice.proto\x12\x06choice\"\xb2\x01\n" +
	"\x06Choice\x12\x16\n" +
	"\x06number\x18\x01 \x01(\tR\x06number\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12%\n" +
	"\vinformation\x18\x04 \x01(\tH\x00R\vinformation\x88\x01\x01\x12\x19\n" +
	"\x05image\x18\x05 \x01(\tH\x01R\x05image\x88\x01\x01B\x0e\n" +
	"\f_informationB\b\n" +
	"\x06_imageB7Z5github.com/esc-chula/intania-vote/libs/grpc-go/choiceb\x06proto3"

var (
	file_proto_choice_choice_proto_rawDescOnce sync.Once
	file_proto_choice_choice_proto_rawDescData []byte
)

func file_proto_choice_choice_proto_rawDescGZIP() []byte {
	file_proto_choice_choice_proto_rawDescOnce.Do(func() {
		file_proto_choice_choice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_choice_choice_proto_rawDesc), len(file_proto_choice_choice_proto_rawDesc)))
	})
	return file_proto_choice_choice_proto_rawDescData
}

var file_proto_choice_choice_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_choice_choice_proto_goTypes = []any{
	(*Choice)(nil), // 0: choice.Choice
}
var file_proto_choice_choice_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_choice_choice_proto_init() }
func file_proto_choice_choice_proto_init() {
	if File_proto_choice_choice_proto != nil {
		return
	}
	file_proto_choice_choice_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_choice_choice_proto_rawDesc), len(file_proto_choice_choice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_choice_choice_proto_goTypes,
		DependencyIndexes: file_proto_choice_choice_proto_depIdxs,
		MessageInfos:      file_proto_choice_choice_proto_msgTypes,
	}.Build()
	File_proto_choice_choice_proto = out.File
	file_proto_choice_choice_proto_goTypes = nil
	file_proto_choice_choice_proto_depIdxs = nil
}
