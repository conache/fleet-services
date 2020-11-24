// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.0
// source: proto/test-run-service/test-run-service.proto

package go_micro_service_testrunservice

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateTestRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRun  *TestRun  `protobuf:"bytes,1,opt,name=testRun,proto3" json:"testRun,omitempty"`
	FileSpec *FileSpec `protobuf:"bytes,2,opt,name=fileSpec,proto3" json:"fileSpec,omitempty"`
}

func (x *CreateTestRunRequest) Reset() {
	*x = CreateTestRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestRunRequest) ProtoMessage() {}

func (x *CreateTestRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestRunRequest.ProtoReflect.Descriptor instead.
func (*CreateTestRunRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestRunRequest) GetTestRun() *TestRun {
	if x != nil {
		return x.TestRun
	}
	return nil
}

func (x *CreateTestRunRequest) GetFileSpec() *FileSpec {
	if x != nil {
		return x.FileSpec
	}
	return nil
}

type FileSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size         int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	MaxChunkSize uint32 `protobuf:"varint,4,opt,name=maxChunkSize,proto3" json:"maxChunkSize,omitempty"`
}

func (x *FileSpec) Reset() {
	*x = FileSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSpec) ProtoMessage() {}

func (x *FileSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSpec.ProtoReflect.Descriptor instead.
func (*FileSpec) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileSpec) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileSpec) GetMaxChunkSize() uint32 {
	if x != nil {
		return x.MaxChunkSize
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TestRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FileId string `protobuf:"bytes,3,opt,name=fileId,proto3" json:"fileId,omitempty"`
	UserId uint32 `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	User   *User  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *TestRun) Reset() {
	*x = TestRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRun) ProtoMessage() {}

func (x *TestRun) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRun.ProtoReflect.Descriptor instead.
func (*TestRun) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{3}
}

func (x *TestRun) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestRun) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestRun) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *TestRun) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TestRun) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type TestRunDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRun *TestRun `protobuf:"bytes,1,opt,name=testRun,proto3" json:"testRun,omitempty"`
}

func (x *TestRunDetails) Reset() {
	*x = TestRunDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunDetails) ProtoMessage() {}

func (x *TestRunDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunDetails.ProtoReflect.Descriptor instead.
func (*TestRunDetails) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{4}
}

func (x *TestRunDetails) GetTestRun() *TestRun {
	if x != nil {
		return x.TestRun
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRuns []*TestRun `protobuf:"bytes,1,rep,name=testRuns,proto3" json:"testRuns,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetTestRuns() []*TestRun {
	if x != nil {
		return x.TestRuns
	}
	return nil
}

type AssignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	FileId    string `protobuf:"bytes,2,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *AssignRequest) Reset() {
	*x = AssignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRequest) ProtoMessage() {}

func (x *AssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRequest.ProtoReflect.Descriptor instead.
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{6}
}

func (x *AssignRequest) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *AssignRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{7}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{8}
}

var File_proto_test_run_service_test_run_service_proto protoreflect.FileDescriptor

var file_proto_test_run_service_test_run_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x72, 0x75, 0x6e,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x72, 0x75,
	0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0xa1, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x45, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x66, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x61, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5a, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x98, 0x01, 0x0a, 0x07, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x22, 0x54, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x22,
	0x45, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf9, 0x05, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x62,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x1a,
	0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x2f, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12,
	0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x66, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_test_run_service_test_run_service_proto_rawDescOnce sync.Once
	file_proto_test_run_service_test_run_service_proto_rawDescData = file_proto_test_run_service_test_run_service_proto_rawDesc
)

func file_proto_test_run_service_test_run_service_proto_rawDescGZIP() []byte {
	file_proto_test_run_service_test_run_service_proto_rawDescOnce.Do(func() {
		file_proto_test_run_service_test_run_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_run_service_test_run_service_proto_rawDescData)
	})
	return file_proto_test_run_service_test_run_service_proto_rawDescData
}

var file_proto_test_run_service_test_run_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_test_run_service_test_run_service_proto_goTypes = []interface{}{
	(*CreateTestRunRequest)(nil), // 0: go.micro.service.testrunservice.CreateTestRunRequest
	(*FileSpec)(nil),             // 1: go.micro.service.testrunservice.FileSpec
	(*User)(nil),                 // 2: go.micro.service.testrunservice.User
	(*TestRun)(nil),              // 3: go.micro.service.testrunservice.TestRun
	(*TestRunDetails)(nil),       // 4: go.micro.service.testrunservice.TestRunDetails
	(*ListResponse)(nil),         // 5: go.micro.service.testrunservice.ListResponse
	(*AssignRequest)(nil),        // 6: go.micro.service.testrunservice.AssignRequest
	(*EmptyRequest)(nil),         // 7: go.micro.service.testrunservice.EmptyRequest
	(*EmptyResponse)(nil),        // 8: go.micro.service.testrunservice.EmptyResponse
}
var file_proto_test_run_service_test_run_service_proto_depIdxs = []int32{
	3,  // 0: go.micro.service.testrunservice.CreateTestRunRequest.testRun:type_name -> go.micro.service.testrunservice.TestRun
	1,  // 1: go.micro.service.testrunservice.CreateTestRunRequest.fileSpec:type_name -> go.micro.service.testrunservice.FileSpec
	2,  // 2: go.micro.service.testrunservice.TestRun.user:type_name -> go.micro.service.testrunservice.User
	3,  // 3: go.micro.service.testrunservice.TestRunDetails.testRun:type_name -> go.micro.service.testrunservice.TestRun
	3,  // 4: go.micro.service.testrunservice.ListResponse.testRuns:type_name -> go.micro.service.testrunservice.TestRun
	0,  // 5: go.micro.service.testrunservice.TestRunService.Create:input_type -> go.micro.service.testrunservice.CreateTestRunRequest
	3,  // 6: go.micro.service.testrunservice.TestRunService.Get:input_type -> go.micro.service.testrunservice.TestRun
	1,  // 7: go.micro.service.testrunservice.TestRunService.GetByFileId:input_type -> go.micro.service.testrunservice.FileSpec
	3,  // 8: go.micro.service.testrunservice.TestRunService.GetById:input_type -> go.micro.service.testrunservice.TestRun
	7,  // 9: go.micro.service.testrunservice.TestRunService.List:input_type -> go.micro.service.testrunservice.EmptyRequest
	3,  // 10: go.micro.service.testrunservice.TestRunService.Delete:input_type -> go.micro.service.testrunservice.TestRun
	6,  // 11: go.micro.service.testrunservice.TestRunService.AssignFile:input_type -> go.micro.service.testrunservice.AssignRequest
	4,  // 12: go.micro.service.testrunservice.TestRunService.Create:output_type -> go.micro.service.testrunservice.TestRunDetails
	4,  // 13: go.micro.service.testrunservice.TestRunService.Get:output_type -> go.micro.service.testrunservice.TestRunDetails
	4,  // 14: go.micro.service.testrunservice.TestRunService.GetByFileId:output_type -> go.micro.service.testrunservice.TestRunDetails
	4,  // 15: go.micro.service.testrunservice.TestRunService.GetById:output_type -> go.micro.service.testrunservice.TestRunDetails
	5,  // 16: go.micro.service.testrunservice.TestRunService.List:output_type -> go.micro.service.testrunservice.ListResponse
	8,  // 17: go.micro.service.testrunservice.TestRunService.Delete:output_type -> go.micro.service.testrunservice.EmptyResponse
	8,  // 18: go.micro.service.testrunservice.TestRunService.AssignFile:output_type -> go.micro.service.testrunservice.EmptyResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_test_run_service_test_run_service_proto_init() }
func file_proto_test_run_service_test_run_service_proto_init() {
	if File_proto_test_run_service_test_run_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_run_service_test_run_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestRunRequest); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSpec); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRun); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunDetails); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignRequest); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_proto_test_run_service_test_run_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_proto_test_run_service_test_run_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_test_run_service_test_run_service_proto_goTypes,
		DependencyIndexes: file_proto_test_run_service_test_run_service_proto_depIdxs,
		MessageInfos:      file_proto_test_run_service_test_run_service_proto_msgTypes,
	}.Build()
	File_proto_test_run_service_test_run_service_proto = out.File
	file_proto_test_run_service_test_run_service_proto_rawDesc = nil
	file_proto_test_run_service_test_run_service_proto_goTypes = nil
	file_proto_test_run_service_test_run_service_proto_depIdxs = nil
}
