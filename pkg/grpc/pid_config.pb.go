// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: pkg/grpc/pid_config.proto

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

// PidConfig is used for Admin gRPC requests to edit templates
type PidConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TemplateName         string                 `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	Header               []byte                 `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Mode                 []byte                 `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	Pid                  []byte                 `protobuf:"bytes,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Formula              string                 `protobuf:"bytes,6,opt,name=formula,proto3" json:"formula,omitempty"`
	IntervalSeconds      int32                  `protobuf:"varint,7,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	Protocol             *string                `protobuf:"bytes,8,opt,name=protocol,proto3,oneof" json:"protocol,omitempty"`
	SignalName           string                 `protobuf:"bytes,9,opt,name=signal_name,json=signalName,proto3" json:"signal_name,omitempty"`
	BytesReturned        int32                  `protobuf:"varint,10,opt,name=bytes_returned,json=bytesReturned,proto3" json:"bytes_returned,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CanFlowControlClear  *bool                  `protobuf:"varint,13,opt,name=can_flow_control_clear,json=canFlowControlClear,proto3,oneof" json:"can_flow_control_clear,omitempty"`
	CanFlowControlIdPair *string                `protobuf:"bytes,14,opt,name=can_flow_control_id_pair,json=canFlowControlIdPair,proto3,oneof" json:"can_flow_control_id_pair,omitempty"`
	Enabled              bool                   `protobuf:"varint,15,opt,name=enabled,proto3" json:"enabled,omitempty"`
	VssCovesaName        *string                `protobuf:"bytes,16,opt,name=vss_covesa_name,json=vssCovesaName,proto3,oneof" json:"vss_covesa_name,omitempty"`
}

func (x *PidConfig) Reset() {
	*x = PidConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PidConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PidConfig) ProtoMessage() {}

func (x *PidConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PidConfig.ProtoReflect.Descriptor instead.
func (*PidConfig) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{0}
}

func (x *PidConfig) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PidConfig) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *PidConfig) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PidConfig) GetMode() []byte {
	if x != nil {
		return x.Mode
	}
	return nil
}

func (x *PidConfig) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *PidConfig) GetFormula() string {
	if x != nil {
		return x.Formula
	}
	return ""
}

func (x *PidConfig) GetIntervalSeconds() int32 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

func (x *PidConfig) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *PidConfig) GetSignalName() string {
	if x != nil {
		return x.SignalName
	}
	return ""
}

func (x *PidConfig) GetBytesReturned() int32 {
	if x != nil {
		return x.BytesReturned
	}
	return 0
}

func (x *PidConfig) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PidConfig) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PidConfig) GetCanFlowControlClear() bool {
	if x != nil && x.CanFlowControlClear != nil {
		return *x.CanFlowControlClear
	}
	return false
}

func (x *PidConfig) GetCanFlowControlIdPair() string {
	if x != nil && x.CanFlowControlIdPair != nil {
		return *x.CanFlowControlIdPair
	}
	return ""
}

func (x *PidConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PidConfig) GetVssCovesaName() string {
	if x != nil && x.VssCovesaName != nil {
		return *x.VssCovesaName
	}
	return ""
}

type PidSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TemplateName         string `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	Header               []byte `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Mode                 []byte `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	Pid                  []byte `protobuf:"bytes,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Formula              string `protobuf:"bytes,6,opt,name=formula,proto3" json:"formula,omitempty"`
	IntervalSeconds      int32  `protobuf:"varint,7,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	Protocol             string `protobuf:"bytes,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SignalName           string `protobuf:"bytes,9,opt,name=signal_name,json=signalName,proto3" json:"signal_name,omitempty"`
	IsParentPid          bool   `protobuf:"varint,10,opt,name=is_parent_pid,json=isParentPid,proto3" json:"is_parent_pid,omitempty"`
	CanFlowControlClear  bool   `protobuf:"varint,11,opt,name=can_flow_control_clear,json=canFlowControlClear,proto3" json:"can_flow_control_clear,omitempty"`
	CanFlowControlIdPair string `protobuf:"bytes,12,opt,name=can_flow_control_id_pair,json=canFlowControlIdPair,proto3" json:"can_flow_control_id_pair,omitempty"`
	Enabled              bool   `protobuf:"varint,13,opt,name=enabled,proto3" json:"enabled,omitempty"`
	VssCovesaName        string `protobuf:"bytes,14,opt,name=vss_covesa_name,json=vssCovesaName,proto3" json:"vss_covesa_name,omitempty"`
}

func (x *PidSummary) Reset() {
	*x = PidSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PidSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PidSummary) ProtoMessage() {}

func (x *PidSummary) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PidSummary.ProtoReflect.Descriptor instead.
func (*PidSummary) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{1}
}

func (x *PidSummary) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PidSummary) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *PidSummary) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *PidSummary) GetMode() []byte {
	if x != nil {
		return x.Mode
	}
	return nil
}

func (x *PidSummary) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *PidSummary) GetFormula() string {
	if x != nil {
		return x.Formula
	}
	return ""
}

func (x *PidSummary) GetIntervalSeconds() int32 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

func (x *PidSummary) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PidSummary) GetSignalName() string {
	if x != nil {
		return x.SignalName
	}
	return ""
}

func (x *PidSummary) GetIsParentPid() bool {
	if x != nil {
		return x.IsParentPid
	}
	return false
}

func (x *PidSummary) GetCanFlowControlClear() bool {
	if x != nil {
		return x.CanFlowControlClear
	}
	return false
}

func (x *PidSummary) GetCanFlowControlIdPair() string {
	if x != nil {
		return x.CanFlowControlIdPair
	}
	return ""
}

func (x *PidSummary) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PidSummary) GetVssCovesaName() string {
	if x != nil {
		return x.VssCovesaName
	}
	return ""
}

type GetPidListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *GetPidListRequest) Reset() {
	*x = GetPidListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPidListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPidListRequest) ProtoMessage() {}

func (x *GetPidListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPidListRequest.ProtoReflect.Descriptor instead.
func (*GetPidListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{2}
}

func (x *GetPidListRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type GetPidListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid []*PidSummary `protobuf:"bytes,1,rep,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetPidListResponse) Reset() {
	*x = GetPidListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPidListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPidListResponse) ProtoMessage() {}

func (x *GetPidListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPidListResponse.ProtoReflect.Descriptor instead.
func (*GetPidListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{3}
}

func (x *GetPidListResponse) GetPid() []*PidSummary {
	if x != nil {
		return x.Pid
	}
	return nil
}

type GetPidByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPidByIDRequest) Reset() {
	*x = GetPidByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPidByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPidByIDRequest) ProtoMessage() {}

func (x *GetPidByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPidByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPidByIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetPidByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPidByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid *PidConfig `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetPidByIDResponse) Reset() {
	*x = GetPidByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPidByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPidByIDResponse) ProtoMessage() {}

func (x *GetPidByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPidByIDResponse.ProtoReflect.Descriptor instead.
func (*GetPidByIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetPidByIDResponse) GetPid() *PidConfig {
	if x != nil {
		return x.Pid
	}
	return nil
}

type UpdatePidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid *PidConfig `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *UpdatePidRequest) Reset() {
	*x = UpdatePidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePidRequest) ProtoMessage() {}

func (x *UpdatePidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePidRequest.ProtoReflect.Descriptor instead.
func (*UpdatePidRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePidRequest) GetPid() *PidConfig {
	if x != nil {
		return x.Pid
	}
	return nil
}

type DeletePidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TemplateName string `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
}

func (x *DeletePidRequest) Reset() {
	*x = DeletePidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePidRequest) ProtoMessage() {}

func (x *DeletePidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePidRequest.ProtoReflect.Descriptor instead.
func (*DeletePidRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePidRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeletePidRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type ChangePidEnableStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChangePidEnableStatusRequest) Reset() {
	*x = ChangePidEnableStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_pid_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePidEnableStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePidEnableStatusRequest) ProtoMessage() {}

func (x *ChangePidEnableStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_pid_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePidEnableStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangePidEnableStatusRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_pid_config_proto_rawDescGZIP(), []int{8}
}

func (x *ChangePidEnableStatusRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pkg_grpc_pid_config_proto protoreflect.FileDescriptor

var file_pkg_grpc_pid_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x69, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x05, 0x0a, 0x09, 0x50, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x16, 0x63, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x13, 0x63, 0x61, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b,
	0x0a, 0x18, 0x63, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x14, 0x63, 0x61, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x49, 0x64, 0x50, 0x61, 0x69, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x0f, 0x76, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x76,
	0x65, 0x73, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0d, 0x76, 0x73, 0x73, 0x43, 0x6f, 0x76, 0x65, 0x73, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x19, 0x0a, 0x17, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x63,
	0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x69, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x76, 0x73, 0x73, 0x5f,
	0x63, 0x6f, 0x76, 0x65, 0x73, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd4, 0x03, 0x0a, 0x0a,
	0x50, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x63, 0x61, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x36, 0x0a, 0x18, 0x63, 0x61, 0x6e, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x5f,
	0x70, 0x61, 0x69, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x6e, 0x46,
	0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x73,
	0x73, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x73, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x73, 0x73, 0x43, 0x6f, 0x76, 0x65, 0x73, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x69, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x69,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xa0, 0x03, 0x0a, 0x10, 0x50, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x53, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x69, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x69, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_pid_config_proto_rawDescOnce sync.Once
	file_pkg_grpc_pid_config_proto_rawDescData = file_pkg_grpc_pid_config_proto_rawDesc
)

func file_pkg_grpc_pid_config_proto_rawDescGZIP() []byte {
	file_pkg_grpc_pid_config_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_pid_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_pid_config_proto_rawDescData)
	})
	return file_pkg_grpc_pid_config_proto_rawDescData
}

var file_pkg_grpc_pid_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_grpc_pid_config_proto_goTypes = []interface{}{
	(*PidConfig)(nil),                    // 0: grpc.PidConfig
	(*PidSummary)(nil),                   // 1: grpc.PidSummary
	(*GetPidListRequest)(nil),            // 2: grpc.GetPidListRequest
	(*GetPidListResponse)(nil),           // 3: grpc.GetPidListResponse
	(*GetPidByIDRequest)(nil),            // 4: grpc.GetPidByIDRequest
	(*GetPidByIDResponse)(nil),           // 5: grpc.GetPidByIDResponse
	(*UpdatePidRequest)(nil),             // 6: grpc.UpdatePidRequest
	(*DeletePidRequest)(nil),             // 7: grpc.DeletePidRequest
	(*ChangePidEnableStatusRequest)(nil), // 8: grpc.ChangePidEnableStatusRequest
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_pkg_grpc_pid_config_proto_depIdxs = []int32{
	9,  // 0: grpc.PidConfig.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: grpc.PidConfig.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: grpc.GetPidListResponse.pid:type_name -> grpc.PidSummary
	0,  // 3: grpc.GetPidByIDResponse.pid:type_name -> grpc.PidConfig
	0,  // 4: grpc.UpdatePidRequest.pid:type_name -> grpc.PidConfig
	6,  // 5: grpc.PidConfigService.CreatePid:input_type -> grpc.UpdatePidRequest
	6,  // 6: grpc.PidConfigService.UpdatePid:input_type -> grpc.UpdatePidRequest
	2,  // 7: grpc.PidConfigService.GetPidList:input_type -> grpc.GetPidListRequest
	4,  // 8: grpc.PidConfigService.GetPidByID:input_type -> grpc.GetPidByIDRequest
	7,  // 9: grpc.PidConfigService.DeletePid:input_type -> grpc.DeletePidRequest
	8,  // 10: grpc.PidConfigService.ChangePidEnableStatus:input_type -> grpc.ChangePidEnableStatusRequest
	10, // 11: grpc.PidConfigService.CreatePid:output_type -> google.protobuf.Empty
	10, // 12: grpc.PidConfigService.UpdatePid:output_type -> google.protobuf.Empty
	3,  // 13: grpc.PidConfigService.GetPidList:output_type -> grpc.GetPidListResponse
	5,  // 14: grpc.PidConfigService.GetPidByID:output_type -> grpc.GetPidByIDResponse
	10, // 15: grpc.PidConfigService.DeletePid:output_type -> google.protobuf.Empty
	10, // 16: grpc.PidConfigService.ChangePidEnableStatus:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_grpc_pid_config_proto_init() }
func file_pkg_grpc_pid_config_proto_init() {
	if File_pkg_grpc_pid_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_pid_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PidConfig); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PidSummary); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPidListRequest); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPidListResponse); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPidByIDRequest); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPidByIDResponse); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePidRequest); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePidRequest); i {
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
		file_pkg_grpc_pid_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePidEnableStatusRequest); i {
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
	file_pkg_grpc_pid_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_pid_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_pid_config_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_pid_config_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_pid_config_proto_msgTypes,
	}.Build()
	File_pkg_grpc_pid_config_proto = out.File
	file_pkg_grpc_pid_config_proto_rawDesc = nil
	file_pkg_grpc_pid_config_proto_goTypes = nil
	file_pkg_grpc_pid_config_proto_depIdxs = nil
}
