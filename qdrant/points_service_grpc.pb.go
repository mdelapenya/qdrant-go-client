// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: points_service.proto

package go_client

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

// PointsClient is the client API for Points service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointsClient interface {
	//
	//Perform insert + updates on points. If point with given ID already exists - it will be overwritten.
	Upsert(ctx context.Context, in *UpsertPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete points
	Delete(ctx context.Context, in *DeletePoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Retrieve points
	Get(ctx context.Context, in *GetPoints, opts ...grpc.CallOption) (*GetResponse, error)
	//
	//Set payload for points
	SetPayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete specified key payload for points
	DeletePayload(ctx context.Context, in *DeletePayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Remove all payload for specified points
	ClearPayload(ctx context.Context, in *ClearPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Create index for field in collection
	CreateFieldIndex(ctx context.Context, in *CreateFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete field index for collection
	DeleteFieldIndex(ctx context.Context, in *DeleteFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	Search(ctx context.Context, in *SearchPoints, opts ...grpc.CallOption) (*SearchResponse, error)
	//
	//Iterate over all or filtered points points
	Scroll(ctx context.Context, in *ScrollPoints, opts ...grpc.CallOption) (*ScrollResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	Recommend(ctx context.Context, in *RecommendPoints, opts ...grpc.CallOption) (*RecommendResponse, error)
	//
	//Count points in collection with given filtering conditions
	Count(ctx context.Context, in *CountPoints, opts ...grpc.CallOption) (*CountResponse, error)
}

type pointsClient struct {
	cc grpc.ClientConnInterface
}

func NewPointsClient(cc grpc.ClientConnInterface) PointsClient {
	return &pointsClient{cc}
}

func (c *pointsClient) Upsert(ctx context.Context, in *UpsertPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Delete(ctx context.Context, in *DeletePoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Get(ctx context.Context, in *GetPoints, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) SetPayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/SetPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) DeletePayload(ctx context.Context, in *DeletePayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/DeletePayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) ClearPayload(ctx context.Context, in *ClearPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/ClearPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) CreateFieldIndex(ctx context.Context, in *CreateFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/CreateFieldIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) DeleteFieldIndex(ctx context.Context, in *DeleteFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/DeleteFieldIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Search(ctx context.Context, in *SearchPoints, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Scroll(ctx context.Context, in *ScrollPoints, opts ...grpc.CallOption) (*ScrollResponse, error) {
	out := new(ScrollResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Scroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Recommend(ctx context.Context, in *RecommendPoints, opts ...grpc.CallOption) (*RecommendResponse, error) {
	out := new(RecommendResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Recommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Count(ctx context.Context, in *CountPoints, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Points/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointsServer is the server API for Points service.
// All implementations must embed UnimplementedPointsServer
// for forward compatibility
type PointsServer interface {
	//
	//Perform insert + updates on points. If point with given ID already exists - it will be overwritten.
	Upsert(context.Context, *UpsertPoints) (*PointsOperationResponse, error)
	//
	//Delete points
	Delete(context.Context, *DeletePoints) (*PointsOperationResponse, error)
	//
	//Retrieve points
	Get(context.Context, *GetPoints) (*GetResponse, error)
	//
	//Set payload for points
	SetPayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error)
	//
	//Delete specified key payload for points
	DeletePayload(context.Context, *DeletePayloadPoints) (*PointsOperationResponse, error)
	//
	//Remove all payload for specified points
	ClearPayload(context.Context, *ClearPayloadPoints) (*PointsOperationResponse, error)
	//
	//Create index for field in collection
	CreateFieldIndex(context.Context, *CreateFieldIndexCollection) (*PointsOperationResponse, error)
	//
	//Delete field index for collection
	DeleteFieldIndex(context.Context, *DeleteFieldIndexCollection) (*PointsOperationResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	Search(context.Context, *SearchPoints) (*SearchResponse, error)
	//
	//Iterate over all or filtered points points
	Scroll(context.Context, *ScrollPoints) (*ScrollResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	Recommend(context.Context, *RecommendPoints) (*RecommendResponse, error)
	//
	//Count points in collection with given filtering conditions
	Count(context.Context, *CountPoints) (*CountResponse, error)
	mustEmbedUnimplementedPointsServer()
}

// UnimplementedPointsServer must be embedded to have forward compatible implementations.
type UnimplementedPointsServer struct {
}

func (UnimplementedPointsServer) Upsert(context.Context, *UpsertPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedPointsServer) Delete(context.Context, *DeletePoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPointsServer) Get(context.Context, *GetPoints) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPointsServer) SetPayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPayload not implemented")
}
func (UnimplementedPointsServer) DeletePayload(context.Context, *DeletePayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayload not implemented")
}
func (UnimplementedPointsServer) ClearPayload(context.Context, *ClearPayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearPayload not implemented")
}
func (UnimplementedPointsServer) CreateFieldIndex(context.Context, *CreateFieldIndexCollection) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFieldIndex not implemented")
}
func (UnimplementedPointsServer) DeleteFieldIndex(context.Context, *DeleteFieldIndexCollection) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFieldIndex not implemented")
}
func (UnimplementedPointsServer) Search(context.Context, *SearchPoints) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedPointsServer) Scroll(context.Context, *ScrollPoints) (*ScrollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scroll not implemented")
}
func (UnimplementedPointsServer) Recommend(context.Context, *RecommendPoints) (*RecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedPointsServer) Count(context.Context, *CountPoints) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedPointsServer) mustEmbedUnimplementedPointsServer() {}

// UnsafePointsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointsServer will
// result in compilation errors.
type UnsafePointsServer interface {
	mustEmbedUnimplementedPointsServer()
}

func RegisterPointsServer(s grpc.ServiceRegistrar, srv PointsServer) {
	s.RegisterService(&Points_ServiceDesc, srv)
}

func _Points_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Upsert(ctx, req.(*UpsertPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Delete(ctx, req.(*DeletePoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Get(ctx, req.(*GetPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_SetPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).SetPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/SetPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).SetPayload(ctx, req.(*SetPayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_DeletePayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).DeletePayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/DeletePayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).DeletePayload(ctx, req.(*DeletePayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_ClearPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearPayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).ClearPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/ClearPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).ClearPayload(ctx, req.(*ClearPayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_CreateFieldIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldIndexCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).CreateFieldIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/CreateFieldIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).CreateFieldIndex(ctx, req.(*CreateFieldIndexCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_DeleteFieldIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFieldIndexCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).DeleteFieldIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/DeleteFieldIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).DeleteFieldIndex(ctx, req.(*DeleteFieldIndexCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Search(ctx, req.(*SearchPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Scroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrollPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Scroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Scroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Scroll(ctx, req.(*ScrollPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Recommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Recommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Recommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Recommend(ctx, req.(*RecommendPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Points/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Count(ctx, req.(*CountPoints))
	}
	return interceptor(ctx, in, info, handler)
}

// Points_ServiceDesc is the grpc.ServiceDesc for Points service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Points_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qdrant.Points",
	HandlerType: (*PointsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _Points_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Points_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Points_Get_Handler,
		},
		{
			MethodName: "SetPayload",
			Handler:    _Points_SetPayload_Handler,
		},
		{
			MethodName: "DeletePayload",
			Handler:    _Points_DeletePayload_Handler,
		},
		{
			MethodName: "ClearPayload",
			Handler:    _Points_ClearPayload_Handler,
		},
		{
			MethodName: "CreateFieldIndex",
			Handler:    _Points_CreateFieldIndex_Handler,
		},
		{
			MethodName: "DeleteFieldIndex",
			Handler:    _Points_DeleteFieldIndex_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Points_Search_Handler,
		},
		{
			MethodName: "Scroll",
			Handler:    _Points_Scroll_Handler,
		},
		{
			MethodName: "Recommend",
			Handler:    _Points_Recommend_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _Points_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "points_service.proto",
}
