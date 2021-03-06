// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_micro_api_resourceManagerService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ResourceManagerServiceClient is the client API for ResourceManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceManagerServiceClient interface {
	ProvisionFileSystem(ctx context.Context, in *FileSystemSpec, opts ...grpc.CallOption) (*EmptyResponse, error)
	ProvisionExecutorInstance(ctx context.Context, in *ExecutorInstanceSpec, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetFileSystem(ctx context.Context, in *FileSystemSpec, opts ...grpc.CallOption) (*FileSystemDetails, error)
}

type resourceManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceManagerServiceClient(cc grpc.ClientConnInterface) ResourceManagerServiceClient {
	return &resourceManagerServiceClient{cc}
}

func (c *resourceManagerServiceClient) ProvisionFileSystem(ctx context.Context, in *FileSystemSpec, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/go.micro.api.resourceManagerService.ResourceManagerService/ProvisionFileSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceManagerServiceClient) ProvisionExecutorInstance(ctx context.Context, in *ExecutorInstanceSpec, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/go.micro.api.resourceManagerService.ResourceManagerService/ProvisionExecutorInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceManagerServiceClient) GetFileSystem(ctx context.Context, in *FileSystemSpec, opts ...grpc.CallOption) (*FileSystemDetails, error) {
	out := new(FileSystemDetails)
	err := c.cc.Invoke(ctx, "/go.micro.api.resourceManagerService.ResourceManagerService/GetFileSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceManagerServiceServer is the server API for ResourceManagerService service.
// All implementations must embed UnimplementedResourceManagerServiceServer
// for forward compatibility
type ResourceManagerServiceServer interface {
	ProvisionFileSystem(context.Context, *FileSystemSpec) (*EmptyResponse, error)
	ProvisionExecutorInstance(context.Context, *ExecutorInstanceSpec) (*EmptyResponse, error)
	GetFileSystem(context.Context, *FileSystemSpec) (*FileSystemDetails, error)
	mustEmbedUnimplementedResourceManagerServiceServer()
}

// UnimplementedResourceManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceManagerServiceServer struct {
}

func (UnimplementedResourceManagerServiceServer) ProvisionFileSystem(context.Context, *FileSystemSpec) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionFileSystem not implemented")
}
func (UnimplementedResourceManagerServiceServer) ProvisionExecutorInstance(context.Context, *ExecutorInstanceSpec) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionExecutorInstance not implemented")
}
func (UnimplementedResourceManagerServiceServer) GetFileSystem(context.Context, *FileSystemSpec) (*FileSystemDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileSystem not implemented")
}
func (UnimplementedResourceManagerServiceServer) mustEmbedUnimplementedResourceManagerServiceServer() {
}

// UnsafeResourceManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceManagerServiceServer will
// result in compilation errors.
type UnsafeResourceManagerServiceServer interface {
	mustEmbedUnimplementedResourceManagerServiceServer()
}

func RegisterResourceManagerServiceServer(s grpc.ServiceRegistrar, srv ResourceManagerServiceServer) {
	s.RegisterService(&_ResourceManagerService_serviceDesc, srv)
}

func _ResourceManagerService_ProvisionFileSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSystemSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceManagerServiceServer).ProvisionFileSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.resourceManagerService.ResourceManagerService/ProvisionFileSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceManagerServiceServer).ProvisionFileSystem(ctx, req.(*FileSystemSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceManagerService_ProvisionExecutorInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutorInstanceSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceManagerServiceServer).ProvisionExecutorInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.resourceManagerService.ResourceManagerService/ProvisionExecutorInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceManagerServiceServer).ProvisionExecutorInstance(ctx, req.(*ExecutorInstanceSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceManagerService_GetFileSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSystemSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceManagerServiceServer).GetFileSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.resourceManagerService.ResourceManagerService/GetFileSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceManagerServiceServer).GetFileSystem(ctx, req.(*FileSystemSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.api.resourceManagerService.ResourceManagerService",
	HandlerType: (*ResourceManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProvisionFileSystem",
			Handler:    _ResourceManagerService_ProvisionFileSystem_Handler,
		},
		{
			MethodName: "ProvisionExecutorInstance",
			Handler:    _ResourceManagerService_ProvisionExecutorInstance_Handler,
		},
		{
			MethodName: "GetFileSystem",
			Handler:    _ResourceManagerService_GetFileSystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource-manager-service.proto",
}
