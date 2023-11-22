// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pkg/grpc/vehicle_template.proto

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
	VehicleTemplateService_GetVehicleTemplates_FullMethodName   = "/grpc.VehicleTemplateService/GetVehicleTemplates"
	VehicleTemplateService_GetVehicleTemplate_FullMethodName    = "/grpc.VehicleTemplateService/GetVehicleTemplate"
	VehicleTemplateService_CreateVehicleTemplate_FullMethodName = "/grpc.VehicleTemplateService/CreateVehicleTemplate"
)

// VehicleTemplateServiceClient is the client API for VehicleTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleTemplateServiceClient interface {
	GetVehicleTemplates(ctx context.Context, in *GetVehicleTemplatesRequest, opts ...grpc.CallOption) (*GetVehicleTemplatesResponse, error)
	GetVehicleTemplate(ctx context.Context, in *GetVehicleTemplateRequest, opts ...grpc.CallOption) (*VehicleTemplate, error)
	CreateVehicleTemplate(ctx context.Context, in *VehicleTemplate, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type vehicleTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleTemplateServiceClient(cc grpc.ClientConnInterface) VehicleTemplateServiceClient {
	return &vehicleTemplateServiceClient{cc}
}

func (c *vehicleTemplateServiceClient) GetVehicleTemplates(ctx context.Context, in *GetVehicleTemplatesRequest, opts ...grpc.CallOption) (*GetVehicleTemplatesResponse, error) {
	out := new(GetVehicleTemplatesResponse)
	err := c.cc.Invoke(ctx, VehicleTemplateService_GetVehicleTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleTemplateServiceClient) GetVehicleTemplate(ctx context.Context, in *GetVehicleTemplateRequest, opts ...grpc.CallOption) (*VehicleTemplate, error) {
	out := new(VehicleTemplate)
	err := c.cc.Invoke(ctx, VehicleTemplateService_GetVehicleTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleTemplateServiceClient) CreateVehicleTemplate(ctx context.Context, in *VehicleTemplate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VehicleTemplateService_CreateVehicleTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleTemplateServiceServer is the server API for VehicleTemplateService service.
// All implementations must embed UnimplementedVehicleTemplateServiceServer
// for forward compatibility
type VehicleTemplateServiceServer interface {
	GetVehicleTemplates(context.Context, *GetVehicleTemplatesRequest) (*GetVehicleTemplatesResponse, error)
	GetVehicleTemplate(context.Context, *GetVehicleTemplateRequest) (*VehicleTemplate, error)
	CreateVehicleTemplate(context.Context, *VehicleTemplate) (*emptypb.Empty, error)
	mustEmbedUnimplementedVehicleTemplateServiceServer()
}

// UnimplementedVehicleTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleTemplateServiceServer struct {
}

func (UnimplementedVehicleTemplateServiceServer) GetVehicleTemplates(context.Context, *GetVehicleTemplatesRequest) (*GetVehicleTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleTemplates not implemented")
}
func (UnimplementedVehicleTemplateServiceServer) GetVehicleTemplate(context.Context, *GetVehicleTemplateRequest) (*VehicleTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicleTemplate not implemented")
}
func (UnimplementedVehicleTemplateServiceServer) CreateVehicleTemplate(context.Context, *VehicleTemplate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicleTemplate not implemented")
}
func (UnimplementedVehicleTemplateServiceServer) mustEmbedUnimplementedVehicleTemplateServiceServer() {
}

// UnsafeVehicleTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleTemplateServiceServer will
// result in compilation errors.
type UnsafeVehicleTemplateServiceServer interface {
	mustEmbedUnimplementedVehicleTemplateServiceServer()
}

func RegisterVehicleTemplateServiceServer(s grpc.ServiceRegistrar, srv VehicleTemplateServiceServer) {
	s.RegisterService(&VehicleTemplateService_ServiceDesc, srv)
}

func _VehicleTemplateService_GetVehicleTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleTemplateServiceServer).GetVehicleTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleTemplateService_GetVehicleTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleTemplateServiceServer).GetVehicleTemplates(ctx, req.(*GetVehicleTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleTemplateService_GetVehicleTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleTemplateServiceServer).GetVehicleTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleTemplateService_GetVehicleTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleTemplateServiceServer).GetVehicleTemplate(ctx, req.(*GetVehicleTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleTemplateService_CreateVehicleTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehicleTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleTemplateServiceServer).CreateVehicleTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleTemplateService_CreateVehicleTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleTemplateServiceServer).CreateVehicleTemplate(ctx, req.(*VehicleTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleTemplateService_ServiceDesc is the grpc.ServiceDesc for VehicleTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.VehicleTemplateService",
	HandlerType: (*VehicleTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVehicleTemplates",
			Handler:    _VehicleTemplateService_GetVehicleTemplates_Handler,
		},
		{
			MethodName: "GetVehicleTemplate",
			Handler:    _VehicleTemplateService_GetVehicleTemplate_Handler,
		},
		{
			MethodName: "CreateVehicleTemplate",
			Handler:    _VehicleTemplateService_CreateVehicleTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/vehicle_template.proto",
}
