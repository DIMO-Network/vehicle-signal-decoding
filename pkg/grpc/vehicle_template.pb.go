// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: pkg/grpc/vehicle_template.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VehicleTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YearStart int32    `protobuf:"varint,1,opt,name=YearStart,proto3" json:"YearStart,omitempty"`
	YearEnd   int32    `protobuf:"varint,2,opt,name=YearEnd,proto3" json:"YearEnd,omitempty"`
	Make      string   `protobuf:"bytes,3,opt,name=Make,proto3" json:"Make,omitempty"`
	Template  string   `protobuf:"bytes,4,opt,name=Template,proto3" json:"Template,omitempty"`
	Models    []string `protobuf:"bytes,5,rep,name=Models,proto3" json:"Models,omitempty"`
	Id        int64    `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VehicleTemplate) Reset() {
	*x = VehicleTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleTemplate) ProtoMessage() {}

func (x *VehicleTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleTemplate.ProtoReflect.Descriptor instead.
func (*VehicleTemplate) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_vehicle_template_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleTemplate) GetYearStart() int32 {
	if x != nil {
		return x.YearStart
	}
	return 0
}

func (x *VehicleTemplate) GetYearEnd() int32 {
	if x != nil {
		return x.YearEnd
	}
	return 0
}

func (x *VehicleTemplate) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *VehicleTemplate) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *VehicleTemplate) GetModels() []string {
	if x != nil {
		return x.Models
	}
	return nil
}

func (x *VehicleTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// used for listing items in admin
type GetVehicleTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YearStart int32  `protobuf:"varint,1,opt,name=YearStart,proto3" json:"YearStart,omitempty"`
	YearEnd   int32  `protobuf:"varint,2,opt,name=YearEnd,proto3" json:"YearEnd,omitempty"`
	Make      string `protobuf:"bytes,3,opt,name=Make,proto3" json:"Make,omitempty"`
	Template  string `protobuf:"bytes,4,opt,name=Template,proto3" json:"Template,omitempty"`
}

func (x *GetVehicleTemplatesRequest) Reset() {
	*x = GetVehicleTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleTemplatesRequest) ProtoMessage() {}

func (x *GetVehicleTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleTemplatesRequest.ProtoReflect.Descriptor instead.
func (*GetVehicleTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_vehicle_template_proto_rawDescGZIP(), []int{1}
}

func (x *GetVehicleTemplatesRequest) GetYearStart() int32 {
	if x != nil {
		return x.YearStart
	}
	return 0
}

func (x *GetVehicleTemplatesRequest) GetYearEnd() int32 {
	if x != nil {
		return x.YearEnd
	}
	return 0
}

func (x *GetVehicleTemplatesRequest) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *GetVehicleTemplatesRequest) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type GetVehicleTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates []*VehicleTemplate `protobuf:"bytes,1,rep,name=Templates,proto3" json:"Templates,omitempty"`
}

func (x *GetVehicleTemplatesResponse) Reset() {
	*x = GetVehicleTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleTemplatesResponse) ProtoMessage() {}

func (x *GetVehicleTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetVehicleTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_vehicle_template_proto_rawDescGZIP(), []int{2}
}

func (x *GetVehicleTemplatesResponse) GetTemplates() []*VehicleTemplate {
	if x != nil {
		return x.Templates
	}
	return nil
}

// used to view a specific vehicles to template setup
type GetVehicleTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVehicleTemplateRequest) Reset() {
	*x = GetVehicleTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleTemplateRequest) ProtoMessage() {}

func (x *GetVehicleTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetVehicleTemplateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_vehicle_template_proto_rawDescGZIP(), []int{3}
}

func (x *GetVehicleTemplateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateVehicleTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateVehicleTemplateResponse) Reset() {
	*x = CreateVehicleTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVehicleTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVehicleTemplateResponse) ProtoMessage() {}

func (x *CreateVehicleTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_vehicle_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVehicleTemplateResponse.ProtoReflect.Descriptor instead.
func (*CreateVehicleTemplateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_vehicle_template_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVehicleTemplateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pkg_grpc_vehicle_template_proto protoreflect.FileDescriptor

var file_pkg_grpc_vehicle_template_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x59, 0x65, 0x61, 0x72,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x59, 0x65, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4d, 0x61, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x59, 0x65, 0x61, 0x72, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x59, 0x65, 0x61, 0x72,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4d, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d,
	0x61, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22,
	0x52, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xe7, 0x02, 0x0a, 0x16, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_vehicle_template_proto_rawDescOnce sync.Once
	file_pkg_grpc_vehicle_template_proto_rawDescData = file_pkg_grpc_vehicle_template_proto_rawDesc
)

func file_pkg_grpc_vehicle_template_proto_rawDescGZIP() []byte {
	file_pkg_grpc_vehicle_template_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_vehicle_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_vehicle_template_proto_rawDescData)
	})
	return file_pkg_grpc_vehicle_template_proto_rawDescData
}

var file_pkg_grpc_vehicle_template_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_grpc_vehicle_template_proto_goTypes = []interface{}{
	(*VehicleTemplate)(nil),               // 0: grpc.VehicleTemplate
	(*GetVehicleTemplatesRequest)(nil),    // 1: grpc.GetVehicleTemplatesRequest
	(*GetVehicleTemplatesResponse)(nil),   // 2: grpc.GetVehicleTemplatesResponse
	(*GetVehicleTemplateRequest)(nil),     // 3: grpc.GetVehicleTemplateRequest
	(*CreateVehicleTemplateResponse)(nil), // 4: grpc.CreateVehicleTemplateResponse
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
}
var file_pkg_grpc_vehicle_template_proto_depIdxs = []int32{
	0, // 0: grpc.GetVehicleTemplatesResponse.Templates:type_name -> grpc.VehicleTemplate
	1, // 1: grpc.VehicleTemplateService.GetVehicleTemplates:input_type -> grpc.GetVehicleTemplatesRequest
	3, // 2: grpc.VehicleTemplateService.GetVehicleTemplate:input_type -> grpc.GetVehicleTemplateRequest
	0, // 3: grpc.VehicleTemplateService.CreateVehicleTemplate:input_type -> grpc.VehicleTemplate
	0, // 4: grpc.VehicleTemplateService.UpdateVehicleTemplate:input_type -> grpc.VehicleTemplate
	2, // 5: grpc.VehicleTemplateService.GetVehicleTemplates:output_type -> grpc.GetVehicleTemplatesResponse
	0, // 6: grpc.VehicleTemplateService.GetVehicleTemplate:output_type -> grpc.VehicleTemplate
	4, // 7: grpc.VehicleTemplateService.CreateVehicleTemplate:output_type -> grpc.CreateVehicleTemplateResponse
	5, // 8: grpc.VehicleTemplateService.UpdateVehicleTemplate:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpc_vehicle_template_proto_init() }
func file_pkg_grpc_vehicle_template_proto_init() {
	if File_pkg_grpc_vehicle_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_vehicle_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleTemplate); i {
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
		file_pkg_grpc_vehicle_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleTemplatesRequest); i {
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
		file_pkg_grpc_vehicle_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleTemplatesResponse); i {
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
		file_pkg_grpc_vehicle_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleTemplateRequest); i {
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
		file_pkg_grpc_vehicle_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVehicleTemplateResponse); i {
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
			RawDescriptor: file_pkg_grpc_vehicle_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_vehicle_template_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_vehicle_template_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_vehicle_template_proto_msgTypes,
	}.Build()
	File_pkg_grpc_vehicle_template_proto = out.File
	file_pkg_grpc_vehicle_template_proto_rawDesc = nil
	file_pkg_grpc_vehicle_template_proto_goTypes = nil
	file_pkg_grpc_vehicle_template_proto_depIdxs = nil
}
