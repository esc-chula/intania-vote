// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: user/user.proto

package user

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

type GetUserByOidcIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByOidcIdRequest) Reset() {
	*x = GetUserByOidcIdRequest{}
	mi := &file_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByOidcIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOidcIdRequest) ProtoMessage() {}

func (x *GetUserByOidcIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByOidcIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByOidcIdRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByOidcIdRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

type GetUserByOidcIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OidcId        string                 `protobuf:"bytes,2,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	StudentId     string                 `protobuf:"bytes,3,opt,name=studentId,proto3" json:"studentId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByOidcIdResponse) Reset() {
	*x = GetUserByOidcIdResponse{}
	mi := &file_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByOidcIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOidcIdResponse) ProtoMessage() {}

func (x *GetUserByOidcIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByOidcIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserByOidcIdResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByOidcIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserByOidcIdResponse) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *GetUserByOidcIdResponse) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

var File_user_user_proto protoreflect.FileDescriptor

const file_user_user_proto_rawDesc = "" +
	"\n" +
	"\x0fuser/user.proto\x12\x04user\"0\n" +
	"\x16GetUserByOidcIdRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\"_\n" +
	"\x17GetUserByOidcIdResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06oidcId\x18\x02 \x01(\tR\x06oidcId\x12\x1c\n" +
	"\tstudentId\x18\x03 \x01(\tR\tstudentId2_\n" +
	"\vUserService\x12P\n" +
	"\x0fGetUserByOidcId\x12\x1c.user.GetUserByOidcIdRequest\x1a\x1d.user.GetUserByOidcIdResponse\"\x00B5Z3github.com/esc-chula/intania-vote/libs/grpc-go/userb\x06proto3"

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData []byte
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)))
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_user_proto_goTypes = []any{
	(*GetUserByOidcIdRequest)(nil),  // 0: user.GetUserByOidcIdRequest
	(*GetUserByOidcIdResponse)(nil), // 1: user.GetUserByOidcIdResponse
}
var file_user_user_proto_depIdxs = []int32{
	0, // 0: user.UserService.GetUserByOidcId:input_type -> user.GetUserByOidcIdRequest
	1, // 1: user.UserService.GetUserByOidcId:output_type -> user.GetUserByOidcIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_user_proto_rawDesc), len(file_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
