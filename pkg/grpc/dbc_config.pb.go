// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pkg/grpc/dbc_config.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DbcConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string                 `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	DbcFile      string                 `protobuf:"bytes,2,opt,name=dbc_file,json=dbcFile,proto3" json:"dbc_file,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DbcConfig) Reset() {
	*x = DbcConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbcConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbcConfig) ProtoMessage() {}

func (x *DbcConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbcConfig.ProtoReflect.Descriptor instead.
func (*DbcConfig) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{0}
}

func (x *DbcConfig) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *DbcConfig) GetDbcFile() string {
	if x != nil {
		return x.DbcFile
	}
	return ""
}

func (x *DbcConfig) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DbcConfig) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DbcSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	DbcFile      string `protobuf:"bytes,2,opt,name=dbc_file,json=dbcFile,proto3" json:"dbc_file,omitempty"`
}

func (x *DbcSummary) Reset() {
	*x = DbcSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbcSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbcSummary) ProtoMessage() {}

func (x *DbcSummary) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbcSummary.ProtoReflect.Descriptor instead.
func (*DbcSummary) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{1}
}

func (x *DbcSummary) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *DbcSummary) GetDbcFile() string {
	if x != nil {
		return x.DbcFile
	}
	return ""
}

type GetDbcListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName *string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3,oneof" json:"template_name,omitempty"`
}

func (x *GetDbcListRequest) Reset() {
	*x = GetDbcListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDbcListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDbcListRequest) ProtoMessage() {}

func (x *GetDbcListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDbcListRequest.ProtoReflect.Descriptor instead.
func (*GetDbcListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{2}
}

func (x *GetDbcListRequest) GetTemplateName() string {
	if x != nil && x.TemplateName != nil {
		return *x.TemplateName
	}
	return ""
}

type GetDbcListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dbc []*DbcSummary `protobuf:"bytes,1,rep,name=dbc,proto3" json:"dbc,omitempty"`
}

func (x *GetDbcListResponse) Reset() {
	*x = GetDbcListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDbcListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDbcListResponse) ProtoMessage() {}

func (x *GetDbcListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDbcListResponse.ProtoReflect.Descriptor instead.
func (*GetDbcListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{3}
}

func (x *GetDbcListResponse) GetDbc() []*DbcSummary {
	if x != nil {
		return x.Dbc
	}
	return nil
}

type GetDbcByTemplateNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *GetDbcByTemplateNameRequest) Reset() {
	*x = GetDbcByTemplateNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDbcByTemplateNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDbcByTemplateNameRequest) ProtoMessage() {}

func (x *GetDbcByTemplateNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDbcByTemplateNameRequest.ProtoReflect.Descriptor instead.
func (*GetDbcByTemplateNameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetDbcByTemplateNameRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type GetDbcByTemplateNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dbc *DbcConfig `protobuf:"bytes,1,opt,name=dbc,proto3" json:"dbc,omitempty"`
}

func (x *GetDbcByTemplateNameResponse) Reset() {
	*x = GetDbcByTemplateNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDbcByTemplateNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDbcByTemplateNameResponse) ProtoMessage() {}

func (x *GetDbcByTemplateNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDbcByTemplateNameResponse.ProtoReflect.Descriptor instead.
func (*GetDbcByTemplateNameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetDbcByTemplateNameResponse) GetDbc() *DbcConfig {
	if x != nil {
		return x.Dbc
	}
	return nil
}

type UpdateDbcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dbc *DbcConfig `protobuf:"bytes,1,opt,name=dbc,proto3" json:"dbc,omitempty"`
}

func (x *UpdateDbcRequest) Reset() {
	*x = UpdateDbcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_dbc_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDbcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDbcRequest) ProtoMessage() {}

func (x *UpdateDbcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_dbc_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDbcRequest.ProtoReflect.Descriptor instead.
func (*UpdateDbcRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_dbc_config_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDbcRequest) GetDbc() *DbcConfig {
	if x != nil {
		return x.Dbc
	}
	return nil
}

var File_pkg_grpc_dbc_config_proto protoreflect.FileDescriptor

var file_pkg_grpc_dbc_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x62, 0x63, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc1, 0x01, 0x0a, 0x09, 0x44, 0x62, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x62, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x4c, 0x0a, 0x0a, 0x44, 0x62, 0x63, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x62, 0x63, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x63, 0x46, 0x69, 0x6c,
	0x65, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x62, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x62, 0x63, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x64, 0x62, 0x63, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x62, 0x63,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x03, 0x64, 0x62, 0x63, 0x22, 0x42, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x44, 0x62, 0x63, 0x42, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x41, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x62, 0x63, 0x42, 0x79, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x03, 0x64, 0x62, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x62, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03,
	0x64, 0x62, 0x63, 0x22, 0x35, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x62, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x64, 0x62, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x62, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x64, 0x62, 0x63, 0x32, 0xac, 0x02, 0x0a, 0x10, 0x44,
	0x62, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x62, 0x63, 0x12, 0x16, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x62, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x62, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x62, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x44, 0x62, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x62, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x62, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x44, 0x62, 0x63, 0x42, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x62, 0x63,
	0x42, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x62, 0x63, 0x42, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_dbc_config_proto_rawDescOnce sync.Once
	file_pkg_grpc_dbc_config_proto_rawDescData = file_pkg_grpc_dbc_config_proto_rawDesc
)

func file_pkg_grpc_dbc_config_proto_rawDescGZIP() []byte {
	file_pkg_grpc_dbc_config_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_dbc_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_dbc_config_proto_rawDescData)
	})
	return file_pkg_grpc_dbc_config_proto_rawDescData
}

var file_pkg_grpc_dbc_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_grpc_dbc_config_proto_goTypes = []interface{}{
	(*DbcConfig)(nil),                    // 0: grpc.DbcConfig
	(*DbcSummary)(nil),                   // 1: grpc.DbcSummary
	(*GetDbcListRequest)(nil),            // 2: grpc.GetDbcListRequest
	(*GetDbcListResponse)(nil),           // 3: grpc.GetDbcListResponse
	(*GetDbcByTemplateNameRequest)(nil),  // 4: grpc.GetDbcByTemplateNameRequest
	(*GetDbcByTemplateNameResponse)(nil), // 5: grpc.GetDbcByTemplateNameResponse
	(*UpdateDbcRequest)(nil),             // 6: grpc.UpdateDbcRequest
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                // 8: google.protobuf.Empty
}
var file_pkg_grpc_dbc_config_proto_depIdxs = []int32{
	7, // 0: grpc.DbcConfig.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: grpc.DbcConfig.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: grpc.GetDbcListResponse.dbc:type_name -> grpc.DbcSummary
	0, // 3: grpc.GetDbcByTemplateNameResponse.dbc:type_name -> grpc.DbcConfig
	0, // 4: grpc.UpdateDbcRequest.dbc:type_name -> grpc.DbcConfig
	6, // 5: grpc.DbcConfigService.CreateDbc:input_type -> grpc.UpdateDbcRequest
	6, // 6: grpc.DbcConfigService.UpdateDbc:input_type -> grpc.UpdateDbcRequest
	2, // 7: grpc.DbcConfigService.GetDbcList:input_type -> grpc.GetDbcListRequest
	4, // 8: grpc.DbcConfigService.GetDbcByTemplateName:input_type -> grpc.GetDbcByTemplateNameRequest
	8, // 9: grpc.DbcConfigService.CreateDbc:output_type -> google.protobuf.Empty
	8, // 10: grpc.DbcConfigService.UpdateDbc:output_type -> google.protobuf.Empty
	3, // 11: grpc.DbcConfigService.GetDbcList:output_type -> grpc.GetDbcListResponse
	5, // 12: grpc.DbcConfigService.GetDbcByTemplateName:output_type -> grpc.GetDbcByTemplateNameResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_grpc_dbc_config_proto_init() }
func file_pkg_grpc_dbc_config_proto_init() {
	if File_pkg_grpc_dbc_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_dbc_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbcConfig); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbcSummary); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDbcListRequest); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDbcListResponse); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDbcByTemplateNameRequest); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDbcByTemplateNameResponse); i {
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
		file_pkg_grpc_dbc_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDbcRequest); i {
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
	file_pkg_grpc_dbc_config_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_dbc_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_dbc_config_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_dbc_config_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_dbc_config_proto_msgTypes,
	}.Build()
	File_pkg_grpc_dbc_config_proto = out.File
	file_pkg_grpc_dbc_config_proto_rawDesc = nil
	file_pkg_grpc_dbc_config_proto_goTypes = nil
	file_pkg_grpc_dbc_config_proto_depIdxs = nil
}
