// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: proto/file-service/file-service.proto

package file_service

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

type ChunkSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	Index  uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ChunkSpec) Reset() {
	*x = ChunkSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkSpec) ProtoMessage() {}

func (x *ChunkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkSpec.ProtoReflect.Descriptor instead.
func (*ChunkSpec) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkSpec) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ChunkSpec) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ChunkSpec) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha2 string `protobuf:"bytes,1,opt,name=sha2,proto3" json:"sha2,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Chunk) GetSha2() string {
	if x != nil {
		return x.Sha2
	}
	return ""
}

func (x *Chunk) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ChunkDataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId        string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	Sha2          string `protobuf:"bytes,2,opt,name=sha2,proto3" json:"sha2,omitempty"`
	Data          []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Authorization []byte `protobuf:"bytes,4,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *ChunkDataMessage) Reset() {
	*x = ChunkDataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkDataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkDataMessage) ProtoMessage() {}

func (x *ChunkDataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkDataMessage.ProtoReflect.Descriptor instead.
func (*ChunkDataMessage) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkDataMessage) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ChunkDataMessage) GetSha2() string {
	if x != nil {
		return x.Sha2
	}
	return ""
}

func (x *ChunkDataMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChunkDataMessage) GetAuthorization() []byte {
	if x != nil {
		return x.Authorization
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size              int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	MaxChunkSize      uint32 `protobuf:"varint,4,opt,name=maxChunkSize,proto3" json:"maxChunkSize,omitempty"`
	ChunksStoresCount uint32 `protobuf:"varint,5,opt,name=chunksStoresCount,proto3" json:"chunksStoresCount,omitempty"`
	TotalChunksCount  uint64 `protobuf:"varint,6,opt,name=totalChunksCount,proto3" json:"totalChunksCount,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetMaxChunkSize() uint32 {
	if x != nil {
		return x.MaxChunkSize
	}
	return 0
}

func (x *File) GetChunksStoresCount() uint32 {
	if x != nil {
		return x.ChunksStoresCount
	}
	return 0
}

func (x *File) GetTotalChunksCount() uint64 {
	if x != nil {
		return x.TotalChunksCount
	}
	return 0
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[4]
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
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{4}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *File `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

type FileChunkUploadedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId              string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	TotalChunksCount    uint64 `protobuf:"varint,2,opt,name=totalChunksCount,proto3" json:"totalChunksCount,omitempty"`
	UploadedChunksCount uint64 `protobuf:"varint,3,opt,name=uploadedChunksCount,proto3" json:"uploadedChunksCount,omitempty"`
}

func (x *FileChunkUploadedEventData) Reset() {
	*x = FileChunkUploadedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_service_file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunkUploadedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunkUploadedEventData) ProtoMessage() {}

func (x *FileChunkUploadedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_service_file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunkUploadedEventData.ProtoReflect.Descriptor instead.
func (*FileChunkUploadedEventData) Descriptor() ([]byte, []int) {
	return file_proto_file_service_file_service_proto_rawDescGZIP(), []int{6}
}

func (x *FileChunkUploadedEventData) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileChunkUploadedEventData) GetTotalChunksCount() uint64 {
	if x != nil {
		return x.TotalChunksCount
	}
	return 0
}

func (x *FileChunkUploadedEventData) GetUploadedChunksCount() uint64 {
	if x != nil {
		return x.UploadedChunksCount
	}
	return 0
}

var File_proto_file_service_file_service_proto protoreflect.FileDescriptor

var file_proto_file_service_file_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x68, 0x61, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x78, 0x0a, 0x10, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x68, 0x61, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x7c, 0x0a,
	0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0a, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x08, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_file_service_file_service_proto_rawDescOnce sync.Once
	file_proto_file_service_file_service_proto_rawDescData = file_proto_file_service_file_service_proto_rawDesc
)

func file_proto_file_service_file_service_proto_rawDescGZIP() []byte {
	file_proto_file_service_file_service_proto_rawDescOnce.Do(func() {
		file_proto_file_service_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_service_file_service_proto_rawDescData)
	})
	return file_proto_file_service_file_service_proto_rawDescData
}

var file_proto_file_service_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_file_service_file_service_proto_goTypes = []interface{}{
	(*ChunkSpec)(nil),                  // 0: ChunkSpec
	(*Chunk)(nil),                      // 1: Chunk
	(*ChunkDataMessage)(nil),           // 2: ChunkDataMessage
	(*File)(nil),                       // 3: File
	(*EmptyResponse)(nil),              // 4: EmptyResponse
	(*Response)(nil),                   // 5: Response
	(*FileChunkUploadedEventData)(nil), // 6: FileChunkUploadedEventData
}
var file_proto_file_service_file_service_proto_depIdxs = []int32{
	3, // 0: Response.file:type_name -> File
	0, // 1: FileService.CreateChunk:input_type -> ChunkSpec
	3, // 2: FileService.CreateFile:input_type -> File
	3, // 3: FileService.ReadFile:input_type -> File
	4, // 4: FileService.CreateChunk:output_type -> EmptyResponse
	5, // 5: FileService.CreateFile:output_type -> Response
	5, // 6: FileService.ReadFile:output_type -> Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_file_service_file_service_proto_init() }
func file_proto_file_service_file_service_proto_init() {
	if File_proto_file_service_file_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_service_file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkSpec); i {
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
		file_proto_file_service_file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_proto_file_service_file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkDataMessage); i {
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
		file_proto_file_service_file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_proto_file_service_file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_file_service_file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_file_service_file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunkUploadedEventData); i {
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
			RawDescriptor: file_proto_file_service_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_service_file_service_proto_goTypes,
		DependencyIndexes: file_proto_file_service_file_service_proto_depIdxs,
		MessageInfos:      file_proto_file_service_file_service_proto_msgTypes,
	}.Build()
	File_proto_file_service_file_service_proto = out.File
	file_proto_file_service_file_service_proto_rawDesc = nil
	file_proto_file_service_file_service_proto_goTypes = nil
	file_proto_file_service_file_service_proto_depIdxs = nil
}
