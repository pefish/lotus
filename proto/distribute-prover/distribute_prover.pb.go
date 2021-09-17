// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: distribute_prover.proto

package distribute_prover

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{0}
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GenerateWindowPoStReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GenerateWindowPoStReply) Reset() {
	*x = GenerateWindowPoStReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateWindowPoStReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateWindowPoStReply) ProtoMessage() {}

func (x *GenerateWindowPoStReply) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateWindowPoStReply.ProtoReflect.Descriptor instead.
func (*GenerateWindowPoStReply) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateWindowPoStReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GenerateWindowPoStRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId        uint64 `protobuf:"varint,1,opt,name=actorId,proto3" json:"actorId,omitempty"`
	Sinfos         []byte `protobuf:"bytes,2,opt,name=sinfos,proto3" json:"sinfos,omitempty"`
	Randomness     []byte `protobuf:"bytes,3,opt,name=randomness,proto3" json:"randomness,omitempty"`
	PartitionIndex uint64 `protobuf:"varint,4,opt,name=partitionIndex,proto3" json:"partitionIndex,omitempty"`
}

func (x *GenerateWindowPoStRequest) Reset() {
	*x = GenerateWindowPoStRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateWindowPoStRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateWindowPoStRequest) ProtoMessage() {}

func (x *GenerateWindowPoStRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateWindowPoStRequest.ProtoReflect.Descriptor instead.
func (*GenerateWindowPoStRequest) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateWindowPoStRequest) GetActorId() uint64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *GenerateWindowPoStRequest) GetSinfos() []byte {
	if x != nil {
		return x.Sinfos
	}
	return nil
}

func (x *GenerateWindowPoStRequest) GetRandomness() []byte {
	if x != nil {
		return x.Randomness
	}
	return nil
}

func (x *GenerateWindowPoStRequest) GetPartitionIndex() uint64 {
	if x != nil {
		return x.PartitionIndex
	}
	return 0
}

type GetGenerateWindowPoStResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionIndex uint64 `protobuf:"varint,1,opt,name=partitionIndex,proto3" json:"partitionIndex,omitempty"`
}

func (x *GetGenerateWindowPoStResultRequest) Reset() {
	*x = GetGenerateWindowPoStResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGenerateWindowPoStResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenerateWindowPoStResultRequest) ProtoMessage() {}

func (x *GetGenerateWindowPoStResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenerateWindowPoStResultRequest.ProtoReflect.Descriptor instead.
func (*GetGenerateWindowPoStResultRequest) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{4}
}

func (x *GetGenerateWindowPoStResultRequest) GetPartitionIndex() uint64 {
	if x != nil {
		return x.PartitionIndex
	}
	return 0
}

type GetGenerateWindowPoStResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof    []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	Skipped  []byte `protobuf:"bytes,2,opt,name=skipped,proto3" json:"skipped,omitempty"`
	Finished bool   `protobuf:"varint,3,opt,name=finished,proto3" json:"finished,omitempty"`
}

func (x *GetGenerateWindowPoStResultReply) Reset() {
	*x = GetGenerateWindowPoStResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distribute_prover_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGenerateWindowPoStResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenerateWindowPoStResultReply) ProtoMessage() {}

func (x *GetGenerateWindowPoStResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_distribute_prover_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenerateWindowPoStResultReply.ProtoReflect.Descriptor instead.
func (*GetGenerateWindowPoStResultReply) Descriptor() ([]byte, []int) {
	return file_distribute_prover_proto_rawDescGZIP(), []int{5}
}

func (x *GetGenerateWindowPoStResultReply) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *GetGenerateWindowPoStResultReply) GetSkipped() []byte {
	if x != nil {
		return x.Skipped
	}
	return nil
}

func (x *GetGenerateWindowPoStResultReply) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

var File_distribute_prover_proto protoreflect.FileDescriptor

var file_distribute_prover_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x0d, 0x0a, 0x0b,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x09, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2b, 0x0a, 0x17, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x95, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x4c, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6e, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x32, 0xda, 0x02,
	0x0a, 0x10, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x12, 0x2c, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x50, 0x6f, 0x53, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f,
	0x3b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distribute_prover_proto_rawDescOnce sync.Once
	file_distribute_prover_proto_rawDescData = file_distribute_prover_proto_rawDesc
)

func file_distribute_prover_proto_rawDescGZIP() []byte {
	file_distribute_prover_proto_rawDescOnce.Do(func() {
		file_distribute_prover_proto_rawDescData = protoimpl.X.CompressGZIP(file_distribute_prover_proto_rawDescData)
	})
	return file_distribute_prover_proto_rawDescData
}

var file_distribute_prover_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_distribute_prover_proto_goTypes = []interface{}{
	(*PingRequest)(nil),                        // 0: distribute_prover.PingRequest
	(*PingReply)(nil),                          // 1: distribute_prover.PingReply
	(*GenerateWindowPoStReply)(nil),            // 2: distribute_prover.GenerateWindowPoStReply
	(*GenerateWindowPoStRequest)(nil),          // 3: distribute_prover.GenerateWindowPoStRequest
	(*GetGenerateWindowPoStResultRequest)(nil), // 4: distribute_prover.GetGenerateWindowPoStResultRequest
	(*GetGenerateWindowPoStResultReply)(nil),   // 5: distribute_prover.GetGenerateWindowPoStResultReply
}
var file_distribute_prover_proto_depIdxs = []int32{
	3, // 0: distribute_prover.DistributeProver.GenerateWindowPoSt:input_type -> distribute_prover.GenerateWindowPoStRequest
	4, // 1: distribute_prover.DistributeProver.GetGenerateWindowPoStResult:input_type -> distribute_prover.GetGenerateWindowPoStResultRequest
	0, // 2: distribute_prover.DistributeProver.Ping:input_type -> distribute_prover.PingRequest
	2, // 3: distribute_prover.DistributeProver.GenerateWindowPoSt:output_type -> distribute_prover.GenerateWindowPoStReply
	5, // 4: distribute_prover.DistributeProver.GetGenerateWindowPoStResult:output_type -> distribute_prover.GetGenerateWindowPoStResultReply
	1, // 5: distribute_prover.DistributeProver.Ping:output_type -> distribute_prover.PingReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distribute_prover_proto_init() }
func file_distribute_prover_proto_init() {
	if File_distribute_prover_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_distribute_prover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_distribute_prover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
		file_distribute_prover_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateWindowPoStReply); i {
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
		file_distribute_prover_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateWindowPoStRequest); i {
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
		file_distribute_prover_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGenerateWindowPoStResultRequest); i {
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
		file_distribute_prover_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGenerateWindowPoStResultReply); i {
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
			RawDescriptor: file_distribute_prover_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distribute_prover_proto_goTypes,
		DependencyIndexes: file_distribute_prover_proto_depIdxs,
		MessageInfos:      file_distribute_prover_proto_msgTypes,
	}.Build()
	File_distribute_prover_proto = out.File
	file_distribute_prover_proto_rawDesc = nil
	file_distribute_prover_proto_goTypes = nil
	file_distribute_prover_proto_depIdxs = nil
}
