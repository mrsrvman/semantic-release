// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pkg/generator/changelog_generator.proto

package generator

import (
	semrel "github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
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

type ChangelogGeneratorInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangelogGeneratorInit) Reset() {
	*x = ChangelogGeneratorInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorInit) ProtoMessage() {}

func (x *ChangelogGeneratorInit) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorInit.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorInit) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{0}
}

type ChangelogGeneratorName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangelogGeneratorName) Reset() {
	*x = ChangelogGeneratorName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorName) ProtoMessage() {}

func (x *ChangelogGeneratorName) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorName.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorName) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1}
}

type ChangelogGeneratorVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangelogGeneratorVersion) Reset() {
	*x = ChangelogGeneratorVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorVersion) ProtoMessage() {}

func (x *ChangelogGeneratorVersion) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorVersion.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorVersion) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{2}
}

type ChangelogGeneratorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commits       []*semrel.Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
	LatestRelease *semrel.Release  `protobuf:"bytes,2,opt,name=latest_release,json=latestRelease,proto3" json:"latest_release,omitempty"`
	NewVersion    string           `protobuf:"bytes,3,opt,name=new_version,json=newVersion,proto3" json:"new_version,omitempty"`
}

func (x *ChangelogGeneratorConfig) Reset() {
	*x = ChangelogGeneratorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorConfig) ProtoMessage() {}

func (x *ChangelogGeneratorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorConfig.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorConfig) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{3}
}

func (x *ChangelogGeneratorConfig) GetCommits() []*semrel.Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

func (x *ChangelogGeneratorConfig) GetLatestRelease() *semrel.Release {
	if x != nil {
		return x.LatestRelease
	}
	return nil
}

func (x *ChangelogGeneratorConfig) GetNewVersion() string {
	if x != nil {
		return x.NewVersion
	}
	return ""
}

type GenerateChangelog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateChangelog) Reset() {
	*x = GenerateChangelog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog) ProtoMessage() {}

func (x *GenerateChangelog) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog.ProtoReflect.Descriptor instead.
func (*GenerateChangelog) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{4}
}

type ChangelogGeneratorInit_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChangelogGeneratorInit_Request) Reset() {
	*x = ChangelogGeneratorInit_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorInit_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorInit_Request) ProtoMessage() {}

func (x *ChangelogGeneratorInit_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorInit_Request.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorInit_Request) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChangelogGeneratorInit_Request) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type ChangelogGeneratorInit_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ChangelogGeneratorInit_Response) Reset() {
	*x = ChangelogGeneratorInit_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorInit_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorInit_Response) ProtoMessage() {}

func (x *ChangelogGeneratorInit_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorInit_Response.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorInit_Response) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ChangelogGeneratorInit_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ChangelogGeneratorName_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangelogGeneratorName_Request) Reset() {
	*x = ChangelogGeneratorName_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorName_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorName_Request) ProtoMessage() {}

func (x *ChangelogGeneratorName_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorName_Request.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorName_Request) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1, 0}
}

type ChangelogGeneratorName_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ChangelogGeneratorName_Response) Reset() {
	*x = ChangelogGeneratorName_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorName_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorName_Response) ProtoMessage() {}

func (x *ChangelogGeneratorName_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorName_Response.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorName_Response) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ChangelogGeneratorName_Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangelogGeneratorVersion_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangelogGeneratorVersion_Request) Reset() {
	*x = ChangelogGeneratorVersion_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorVersion_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorVersion_Request) ProtoMessage() {}

func (x *ChangelogGeneratorVersion_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorVersion_Request.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorVersion_Request) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{2, 0}
}

type ChangelogGeneratorVersion_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ChangelogGeneratorVersion_Response) Reset() {
	*x = ChangelogGeneratorVersion_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangelogGeneratorVersion_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogGeneratorVersion_Response) ProtoMessage() {}

func (x *ChangelogGeneratorVersion_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogGeneratorVersion_Response.ProtoReflect.Descriptor instead.
func (*ChangelogGeneratorVersion_Response) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ChangelogGeneratorVersion_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GenerateChangelog_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ChangelogGeneratorConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GenerateChangelog_Request) Reset() {
	*x = GenerateChangelog_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog_Request) ProtoMessage() {}

func (x *GenerateChangelog_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog_Request.ProtoReflect.Descriptor instead.
func (*GenerateChangelog_Request) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GenerateChangelog_Request) GetConfig() *ChangelogGeneratorConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type GenerateChangelog_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changelog string `protobuf:"bytes,1,opt,name=changelog,proto3" json:"changelog,omitempty"`
}

func (x *GenerateChangelog_Response) Reset() {
	*x = GenerateChangelog_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_generator_changelog_generator_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateChangelog_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateChangelog_Response) ProtoMessage() {}

func (x *GenerateChangelog_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_generator_changelog_generator_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateChangelog_Response.ProtoReflect.Descriptor instead.
func (*GenerateChangelog_Response) Descriptor() ([]byte, []int) {
	return file_pkg_generator_changelog_generator_proto_rawDescGZIP(), []int{4, 1}
}

func (x *GenerateChangelog_Response) GetChangelog() string {
	if x != nil {
		return x.Changelog
	}
	return ""
}

var File_pkg_generator_changelog_generator_proto protoreflect.FileDescriptor

var file_pkg_generator_changelog_generator_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x6d, 0x72, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x1a, 0x89,
	0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x16,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4c, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x09,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x8f, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x2f, 0x0a, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x7b, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x1a, 0x3c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x32, 0xc9,
	0x02, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x49, 0x0a, 0x04, 0x49,
	0x6e, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x6d,
	0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_generator_changelog_generator_proto_rawDescOnce sync.Once
	file_pkg_generator_changelog_generator_proto_rawDescData = file_pkg_generator_changelog_generator_proto_rawDesc
)

func file_pkg_generator_changelog_generator_proto_rawDescGZIP() []byte {
	file_pkg_generator_changelog_generator_proto_rawDescOnce.Do(func() {
		file_pkg_generator_changelog_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_generator_changelog_generator_proto_rawDescData)
	})
	return file_pkg_generator_changelog_generator_proto_rawDescData
}

var file_pkg_generator_changelog_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pkg_generator_changelog_generator_proto_goTypes = []interface{}{
	(*ChangelogGeneratorInit)(nil),             // 0: ChangelogGeneratorInit
	(*ChangelogGeneratorName)(nil),             // 1: ChangelogGeneratorName
	(*ChangelogGeneratorVersion)(nil),          // 2: ChangelogGeneratorVersion
	(*ChangelogGeneratorConfig)(nil),           // 3: ChangelogGeneratorConfig
	(*GenerateChangelog)(nil),                  // 4: GenerateChangelog
	(*ChangelogGeneratorInit_Request)(nil),     // 5: ChangelogGeneratorInit.Request
	(*ChangelogGeneratorInit_Response)(nil),    // 6: ChangelogGeneratorInit.Response
	nil,                                        // 7: ChangelogGeneratorInit.Request.ConfigEntry
	(*ChangelogGeneratorName_Request)(nil),     // 8: ChangelogGeneratorName.Request
	(*ChangelogGeneratorName_Response)(nil),    // 9: ChangelogGeneratorName.Response
	(*ChangelogGeneratorVersion_Request)(nil),  // 10: ChangelogGeneratorVersion.Request
	(*ChangelogGeneratorVersion_Response)(nil), // 11: ChangelogGeneratorVersion.Response
	(*GenerateChangelog_Request)(nil),          // 12: GenerateChangelog.Request
	(*GenerateChangelog_Response)(nil),         // 13: GenerateChangelog.Response
	(*semrel.Commit)(nil),                      // 14: Commit
	(*semrel.Release)(nil),                     // 15: Release
}
var file_pkg_generator_changelog_generator_proto_depIdxs = []int32{
	14, // 0: ChangelogGeneratorConfig.commits:type_name -> Commit
	15, // 1: ChangelogGeneratorConfig.latest_release:type_name -> Release
	7,  // 2: ChangelogGeneratorInit.Request.config:type_name -> ChangelogGeneratorInit.Request.ConfigEntry
	3,  // 3: GenerateChangelog.Request.config:type_name -> ChangelogGeneratorConfig
	5,  // 4: ChangelogGeneratorPlugin.Init:input_type -> ChangelogGeneratorInit.Request
	8,  // 5: ChangelogGeneratorPlugin.Name:input_type -> ChangelogGeneratorName.Request
	10, // 6: ChangelogGeneratorPlugin.Version:input_type -> ChangelogGeneratorVersion.Request
	12, // 7: ChangelogGeneratorPlugin.Generate:input_type -> GenerateChangelog.Request
	6,  // 8: ChangelogGeneratorPlugin.Init:output_type -> ChangelogGeneratorInit.Response
	9,  // 9: ChangelogGeneratorPlugin.Name:output_type -> ChangelogGeneratorName.Response
	11, // 10: ChangelogGeneratorPlugin.Version:output_type -> ChangelogGeneratorVersion.Response
	13, // 11: ChangelogGeneratorPlugin.Generate:output_type -> GenerateChangelog.Response
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_generator_changelog_generator_proto_init() }
func file_pkg_generator_changelog_generator_proto_init() {
	if File_pkg_generator_changelog_generator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_generator_changelog_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorInit); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorName); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorVersion); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorConfig); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorInit_Request); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorInit_Response); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorName_Request); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorName_Response); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorVersion_Request); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangelogGeneratorVersion_Response); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog_Request); i {
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
		file_pkg_generator_changelog_generator_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateChangelog_Response); i {
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
			RawDescriptor: file_pkg_generator_changelog_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_generator_changelog_generator_proto_goTypes,
		DependencyIndexes: file_pkg_generator_changelog_generator_proto_depIdxs,
		MessageInfos:      file_pkg_generator_changelog_generator_proto_msgTypes,
	}.Build()
	File_pkg_generator_changelog_generator_proto = out.File
	file_pkg_generator_changelog_generator_proto_rawDesc = nil
	file_pkg_generator_changelog_generator_proto_goTypes = nil
	file_pkg_generator_changelog_generator_proto_depIdxs = nil
}
