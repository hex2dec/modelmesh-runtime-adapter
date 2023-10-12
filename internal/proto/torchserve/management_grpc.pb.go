// Copied from https://github.com/pytorch/serve/blob/8c23585d2453f230c411721028ad4b07e58cc7dd/frontend/server/src/main/resources/proto/management.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.5
// source: management.proto

package torchserve

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

const (
	ManagementAPIsService_DescribeModel_FullMethodName   = "/org.pytorch.serve.grpc.management.ManagementAPIsService/DescribeModel"
	ManagementAPIsService_ListModels_FullMethodName      = "/org.pytorch.serve.grpc.management.ManagementAPIsService/ListModels"
	ManagementAPIsService_RegisterModel_FullMethodName   = "/org.pytorch.serve.grpc.management.ManagementAPIsService/RegisterModel"
	ManagementAPIsService_ScaleWorker_FullMethodName     = "/org.pytorch.serve.grpc.management.ManagementAPIsService/ScaleWorker"
	ManagementAPIsService_SetDefault_FullMethodName      = "/org.pytorch.serve.grpc.management.ManagementAPIsService/SetDefault"
	ManagementAPIsService_UnregisterModel_FullMethodName = "/org.pytorch.serve.grpc.management.ManagementAPIsService/UnregisterModel"
)

// ManagementAPIsServiceClient is the client API for ManagementAPIsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementAPIsServiceClient interface {
	// Provides detailed information about the default version of a model.
	DescribeModel(ctx context.Context, in *DescribeModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
	// List registered models in TorchServe.
	ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
	// Register a new model in TorchServe.
	RegisterModel(ctx context.Context, in *RegisterModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
	// Configure number of workers for a default version of a model.This is a asynchronous call by default. Caller need to call describeModel to check if the model workers has been changed.
	ScaleWorker(ctx context.Context, in *ScaleWorkerRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
	// Set default version of a model
	SetDefault(ctx context.Context, in *SetDefaultRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
	// Unregister the default version of a model from TorchServe if it is the only version available.This is a asynchronous call by default. Caller can call listModels to confirm model is unregistered
	UnregisterModel(ctx context.Context, in *UnregisterModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error)
}

type managementAPIsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementAPIsServiceClient(cc grpc.ClientConnInterface) ManagementAPIsServiceClient {
	return &managementAPIsServiceClient{cc}
}

func (c *managementAPIsServiceClient) DescribeModel(ctx context.Context, in *DescribeModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_DescribeModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementAPIsServiceClient) ListModels(ctx context.Context, in *ListModelsRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_ListModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementAPIsServiceClient) RegisterModel(ctx context.Context, in *RegisterModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_RegisterModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementAPIsServiceClient) ScaleWorker(ctx context.Context, in *ScaleWorkerRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_ScaleWorker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementAPIsServiceClient) SetDefault(ctx context.Context, in *SetDefaultRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_SetDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementAPIsServiceClient) UnregisterModel(ctx context.Context, in *UnregisterModelRequest, opts ...grpc.CallOption) (*ManagementResponse, error) {
	out := new(ManagementResponse)
	err := c.cc.Invoke(ctx, ManagementAPIsService_UnregisterModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementAPIsServiceServer is the server API for ManagementAPIsService service.
// All implementations must embed UnimplementedManagementAPIsServiceServer
// for forward compatibility
type ManagementAPIsServiceServer interface {
	// Provides detailed information about the default version of a model.
	DescribeModel(context.Context, *DescribeModelRequest) (*ManagementResponse, error)
	// List registered models in TorchServe.
	ListModels(context.Context, *ListModelsRequest) (*ManagementResponse, error)
	// Register a new model in TorchServe.
	RegisterModel(context.Context, *RegisterModelRequest) (*ManagementResponse, error)
	// Configure number of workers for a default version of a model.This is a asynchronous call by default. Caller need to call describeModel to check if the model workers has been changed.
	ScaleWorker(context.Context, *ScaleWorkerRequest) (*ManagementResponse, error)
	// Set default version of a model
	SetDefault(context.Context, *SetDefaultRequest) (*ManagementResponse, error)
	// Unregister the default version of a model from TorchServe if it is the only version available.This is a asynchronous call by default. Caller can call listModels to confirm model is unregistered
	UnregisterModel(context.Context, *UnregisterModelRequest) (*ManagementResponse, error)
	mustEmbedUnimplementedManagementAPIsServiceServer()
}

// UnimplementedManagementAPIsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagementAPIsServiceServer struct {
}

func (UnimplementedManagementAPIsServiceServer) DescribeModel(context.Context, *DescribeModelRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeModel not implemented")
}
func (UnimplementedManagementAPIsServiceServer) ListModels(context.Context, *ListModelsRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedManagementAPIsServiceServer) RegisterModel(context.Context, *RegisterModelRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterModel not implemented")
}
func (UnimplementedManagementAPIsServiceServer) ScaleWorker(context.Context, *ScaleWorkerRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleWorker not implemented")
}
func (UnimplementedManagementAPIsServiceServer) SetDefault(context.Context, *SetDefaultRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefault not implemented")
}
func (UnimplementedManagementAPIsServiceServer) UnregisterModel(context.Context, *UnregisterModelRequest) (*ManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterModel not implemented")
}
func (UnimplementedManagementAPIsServiceServer) mustEmbedUnimplementedManagementAPIsServiceServer() {}

// UnsafeManagementAPIsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementAPIsServiceServer will
// result in compilation errors.
type UnsafeManagementAPIsServiceServer interface {
	mustEmbedUnimplementedManagementAPIsServiceServer()
}

func RegisterManagementAPIsServiceServer(s grpc.ServiceRegistrar, srv ManagementAPIsServiceServer) {
	s.RegisterService(&ManagementAPIsService_ServiceDesc, srv)
}

func _ManagementAPIsService_DescribeModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).DescribeModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_DescribeModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).DescribeModel(ctx, req.(*DescribeModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementAPIsService_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).ListModels(ctx, req.(*ListModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementAPIsService_RegisterModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).RegisterModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_RegisterModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).RegisterModel(ctx, req.(*RegisterModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementAPIsService_ScaleWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).ScaleWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_ScaleWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).ScaleWorker(ctx, req.(*ScaleWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementAPIsService_SetDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).SetDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_SetDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).SetDefault(ctx, req.(*SetDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementAPIsService_UnregisterModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementAPIsServiceServer).UnregisterModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementAPIsService_UnregisterModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementAPIsServiceServer).UnregisterModel(ctx, req.(*UnregisterModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagementAPIsService_ServiceDesc is the grpc.ServiceDesc for ManagementAPIsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementAPIsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.pytorch.serve.grpc.management.ManagementAPIsService",
	HandlerType: (*ManagementAPIsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeModel",
			Handler:    _ManagementAPIsService_DescribeModel_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _ManagementAPIsService_ListModels_Handler,
		},
		{
			MethodName: "RegisterModel",
			Handler:    _ManagementAPIsService_RegisterModel_Handler,
		},
		{
			MethodName: "ScaleWorker",
			Handler:    _ManagementAPIsService_ScaleWorker_Handler,
		},
		{
			MethodName: "SetDefault",
			Handler:    _ManagementAPIsService_SetDefault_Handler,
		},
		{
			MethodName: "UnregisterModel",
			Handler:    _ManagementAPIsService_UnregisterModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management.proto",
}
