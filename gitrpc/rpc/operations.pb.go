// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: operations.proto

package rpc

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

type CommitFilesActionHeader_ActionType int32

const (
	// CREATE creates a new file.
	CommitFilesActionHeader_CREATE CommitFilesActionHeader_ActionType = 0
	// UPDATE updates an existing file.
	CommitFilesActionHeader_UPDATE CommitFilesActionHeader_ActionType = 1
	// DELETE deletes an existing file or dir.
	CommitFilesActionHeader_DELETE CommitFilesActionHeader_ActionType = 2
	// MOVE moves existing file to another dir.
	CommitFilesActionHeader_MOVE CommitFilesActionHeader_ActionType = 3
)

// Enum value maps for CommitFilesActionHeader_ActionType.
var (
	CommitFilesActionHeader_ActionType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
		3: "MOVE",
	}
	CommitFilesActionHeader_ActionType_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
		"MOVE":   3,
	}
)

func (x CommitFilesActionHeader_ActionType) Enum() *CommitFilesActionHeader_ActionType {
	p := new(CommitFilesActionHeader_ActionType)
	*p = x
	return p
}

func (x CommitFilesActionHeader_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommitFilesActionHeader_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_operations_proto_enumTypes[0].Descriptor()
}

func (CommitFilesActionHeader_ActionType) Type() protoreflect.EnumType {
	return &file_operations_proto_enumTypes[0]
}

func (x CommitFilesActionHeader_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommitFilesActionHeader_ActionType.Descriptor instead.
func (CommitFilesActionHeader_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{1, 0}
}

// CommitFilesRequestHeader is the header of the UserCommitFiles that defines the commit details,
// parent and other information related to the call.
type CommitFilesRequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base          *WriteRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	BranchName    string        `protobuf:"bytes,2,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	NewBranchName string        `protobuf:"bytes,3,opt,name=new_branch_name,json=newBranchName,proto3" json:"new_branch_name,omitempty"`
	Title         string        `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Message       string        `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Author        *Identity     `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	AuthorDate    int64         `protobuf:"varint,7,opt,name=authorDate,proto3" json:"authorDate,omitempty"`
	Committer     *Identity     `protobuf:"bytes,8,opt,name=committer,proto3" json:"committer,omitempty"`
	CommitterDate int64         `protobuf:"varint,9,opt,name=committerDate,proto3" json:"committerDate,omitempty"`
}

func (x *CommitFilesRequestHeader) Reset() {
	*x = CommitFilesRequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitFilesRequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitFilesRequestHeader) ProtoMessage() {}

func (x *CommitFilesRequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitFilesRequestHeader.ProtoReflect.Descriptor instead.
func (*CommitFilesRequestHeader) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{0}
}

func (x *CommitFilesRequestHeader) GetBase() *WriteRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CommitFilesRequestHeader) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *CommitFilesRequestHeader) GetNewBranchName() string {
	if x != nil {
		return x.NewBranchName
	}
	return ""
}

func (x *CommitFilesRequestHeader) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CommitFilesRequestHeader) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommitFilesRequestHeader) GetAuthor() *Identity {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *CommitFilesRequestHeader) GetAuthorDate() int64 {
	if x != nil {
		return x.AuthorDate
	}
	return 0
}

func (x *CommitFilesRequestHeader) GetCommitter() *Identity {
	if x != nil {
		return x.Committer
	}
	return nil
}

func (x *CommitFilesRequestHeader) GetCommitterDate() int64 {
	if x != nil {
		return x.CommitterDate
	}
	return 0
}

// CommitFilesActionHeader contains the details of the action to be performed.
type CommitFilesActionHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// action is the type of the action taken to build a commit. Not all fields are
	// used for all of the actions.
	Action CommitFilesActionHeader_ActionType `protobuf:"varint,1,opt,name=action,proto3,enum=rpc.CommitFilesActionHeader_ActionType" json:"action,omitempty"`
	// path refers to the file or directory being modified.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Sha  string `protobuf:"bytes,3,opt,name=sha,proto3" json:"sha,omitempty"`
}

func (x *CommitFilesActionHeader) Reset() {
	*x = CommitFilesActionHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitFilesActionHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitFilesActionHeader) ProtoMessage() {}

func (x *CommitFilesActionHeader) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitFilesActionHeader.ProtoReflect.Descriptor instead.
func (*CommitFilesActionHeader) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{1}
}

func (x *CommitFilesActionHeader) GetAction() CommitFilesActionHeader_ActionType {
	if x != nil {
		return x.Action
	}
	return CommitFilesActionHeader_CREATE
}

func (x *CommitFilesActionHeader) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CommitFilesActionHeader) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

// CommitFilesAction is the request message used to stream in the actions to build a commit.
type CommitFilesAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*CommitFilesAction_Header
	//	*CommitFilesAction_Content
	Payload isCommitFilesAction_Payload `protobuf_oneof:"payload"`
}

func (x *CommitFilesAction) Reset() {
	*x = CommitFilesAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitFilesAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitFilesAction) ProtoMessage() {}

func (x *CommitFilesAction) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitFilesAction.ProtoReflect.Descriptor instead.
func (*CommitFilesAction) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{2}
}

func (m *CommitFilesAction) GetPayload() isCommitFilesAction_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CommitFilesAction) GetHeader() *CommitFilesActionHeader {
	if x, ok := x.GetPayload().(*CommitFilesAction_Header); ok {
		return x.Header
	}
	return nil
}

func (x *CommitFilesAction) GetContent() []byte {
	if x, ok := x.GetPayload().(*CommitFilesAction_Content); ok {
		return x.Content
	}
	return nil
}

type isCommitFilesAction_Payload interface {
	isCommitFilesAction_Payload()
}

type CommitFilesAction_Header struct {
	// header contains the details of action being performed. Header must be sent before the
	// file if file is used by the action.
	Header *CommitFilesActionHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type CommitFilesAction_Content struct {
	// not used for DELETE action.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*CommitFilesAction_Header) isCommitFilesAction_Payload() {}

func (*CommitFilesAction_Content) isCommitFilesAction_Payload() {}

type CommitFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*CommitFilesRequest_Header
	//	*CommitFilesRequest_Action
	Payload isCommitFilesRequest_Payload `protobuf_oneof:"payload"`
}

func (x *CommitFilesRequest) Reset() {
	*x = CommitFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitFilesRequest) ProtoMessage() {}

func (x *CommitFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitFilesRequest.ProtoReflect.Descriptor instead.
func (*CommitFilesRequest) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{3}
}

func (m *CommitFilesRequest) GetPayload() isCommitFilesRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CommitFilesRequest) GetHeader() *CommitFilesRequestHeader {
	if x, ok := x.GetPayload().(*CommitFilesRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (x *CommitFilesRequest) GetAction() *CommitFilesAction {
	if x, ok := x.GetPayload().(*CommitFilesRequest_Action); ok {
		return x.Action
	}
	return nil
}

type isCommitFilesRequest_Payload interface {
	isCommitFilesRequest_Payload()
}

type CommitFilesRequest_Header struct {
	// header defines the details of where to commit, the details and which commit to use as the parent.
	// header must always be sent as the first request of the stream.
	Header *CommitFilesRequestHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type CommitFilesRequest_Action struct {
	// action contains an action to build a commit. There can be multiple actions per stream.
	Action *CommitFilesAction `protobuf:"bytes,2,opt,name=action,proto3,oneof"`
}

func (*CommitFilesRequest_Header) isCommitFilesRequest_Payload() {}

func (*CommitFilesRequest_Action) isCommitFilesRequest_Payload() {}

type CommitFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitId string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	Branch   string `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *CommitFilesResponse) Reset() {
	*x = CommitFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitFilesResponse) ProtoMessage() {}

func (x *CommitFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitFilesResponse.ProtoReflect.Descriptor instead.
func (*CommitFilesResponse) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{4}
}

func (x *CommitFilesResponse) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *CommitFilesResponse) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

var File_operations_proto protoreflect.FileDescriptor

var file_operations_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x77, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbc, 0x01, 0x0a,
	0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22,
	0x3a, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x22, 0x72, 0x0a, 0x11, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x8a, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4a, 0x0a, 0x13,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x32, 0x58, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x17, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x67, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_operations_proto_rawDescOnce sync.Once
	file_operations_proto_rawDescData = file_operations_proto_rawDesc
)

func file_operations_proto_rawDescGZIP() []byte {
	file_operations_proto_rawDescOnce.Do(func() {
		file_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_operations_proto_rawDescData)
	})
	return file_operations_proto_rawDescData
}

var file_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operations_proto_goTypes = []interface{}{
	(CommitFilesActionHeader_ActionType)(0), // 0: rpc.CommitFilesActionHeader.ActionType
	(*CommitFilesRequestHeader)(nil),        // 1: rpc.CommitFilesRequestHeader
	(*CommitFilesActionHeader)(nil),         // 2: rpc.CommitFilesActionHeader
	(*CommitFilesAction)(nil),               // 3: rpc.CommitFilesAction
	(*CommitFilesRequest)(nil),              // 4: rpc.CommitFilesRequest
	(*CommitFilesResponse)(nil),             // 5: rpc.CommitFilesResponse
	(*WriteRequest)(nil),                    // 6: rpc.WriteRequest
	(*Identity)(nil),                        // 7: rpc.Identity
}
var file_operations_proto_depIdxs = []int32{
	6, // 0: rpc.CommitFilesRequestHeader.base:type_name -> rpc.WriteRequest
	7, // 1: rpc.CommitFilesRequestHeader.author:type_name -> rpc.Identity
	7, // 2: rpc.CommitFilesRequestHeader.committer:type_name -> rpc.Identity
	0, // 3: rpc.CommitFilesActionHeader.action:type_name -> rpc.CommitFilesActionHeader.ActionType
	2, // 4: rpc.CommitFilesAction.header:type_name -> rpc.CommitFilesActionHeader
	1, // 5: rpc.CommitFilesRequest.header:type_name -> rpc.CommitFilesRequestHeader
	3, // 6: rpc.CommitFilesRequest.action:type_name -> rpc.CommitFilesAction
	4, // 7: rpc.CommitFilesService.CommitFiles:input_type -> rpc.CommitFilesRequest
	5, // 8: rpc.CommitFilesService.CommitFiles:output_type -> rpc.CommitFilesResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_operations_proto_init() }
func file_operations_proto_init() {
	if File_operations_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitFilesRequestHeader); i {
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
		file_operations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitFilesActionHeader); i {
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
		file_operations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitFilesAction); i {
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
		file_operations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitFilesRequest); i {
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
		file_operations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitFilesResponse); i {
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
	file_operations_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CommitFilesAction_Header)(nil),
		(*CommitFilesAction_Content)(nil),
	}
	file_operations_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CommitFilesRequest_Header)(nil),
		(*CommitFilesRequest_Action)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operations_proto_goTypes,
		DependencyIndexes: file_operations_proto_depIdxs,
		EnumInfos:         file_operations_proto_enumTypes,
		MessageInfos:      file_operations_proto_msgTypes,
	}.Build()
	File_operations_proto = out.File
	file_operations_proto_rawDesc = nil
	file_operations_proto_goTypes = nil
	file_operations_proto_depIdxs = nil
}
