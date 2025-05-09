// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/user/user.proto

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

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	StudentId     string                 `protobuf:"bytes,2,opt,name=studentId,proto3" json:"studentId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_proto_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *CreateUserRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_proto_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

type GetUserByOidcIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByOidcIdRequest) Reset() {
	*x = GetUserByOidcIdRequest{}
	mi := &file_proto_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByOidcIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOidcIdRequest) ProtoMessage() {}

func (x *GetUserByOidcIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
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
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
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
	mi := &file_proto_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByOidcIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOidcIdResponse) ProtoMessage() {}

func (x *GetUserByOidcIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[3]
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
	return file_proto_user_user_proto_rawDescGZIP(), []int{3}
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

var File_proto_user_user_proto protoreflect.FileDescriptor

const file_proto_user_user_proto_rawDesc = "" +
	"\n" +
	"\x15proto/user/user.proto\x12\x04user\"I\n" +
	"\x11CreateUserRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x1c\n" +
	"\tstudentId\x18\x02 \x01(\tR\tstudentId\"\x14\n" +
	"\x12CreateUserResponse\"0\n" +
	"\x16GetUserByOidcIdRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\"_\n" +
	"\x17GetUserByOidcIdResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06oidcId\x18\x02 \x01(\tR\x06oidcId\x12\x1c\n" +
	"\tstudentId\x18\x03 \x01(\tR\tstudentId2\xa2\x01\n" +
	"\vUserService\x12A\n" +
	"\n" +
	"CreateUser\x12\x17.user.CreateUserRequest\x1a\x18.user.CreateUserResponse\"\x00\x12P\n" +
	"\x0fGetUserByOidcId\x12\x1c.user.GetUserByOidcIdRequest\x1a\x1d.user.GetUserByOidcIdResponse\"\x00B5Z3github.com/esc-chula/intania-vote/libs/grpc-go/userb\x06proto3"

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData []byte
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)))
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_user_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),       // 0: user.CreateUserRequest
	(*CreateUserResponse)(nil),      // 1: user.CreateUserResponse
	(*GetUserByOidcIdRequest)(nil),  // 2: user.GetUserByOidcIdRequest
	(*GetUserByOidcIdResponse)(nil), // 3: user.GetUserByOidcIdResponse
}
var file_proto_user_user_proto_depIdxs = []int32{
	0, // 0: user.UserService.CreateUser:input_type -> user.CreateUserRequest
	2, // 1: user.UserService.GetUserByOidcId:input_type -> user.GetUserByOidcIdRequest
	1, // 2: user.UserService.CreateUser:output_type -> user.CreateUserResponse
	3, // 3: user.UserService.GetUserByOidcId:output_type -> user.GetUserByOidcIdResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_user_proto_rawDesc), len(file_proto_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
