// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: qdrant.proto

package qdrant

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

// QdrantClient is the client API for Qdrant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QdrantClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckReply, error)
}

type qdrantClient struct {
	cc grpc.ClientConnInterface
}

func NewQdrantClient(cc grpc.ClientConnInterface) QdrantClient {
	return &qdrantClient{cc}
}

func (c *qdrantClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckReply, error) {
	out := new(HealthCheckReply)
	err := c.cc.Invoke(ctx, "/qdrant.Qdrant/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QdrantServer is the server API for Qdrant service.
// All implementations must embed UnimplementedQdrantServer
// for forward compatibility
type QdrantServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckReply, error)
	mustEmbedUnimplementedQdrantServer()
}

// UnimplementedQdrantServer must be embedded to have forward compatible implementations.
type UnimplementedQdrantServer struct {
}

func (UnimplementedQdrantServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedQdrantServer) mustEmbedUnimplementedQdrantServer() {}

// UnsafeQdrantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QdrantServer will
// result in compilation errors.
type UnsafeQdrantServer interface {
	mustEmbedUnimplementedQdrantServer()
}

func RegisterQdrantServer(s grpc.ServiceRegistrar, srv QdrantServer) {
	s.RegisterService(&Qdrant_ServiceDesc, srv)
}

func _Qdrant_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QdrantServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Qdrant/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QdrantServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Qdrant_ServiceDesc is the grpc.ServiceDesc for Qdrant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Qdrant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qdrant.Qdrant",
	HandlerType: (*QdrantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Qdrant_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qdrant.proto",
}
