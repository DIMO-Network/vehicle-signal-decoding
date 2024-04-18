// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: pkg/grpc/aftermarket_device_template.proto

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
	AftermarketDeviceTemplateService_CreateAftermarketDeviceTemplate_FullMethodName = "/grpc.AftermarketDeviceTemplateService/CreateAftermarketDeviceTemplate"
	AftermarketDeviceTemplateService_DeleteAftermarketDeviceTemplate_FullMethodName = "/grpc.AftermarketDeviceTemplateService/DeleteAftermarketDeviceTemplate"
	AftermarketDeviceTemplateService_GetAftermarketDeviceTemplates_FullMethodName   = "/grpc.AftermarketDeviceTemplateService/GetAftermarketDeviceTemplates"
	AftermarketDeviceTemplateService_UpdateAftermarketDeviceTemplate_FullMethodName = "/grpc.AftermarketDeviceTemplateService/UpdateAftermarketDeviceTemplate"
)

// AftermarketDeviceTemplateServiceClient is the client API for AftermarketDeviceTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AftermarketDeviceTemplateServiceClient interface {
	CreateAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAftermarketDeviceTemplates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AftermarketDeviceTemplates, error)
	UpdateAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aftermarketDeviceTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAftermarketDeviceTemplateServiceClient(cc grpc.ClientConnInterface) AftermarketDeviceTemplateServiceClient {
	return &aftermarketDeviceTemplateServiceClient{cc}
}

func (c *aftermarketDeviceTemplateServiceClient) CreateAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AftermarketDeviceTemplateService_CreateAftermarketDeviceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftermarketDeviceTemplateServiceClient) DeleteAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AftermarketDeviceTemplateService_DeleteAftermarketDeviceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftermarketDeviceTemplateServiceClient) GetAftermarketDeviceTemplates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AftermarketDeviceTemplates, error) {
	out := new(AftermarketDeviceTemplates)
	err := c.cc.Invoke(ctx, AftermarketDeviceTemplateService_GetAftermarketDeviceTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftermarketDeviceTemplateServiceClient) UpdateAftermarketDeviceTemplate(ctx context.Context, in *AftermarketDeviceTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AftermarketDeviceTemplateService_UpdateAftermarketDeviceTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AftermarketDeviceTemplateServiceServer is the server API for AftermarketDeviceTemplateService service.
// All implementations must embed UnimplementedAftermarketDeviceTemplateServiceServer
// for forward compatibility
type AftermarketDeviceTemplateServiceServer interface {
	CreateAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error)
	DeleteAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error)
	GetAftermarketDeviceTemplates(context.Context, *emptypb.Empty) (*AftermarketDeviceTemplates, error)
	UpdateAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAftermarketDeviceTemplateServiceServer()
}

// UnimplementedAftermarketDeviceTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAftermarketDeviceTemplateServiceServer struct {
}

func (UnimplementedAftermarketDeviceTemplateServiceServer) CreateAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAftermarketDeviceTemplate not implemented")
}
func (UnimplementedAftermarketDeviceTemplateServiceServer) DeleteAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAftermarketDeviceTemplate not implemented")
}
func (UnimplementedAftermarketDeviceTemplateServiceServer) GetAftermarketDeviceTemplates(context.Context, *emptypb.Empty) (*AftermarketDeviceTemplates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAftermarketDeviceTemplates not implemented")
}
func (UnimplementedAftermarketDeviceTemplateServiceServer) UpdateAftermarketDeviceTemplate(context.Context, *AftermarketDeviceTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAftermarketDeviceTemplate not implemented")
}
func (UnimplementedAftermarketDeviceTemplateServiceServer) mustEmbedUnimplementedAftermarketDeviceTemplateServiceServer() {
}

// UnsafeAftermarketDeviceTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AftermarketDeviceTemplateServiceServer will
// result in compilation errors.
type UnsafeAftermarketDeviceTemplateServiceServer interface {
	mustEmbedUnimplementedAftermarketDeviceTemplateServiceServer()
}

func RegisterAftermarketDeviceTemplateServiceServer(s grpc.ServiceRegistrar, srv AftermarketDeviceTemplateServiceServer) {
	s.RegisterService(&AftermarketDeviceTemplateService_ServiceDesc, srv)
}

func _AftermarketDeviceTemplateService_CreateAftermarketDeviceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AftermarketDeviceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AftermarketDeviceTemplateServiceServer).CreateAftermarketDeviceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AftermarketDeviceTemplateService_CreateAftermarketDeviceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AftermarketDeviceTemplateServiceServer).CreateAftermarketDeviceTemplate(ctx, req.(*AftermarketDeviceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AftermarketDeviceTemplateService_DeleteAftermarketDeviceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AftermarketDeviceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AftermarketDeviceTemplateServiceServer).DeleteAftermarketDeviceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AftermarketDeviceTemplateService_DeleteAftermarketDeviceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AftermarketDeviceTemplateServiceServer).DeleteAftermarketDeviceTemplate(ctx, req.(*AftermarketDeviceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AftermarketDeviceTemplateService_GetAftermarketDeviceTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AftermarketDeviceTemplateServiceServer).GetAftermarketDeviceTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AftermarketDeviceTemplateService_GetAftermarketDeviceTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AftermarketDeviceTemplateServiceServer).GetAftermarketDeviceTemplates(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AftermarketDeviceTemplateService_UpdateAftermarketDeviceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AftermarketDeviceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AftermarketDeviceTemplateServiceServer).UpdateAftermarketDeviceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AftermarketDeviceTemplateService_UpdateAftermarketDeviceTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AftermarketDeviceTemplateServiceServer).UpdateAftermarketDeviceTemplate(ctx, req.(*AftermarketDeviceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AftermarketDeviceTemplateService_ServiceDesc is the grpc.ServiceDesc for AftermarketDeviceTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AftermarketDeviceTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AftermarketDeviceTemplateService",
	HandlerType: (*AftermarketDeviceTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAftermarketDeviceTemplate",
			Handler:    _AftermarketDeviceTemplateService_CreateAftermarketDeviceTemplate_Handler,
		},
		{
			MethodName: "DeleteAftermarketDeviceTemplate",
			Handler:    _AftermarketDeviceTemplateService_DeleteAftermarketDeviceTemplate_Handler,
		},
		{
			MethodName: "GetAftermarketDeviceTemplates",
			Handler:    _AftermarketDeviceTemplateService_GetAftermarketDeviceTemplates_Handler,
		},
		{
			MethodName: "UpdateAftermarketDeviceTemplate",
			Handler:    _AftermarketDeviceTemplateService_UpdateAftermarketDeviceTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/aftermarket_device_template.proto",
}
