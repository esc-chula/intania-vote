// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/ballot/ballot.proto

package ballot

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

type CreateBallotProofRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	VoteSlug      string                 `protobuf:"bytes,2,opt,name=voteSlug,proto3" json:"voteSlug,omitempty"`
	ChoiceNumber  int32                  `protobuf:"varint,3,opt,name=choiceNumber,proto3" json:"choiceNumber,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBallotProofRequest) Reset() {
	*x = CreateBallotProofRequest{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBallotProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBallotProofRequest) ProtoMessage() {}

func (x *CreateBallotProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBallotProofRequest.ProtoReflect.Descriptor instead.
func (*CreateBallotProofRequest) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBallotProofRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *CreateBallotProofRequest) GetVoteSlug() string {
	if x != nil {
		return x.VoteSlug
	}
	return ""
}

func (x *CreateBallotProofRequest) GetChoiceNumber() int32 {
	if x != nil {
		return x.ChoiceNumber
	}
	return 0
}

type CreateBallotProofResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Proof         *Proof                 `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBallotProofResponse) Reset() {
	*x = CreateBallotProofResponse{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBallotProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBallotProofResponse) ProtoMessage() {}

func (x *CreateBallotProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBallotProofResponse.ProtoReflect.Descriptor instead.
func (*CreateBallotProofResponse) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBallotProofResponse) GetProof() *Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

type CreateBallotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	VoteSlug      string                 `protobuf:"bytes,2,opt,name=voteSlug,proto3" json:"voteSlug,omitempty"`
	Proof         *Proof                 `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBallotRequest) Reset() {
	*x = CreateBallotRequest{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBallotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBallotRequest) ProtoMessage() {}

func (x *CreateBallotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBallotRequest.ProtoReflect.Descriptor instead.
func (*CreateBallotRequest) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBallotRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *CreateBallotRequest) GetVoteSlug() string {
	if x != nil {
		return x.VoteSlug
	}
	return ""
}

func (x *CreateBallotRequest) GetProof() *Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

type CreateBallotResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BallotKey     string                 `protobuf:"bytes,1,opt,name=ballotKey,proto3" json:"ballotKey,omitempty"` // voteId:receipt:encryptedChoiceId
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBallotResponse) Reset() {
	*x = CreateBallotResponse{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBallotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBallotResponse) ProtoMessage() {}

func (x *CreateBallotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBallotResponse.ProtoReflect.Descriptor instead.
func (*CreateBallotResponse) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBallotResponse) GetBallotKey() string {
	if x != nil {
		return x.BallotKey
	}
	return ""
}

type VerifyBallotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	BallotKey     string                 `protobuf:"bytes,2,opt,name=ballotKey,proto3" json:"ballotKey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyBallotRequest) Reset() {
	*x = VerifyBallotRequest{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyBallotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBallotRequest) ProtoMessage() {}

func (x *VerifyBallotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBallotRequest.ProtoReflect.Descriptor instead.
func (*VerifyBallotRequest) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyBallotRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *VerifyBallotRequest) GetBallotKey() string {
	if x != nil {
		return x.BallotKey
	}
	return ""
}

type VerifyBallotResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsValid       bool                   `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	ChoiceNumber  int32                  `protobuf:"varint,2,opt,name=choiceNumber,proto3" json:"choiceNumber,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyBallotResponse) Reset() {
	*x = VerifyBallotResponse{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyBallotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyBallotResponse) ProtoMessage() {}

func (x *VerifyBallotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyBallotResponse.ProtoReflect.Descriptor instead.
func (*VerifyBallotResponse) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyBallotResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *VerifyBallotResponse) GetChoiceNumber() int32 {
	if x != nil {
		return x.ChoiceNumber
	}
	return 0
}

func (x *VerifyBallotResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type Proof struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Commitment     string                 `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	BlindingFactor string                 `protobuf:"bytes,2,opt,name=blindingFactor,proto3" json:"blindingFactor,omitempty"`
	Challenge      string                 `protobuf:"bytes,3,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Response       string                 `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	Nullifier      string                 `protobuf:"bytes,5,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	Receipt        string                 `protobuf:"bytes,6,opt,name=receipt,proto3" json:"receipt,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Proof) Reset() {
	*x = Proof{}
	mi := &file_proto_ballot_ballot_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ballot_ballot_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_proto_ballot_ballot_proto_rawDescGZIP(), []int{6}
}

func (x *Proof) GetCommitment() string {
	if x != nil {
		return x.Commitment
	}
	return ""
}

func (x *Proof) GetBlindingFactor() string {
	if x != nil {
		return x.BlindingFactor
	}
	return ""
}

func (x *Proof) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *Proof) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *Proof) GetNullifier() string {
	if x != nil {
		return x.Nullifier
	}
	return ""
}

func (x *Proof) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

var File_proto_ballot_ballot_proto protoreflect.FileDescriptor

const file_proto_ballot_ballot_proto_rawDesc = "" +
	"\n" +
	"\x19proto/ballot/ballot.proto\x12\x06ballot\"r\n" +
	"\x18CreateBallotProofRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x1a\n" +
	"\bvoteSlug\x18\x02 \x01(\tR\bvoteSlug\x12\"\n" +
	"\fchoiceNumber\x18\x03 \x01(\x05R\fchoiceNumber\"@\n" +
	"\x19CreateBallotProofResponse\x12#\n" +
	"\x05proof\x18\x01 \x01(\v2\r.ballot.ProofR\x05proof\"n\n" +
	"\x13CreateBallotRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x1a\n" +
	"\bvoteSlug\x18\x02 \x01(\tR\bvoteSlug\x12#\n" +
	"\x05proof\x18\x03 \x01(\v2\r.ballot.ProofR\x05proof\"4\n" +
	"\x14CreateBallotResponse\x12\x1c\n" +
	"\tballotKey\x18\x01 \x01(\tR\tballotKey\"K\n" +
	"\x13VerifyBallotRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x1c\n" +
	"\tballotKey\x18\x02 \x01(\tR\tballotKey\"r\n" +
	"\x14VerifyBallotResponse\x12\x18\n" +
	"\aisValid\x18\x01 \x01(\bR\aisValid\x12\"\n" +
	"\fchoiceNumber\x18\x02 \x01(\x05R\fchoiceNumber\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\tR\ttimestamp\"\xc1\x01\n" +
	"\x05Proof\x12\x1e\n" +
	"\n" +
	"commitment\x18\x01 \x01(\tR\n" +
	"commitment\x12&\n" +
	"\x0eblindingFactor\x18\x02 \x01(\tR\x0eblindingFactor\x12\x1c\n" +
	"\tchallenge\x18\x03 \x01(\tR\tchallenge\x12\x1a\n" +
	"\bresponse\x18\x04 \x01(\tR\bresponse\x12\x1c\n" +
	"\tnullifier\x18\x05 \x01(\tR\tnullifier\x12\x18\n" +
	"\areceipt\x18\x06 \x01(\tR\areceipt2\x85\x02\n" +
	"\rBallotService\x12Z\n" +
	"\x11CreateBallotProof\x12 .ballot.CreateBallotProofRequest\x1a!.ballot.CreateBallotProofResponse\"\x00\x12K\n" +
	"\fCreateBallot\x12\x1b.ballot.CreateBallotRequest\x1a\x1c.ballot.CreateBallotResponse\"\x00\x12K\n" +
	"\fVerifyBallot\x12\x1b.ballot.VerifyBallotRequest\x1a\x1c.ballot.VerifyBallotResponse\"\x00B7Z5github.com/esc-chula/intania-vote/libs/grpc-go/ballotb\x06proto3"

var (
	file_proto_ballot_ballot_proto_rawDescOnce sync.Once
	file_proto_ballot_ballot_proto_rawDescData []byte
)

func file_proto_ballot_ballot_proto_rawDescGZIP() []byte {
	file_proto_ballot_ballot_proto_rawDescOnce.Do(func() {
		file_proto_ballot_ballot_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_ballot_ballot_proto_rawDesc), len(file_proto_ballot_ballot_proto_rawDesc)))
	})
	return file_proto_ballot_ballot_proto_rawDescData
}

var file_proto_ballot_ballot_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_ballot_ballot_proto_goTypes = []any{
	(*CreateBallotProofRequest)(nil),  // 0: ballot.CreateBallotProofRequest
	(*CreateBallotProofResponse)(nil), // 1: ballot.CreateBallotProofResponse
	(*CreateBallotRequest)(nil),       // 2: ballot.CreateBallotRequest
	(*CreateBallotResponse)(nil),      // 3: ballot.CreateBallotResponse
	(*VerifyBallotRequest)(nil),       // 4: ballot.VerifyBallotRequest
	(*VerifyBallotResponse)(nil),      // 5: ballot.VerifyBallotResponse
	(*Proof)(nil),                     // 6: ballot.Proof
}
var file_proto_ballot_ballot_proto_depIdxs = []int32{
	6, // 0: ballot.CreateBallotProofResponse.proof:type_name -> ballot.Proof
	6, // 1: ballot.CreateBallotRequest.proof:type_name -> ballot.Proof
	0, // 2: ballot.BallotService.CreateBallotProof:input_type -> ballot.CreateBallotProofRequest
	2, // 3: ballot.BallotService.CreateBallot:input_type -> ballot.CreateBallotRequest
	4, // 4: ballot.BallotService.VerifyBallot:input_type -> ballot.VerifyBallotRequest
	1, // 5: ballot.BallotService.CreateBallotProof:output_type -> ballot.CreateBallotProofResponse
	3, // 6: ballot.BallotService.CreateBallot:output_type -> ballot.CreateBallotResponse
	5, // 7: ballot.BallotService.VerifyBallot:output_type -> ballot.VerifyBallotResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_ballot_ballot_proto_init() }
func file_proto_ballot_ballot_proto_init() {
	if File_proto_ballot_ballot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_ballot_ballot_proto_rawDesc), len(file_proto_ballot_ballot_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ballot_ballot_proto_goTypes,
		DependencyIndexes: file_proto_ballot_ballot_proto_depIdxs,
		MessageInfos:      file_proto_ballot_ballot_proto_msgTypes,
	}.Build()
	File_proto_ballot_ballot_proto = out.File
	file_proto_ballot_ballot_proto_goTypes = nil
	file_proto_ballot_ballot_proto_depIdxs = nil
}
