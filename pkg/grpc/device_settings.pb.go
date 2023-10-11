// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pkg/grpc/device_settings.proto

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

type DeviceSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                                     int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TemplateName                           string                 `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	BatteryCriticalLevelVoltage            float64                `protobuf:"fixed64,3,opt,name=battery_critical_level_voltage,json=batteryCriticalLevelVoltage,proto3" json:"battery_critical_level_voltage,omitempty"`
	SafetyCutOutVoltage                    float64                `protobuf:"fixed64,4,opt,name=safety_cut_out_voltage,json=safetyCutOutVoltage,proto3" json:"safety_cut_out_voltage,omitempty"`
	SleepTimerEventDrivenInterval          float64                `protobuf:"fixed64,5,opt,name=sleep_timer_event_driven_interval,json=sleepTimerEventDrivenInterval,proto3" json:"sleep_timer_event_driven_interval,omitempty"`
	SleepTimerEventDrivenPeriod            float64                `protobuf:"fixed64,6,opt,name=sleep_timer_event_driven_period,json=sleepTimerEventDrivenPeriod,proto3" json:"sleep_timer_event_driven_period,omitempty"`
	SleepTimerInactivityAfterSleepInterval float64                `protobuf:"fixed64,7,opt,name=sleep_timer_inactivity_after_sleep_interval,json=sleepTimerInactivityAfterSleepInterval,proto3" json:"sleep_timer_inactivity_after_sleep_interval,omitempty"`
	SleepTimerInactivityFallbackInterval   float64                `protobuf:"fixed64,8,opt,name=sleep_timer_inactivity_fallback_interval,json=sleepTimerInactivityFallbackInterval,proto3" json:"sleep_timer_inactivity_fallback_interval,omitempty"`
	WakeTriggerVoltageLevel                float64                `protobuf:"fixed64,9,opt,name=wake_trigger_voltage_level,json=wakeTriggerVoltageLevel,proto3" json:"wake_trigger_voltage_level,omitempty"`
	CreatedAt                              *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                              *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (x *DeviceSettings) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceSettings) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *DeviceSettings) GetBatteryCriticalLevelVoltage() float64 {
	if x != nil {
		return x.BatteryCriticalLevelVoltage
	}
	return 0
}

func (x *DeviceSettings) GetSafetyCutOutVoltage() float64 {
	if x != nil {
		return x.SafetyCutOutVoltage
	}
	return 0
}

func (x *DeviceSettings) GetSleepTimerEventDrivenInterval() float64 {
	if x != nil {
		return x.SleepTimerEventDrivenInterval
	}
	return 0
}

func (x *DeviceSettings) GetSleepTimerEventDrivenPeriod() float64 {
	if x != nil {
		return x.SleepTimerEventDrivenPeriod
	}
	return 0
}

func (x *DeviceSettings) GetSleepTimerInactivityAfterSleepInterval() float64 {
	if x != nil {
		return x.SleepTimerInactivityAfterSleepInterval
	}
	return 0
}

func (x *DeviceSettings) GetSleepTimerInactivityFallbackInterval() float64 {
	if x != nil {
		return x.SleepTimerInactivityFallbackInterval
	}
	return 0
}

func (x *DeviceSettings) GetWakeTriggerVoltageLevel() float64 {
	if x != nil {
		return x.WakeTriggerVoltageLevel
	}
	return 0
}

func (x *DeviceSettings) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeviceSettings) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DeviceSettingsSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TemplateName string `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *DeviceSettingsSummary) Reset() {
	*x = DeviceSettingsSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSettingsSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSettingsSummary) ProtoMessage() {}

func (x *DeviceSettingsSummary) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeviceSettingsSummary.ProtoReflect.Descriptor instead.
func (*DeviceSettingsSummary) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceSettingsSummary) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceSettingsSummary) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type GetDeviceSettingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName *string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3,oneof" json:"template_name,omitempty"`
}

func (x *GetDeviceSettingListRequest) Reset() {
	*x = GetDeviceSettingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingListRequest) ProtoMessage() {}

func (x *GetDeviceSettingListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetDeviceSettingListRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceSettingListRequest) GetTemplateName() string {
	if x != nil && x.TemplateName != nil {
		return *x.TemplateName
	}
	return ""
}

type GetDeviceSettingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSettings []*DeviceSettingsSummary `protobuf:"bytes,1,rep,name=device_settings,json=deviceSettings,proto3" json:"device_settings,omitempty"`
}

func (x *GetDeviceSettingListResponse) Reset() {
	*x = GetDeviceSettingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingListResponse) ProtoMessage() {}

func (x *GetDeviceSettingListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetDeviceSettingListResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceSettingListResponse) GetDeviceSettings() []*DeviceSettingsSummary {
	if x != nil {
		return x.DeviceSettings
	}
	return nil
}

type GetDeviceSettingByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeviceSettingByIDRequest) Reset() {
	*x = GetDeviceSettingByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingByIDRequest) ProtoMessage() {}

func (x *GetDeviceSettingByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetDeviceSettingByIDRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingByIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceSettingByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDeviceSettingByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSettings *DeviceSettings `protobuf:"bytes,1,opt,name=device_settings,json=deviceSettings,proto3" json:"device_settings,omitempty"`
}

func (x *GetDeviceSettingByIDResponse) Reset() {
	*x = GetDeviceSettingByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceSettingByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceSettingByIDResponse) ProtoMessage() {}

func (x *GetDeviceSettingByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceSettingByIDResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceSettingByIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeviceSettingByIDResponse) GetDeviceSettings() *DeviceSettings {
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
		mi := &file_pkg_grpc_device_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceSettingsRequest) ProtoMessage() {}

func (x *UpdateDeviceSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_device_settings_proto_msgTypes[6]
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
	return file_pkg_grpc_device_settings_proto_rawDescGZIP(), []int{6}
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
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x05, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x1e,
	0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x1b, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x63, 0x75, 0x74, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x13, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x43, 0x75, 0x74, 0x4f, 0x75, 0x74, 0x56,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x21, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x1d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x44, 0x0a, 0x1f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1b, 0x73, 0x6c, 0x65, 0x65, 0x70,
	0x54, 0x69, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x5b, 0x0a, 0x2b, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x26, 0x73, 0x6c, 0x65,
	0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x56, 0x0a, 0x28, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x24, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x77,
	0x61, 0x6b, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x17, 0x77, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4c,
	0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x2d, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x5c, 0x0a, 0x1b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x15, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52,
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

var file_pkg_grpc_device_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_grpc_device_settings_proto_goTypes = []interface{}{
	(*DeviceSettings)(nil),               // 0: grpc.DeviceSettings
	(*DeviceSettingsSummary)(nil),        // 1: grpc.DeviceSettingsSummary
	(*GetDeviceSettingListRequest)(nil),  // 2: grpc.GetDeviceSettingListRequest
	(*GetDeviceSettingListResponse)(nil), // 3: grpc.GetDeviceSettingListResponse
	(*GetDeviceSettingByIDRequest)(nil),  // 4: grpc.GetDeviceSettingByIDRequest
	(*GetDeviceSettingByIDResponse)(nil), // 5: grpc.GetDeviceSettingByIDResponse
	(*UpdateDeviceSettingsRequest)(nil),  // 6: grpc.UpdateDeviceSettingsRequest
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                // 8: google.protobuf.Empty
}
var file_pkg_grpc_device_settings_proto_depIdxs = []int32{
	7, // 0: grpc.DeviceSettings.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: grpc.DeviceSettings.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: grpc.GetDeviceSettingListResponse.device_settings:type_name -> grpc.DeviceSettingsSummary
	0, // 3: grpc.GetDeviceSettingByIDResponse.device_settings:type_name -> grpc.DeviceSettings
	0, // 4: grpc.UpdateDeviceSettingsRequest.device_settings:type_name -> grpc.DeviceSettings
	6, // 5: grpc.DeviceSettingsService.CreateDeviceSettings:input_type -> grpc.UpdateDeviceSettingsRequest
	6, // 6: grpc.DeviceSettingsService.UpdateDeviceSettings:input_type -> grpc.UpdateDeviceSettingsRequest
	2, // 7: grpc.DeviceSettingsService.GetDeviceSettingList:input_type -> grpc.GetDeviceSettingListRequest
	4, // 8: grpc.DeviceSettingsService.GetDeviceSettingByID:input_type -> grpc.GetDeviceSettingByIDRequest
	8, // 9: grpc.DeviceSettingsService.CreateDeviceSettings:output_type -> google.protobuf.Empty
	8, // 10: grpc.DeviceSettingsService.UpdateDeviceSettings:output_type -> google.protobuf.Empty
	3, // 11: grpc.DeviceSettingsService.GetDeviceSettingList:output_type -> grpc.GetDeviceSettingListResponse
	5, // 12: grpc.DeviceSettingsService.GetDeviceSettingByID:output_type -> grpc.GetDeviceSettingByIDResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			switch v := v.(*DeviceSettingsSummary); i {
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
			switch v := v.(*GetDeviceSettingListRequest); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceSettingByIDRequest); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceSettingByIDResponse); i {
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
		file_pkg_grpc_device_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	file_pkg_grpc_device_settings_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_device_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
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
