// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: pkg/grpc/hardware_v1.proto

package grpc

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

// mostly power settings
type DeviceSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName                             string  `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	BatteryCriticalLevelVoltage              float32 `protobuf:"fixed32,2,opt,name=battery_critical_level_voltage,json=batteryCriticalLevelVoltage,proto3" json:"battery_critical_level_voltage,omitempty"`
	SafetyCutOutVoltage                      float32 `protobuf:"fixed32,3,opt,name=safety_cut_out_voltage,json=safetyCutOutVoltage,proto3" json:"safety_cut_out_voltage,omitempty"`
	SleepTimerEventDrivenIntervalSecs        float32 `protobuf:"fixed32,4,opt,name=sleep_timer_event_driven_interval_secs,json=sleepTimerEventDrivenIntervalSecs,proto3" json:"sleep_timer_event_driven_interval_secs,omitempty"`
	SleepTimerEventDrivenPeriodSecs          float32 `protobuf:"fixed32,5,opt,name=sleep_timer_event_driven_period_secs,json=sleepTimerEventDrivenPeriodSecs,proto3" json:"sleep_timer_event_driven_period_secs,omitempty"`
	SleepTimerInactivityAfterSleepSecs       float32 `protobuf:"fixed32,6,opt,name=sleep_timer_inactivity_after_sleep_secs,json=sleepTimerInactivityAfterSleepSecs,proto3" json:"sleep_timer_inactivity_after_sleep_secs,omitempty"`
	SleepTimerInactivityFallbackIntervalSecs float32 `protobuf:"fixed32,7,opt,name=sleep_timer_inactivity_fallback_interval_secs,json=sleepTimerInactivityFallbackIntervalSecs,proto3" json:"sleep_timer_inactivity_fallback_interval_secs,omitempty"`
	WakeTriggerVoltageLevel                  float32 `protobuf:"fixed32,8,opt,name=wake_trigger_voltage_level,json=wakeTriggerVoltageLevel,proto3" json:"wake_trigger_voltage_level,omitempty"`
	MinVoltageObdLoggers                     float32 `protobuf:"fixed32,9,opt,name=min_voltage_obd_loggers,json=minVoltageObdLoggers,proto3" json:"min_voltage_obd_loggers,omitempty"`
	LocationFrequencySecs                    float32 `protobuf:"fixed32,10,opt,name=location_frequency_secs,json=locationFrequencySecs,proto3" json:"location_frequency_secs,omitempty"`
}

func (x *DeviceSetting) Reset() {
	*x = DeviceSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSetting) ProtoMessage() {}

func (x *DeviceSetting) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSetting.ProtoReflect.Descriptor instead.
func (*DeviceSetting) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_hardware_v1_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceSetting) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *DeviceSetting) GetBatteryCriticalLevelVoltage() float32 {
	if x != nil {
		return x.BatteryCriticalLevelVoltage
	}
	return 0
}

func (x *DeviceSetting) GetSafetyCutOutVoltage() float32 {
	if x != nil {
		return x.SafetyCutOutVoltage
	}
	return 0
}

func (x *DeviceSetting) GetSleepTimerEventDrivenIntervalSecs() float32 {
	if x != nil {
		return x.SleepTimerEventDrivenIntervalSecs
	}
	return 0
}

func (x *DeviceSetting) GetSleepTimerEventDrivenPeriodSecs() float32 {
	if x != nil {
		return x.SleepTimerEventDrivenPeriodSecs
	}
	return 0
}

func (x *DeviceSetting) GetSleepTimerInactivityAfterSleepSecs() float32 {
	if x != nil {
		return x.SleepTimerInactivityAfterSleepSecs
	}
	return 0
}

func (x *DeviceSetting) GetSleepTimerInactivityFallbackIntervalSecs() float32 {
	if x != nil {
		return x.SleepTimerInactivityFallbackIntervalSecs
	}
	return 0
}

func (x *DeviceSetting) GetWakeTriggerVoltageLevel() float32 {
	if x != nil {
		return x.WakeTriggerVoltageLevel
	}
	return 0
}

func (x *DeviceSetting) GetMinVoltageObdLoggers() float32 {
	if x != nil {
		return x.MinVoltageObdLoggers
	}
	return 0
}

func (x *DeviceSetting) GetLocationFrequencySecs() float32 {
	if x != nil {
		return x.LocationFrequencySecs
	}
	return 0
}

// PIDs configuration for device to query obd
type PIDConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Header              uint32 `protobuf:"varint,2,opt,name=header,proto3" json:"header,omitempty"`
	Mode                uint32 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Pid                 uint32 `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	Formula             string `protobuf:"bytes,5,opt,name=formula,proto3" json:"formula,omitempty"`
	IntervalSeconds     uint32 `protobuf:"varint,6,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	Protocol            string `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	CanFlowControlClear bool   `protobuf:"varint,8,opt,name=can_flow_control_clear,json=canFlowControlClear,proto3" json:"can_flow_control_clear,omitempty"`
	// can be used to specify a custom response header. format is hex,hex - second hex is the resp header
	CanFlowControlIdPair string `protobuf:"bytes,9,opt,name=can_flow_control_id_pair,json=canFlowControlIdPair,proto3" json:"can_flow_control_id_pair,omitempty"`
}

func (x *PIDConfig) Reset() {
	*x = PIDConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIDConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIDConfig) ProtoMessage() {}

func (x *PIDConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIDConfig.ProtoReflect.Descriptor instead.
func (*PIDConfig) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_hardware_v1_proto_rawDescGZIP(), []int{1}
}

func (x *PIDConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PIDConfig) GetHeader() uint32 {
	if x != nil {
		return x.Header
	}
	return 0
}

func (x *PIDConfig) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *PIDConfig) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *PIDConfig) GetFormula() string {
	if x != nil {
		return x.Formula
	}
	return ""
}

func (x *PIDConfig) GetIntervalSeconds() uint32 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

func (x *PIDConfig) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PIDConfig) GetCanFlowControlClear() bool {
	if x != nil {
		return x.CanFlowControlClear
	}
	return false
}

func (x *PIDConfig) GetCanFlowControlIdPair() string {
	if x != nil {
		return x.CanFlowControlIdPair
	}
	return ""
}

type PIDRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string       `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	Version      string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Requests     []*PIDConfig `protobuf:"bytes,3,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *PIDRequests) Reset() {
	*x = PIDRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIDRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIDRequests) ProtoMessage() {}

func (x *PIDRequests) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_hardware_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIDRequests.ProtoReflect.Descriptor instead.
func (*PIDRequests) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_hardware_v1_proto_rawDescGZIP(), []int{2}
}

func (x *PIDRequests) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *PIDRequests) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PIDRequests) GetRequests() []*PIDConfig {
	if x != nil {
		return x.Requests
	}
	return nil
}

var File_pkg_grpc_hardware_v1_proto protoreflect.FileDescriptor

var file_pkg_grpc_hardware_v1_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0xb2, 0x05, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x1e, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x1b, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x16, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x63, 0x75, 0x74, 0x5f, 0x6f, 0x75, 0x74,
	0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13,
	0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x43, 0x75, 0x74, 0x4f, 0x75, 0x74, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x26, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x21, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x53, 0x65, 0x63, 0x73, 0x12, 0x4d, 0x0a, 0x24, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x1f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x53, 0x65, 0x63, 0x73, 0x12, 0x53, 0x0a, 0x27, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x22, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x53, 0x65, 0x63, 0x73, 0x12, 0x5f, 0x0a, 0x2d, 0x73, 0x6c,
	0x65, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x28, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x77,
	0x61, 0x6b, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x17, 0x77, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x5f,
	0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x62, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x6d, 0x69, 0x6e, 0x56, 0x6f,
	0x6c, 0x74, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12,
	0x36, 0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x53, 0x65, 0x63, 0x73, 0x22, 0xab, 0x02, 0x0a, 0x09, 0x50, 0x49, 0x44, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75,
	0x6c, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c,
	0x61, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x61, 0x6e, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x63, 0x61, 0x6e, 0x46, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x36, 0x0a,
	0x18, 0x63, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x63, 0x61, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49,
	0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0x79, 0x0a, 0x0b, 0x50, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x49, 0x44,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_hardware_v1_proto_rawDescOnce sync.Once
	file_pkg_grpc_hardware_v1_proto_rawDescData = file_pkg_grpc_hardware_v1_proto_rawDesc
)

func file_pkg_grpc_hardware_v1_proto_rawDescGZIP() []byte {
	file_pkg_grpc_hardware_v1_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_hardware_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_hardware_v1_proto_rawDescData)
	})
	return file_pkg_grpc_hardware_v1_proto_rawDescData
}

var file_pkg_grpc_hardware_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_grpc_hardware_v1_proto_goTypes = []interface{}{
	(*DeviceSetting)(nil), // 0: grpc.DeviceSetting
	(*PIDConfig)(nil),     // 1: grpc.PIDConfig
	(*PIDRequests)(nil),   // 2: grpc.PIDRequests
}
var file_pkg_grpc_hardware_v1_proto_depIdxs = []int32{
	1, // 0: grpc.PIDRequests.requests:type_name -> grpc.PIDConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpc_hardware_v1_proto_init() }
func file_pkg_grpc_hardware_v1_proto_init() {
	if File_pkg_grpc_hardware_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_hardware_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceSetting); i {
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
		file_pkg_grpc_hardware_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIDConfig); i {
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
		file_pkg_grpc_hardware_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIDRequests); i {
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
			RawDescriptor: file_pkg_grpc_hardware_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_grpc_hardware_v1_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_hardware_v1_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_hardware_v1_proto_msgTypes,
	}.Build()
	File_pkg_grpc_hardware_v1_proto = out.File
	file_pkg_grpc_hardware_v1_proto_rawDesc = nil
	file_pkg_grpc_hardware_v1_proto_goTypes = nil
	file_pkg_grpc_hardware_v1_proto_depIdxs = nil
}
