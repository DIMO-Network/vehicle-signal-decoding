// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pkg/grpc/dbc_config.proto

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
	DbcConfigService_CreateDbc_FullMethodName            = "/grpc.DbcConfigService/CreateDbc"
	DbcConfigService_UpdateDbc_FullMethodName            = "/grpc.DbcConfigService/UpdateDbc"
	DbcConfigService_GetDbcList_FullMethodName           = "/grpc.DbcConfigService/GetDbcList"
	DbcConfigService_GetDbcByTemplateName_FullMethodName = "/grpc.DbcConfigService/GetDbcByTemplateName"
)

// DbcConfigServiceClient is the client API for DbcConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbcConfigServiceClient interface {
	CreateDbc(ctx context.Context, in *UpdateDbcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateDbc(ctx context.Context, in *UpdateDbcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDbcList(ctx context.Context, in *GetDbcListRequest, opts ...grpc.CallOption) (*GetDbcListResponse, error)
	GetDbcByTemplateName(ctx context.Context, in *GetDbcByTemplateNameRequest, opts ...grpc.CallOption) (*GetDbcByTemplateNameResponse, error)
}

type dbcConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDbcConfigServiceClient(cc grpc.ClientConnInterface) DbcConfigServiceClient {
	return &dbcConfigServiceClient{cc}
}

func (c *dbcConfigServiceClient) CreateDbc(ctx context.Context, in *UpdateDbcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DbcConfigService_CreateDbc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbcConfigServiceClient) UpdateDbc(ctx context.Context, in *UpdateDbcRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DbcConfigService_UpdateDbc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbcConfigServiceClient) GetDbcList(ctx context.Context, in *GetDbcListRequest, opts ...grpc.CallOption) (*GetDbcListResponse, error) {
	out := new(GetDbcListResponse)
	err := c.cc.Invoke(ctx, DbcConfigService_GetDbcList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbcConfigServiceClient) GetDbcByTemplateName(ctx context.Context, in *GetDbcByTemplateNameRequest, opts ...grpc.CallOption) (*GetDbcByTemplateNameResponse, error) {
	out := new(GetDbcByTemplateNameResponse)
	err := c.cc.Invoke(ctx, DbcConfigService_GetDbcByTemplateName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbcConfigServiceServer is the server API for DbcConfigService service.
// All implementations must embed UnimplementedDbcConfigServiceServer
// for forward compatibility
type DbcConfigServiceServer interface {
	CreateDbc(context.Context, *UpdateDbcRequest) (*emptypb.Empty, error)
	UpdateDbc(context.Context, *UpdateDbcRequest) (*emptypb.Empty, error)
	GetDbcList(context.Context, *GetDbcListRequest) (*GetDbcListResponse, error)
	GetDbcByTemplateName(context.Context, *GetDbcByTemplateNameRequest) (*GetDbcByTemplateNameResponse, error)
	mustEmbedUnimplementedDbcConfigServiceServer()
}

// UnimplementedDbcConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDbcConfigServiceServer struct {
}

func (UnimplementedDbcConfigServiceServer) CreateDbc(context.Context, *UpdateDbcRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDbc not implemented")
}
func (UnimplementedDbcConfigServiceServer) UpdateDbc(context.Context, *UpdateDbcRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDbc not implemented")
}
func (UnimplementedDbcConfigServiceServer) GetDbcList(context.Context, *GetDbcListRequest) (*GetDbcListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDbcList not implemented")
}
func (UnimplementedDbcConfigServiceServer) GetDbcByTemplateName(context.Context, *GetDbcByTemplateNameRequest) (*GetDbcByTemplateNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDbcByTemplateName not implemented")
}
func (UnimplementedDbcConfigServiceServer) mustEmbedUnimplementedDbcConfigServiceServer() {}

// UnsafeDbcConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbcConfigServiceServer will
// result in compilation errors.
type UnsafeDbcConfigServiceServer interface {
	mustEmbedUnimplementedDbcConfigServiceServer()
}

func RegisterDbcConfigServiceServer(s grpc.ServiceRegistrar, srv DbcConfigServiceServer) {
	s.RegisterService(&DbcConfigService_ServiceDesc, srv)
}

func _DbcConfigService_CreateDbc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDbcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbcConfigServiceServer).CreateDbc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbcConfigService_CreateDbc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbcConfigServiceServer).CreateDbc(ctx, req.(*UpdateDbcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbcConfigService_UpdateDbc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDbcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbcConfigServiceServer).UpdateDbc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbcConfigService_UpdateDbc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbcConfigServiceServer).UpdateDbc(ctx, req.(*UpdateDbcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbcConfigService_GetDbcList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDbcListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbcConfigServiceServer).GetDbcList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbcConfigService_GetDbcList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbcConfigServiceServer).GetDbcList(ctx, req.(*GetDbcListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbcConfigService_GetDbcByTemplateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDbcByTemplateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbcConfigServiceServer).GetDbcByTemplateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DbcConfigService_GetDbcByTemplateName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbcConfigServiceServer).GetDbcByTemplateName(ctx, req.(*GetDbcByTemplateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DbcConfigService_ServiceDesc is the grpc.ServiceDesc for DbcConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DbcConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DbcConfigService",
	HandlerType: (*DbcConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDbc",
			Handler:    _DbcConfigService_CreateDbc_Handler,
		},
		{
			MethodName: "UpdateDbc",
			Handler:    _DbcConfigService_UpdateDbc_Handler,
		},
		{
			MethodName: "GetDbcList",
			Handler:    _DbcConfigService_GetDbcList_Handler,
		},
		{
			MethodName: "GetDbcByTemplateName",
			Handler:    _DbcConfigService_GetDbcByTemplateName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/dbc_config.proto",
}
