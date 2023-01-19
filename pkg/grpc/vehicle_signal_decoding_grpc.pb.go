// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: pkg/grpc/vehicle_signal_decoding.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VehicleSignalDecodingServiceClient is the client API for VehicleSignalDecodingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleSignalDecodingServiceClient interface {
	ToDo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type vehicleSignalDecodingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleSignalDecodingServiceClient(cc grpc.ClientConnInterface) VehicleSignalDecodingServiceClient {
	return &vehicleSignalDecodingServiceClient{cc}
}

func (c *vehicleSignalDecodingServiceClient) ToDo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/ToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleSignalDecodingServiceServer is the server API for VehicleSignalDecodingService service.
// All implementations must embed UnimplementedVehicleSignalDecodingServiceServer
// for forward compatibility
type VehicleSignalDecodingServiceServer interface {
	ToDo(context.Context, *BaseRequest) (*BaseResponse, error)
	mustEmbedUnimplementedVehicleSignalDecodingServiceServer()
}

// UnimplementedVehicleSignalDecodingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleSignalDecodingServiceServer struct {
}

func (UnimplementedVehicleSignalDecodingServiceServer) ToDo(context.Context, *BaseRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToDo not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) mustEmbedUnimplementedVehicleSignalDecodingServiceServer() {
}

// UnsafeVehicleSignalDecodingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleSignalDecodingServiceServer will
// result in compilation errors.
type UnsafeVehicleSignalDecodingServiceServer interface {
	mustEmbedUnimplementedVehicleSignalDecodingServiceServer()
}

func RegisterVehicleSignalDecodingServiceServer(s grpc.ServiceRegistrar, srv VehicleSignalDecodingServiceServer) {
	s.RegisterService(&VehicleSignalDecodingService_ServiceDesc, srv)
}

func _VehicleSignalDecodingService_ToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).ToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/ToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).ToDo(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleSignalDecodingService_ServiceDesc is the grpc.ServiceDesc for VehicleSignalDecodingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleSignalDecodingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.VehicleSignalDecodingService",
	HandlerType: (*VehicleSignalDecodingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToDo",
			Handler:    _VehicleSignalDecodingService_ToDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/vehicle_signal_decoding.proto",
}
