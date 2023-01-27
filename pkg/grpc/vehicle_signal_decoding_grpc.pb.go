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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VehicleSignalDecodingServiceClient is the client API for VehicleSignalDecodingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleSignalDecodingServiceClient interface {
	CreateDBCCode(ctx context.Context, in *CreateDBCCodeRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error)
	UpdateDBCCode(ctx context.Context, in *UpdateDBCCodeRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error)
	GetDBCCodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDBCCodeListResponse, error)
	GetDBCCodesByID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetDBCCodeResponse, error)
	CreateTestSignal(ctx context.Context, in *CreateTestSignalRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error)
	UpdateTestSignal(ctx context.Context, in *UpdateTestSignalRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error)
	GetTestSignals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTestSignalListResponse, error)
	GetTestSignalsByDeviceDefinitionID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error)
	GetTestSignalsByUserDeviceID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error)
	GetTestSignalsByDBCCodeID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error)
	GetTestSignalByID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalResponse, error)
}

type vehicleSignalDecodingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleSignalDecodingServiceClient(cc grpc.ClientConnInterface) VehicleSignalDecodingServiceClient {
	return &vehicleSignalDecodingServiceClient{cc}
}

func (c *vehicleSignalDecodingServiceClient) CreateDBCCode(ctx context.Context, in *CreateDBCCodeRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error) {
	out := new(VehicleSignalBaseResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/CreateDBCCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) UpdateDBCCode(ctx context.Context, in *UpdateDBCCodeRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error) {
	out := new(VehicleSignalBaseResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/UpdateDBCCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetDBCCodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDBCCodeListResponse, error) {
	out := new(GetDBCCodeListResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetDBCCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetDBCCodesByID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetDBCCodeResponse, error) {
	out := new(GetDBCCodeResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetDBCCodesByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) CreateTestSignal(ctx context.Context, in *CreateTestSignalRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error) {
	out := new(VehicleSignalBaseResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/CreateTestSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) UpdateTestSignal(ctx context.Context, in *UpdateTestSignalRequest, opts ...grpc.CallOption) (*VehicleSignalBaseResponse, error) {
	out := new(VehicleSignalBaseResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/UpdateTestSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetTestSignals(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTestSignalListResponse, error) {
	out := new(GetTestSignalListResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetTestSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetTestSignalsByDeviceDefinitionID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error) {
	out := new(GetTestSignalListResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetTestSignalsByDeviceDefinitionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetTestSignalsByUserDeviceID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error) {
	out := new(GetTestSignalListResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetTestSignalsByUserDeviceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetTestSignalsByDBCCodeID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalListResponse, error) {
	out := new(GetTestSignalListResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetTestSignalsByDBCCodeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleSignalDecodingServiceClient) GetTestSignalByID(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTestSignalResponse, error) {
	out := new(GetTestSignalResponse)
	err := c.cc.Invoke(ctx, "/grpc.VehicleSignalDecodingService/GetTestSignalByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleSignalDecodingServiceServer is the server API for VehicleSignalDecodingService service.
// All implementations must embed UnimplementedVehicleSignalDecodingServiceServer
// for forward compatibility
type VehicleSignalDecodingServiceServer interface {
	CreateDBCCode(context.Context, *CreateDBCCodeRequest) (*VehicleSignalBaseResponse, error)
	UpdateDBCCode(context.Context, *UpdateDBCCodeRequest) (*VehicleSignalBaseResponse, error)
	GetDBCCodes(context.Context, *emptypb.Empty) (*GetDBCCodeListResponse, error)
	GetDBCCodesByID(context.Context, *GetByIdRequest) (*GetDBCCodeResponse, error)
	CreateTestSignal(context.Context, *CreateTestSignalRequest) (*VehicleSignalBaseResponse, error)
	UpdateTestSignal(context.Context, *UpdateTestSignalRequest) (*VehicleSignalBaseResponse, error)
	GetTestSignals(context.Context, *emptypb.Empty) (*GetTestSignalListResponse, error)
	GetTestSignalsByDeviceDefinitionID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error)
	GetTestSignalsByUserDeviceID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error)
	GetTestSignalsByDBCCodeID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error)
	GetTestSignalByID(context.Context, *GetByIdRequest) (*GetTestSignalResponse, error)
	mustEmbedUnimplementedVehicleSignalDecodingServiceServer()
}

// UnimplementedVehicleSignalDecodingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleSignalDecodingServiceServer struct {
}

func (UnimplementedVehicleSignalDecodingServiceServer) CreateDBCCode(context.Context, *CreateDBCCodeRequest) (*VehicleSignalBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDBCCode not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) UpdateDBCCode(context.Context, *UpdateDBCCodeRequest) (*VehicleSignalBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDBCCode not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetDBCCodes(context.Context, *emptypb.Empty) (*GetDBCCodeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDBCCodes not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetDBCCodesByID(context.Context, *GetByIdRequest) (*GetDBCCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDBCCodesByID not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) CreateTestSignal(context.Context, *CreateTestSignalRequest) (*VehicleSignalBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestSignal not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) UpdateTestSignal(context.Context, *UpdateTestSignalRequest) (*VehicleSignalBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestSignal not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetTestSignals(context.Context, *emptypb.Empty) (*GetTestSignalListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSignals not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetTestSignalsByDeviceDefinitionID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSignalsByDeviceDefinitionID not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetTestSignalsByUserDeviceID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSignalsByUserDeviceID not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetTestSignalsByDBCCodeID(context.Context, *GetByIdRequest) (*GetTestSignalListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSignalsByDBCCodeID not implemented")
}
func (UnimplementedVehicleSignalDecodingServiceServer) GetTestSignalByID(context.Context, *GetByIdRequest) (*GetTestSignalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSignalByID not implemented")
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

func _VehicleSignalDecodingService_CreateDBCCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDBCCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).CreateDBCCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/CreateDBCCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).CreateDBCCode(ctx, req.(*CreateDBCCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_UpdateDBCCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDBCCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).UpdateDBCCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/UpdateDBCCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).UpdateDBCCode(ctx, req.(*UpdateDBCCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetDBCCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetDBCCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetDBCCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetDBCCodes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetDBCCodesByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetDBCCodesByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetDBCCodesByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetDBCCodesByID(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_CreateTestSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestSignalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).CreateTestSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/CreateTestSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).CreateTestSignal(ctx, req.(*CreateTestSignalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_UpdateTestSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestSignalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).UpdateTestSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/UpdateTestSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).UpdateTestSignal(ctx, req.(*UpdateTestSignalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetTestSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetTestSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignals(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetTestSignalsByDeviceDefinitionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByDeviceDefinitionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetTestSignalsByDeviceDefinitionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByDeviceDefinitionID(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetTestSignalsByUserDeviceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByUserDeviceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetTestSignalsByUserDeviceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByUserDeviceID(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetTestSignalsByDBCCodeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByDBCCodeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetTestSignalsByDBCCodeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalsByDBCCodeID(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleSignalDecodingService_GetTestSignalByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VehicleSignalDecodingService/GetTestSignalByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleSignalDecodingServiceServer).GetTestSignalByID(ctx, req.(*GetByIdRequest))
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
			MethodName: "CreateDBCCode",
			Handler:    _VehicleSignalDecodingService_CreateDBCCode_Handler,
		},
		{
			MethodName: "UpdateDBCCode",
			Handler:    _VehicleSignalDecodingService_UpdateDBCCode_Handler,
		},
		{
			MethodName: "GetDBCCodes",
			Handler:    _VehicleSignalDecodingService_GetDBCCodes_Handler,
		},
		{
			MethodName: "GetDBCCodesByID",
			Handler:    _VehicleSignalDecodingService_GetDBCCodesByID_Handler,
		},
		{
			MethodName: "CreateTestSignal",
			Handler:    _VehicleSignalDecodingService_CreateTestSignal_Handler,
		},
		{
			MethodName: "UpdateTestSignal",
			Handler:    _VehicleSignalDecodingService_UpdateTestSignal_Handler,
		},
		{
			MethodName: "GetTestSignals",
			Handler:    _VehicleSignalDecodingService_GetTestSignals_Handler,
		},
		{
			MethodName: "GetTestSignalsByDeviceDefinitionID",
			Handler:    _VehicleSignalDecodingService_GetTestSignalsByDeviceDefinitionID_Handler,
		},
		{
			MethodName: "GetTestSignalsByUserDeviceID",
			Handler:    _VehicleSignalDecodingService_GetTestSignalsByUserDeviceID_Handler,
		},
		{
			MethodName: "GetTestSignalsByDBCCodeID",
			Handler:    _VehicleSignalDecodingService_GetTestSignalsByDBCCodeID_Handler,
		},
		{
			MethodName: "GetTestSignalByID",
			Handler:    _VehicleSignalDecodingService_GetTestSignalByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/vehicle_signal_decoding.proto",
}
