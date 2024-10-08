// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: pkg/grpc/template_config.proto

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

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentTemplateName string                 `protobuf:"bytes,2,opt,name=parent_template_name,json=parentTemplateName,proto3" json:"parent_template_name,omitempty"`
	Version            string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Protocol           string                 `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Powertrain         string                 `protobuf:"bytes,5,opt,name=powertrain,proto3" json:"powertrain,omitempty"`
	HasDbc             bool                   `protobuf:"varint,6,opt,name=has_dbc,json=hasDbc,proto3" json:"has_dbc,omitempty"`
	PidsCount          int32                  `protobuf:"varint,7,opt,name=pids_count,json=pidsCount,proto3" json:"pids_count,omitempty"`
	Pids               []*PID                 `protobuf:"bytes,8,rep,name=pids,proto3" json:"pids,omitempty"`
	Dbc                string                 `protobuf:"bytes,9,opt,name=dbc,proto3" json:"dbc,omitempty"`
	TemplateVehicles   []string               `protobuf:"bytes,10,rep,name=template_vehicles,json=templateVehicles,proto3" json:"template_vehicles,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Comments           *string                `protobuf:"bytes,13,opt,name=comments,proto3,oneof" json:"comments,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetParentTemplateName() string {
	if x != nil {
		return x.ParentTemplateName
	}
	return ""
}

func (x *Template) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Template) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Template) GetPowertrain() string {
	if x != nil {
		return x.Powertrain
	}
	return ""
}

func (x *Template) GetHasDbc() bool {
	if x != nil {
		return x.HasDbc
	}
	return false
}

func (x *Template) GetPidsCount() int32 {
	if x != nil {
		return x.PidsCount
	}
	return 0
}

func (x *Template) GetPids() []*PID {
	if x != nil {
		return x.Pids
	}
	return nil
}

func (x *Template) GetDbc() string {
	if x != nil {
		return x.Dbc
	}
	return ""
}

func (x *Template) GetTemplateVehicles() []string {
	if x != nil {
		return x.TemplateVehicles
	}
	return nil
}

func (x *Template) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Template) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Template) GetComments() string {
	if x != nil && x.Comments != nil {
		return *x.Comments
	}
	return ""
}

type TemplateSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version            string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Protocol           string                 `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Powertrain         string                 `protobuf:"bytes,4,opt,name=powertrain,proto3" json:"powertrain,omitempty"`
	HasDbc             bool                   `protobuf:"varint,5,opt,name=has_dbc,json=hasDbc,proto3" json:"has_dbc,omitempty"`
	PidsCount          int32                  `protobuf:"varint,6,opt,name=pids_count,json=pidsCount,proto3" json:"pids_count,omitempty"`
	ParentTemplateName string                 `protobuf:"bytes,7,opt,name=parent_template_name,json=parentTemplateName,proto3" json:"parent_template_name,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TemplateSummary) Reset() {
	*x = TemplateSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSummary) ProtoMessage() {}

func (x *TemplateSummary) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSummary.ProtoReflect.Descriptor instead.
func (*TemplateSummary) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateSummary) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TemplateSummary) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *TemplateSummary) GetPowertrain() string {
	if x != nil {
		return x.Powertrain
	}
	return ""
}

func (x *TemplateSummary) GetHasDbc() bool {
	if x != nil {
		return x.HasDbc
	}
	return false
}

func (x *TemplateSummary) GetPidsCount() int32 {
	if x != nil {
		return x.PidsCount
	}
	return 0
}

func (x *TemplateSummary) GetParentTemplateName() string {
	if x != nil {
		return x.ParentTemplateName
	}
	return ""
}

func (x *TemplateSummary) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PID) Reset() {
	*x = PID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PID) ProtoMessage() {}

func (x *PID) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PID.ProtoReflect.Descriptor instead.
func (*PID) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{2}
}

func (x *PID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTemplateListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Powertrain *string `protobuf:"bytes,1,opt,name=powertrain,proto3,oneof" json:"powertrain,omitempty"`
	Protocol   *string `protobuf:"bytes,2,opt,name=protocol,proto3,oneof" json:"protocol,omitempty"`
}

func (x *GetTemplateListRequest) Reset() {
	*x = GetTemplateListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateListRequest) ProtoMessage() {}

func (x *GetTemplateListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateListRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{3}
}

func (x *GetTemplateListRequest) GetPowertrain() string {
	if x != nil && x.Powertrain != nil {
		return *x.Powertrain
	}
	return ""
}

func (x *GetTemplateListRequest) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

type GetTemplateListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates []*TemplateSummary `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
}

func (x *GetTemplateListResponse) Reset() {
	*x = GetTemplateListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateListResponse) ProtoMessage() {}

func (x *GetTemplateListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateListResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemplateListResponse) GetTemplates() []*TemplateSummary {
	if x != nil {
		return x.Templates
	}
	return nil
}

type GetTemplateByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTemplateByNameRequest) Reset() {
	*x = GetTemplateByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateByNameRequest) ProtoMessage() {}

func (x *GetTemplateByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateByNameRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateByNameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetTemplateByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTemplateByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Pids     []*PID    `protobuf:"bytes,2,rep,name=pids,proto3" json:"pids,omitempty"`
}

func (x *GetTemplateByNameResponse) Reset() {
	*x = GetTemplateByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateByNameResponse) ProtoMessage() {}

func (x *GetTemplateByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateByNameResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateByNameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{6}
}

func (x *GetTemplateByNameResponse) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *GetTemplateByNameResponse) GetPids() []*PID {
	if x != nil {
		return x.Pids
	}
	return nil
}

type UpdateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Template *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *UpdateTemplateRequest) Reset() {
	*x = UpdateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_template_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateRequest) ProtoMessage() {}

func (x *UpdateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_template_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_template_config_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTemplateRequest) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

var File_pkg_grpc_template_config_proto protoreflect.FileDescriptor

var file_pkg_grpc_template_config_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x03, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x61, 0x73, 0x5f, 0x64, 0x62, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x44, 0x62, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x69, 0x64, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x69, 0x64, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x04,
	0x70, 0x69, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x62, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x62, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x5f, 0x64, 0x62, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x44, 0x62, 0x63, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x69, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x69, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x03, 0x50, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x7a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x4e, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22, 0x2e, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x49, 0x44, 0x52,
	0x04, 0x70, 0x69, 0x64, 0x73, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0xcb, 0x02, 0x0a, 0x15, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_template_config_proto_rawDescOnce sync.Once
	file_pkg_grpc_template_config_proto_rawDescData = file_pkg_grpc_template_config_proto_rawDesc
)

func file_pkg_grpc_template_config_proto_rawDescGZIP() []byte {
	file_pkg_grpc_template_config_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_template_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_template_config_proto_rawDescData)
	})
	return file_pkg_grpc_template_config_proto_rawDescData
}

var file_pkg_grpc_template_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_grpc_template_config_proto_goTypes = []interface{}{
	(*Template)(nil),                  // 0: grpc.Template
	(*TemplateSummary)(nil),           // 1: grpc.TemplateSummary
	(*PID)(nil),                       // 2: grpc.PID
	(*GetTemplateListRequest)(nil),    // 3: grpc.GetTemplateListRequest
	(*GetTemplateListResponse)(nil),   // 4: grpc.GetTemplateListResponse
	(*GetTemplateByNameRequest)(nil),  // 5: grpc.GetTemplateByNameRequest
	(*GetTemplateByNameResponse)(nil), // 6: grpc.GetTemplateByNameResponse
	(*UpdateTemplateRequest)(nil),     // 7: grpc.UpdateTemplateRequest
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_pkg_grpc_template_config_proto_depIdxs = []int32{
	2,  // 0: grpc.Template.pids:type_name -> grpc.PID
	8,  // 1: grpc.Template.created_at:type_name -> google.protobuf.Timestamp
	8,  // 2: grpc.Template.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 3: grpc.TemplateSummary.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 4: grpc.GetTemplateListResponse.templates:type_name -> grpc.TemplateSummary
	0,  // 5: grpc.GetTemplateByNameResponse.template:type_name -> grpc.Template
	2,  // 6: grpc.GetTemplateByNameResponse.pids:type_name -> grpc.PID
	0,  // 7: grpc.UpdateTemplateRequest.template:type_name -> grpc.Template
	7,  // 8: grpc.TemplateConfigService.CreateTemplate:input_type -> grpc.UpdateTemplateRequest
	7,  // 9: grpc.TemplateConfigService.UpdateTemplate:input_type -> grpc.UpdateTemplateRequest
	3,  // 10: grpc.TemplateConfigService.GetTemplateList:input_type -> grpc.GetTemplateListRequest
	5,  // 11: grpc.TemplateConfigService.GetTemplateByName:input_type -> grpc.GetTemplateByNameRequest
	9,  // 12: grpc.TemplateConfigService.CreateTemplate:output_type -> google.protobuf.Empty
	9,  // 13: grpc.TemplateConfigService.UpdateTemplate:output_type -> google.protobuf.Empty
	4,  // 14: grpc.TemplateConfigService.GetTemplateList:output_type -> grpc.GetTemplateListResponse
	6,  // 15: grpc.TemplateConfigService.GetTemplateByName:output_type -> grpc.GetTemplateByNameResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_grpc_template_config_proto_init() }
func file_pkg_grpc_template_config_proto_init() {
	if File_pkg_grpc_template_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_template_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSummary); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PID); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateListRequest); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateListResponse); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateByNameRequest); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateByNameResponse); i {
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
		file_pkg_grpc_template_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateRequest); i {
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
	file_pkg_grpc_template_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_grpc_template_config_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_template_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_template_config_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_template_config_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_template_config_proto_msgTypes,
	}.Build()
	File_pkg_grpc_template_config_proto = out.File
	file_pkg_grpc_template_config_proto_rawDesc = nil
	file_pkg_grpc_template_config_proto_goTypes = nil
	file_pkg_grpc_template_config_proto_depIdxs = nil
}
