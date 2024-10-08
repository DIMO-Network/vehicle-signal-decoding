// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: pkg/grpc/pid_config.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PidConfigService_CreatePid_FullMethodName             = "/grpc.PidConfigService/CreatePid"
	PidConfigService_UpdatePid_FullMethodName             = "/grpc.PidConfigService/UpdatePid"
	PidConfigService_GetPidList_FullMethodName            = "/grpc.PidConfigService/GetPidList"
	PidConfigService_GetPidByID_FullMethodName            = "/grpc.PidConfigService/GetPidByID"
	PidConfigService_DeletePid_FullMethodName             = "/grpc.PidConfigService/DeletePid"
	PidConfigService_ChangePidEnableStatus_FullMethodName = "/grpc.PidConfigService/ChangePidEnableStatus"
	PidConfigService_GetSignalNames_FullMethodName        = "/grpc.PidConfigService/GetSignalNames"
)

// PidConfigServiceClient is the client API for PidConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PidConfigServiceClient interface {
	CreatePid(ctx context.Context, in *UpdatePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdatePid(ctx context.Context, in *UpdatePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPidList(ctx context.Context, in *GetPidListRequest, opts ...grpc.CallOption) (*GetPidListResponse, error)
	GetPidByID(ctx context.Context, in *GetPidByIDRequest, opts ...grpc.CallOption) (*GetPidByIDResponse, error)
	DeletePid(ctx context.Context, in *DeletePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangePidEnableStatus(ctx context.Context, in *ChangePidEnableStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSignalNames(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SignalNames, error)
}

type pidConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPidConfigServiceClient(cc grpc.ClientConnInterface) PidConfigServiceClient {
	return &pidConfigServiceClient{cc}
}

func (c *pidConfigServiceClient) CreatePid(ctx context.Context, in *UpdatePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PidConfigService_CreatePid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) UpdatePid(ctx context.Context, in *UpdatePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PidConfigService_UpdatePid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) GetPidList(ctx context.Context, in *GetPidListRequest, opts ...grpc.CallOption) (*GetPidListResponse, error) {
	out := new(GetPidListResponse)
	err := c.cc.Invoke(ctx, PidConfigService_GetPidList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) GetPidByID(ctx context.Context, in *GetPidByIDRequest, opts ...grpc.CallOption) (*GetPidByIDResponse, error) {
	out := new(GetPidByIDResponse)
	err := c.cc.Invoke(ctx, PidConfigService_GetPidByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) DeletePid(ctx context.Context, in *DeletePidRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PidConfigService_DeletePid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) ChangePidEnableStatus(ctx context.Context, in *ChangePidEnableStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PidConfigService_ChangePidEnableStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pidConfigServiceClient) GetSignalNames(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SignalNames, error) {
	out := new(SignalNames)
	err := c.cc.Invoke(ctx, PidConfigService_GetSignalNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PidConfigServiceServer is the server API for PidConfigService service.
// All implementations must embed UnimplementedPidConfigServiceServer
// for forward compatibility
type PidConfigServiceServer interface {
	CreatePid(context.Context, *UpdatePidRequest) (*emptypb.Empty, error)
	UpdatePid(context.Context, *UpdatePidRequest) (*emptypb.Empty, error)
	GetPidList(context.Context, *GetPidListRequest) (*GetPidListResponse, error)
	GetPidByID(context.Context, *GetPidByIDRequest) (*GetPidByIDResponse, error)
	DeletePid(context.Context, *DeletePidRequest) (*emptypb.Empty, error)
	ChangePidEnableStatus(context.Context, *ChangePidEnableStatusRequest) (*emptypb.Empty, error)
	GetSignalNames(context.Context, *emptypb.Empty) (*SignalNames, error)
	mustEmbedUnimplementedPidConfigServiceServer()
}

// UnimplementedPidConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPidConfigServiceServer struct {
}

func (UnimplementedPidConfigServiceServer) CreatePid(context.Context, *UpdatePidRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePid not implemented")
}
func (UnimplementedPidConfigServiceServer) UpdatePid(context.Context, *UpdatePidRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePid not implemented")
}
func (UnimplementedPidConfigServiceServer) GetPidList(context.Context, *GetPidListRequest) (*GetPidListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPidList not implemented")
}
func (UnimplementedPidConfigServiceServer) GetPidByID(context.Context, *GetPidByIDRequest) (*GetPidByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPidByID not implemented")
}
func (UnimplementedPidConfigServiceServer) DeletePid(context.Context, *DeletePidRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePid not implemented")
}
func (UnimplementedPidConfigServiceServer) ChangePidEnableStatus(context.Context, *ChangePidEnableStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePidEnableStatus not implemented")
}
func (UnimplementedPidConfigServiceServer) GetSignalNames(context.Context, *emptypb.Empty) (*SignalNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalNames not implemented")
}
func (UnimplementedPidConfigServiceServer) mustEmbedUnimplementedPidConfigServiceServer() {}

// UnsafePidConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PidConfigServiceServer will
// result in compilation errors.
type UnsafePidConfigServiceServer interface {
	mustEmbedUnimplementedPidConfigServiceServer()
}

func RegisterPidConfigServiceServer(s grpc.ServiceRegistrar, srv PidConfigServiceServer) {
	s.RegisterService(&PidConfigService_ServiceDesc, srv)
}

func _PidConfigService_CreatePid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).CreatePid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_CreatePid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).CreatePid(ctx, req.(*UpdatePidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_UpdatePid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).UpdatePid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_UpdatePid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).UpdatePid(ctx, req.(*UpdatePidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_GetPidList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPidListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).GetPidList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_GetPidList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).GetPidList(ctx, req.(*GetPidListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_GetPidByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPidByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).GetPidByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_GetPidByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).GetPidByID(ctx, req.(*GetPidByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_DeletePid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).DeletePid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_DeletePid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).DeletePid(ctx, req.(*DeletePidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_ChangePidEnableStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePidEnableStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).ChangePidEnableStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_ChangePidEnableStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).ChangePidEnableStatus(ctx, req.(*ChangePidEnableStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PidConfigService_GetSignalNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidConfigServiceServer).GetSignalNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PidConfigService_GetSignalNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidConfigServiceServer).GetSignalNames(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PidConfigService_ServiceDesc is the grpc.ServiceDesc for PidConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PidConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PidConfigService",
	HandlerType: (*PidConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePid",
			Handler:    _PidConfigService_CreatePid_Handler,
		},
		{
			MethodName: "UpdatePid",
			Handler:    _PidConfigService_UpdatePid_Handler,
		},
		{
			MethodName: "GetPidList",
			Handler:    _PidConfigService_GetPidList_Handler,
		},
		{
			MethodName: "GetPidByID",
			Handler:    _PidConfigService_GetPidByID_Handler,
		},
		{
			MethodName: "DeletePid",
			Handler:    _PidConfigService_DeletePid_Handler,
		},
		{
			MethodName: "ChangePidEnableStatus",
			Handler:    _PidConfigService_ChangePidEnableStatus_Handler,
		},
		{
			MethodName: "GetSignalNames",
			Handler:    _PidConfigService_GetSignalNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/pid_config.proto",
}
