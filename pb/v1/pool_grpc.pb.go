// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/pool.proto

package v1

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
	PoolService_CreatePool_FullMethodName = "/v1.PoolService/CreatePool"
)

// PoolServiceClient is the client API for PoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolServiceClient interface {
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error)
}

type poolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolServiceClient(cc grpc.ClientConnInterface) PoolServiceClient {
	return &poolServiceClient{cc}
}

func (c *poolServiceClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error) {
	out := new(CreatePoolResponse)
	err := c.cc.Invoke(ctx, PoolService_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolServiceServer is the server API for PoolService service.
// All implementations must embed UnimplementedPoolServiceServer
// for forward compatibility
type PoolServiceServer interface {
	CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error)
	mustEmbedUnimplementedPoolServiceServer()
}

// UnimplementedPoolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoolServiceServer struct {
}

func (UnimplementedPoolServiceServer) CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedPoolServiceServer) mustEmbedUnimplementedPoolServiceServer() {}

// UnsafePoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolServiceServer will
// result in compilation errors.
type UnsafePoolServiceServer interface {
	mustEmbedUnimplementedPoolServiceServer()
}

func RegisterPoolServiceServer(s grpc.ServiceRegistrar, srv PoolServiceServer) {
	s.RegisterService(&PoolService_ServiceDesc, srv)
}

func _PoolService_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoolService_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoolService_ServiceDesc is the grpc.ServiceDesc for PoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PoolService",
	HandlerType: (*PoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _PoolService_CreatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/pool.proto",
}
