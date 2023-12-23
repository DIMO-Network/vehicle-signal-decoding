// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/grpc/device_settings.proto

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

// used in Admin - todo rename this to be more admin specific and update admin after
type DeviceSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TemplateName string `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	Settings     string `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	Powertrain   string `protobuf:"bytes,4,opt,name=powertrain,proto3" json:"powertrain,omitempty"`
}

func (x *DeviceSettings) Reset() {
	*x = DeviceSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSettings) ProtoMessage() {}

func (x *DeviceSettings) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSettings.ProtoReflect.Descriptor instead.
func (*DeviceSettings) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceSettings) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *DeviceSettings) GetSettings() string {
	if x != nil {
		return x.Settings
	}
	return ""
}

func (x *DeviceSettings) GetPowertrain() string {
	if x != nil {
		return x.Powertrain
	}
	return ""
}

type GetDeviceSettingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSettings []*DeviceSettings `protobuf:"bytes,1,rep,name=device_settings,json=deviceSettings,proto3" json:"device_settings,omitempty"`
}

func (x *GetDeviceSettingListResponse) Reset() {
	*x = GetDeviceSettingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingListResponse) ProtoMessage() {}

func (x *GetDeviceSettingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceSettingListResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{1}
}

func (x *GetDeviceSettingListResponse) GetDeviceSettings() []*DeviceSettings {
	if x != nil {
		return x.DeviceSettings
	}
	return nil
}

type GetDeviceSettingByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDeviceSettingByNameRequest) Reset() {
	*x = GetDeviceSettingByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingByNameRequest) ProtoMessage() {}

func (x *GetDeviceSettingByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceSettingByNameRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingByNameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceSettingByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDeviceSettingByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSettings *DeviceSettings `protobuf:"bytes,1,opt,name=device_settings,json=deviceSettings,proto3" json:"device_settings,omitempty"`
}

func (x *GetDeviceSettingByNameResponse) Reset() {
	*x = GetDeviceSettingByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingByNameResponse) ProtoMessage() {}

func (x *GetDeviceSettingByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceSettingByNameResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingByNameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceSettingByNameResponse) GetDeviceSettings() *DeviceSettings {
	if x != nil {
		return x.DeviceSettings
	}
	return nil
}

type UpdateDeviceSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSettings *DeviceSettings `protobuf:"bytes,1,opt,name=device_settings,json=deviceSettings,proto3" json:"device_settings,omitempty"`
}

func (x *UpdateDeviceSettingsRequest) Reset() {
	*x = UpdateDeviceSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceSettingsRequest) ProtoMessage() {}

func (x *UpdateDeviceSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceSettingsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDeviceSettingsRequest) GetDeviceSettings() *DeviceSettings {
	if x != nil {
		return x.DeviceSettings
	}
	return nil
}

var File_pkg_grpc_device_settings_proto protoreflect.FileDescriptor

var file_pkg_grpc_device_settings_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x5d, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x33, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x5f, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x5c, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xf6,
	0x02, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x63, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_device_settings_proto_rawDescOnce sync.Once
	file_pkg_grpc_device_settings_proto_rawDescData = file_pkg_grpc_device_settings_proto_rawDesc
)

func file_pkg_grpc_device_settings_proto_rawDescGZIP() []byte {
	file_pkg_grpc_device_settings_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_device_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_device_settings_proto_rawDescData)
	})
	return file_pkg_grpc_device_settings_proto_rawDescData
}

var file_pkg_grpc_device_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_grpc_device_settings_proto_goTypes = []interface{}{
	(*DeviceSettings)(nil),                 // 0: grpc.DeviceSettings
	(*GetDeviceSettingListResponse)(nil),   // 1: grpc.GetDeviceSettingListResponse
	(*GetDeviceSettingByNameRequest)(nil),  // 2: grpc.GetDeviceSettingByNameRequest
	(*GetDeviceSettingByNameResponse)(nil), // 3: grpc.GetDeviceSettingByNameResponse
	(*UpdateDeviceSettingsRequest)(nil),    // 4: grpc.UpdateDeviceSettingsRequest
	(*emptypb.Empty)(nil),                  // 5: google.protobuf.Empty
}
var file_pkg_grpc_device_settings_proto_depIdxs = []int32{
	0, // 0: grpc.GetDeviceSettingListResponse.device_settings:type_name -> grpc.DeviceSettings
	0, // 1: grpc.GetDeviceSettingByNameResponse.device_settings:type_name -> grpc.DeviceSettings
	0, // 2: grpc.UpdateDeviceSettingsRequest.device_settings:type_name -> grpc.DeviceSettings
	4, // 3: grpc.DeviceSettingsService.CreateDeviceSettings:input_type -> grpc.UpdateDeviceSettingsRequest
	4, // 4: grpc.DeviceSettingsService.UpdateDeviceSettings:input_type -> grpc.UpdateDeviceSettingsRequest
	5, // 5: grpc.DeviceSettingsService.GetDeviceSettingList:input_type -> google.protobuf.Empty
	2, // 6: grpc.DeviceSettingsService.GetDeviceSettingByName:input_type -> grpc.GetDeviceSettingByNameRequest
	5, // 7: grpc.DeviceSettingsService.CreateDeviceSettings:output_type -> google.protobuf.Empty
	5, // 8: grpc.DeviceSettingsService.UpdateDeviceSettings:output_type -> google.protobuf.Empty
	1, // 9: grpc.DeviceSettingsService.GetDeviceSettingList:output_type -> grpc.GetDeviceSettingListResponse
	3, // 10: grpc.DeviceSettingsService.GetDeviceSettingByName:output_type -> grpc.GetDeviceSettingByNameResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_grpc_device_settings_proto_init() }
func file_pkg_grpc_device_settings_proto_init() {
	if File_pkg_grpc_device_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_device_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceSettings); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceSettingListResponse); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceSettingByNameRequest); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceSettingByNameResponse); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceSettingsRequest); i {
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
			RawDescriptor: file_pkg_grpc_device_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_device_settings_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_device_settings_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_device_settings_proto_msgTypes,
	}.Build()
	File_pkg_grpc_device_settings_proto = out.File
	file_pkg_grpc_device_settings_proto_rawDesc = nil
	file_pkg_grpc_device_settings_proto_goTypes = nil
	file_pkg_grpc_device_settings_proto_depIdxs = nil
}
