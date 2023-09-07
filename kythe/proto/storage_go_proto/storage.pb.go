// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: kythe/proto/storage.proto

package storage_go_proto

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

type VName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Corpus    string `protobuf:"bytes,2,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Root      string `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Language  string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *VName) Reset() {
	*x = VName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VName) ProtoMessage() {}

func (x *VName) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VName.ProtoReflect.Descriptor instead.
func (*VName) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{0}
}

func (x *VName) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *VName) GetCorpus() string {
	if x != nil {
		return x.Corpus
	}
	return ""
}

func (x *VName) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *VName) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *VName) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type VNameMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature bool `protobuf:"varint,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Corpus    bool `protobuf:"varint,2,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Root      bool `protobuf:"varint,3,opt,name=root,proto3" json:"root,omitempty"`
	Path      bool `protobuf:"varint,4,opt,name=path,proto3" json:"path,omitempty"`
	Language  bool `protobuf:"varint,5,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *VNameMask) Reset() {
	*x = VNameMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VNameMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VNameMask) ProtoMessage() {}

func (x *VNameMask) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VNameMask.ProtoReflect.Descriptor instead.
func (*VNameMask) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{1}
}

func (x *VNameMask) GetSignature() bool {
	if x != nil {
		return x.Signature
	}
	return false
}

func (x *VNameMask) GetCorpus() bool {
	if x != nil {
		return x.Corpus
	}
	return false
}

func (x *VNameMask) GetRoot() bool {
	if x != nil {
		return x.Root
	}
	return false
}

func (x *VNameMask) GetPath() bool {
	if x != nil {
		return x.Path
	}
	return false
}

func (x *VNameMask) GetLanguage() bool {
	if x != nil {
		return x.Language
	}
	return false
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source    *VName `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	EdgeKind  string `protobuf:"bytes,2,opt,name=edge_kind,json=edgeKind,proto3" json:"edge_kind,omitempty"`
	Target    *VName `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	FactName  string `protobuf:"bytes,4,opt,name=fact_name,json=factName,proto3" json:"fact_name,omitempty"`
	FactValue []byte `protobuf:"bytes,5,opt,name=fact_value,json=factValue,proto3" json:"fact_value,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Entry) GetSource() *VName {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Entry) GetEdgeKind() string {
	if x != nil {
		return x.EdgeKind
	}
	return ""
}

func (x *Entry) GetTarget() *VName {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Entry) GetFactName() string {
	if x != nil {
		return x.FactName
	}
	return ""
}

func (x *Entry) GetFactValue() []byte {
	if x != nil {
		return x.FactValue
	}
	return nil
}

type Entries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Entries) Reset() {
	*x = Entries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entries) ProtoMessage() {}

func (x *Entries) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entries.ProtoReflect.Descriptor instead.
func (*Entries) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{3}
}

func (x *Entries) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source   *VName `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	EdgeKind string `protobuf:"bytes,2,opt,name=edge_kind,json=edgeKind,proto3" json:"edge_kind,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRequest) GetSource() *VName {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ReadRequest) GetEdgeKind() string {
	if x != nil {
		return x.EdgeKind
	}
	return ""
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *VName                 `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Update []*WriteRequest_Update `protobuf:"bytes,2,rep,name=update,proto3" json:"update,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{5}
}

func (x *WriteRequest) GetSource() *VName {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *WriteRequest) GetUpdate() []*WriteRequest_Update {
	if x != nil {
		return x.Update
	}
	return nil
}

type WriteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteReply) Reset() {
	*x = WriteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReply) ProtoMessage() {}

func (x *WriteReply) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReply.ProtoReflect.Descriptor instead.
func (*WriteReply) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{6}
}

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target     *VName `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	EdgeKind   string `protobuf:"bytes,2,opt,name=edge_kind,json=edgeKind,proto3" json:"edge_kind,omitempty"`
	FactPrefix string `protobuf:"bytes,3,opt,name=fact_prefix,json=factPrefix,proto3" json:"fact_prefix,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{7}
}

func (x *ScanRequest) GetTarget() *VName {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ScanRequest) GetEdgeKind() string {
	if x != nil {
		return x.EdgeKind
	}
	return ""
}

func (x *ScanRequest) GetFactPrefix() string {
	if x != nil {
		return x.FactPrefix
	}
	return ""
}

type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Shards int64 `protobuf:"varint,2,opt,name=shards,proto3" json:"shards,omitempty"`
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{8}
}

func (x *CountRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CountRequest) GetShards() int64 {
	if x != nil {
		return x.Shards
	}
	return 0
}

type CountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries int64 `protobuf:"varint,1,opt,name=entries,proto3" json:"entries,omitempty"`
}

func (x *CountReply) Reset() {
	*x = CountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountReply) ProtoMessage() {}

func (x *CountReply) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountReply.ProtoReflect.Descriptor instead.
func (*CountReply) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{9}
}

func (x *CountReply) GetEntries() int64 {
	if x != nil {
		return x.Entries
	}
	return 0
}

type ShardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Shards int64 `protobuf:"varint,2,opt,name=shards,proto3" json:"shards,omitempty"`
}

func (x *ShardRequest) Reset() {
	*x = ShardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardRequest) ProtoMessage() {}

func (x *ShardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardRequest.ProtoReflect.Descriptor instead.
func (*ShardRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{10}
}

func (x *ShardRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ShardRequest) GetShards() int64 {
	if x != nil {
		return x.Shards
	}
	return 0
}

type VNameRewriteRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern string `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	VName   *VName `protobuf:"bytes,2,opt,name=v_name,json=vname,proto3" json:"v_name,omitempty"`
}

func (x *VNameRewriteRule) Reset() {
	*x = VNameRewriteRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VNameRewriteRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VNameRewriteRule) ProtoMessage() {}

func (x *VNameRewriteRule) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VNameRewriteRule.ProtoReflect.Descriptor instead.
func (*VNameRewriteRule) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{11}
}

func (x *VNameRewriteRule) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *VNameRewriteRule) GetVName() *VName {
	if x != nil {
		return x.VName
	}
	return nil
}

type VNameRewriteRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule []*VNameRewriteRule `protobuf:"bytes,1,rep,name=rule,proto3" json:"rule,omitempty"`
}

func (x *VNameRewriteRules) Reset() {
	*x = VNameRewriteRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VNameRewriteRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VNameRewriteRules) ProtoMessage() {}

func (x *VNameRewriteRules) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VNameRewriteRules.ProtoReflect.Descriptor instead.
func (*VNameRewriteRules) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{12}
}

func (x *VNameRewriteRules) GetRule() []*VNameRewriteRule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type WriteRequest_Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EdgeKind  string `protobuf:"bytes,1,opt,name=edge_kind,json=edgeKind,proto3" json:"edge_kind,omitempty"`
	Target    *VName `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	FactName  string `protobuf:"bytes,3,opt,name=fact_name,json=factName,proto3" json:"fact_name,omitempty"`
	FactValue []byte `protobuf:"bytes,4,opt,name=fact_value,json=factValue,proto3" json:"fact_value,omitempty"`
}

func (x *WriteRequest_Update) Reset() {
	*x = WriteRequest_Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_storage_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest_Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest_Update) ProtoMessage() {}

func (x *WriteRequest_Update) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_storage_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest_Update.ProtoReflect.Descriptor instead.
func (*WriteRequest_Update) Descriptor() ([]byte, []int) {
	return file_kythe_proto_storage_proto_rawDescGZIP(), []int{5, 0}
}

func (x *WriteRequest_Update) GetEdgeKind() string {
	if x != nil {
		return x.EdgeKind
	}
	return ""
}

func (x *WriteRequest_Update) GetTarget() *VName {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *WriteRequest_Update) GetFactName() string {
	if x != nil {
		return x.FactName
	}
	return ""
}

func (x *WriteRequest_Update) GetFactValue() []byte {
	if x != nil {
		return x.FactValue
	}
	return nil
}

var File_kythe_proto_storage_proto protoreflect.FileDescriptor

var file_kythe_proto_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x79, 0x74,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x56, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01, 0x0a,
	0x09, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x64,
	0x67, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x64, 0x67, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x37, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x64, 0x67, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x22, 0x84, 0x02, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x8d, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x64, 0x67, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x2a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x61,
	0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x77, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x64, 0x67, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x3c,
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x22, 0x26, 0x0a, 0x0a,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x73, 0x22, 0x57, 0x0a, 0x10, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x29, 0x0a, 0x06, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x76, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x56,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x42, 0x48, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x25, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x69, 0x6f,
	0x2f, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kythe_proto_storage_proto_rawDescOnce sync.Once
	file_kythe_proto_storage_proto_rawDescData = file_kythe_proto_storage_proto_rawDesc
)

func file_kythe_proto_storage_proto_rawDescGZIP() []byte {
	file_kythe_proto_storage_proto_rawDescOnce.Do(func() {
		file_kythe_proto_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_storage_proto_rawDescData)
	})
	return file_kythe_proto_storage_proto_rawDescData
}

var file_kythe_proto_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_kythe_proto_storage_proto_goTypes = []interface{}{
	(*VName)(nil),               // 0: kythe.proto.VName
	(*VNameMask)(nil),           // 1: kythe.proto.VNameMask
	(*Entry)(nil),               // 2: kythe.proto.Entry
	(*Entries)(nil),             // 3: kythe.proto.Entries
	(*ReadRequest)(nil),         // 4: kythe.proto.ReadRequest
	(*WriteRequest)(nil),        // 5: kythe.proto.WriteRequest
	(*WriteReply)(nil),          // 6: kythe.proto.WriteReply
	(*ScanRequest)(nil),         // 7: kythe.proto.ScanRequest
	(*CountRequest)(nil),        // 8: kythe.proto.CountRequest
	(*CountReply)(nil),          // 9: kythe.proto.CountReply
	(*ShardRequest)(nil),        // 10: kythe.proto.ShardRequest
	(*VNameRewriteRule)(nil),    // 11: kythe.proto.VNameRewriteRule
	(*VNameRewriteRules)(nil),   // 12: kythe.proto.VNameRewriteRules
	(*WriteRequest_Update)(nil), // 13: kythe.proto.WriteRequest.Update
}
var file_kythe_proto_storage_proto_depIdxs = []int32{
	0,  // 0: kythe.proto.Entry.source:type_name -> kythe.proto.VName
	0,  // 1: kythe.proto.Entry.target:type_name -> kythe.proto.VName
	2,  // 2: kythe.proto.Entries.entries:type_name -> kythe.proto.Entry
	0,  // 3: kythe.proto.ReadRequest.source:type_name -> kythe.proto.VName
	0,  // 4: kythe.proto.WriteRequest.source:type_name -> kythe.proto.VName
	13, // 5: kythe.proto.WriteRequest.update:type_name -> kythe.proto.WriteRequest.Update
	0,  // 6: kythe.proto.ScanRequest.target:type_name -> kythe.proto.VName
	0,  // 7: kythe.proto.VNameRewriteRule.v_name:type_name -> kythe.proto.VName
	11, // 8: kythe.proto.VNameRewriteRules.rule:type_name -> kythe.proto.VNameRewriteRule
	0,  // 9: kythe.proto.WriteRequest.Update.target:type_name -> kythe.proto.VName
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_kythe_proto_storage_proto_init() }
func file_kythe_proto_storage_proto_init() {
	if File_kythe_proto_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VName); i {
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
		file_kythe_proto_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VNameMask); i {
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
		file_kythe_proto_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_kythe_proto_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entries); i {
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
		file_kythe_proto_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_kythe_proto_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_kythe_proto_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReply); i {
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
		file_kythe_proto_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
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
		file_kythe_proto_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRequest); i {
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
		file_kythe_proto_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountReply); i {
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
		file_kythe_proto_storage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardRequest); i {
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
		file_kythe_proto_storage_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VNameRewriteRule); i {
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
		file_kythe_proto_storage_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VNameRewriteRules); i {
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
		file_kythe_proto_storage_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest_Update); i {
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
			RawDescriptor: file_kythe_proto_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kythe_proto_storage_proto_goTypes,
		DependencyIndexes: file_kythe_proto_storage_proto_depIdxs,
		MessageInfos:      file_kythe_proto_storage_proto_msgTypes,
	}.Build()
	File_kythe_proto_storage_proto = out.File
	file_kythe_proto_storage_proto_rawDesc = nil
	file_kythe_proto_storage_proto_goTypes = nil
	file_kythe_proto_storage_proto_depIdxs = nil
}
