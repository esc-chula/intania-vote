// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/vote/vote.proto

package vote

import (
	choice "github.com/esc-chula/intania-vote/libs/grpc-go/choice"
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

type Owner int32

const (
	Owner_USER  Owner = 0
	Owner_ESC   Owner = 1
	Owner_ISESC Owner = 2
)

// Enum value maps for Owner.
var (
	Owner_name = map[int32]string{
		0: "USER",
		1: "ESC",
		2: "ISESC",
	}
	Owner_value = map[string]int32{
		"USER":  0,
		"ESC":   1,
		"ISESC": 2,
	}
)

func (x Owner) Enum() *Owner {
	p := new(Owner)
	*p = x
	return p
}

func (x Owner) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Owner) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vote_vote_proto_enumTypes[0].Descriptor()
}

func (Owner) Type() protoreflect.EnumType {
	return &file_proto_vote_vote_proto_enumTypes[0]
}

func (x Owner) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Owner.Descriptor instead.
func (Owner) EnumDescriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{0}
}

type CreateVoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	Vote          *Vote                  `protobuf:"bytes,2,opt,name=vote,proto3" json:"vote,omitempty"`
	Choices       []*choice.Choice       `protobuf:"bytes,3,rep,name=choices,proto3" json:"choices,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVoteRequest) Reset() {
	*x = CreateVoteRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteRequest) ProtoMessage() {}

func (x *CreateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVoteRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *CreateVoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *CreateVoteRequest) GetChoices() []*choice.Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type CreateVoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVoteResponse) Reset() {
	*x = CreateVoteResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteResponse) ProtoMessage() {}

func (x *CreateVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteResponse.ProtoReflect.Descriptor instead.
func (*CreateVoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{1}
}

type GetVoteByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVoteByIdRequest) Reset() {
	*x = GetVoteByIdRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteByIdRequest) ProtoMessage() {}

func (x *GetVoteByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteByIdRequest.ProtoReflect.Descriptor instead.
func (*GetVoteByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{2}
}

func (x *GetVoteByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetVoteByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vote          *Vote                  `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Choices       []*choice.Choice       `protobuf:"bytes,2,rep,name=choices,proto3" json:"choices,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVoteByIdResponse) Reset() {
	*x = GetVoteByIdResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteByIdResponse) ProtoMessage() {}

func (x *GetVoteByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteByIdResponse.ProtoReflect.Descriptor instead.
func (*GetVoteByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{3}
}

func (x *GetVoteByIdResponse) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *GetVoteByIdResponse) GetChoices() []*choice.Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type GetVoteBySlugRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slug          string                 `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVoteBySlugRequest) Reset() {
	*x = GetVoteBySlugRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteBySlugRequest) ProtoMessage() {}

func (x *GetVoteBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetVoteBySlugRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{4}
}

func (x *GetVoteBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetVoteBySlugResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vote          *Vote                  `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Choices       []*choice.Choice       `protobuf:"bytes,2,rep,name=choices,proto3" json:"choices,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVoteBySlugResponse) Reset() {
	*x = GetVoteBySlugResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteBySlugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteBySlugResponse) ProtoMessage() {}

func (x *GetVoteBySlugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteBySlugResponse.ProtoReflect.Descriptor instead.
func (*GetVoteBySlugResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{5}
}

func (x *GetVoteBySlugResponse) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *GetVoteBySlugResponse) GetChoices() []*choice.Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type GetVotesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesRequest) Reset() {
	*x = GetVotesRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesRequest) ProtoMessage() {}

func (x *GetVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesRequest.ProtoReflect.Descriptor instead.
func (*GetVotesRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{6}
}

type GetVotesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Votes         []*Votes               `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesResponse) Reset() {
	*x = GetVotesResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesResponse) ProtoMessage() {}

func (x *GetVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesResponse.ProtoReflect.Descriptor instead.
func (*GetVotesResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{7}
}

func (x *GetVotesResponse) GetVotes() []*Votes {
	if x != nil {
		return x.Votes
	}
	return nil
}

type GetVotesByUserEligibilityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesByUserEligibilityRequest) Reset() {
	*x = GetVotesByUserEligibilityRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesByUserEligibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesByUserEligibilityRequest) ProtoMessage() {}

func (x *GetVotesByUserEligibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesByUserEligibilityRequest.ProtoReflect.Descriptor instead.
func (*GetVotesByUserEligibilityRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{8}
}

func (x *GetVotesByUserEligibilityRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

type GetVotesByUserEligibilityResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Votes         []*Votes               `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesByUserEligibilityResponse) Reset() {
	*x = GetVotesByUserEligibilityResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesByUserEligibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesByUserEligibilityResponse) ProtoMessage() {}

func (x *GetVotesByUserEligibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesByUserEligibilityResponse.ProtoReflect.Descriptor instead.
func (*GetVotesByUserEligibilityResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{9}
}

func (x *GetVotesByUserEligibilityResponse) GetVotes() []*Votes {
	if x != nil {
		return x.Votes
	}
	return nil
}

type HasUserVotedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OidcId        string                 `protobuf:"bytes,1,opt,name=oidcId,proto3" json:"oidcId,omitempty"`
	Slug          string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HasUserVotedRequest) Reset() {
	*x = HasUserVotedRequest{}
	mi := &file_proto_vote_vote_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasUserVotedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasUserVotedRequest) ProtoMessage() {}

func (x *HasUserVotedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasUserVotedRequest.ProtoReflect.Descriptor instead.
func (*HasUserVotedRequest) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{10}
}

func (x *HasUserVotedRequest) GetOidcId() string {
	if x != nil {
		return x.OidcId
	}
	return ""
}

func (x *HasUserVotedRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type HasUserVotedResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasVoted      bool                   `protobuf:"varint,1,opt,name=hasVoted,proto3" json:"hasVoted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HasUserVotedResponse) Reset() {
	*x = HasUserVotedResponse{}
	mi := &file_proto_vote_vote_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasUserVotedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasUserVotedResponse) ProtoMessage() {}

func (x *HasUserVotedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasUserVotedResponse.ProtoReflect.Descriptor instead.
func (*HasUserVotedResponse) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{11}
}

func (x *HasUserVotedResponse) GetHasVoted() bool {
	if x != nil {
		return x.HasVoted
	}
	return false
}

type Vote struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Name               string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description        string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Slug               string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Image              *string                `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Owner              Owner                  `protobuf:"varint,5,opt,name=owner,proto3,enum=vote.Owner" json:"owner,omitempty"`
	EligibleStudentId  string                 `protobuf:"bytes,6,opt,name=eligibleStudentId,proto3" json:"eligibleStudentId,omitempty"`
	EligibleDepartment string                 `protobuf:"bytes,7,opt,name=eligibleDepartment,proto3" json:"eligibleDepartment,omitempty"`
	EligibleYear       string                 `protobuf:"bytes,8,opt,name=eligibleYear,proto3" json:"eligibleYear,omitempty"`
	IsPrivate          bool                   `protobuf:"varint,9,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	IsRealTime         bool                   `protobuf:"varint,10,opt,name=isRealTime,proto3" json:"isRealTime,omitempty"`
	IsAllowMultiple    bool                   `protobuf:"varint,11,opt,name=isAllowMultiple,proto3" json:"isAllowMultiple,omitempty"`
	StartAt            string                 `protobuf:"bytes,12,opt,name=startAt,proto3" json:"startAt,omitempty"`
	EndAt              string                 `protobuf:"bytes,13,opt,name=endAt,proto3" json:"endAt,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Vote) Reset() {
	*x = Vote{}
	mi := &file_proto_vote_vote_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{12}
}

func (x *Vote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vote) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vote) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Vote) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Vote) GetOwner() Owner {
	if x != nil {
		return x.Owner
	}
	return Owner_USER
}

func (x *Vote) GetEligibleStudentId() string {
	if x != nil {
		return x.EligibleStudentId
	}
	return ""
}

func (x *Vote) GetEligibleDepartment() string {
	if x != nil {
		return x.EligibleDepartment
	}
	return ""
}

func (x *Vote) GetEligibleYear() string {
	if x != nil {
		return x.EligibleYear
	}
	return ""
}

func (x *Vote) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Vote) GetIsRealTime() bool {
	if x != nil {
		return x.IsRealTime
	}
	return false
}

func (x *Vote) GetIsAllowMultiple() bool {
	if x != nil {
		return x.IsAllowMultiple
	}
	return false
}

func (x *Vote) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *Vote) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

type Votes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vote          *Vote                  `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Choices       []*choice.Choice       `protobuf:"bytes,2,rep,name=choices,proto3" json:"choices,omitempty"`
	TotalBallots  uint32                 `protobuf:"varint,3,opt,name=totalBallots,proto3" json:"totalBallots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Votes) Reset() {
	*x = Votes{}
	mi := &file_proto_vote_vote_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Votes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Votes) ProtoMessage() {}

func (x *Votes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Votes.ProtoReflect.Descriptor instead.
func (*Votes) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{13}
}

func (x *Votes) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

func (x *Votes) GetChoices() []*choice.Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *Votes) GetTotalBallots() uint32 {
	if x != nil {
		return x.TotalBallots
	}
	return 0
}

type Tally struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Choices       []*TallyChoices        `protobuf:"bytes,1,rep,name=choices,proto3" json:"choices,omitempty"`
	Total         uint32                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tally) Reset() {
	*x = Tally{}
	mi := &file_proto_vote_vote_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tally) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tally) ProtoMessage() {}

func (x *Tally) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tally.ProtoReflect.Descriptor instead.
func (*Tally) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{14}
}

func (x *Tally) GetChoices() []*TallyChoices {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *Tally) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type TallyChoices struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        string                 `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Count         uint32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TallyChoices) Reset() {
	*x = TallyChoices{}
	mi := &file_proto_vote_vote_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TallyChoices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TallyChoices) ProtoMessage() {}

func (x *TallyChoices) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vote_vote_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TallyChoices.ProtoReflect.Descriptor instead.
func (*TallyChoices) Descriptor() ([]byte, []int) {
	return file_proto_vote_vote_proto_rawDescGZIP(), []int{15}
}

func (x *TallyChoices) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *TallyChoices) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_vote_vote_proto protoreflect.FileDescriptor

const file_proto_vote_vote_proto_rawDesc = "" +
	"\n" +
	"\x15proto/vote/vote.proto\x12\x04vote\x1a\x19proto/choice/choice.proto\"u\n" +
	"\x11CreateVoteRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x1e\n" +
	"\x04vote\x18\x02 \x01(\v2\n" +
	".vote.VoteR\x04vote\x12(\n" +
	"\achoices\x18\x03 \x03(\v2\x0e.choice.ChoiceR\achoices\"\x14\n" +
	"\x12CreateVoteResponse\"$\n" +
	"\x12GetVoteByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"_\n" +
	"\x13GetVoteByIdResponse\x12\x1e\n" +
	"\x04vote\x18\x01 \x01(\v2\n" +
	".vote.VoteR\x04vote\x12(\n" +
	"\achoices\x18\x02 \x03(\v2\x0e.choice.ChoiceR\achoices\"*\n" +
	"\x14GetVoteBySlugRequest\x12\x12\n" +
	"\x04slug\x18\x01 \x01(\tR\x04slug\"a\n" +
	"\x15GetVoteBySlugResponse\x12\x1e\n" +
	"\x04vote\x18\x01 \x01(\v2\n" +
	".vote.VoteR\x04vote\x12(\n" +
	"\achoices\x18\x02 \x03(\v2\x0e.choice.ChoiceR\achoices\"\x11\n" +
	"\x0fGetVotesRequest\"5\n" +
	"\x10GetVotesResponse\x12!\n" +
	"\x05votes\x18\x01 \x03(\v2\v.vote.VotesR\x05votes\":\n" +
	" GetVotesByUserEligibilityRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\"F\n" +
	"!GetVotesByUserEligibilityResponse\x12!\n" +
	"\x05votes\x18\x01 \x03(\v2\v.vote.VotesR\x05votes\"A\n" +
	"\x13HasUserVotedRequest\x12\x16\n" +
	"\x06oidcId\x18\x01 \x01(\tR\x06oidcId\x12\x12\n" +
	"\x04slug\x18\x02 \x01(\tR\x04slug\"2\n" +
	"\x14HasUserVotedResponse\x12\x1a\n" +
	"\bhasVoted\x18\x01 \x01(\bR\bhasVoted\"\xb2\x03\n" +
	"\x04Vote\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\x12\x19\n" +
	"\x05image\x18\x04 \x01(\tH\x00R\x05image\x88\x01\x01\x12!\n" +
	"\x05owner\x18\x05 \x01(\x0e2\v.vote.OwnerR\x05owner\x12,\n" +
	"\x11eligibleStudentId\x18\x06 \x01(\tR\x11eligibleStudentId\x12.\n" +
	"\x12eligibleDepartment\x18\a \x01(\tR\x12eligibleDepartment\x12\"\n" +
	"\feligibleYear\x18\b \x01(\tR\feligibleYear\x12\x1c\n" +
	"\tisPrivate\x18\t \x01(\bR\tisPrivate\x12\x1e\n" +
	"\n" +
	"isRealTime\x18\n" +
	" \x01(\bR\n" +
	"isRealTime\x12(\n" +
	"\x0fisAllowMultiple\x18\v \x01(\bR\x0fisAllowMultiple\x12\x18\n" +
	"\astartAt\x18\f \x01(\tR\astartAt\x12\x14\n" +
	"\x05endAt\x18\r \x01(\tR\x05endAtB\b\n" +
	"\x06_image\"u\n" +
	"\x05Votes\x12\x1e\n" +
	"\x04vote\x18\x01 \x01(\v2\n" +
	".vote.VoteR\x04vote\x12(\n" +
	"\achoices\x18\x02 \x03(\v2\x0e.choice.ChoiceR\achoices\x12\"\n" +
	"\ftotalBallots\x18\x03 \x01(\rR\ftotalBallots\"K\n" +
	"\x05Tally\x12,\n" +
	"\achoices\x18\x01 \x03(\v2\x12.vote.TallyChoicesR\achoices\x12\x14\n" +
	"\x05total\x18\x02 \x01(\rR\x05total\"<\n" +
	"\fTallyChoices\x12\x16\n" +
	"\x06number\x18\x01 \x01(\tR\x06number\x12\x14\n" +
	"\x05count\x18\x02 \x01(\rR\x05count*%\n" +
	"\x05Owner\x12\b\n" +
	"\x04USER\x10\x00\x12\a\n" +
	"\x03ESC\x10\x01\x12\t\n" +
	"\x05ISESC\x10\x022\xd8\x03\n" +
	"\vVoteService\x12A\n" +
	"\n" +
	"CreateVote\x12\x17.vote.CreateVoteRequest\x1a\x18.vote.CreateVoteResponse\"\x00\x12D\n" +
	"\vGetVoteById\x12\x18.vote.GetVoteByIdRequest\x1a\x19.vote.GetVoteByIdResponse\"\x00\x12J\n" +
	"\rGetVoteBySlug\x12\x1a.vote.GetVoteBySlugRequest\x1a\x1b.vote.GetVoteBySlugResponse\"\x00\x12;\n" +
	"\bGetVotes\x12\x15.vote.GetVotesRequest\x1a\x16.vote.GetVotesResponse\"\x00\x12n\n" +
	"\x19GetVotesByUserEligibility\x12&.vote.GetVotesByUserEligibilityRequest\x1a'.vote.GetVotesByUserEligibilityResponse\"\x00\x12G\n" +
	"\fHasUserVoted\x12\x19.vote.HasUserVotedRequest\x1a\x1a.vote.HasUserVotedResponse\"\x00B5Z3github.com/esc-chula/intania-vote/libs/grpc-go/voteb\x06proto3"

var (
	file_proto_vote_vote_proto_rawDescOnce sync.Once
	file_proto_vote_vote_proto_rawDescData []byte
)

func file_proto_vote_vote_proto_rawDescGZIP() []byte {
	file_proto_vote_vote_proto_rawDescOnce.Do(func() {
		file_proto_vote_vote_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_vote_vote_proto_rawDesc), len(file_proto_vote_vote_proto_rawDesc)))
	})
	return file_proto_vote_vote_proto_rawDescData
}

var file_proto_vote_vote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_vote_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_vote_vote_proto_goTypes = []any{
	(Owner)(0),                                // 0: vote.Owner
	(*CreateVoteRequest)(nil),                 // 1: vote.CreateVoteRequest
	(*CreateVoteResponse)(nil),                // 2: vote.CreateVoteResponse
	(*GetVoteByIdRequest)(nil),                // 3: vote.GetVoteByIdRequest
	(*GetVoteByIdResponse)(nil),               // 4: vote.GetVoteByIdResponse
	(*GetVoteBySlugRequest)(nil),              // 5: vote.GetVoteBySlugRequest
	(*GetVoteBySlugResponse)(nil),             // 6: vote.GetVoteBySlugResponse
	(*GetVotesRequest)(nil),                   // 7: vote.GetVotesRequest
	(*GetVotesResponse)(nil),                  // 8: vote.GetVotesResponse
	(*GetVotesByUserEligibilityRequest)(nil),  // 9: vote.GetVotesByUserEligibilityRequest
	(*GetVotesByUserEligibilityResponse)(nil), // 10: vote.GetVotesByUserEligibilityResponse
	(*HasUserVotedRequest)(nil),               // 11: vote.HasUserVotedRequest
	(*HasUserVotedResponse)(nil),              // 12: vote.HasUserVotedResponse
	(*Vote)(nil),                              // 13: vote.Vote
	(*Votes)(nil),                             // 14: vote.Votes
	(*Tally)(nil),                             // 15: vote.Tally
	(*TallyChoices)(nil),                      // 16: vote.TallyChoices
	(*choice.Choice)(nil),                     // 17: choice.Choice
}
var file_proto_vote_vote_proto_depIdxs = []int32{
	13, // 0: vote.CreateVoteRequest.vote:type_name -> vote.Vote
	17, // 1: vote.CreateVoteRequest.choices:type_name -> choice.Choice
	13, // 2: vote.GetVoteByIdResponse.vote:type_name -> vote.Vote
	17, // 3: vote.GetVoteByIdResponse.choices:type_name -> choice.Choice
	13, // 4: vote.GetVoteBySlugResponse.vote:type_name -> vote.Vote
	17, // 5: vote.GetVoteBySlugResponse.choices:type_name -> choice.Choice
	14, // 6: vote.GetVotesResponse.votes:type_name -> vote.Votes
	14, // 7: vote.GetVotesByUserEligibilityResponse.votes:type_name -> vote.Votes
	0,  // 8: vote.Vote.owner:type_name -> vote.Owner
	13, // 9: vote.Votes.vote:type_name -> vote.Vote
	17, // 10: vote.Votes.choices:type_name -> choice.Choice
	16, // 11: vote.Tally.choices:type_name -> vote.TallyChoices
	1,  // 12: vote.VoteService.CreateVote:input_type -> vote.CreateVoteRequest
	3,  // 13: vote.VoteService.GetVoteById:input_type -> vote.GetVoteByIdRequest
	5,  // 14: vote.VoteService.GetVoteBySlug:input_type -> vote.GetVoteBySlugRequest
	7,  // 15: vote.VoteService.GetVotes:input_type -> vote.GetVotesRequest
	9,  // 16: vote.VoteService.GetVotesByUserEligibility:input_type -> vote.GetVotesByUserEligibilityRequest
	11, // 17: vote.VoteService.HasUserVoted:input_type -> vote.HasUserVotedRequest
	2,  // 18: vote.VoteService.CreateVote:output_type -> vote.CreateVoteResponse
	4,  // 19: vote.VoteService.GetVoteById:output_type -> vote.GetVoteByIdResponse
	6,  // 20: vote.VoteService.GetVoteBySlug:output_type -> vote.GetVoteBySlugResponse
	8,  // 21: vote.VoteService.GetVotes:output_type -> vote.GetVotesResponse
	10, // 22: vote.VoteService.GetVotesByUserEligibility:output_type -> vote.GetVotesByUserEligibilityResponse
	12, // 23: vote.VoteService.HasUserVoted:output_type -> vote.HasUserVotedResponse
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_vote_vote_proto_init() }
func file_proto_vote_vote_proto_init() {
	if File_proto_vote_vote_proto != nil {
		return
	}
	file_proto_vote_vote_proto_msgTypes[12].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_vote_vote_proto_rawDesc), len(file_proto_vote_vote_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vote_vote_proto_goTypes,
		DependencyIndexes: file_proto_vote_vote_proto_depIdxs,
		EnumInfos:         file_proto_vote_vote_proto_enumTypes,
		MessageInfos:      file_proto_vote_vote_proto_msgTypes,
	}.Build()
	File_proto_vote_vote_proto = out.File
	file_proto_vote_vote_proto_goTypes = nil
	file_proto_vote_vote_proto_depIdxs = nil
}
