// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: feature_store.proto

package proto

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

// FeatureStoreClient is the client API for FeatureStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureStoreClient interface {
	HeartBeat(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PongResp, error)
	GetSkuFeatures(ctx context.Context, in *GetSkuFeaturesReq, opts ...grpc.CallOption) (*GetSkuFeaturesResp, error)
	GetSkuFeaturesForAllCatalogues(ctx context.Context, in *GetSkuFeaturesForAllCataloguesReq, opts ...grpc.CallOption) (*GetSkuFeaturesForAllCataloguesResp, error)
	UpsertSkuFeatures(ctx context.Context, in *UpsertSkuFeaturesReq, opts ...grpc.CallOption) (*UpsertSkuFeaturesResp, error)
	Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error)
	Release(ctx context.Context, in *ReleaseReq, opts ...grpc.CallOption) (*ReleaseResp, error)
	OrderExists(ctx context.Context, in *OrderExistsReq, opts ...grpc.CallOption) (*OrderExistsResp, error)
	MarkOrder(ctx context.Context, in *MarkOrderReq, opts ...grpc.CallOption) (*MarkOrderResp, error)
	ImpressionExists(ctx context.Context, in *ImpressionExistsReq, opts ...grpc.CallOption) (*ImpressionExistsResp, error)
	MarkImpression(ctx context.Context, in *MarkImpressionReq, opts ...grpc.CallOption) (*MarkImpressionResp, error)
}

type featureStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureStoreClient(cc grpc.ClientConnInterface) FeatureStoreClient {
	return &featureStoreClient{cc}
}

func (c *featureStoreClient) HeartBeat(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PongResp, error) {
	out := new(PongResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) GetSkuFeatures(ctx context.Context, in *GetSkuFeaturesReq, opts ...grpc.CallOption) (*GetSkuFeaturesResp, error) {
	out := new(GetSkuFeaturesResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/GetSkuFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) GetSkuFeaturesForAllCatalogues(ctx context.Context, in *GetSkuFeaturesForAllCataloguesReq, opts ...grpc.CallOption) (*GetSkuFeaturesForAllCataloguesResp, error) {
	out := new(GetSkuFeaturesForAllCataloguesResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/GetSkuFeaturesForAllCatalogues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) UpsertSkuFeatures(ctx context.Context, in *UpsertSkuFeaturesReq, opts ...grpc.CallOption) (*UpsertSkuFeaturesResp, error) {
	out := new(UpsertSkuFeaturesResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/UpsertSkuFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) Lock(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*LockResp, error) {
	out := new(LockResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) Release(ctx context.Context, in *ReleaseReq, opts ...grpc.CallOption) (*ReleaseResp, error) {
	out := new(ReleaseResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) OrderExists(ctx context.Context, in *OrderExistsReq, opts ...grpc.CallOption) (*OrderExistsResp, error) {
	out := new(OrderExistsResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/OrderExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) MarkOrder(ctx context.Context, in *MarkOrderReq, opts ...grpc.CallOption) (*MarkOrderResp, error) {
	out := new(MarkOrderResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/MarkOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) ImpressionExists(ctx context.Context, in *ImpressionExistsReq, opts ...grpc.CallOption) (*ImpressionExistsResp, error) {
	out := new(ImpressionExistsResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/ImpressionExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureStoreClient) MarkImpression(ctx context.Context, in *MarkImpressionReq, opts ...grpc.CallOption) (*MarkImpressionResp, error) {
	out := new(MarkImpressionResp)
	err := c.cc.Invoke(ctx, "/feature_store.FeatureStore/MarkImpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureStoreServer is the server API for FeatureStore service.
// All implementations must embed UnimplementedFeatureStoreServer
// for forward compatibility
type FeatureStoreServer interface {
	HeartBeat(context.Context, *PingReq) (*PongResp, error)
	GetSkuFeatures(context.Context, *GetSkuFeaturesReq) (*GetSkuFeaturesResp, error)
	GetSkuFeaturesForAllCatalogues(context.Context, *GetSkuFeaturesForAllCataloguesReq) (*GetSkuFeaturesForAllCataloguesResp, error)
	UpsertSkuFeatures(context.Context, *UpsertSkuFeaturesReq) (*UpsertSkuFeaturesResp, error)
	Lock(context.Context, *LockReq) (*LockResp, error)
	Release(context.Context, *ReleaseReq) (*ReleaseResp, error)
	OrderExists(context.Context, *OrderExistsReq) (*OrderExistsResp, error)
	MarkOrder(context.Context, *MarkOrderReq) (*MarkOrderResp, error)
	ImpressionExists(context.Context, *ImpressionExistsReq) (*ImpressionExistsResp, error)
	MarkImpression(context.Context, *MarkImpressionReq) (*MarkImpressionResp, error)
	mustEmbedUnimplementedFeatureStoreServer()
}

// UnimplementedFeatureStoreServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureStoreServer struct {
}

func (UnimplementedFeatureStoreServer) HeartBeat(context.Context, *PingReq) (*PongResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedFeatureStoreServer) GetSkuFeatures(context.Context, *GetSkuFeaturesReq) (*GetSkuFeaturesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkuFeatures not implemented")
}
func (UnimplementedFeatureStoreServer) GetSkuFeaturesForAllCatalogues(context.Context, *GetSkuFeaturesForAllCataloguesReq) (*GetSkuFeaturesForAllCataloguesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkuFeaturesForAllCatalogues not implemented")
}
func (UnimplementedFeatureStoreServer) UpsertSkuFeatures(context.Context, *UpsertSkuFeaturesReq) (*UpsertSkuFeaturesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertSkuFeatures not implemented")
}
func (UnimplementedFeatureStoreServer) Lock(context.Context, *LockReq) (*LockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedFeatureStoreServer) Release(context.Context, *ReleaseReq) (*ReleaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedFeatureStoreServer) OrderExists(context.Context, *OrderExistsReq) (*OrderExistsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderExists not implemented")
}
func (UnimplementedFeatureStoreServer) MarkOrder(context.Context, *MarkOrderReq) (*MarkOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkOrder not implemented")
}
func (UnimplementedFeatureStoreServer) ImpressionExists(context.Context, *ImpressionExistsReq) (*ImpressionExistsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImpressionExists not implemented")
}
func (UnimplementedFeatureStoreServer) MarkImpression(context.Context, *MarkImpressionReq) (*MarkImpressionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkImpression not implemented")
}
func (UnimplementedFeatureStoreServer) mustEmbedUnimplementedFeatureStoreServer() {}

// UnsafeFeatureStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureStoreServer will
// result in compilation errors.
type UnsafeFeatureStoreServer interface {
	mustEmbedUnimplementedFeatureStoreServer()
}

func RegisterFeatureStoreServer(s grpc.ServiceRegistrar, srv FeatureStoreServer) {
	s.RegisterService(&FeatureStore_ServiceDesc, srv)
}

func _FeatureStore_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).HeartBeat(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_GetSkuFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkuFeaturesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).GetSkuFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/GetSkuFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).GetSkuFeatures(ctx, req.(*GetSkuFeaturesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_GetSkuFeaturesForAllCatalogues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkuFeaturesForAllCataloguesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).GetSkuFeaturesForAllCatalogues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/GetSkuFeaturesForAllCatalogues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).GetSkuFeaturesForAllCatalogues(ctx, req.(*GetSkuFeaturesForAllCataloguesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_UpsertSkuFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertSkuFeaturesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).UpsertSkuFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/UpsertSkuFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).UpsertSkuFeatures(ctx, req.(*UpsertSkuFeaturesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).Lock(ctx, req.(*LockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).Release(ctx, req.(*ReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_OrderExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderExistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).OrderExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/OrderExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).OrderExists(ctx, req.(*OrderExistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_MarkOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).MarkOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/MarkOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).MarkOrder(ctx, req.(*MarkOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_ImpressionExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImpressionExistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).ImpressionExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/ImpressionExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).ImpressionExists(ctx, req.(*ImpressionExistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureStore_MarkImpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkImpressionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureStoreServer).MarkImpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feature_store.FeatureStore/MarkImpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureStoreServer).MarkImpression(ctx, req.(*MarkImpressionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureStore_ServiceDesc is the grpc.ServiceDesc for FeatureStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feature_store.FeatureStore",
	HandlerType: (*FeatureStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _FeatureStore_HeartBeat_Handler,
		},
		{
			MethodName: "GetSkuFeatures",
			Handler:    _FeatureStore_GetSkuFeatures_Handler,
		},
		{
			MethodName: "GetSkuFeaturesForAllCatalogues",
			Handler:    _FeatureStore_GetSkuFeaturesForAllCatalogues_Handler,
		},
		{
			MethodName: "UpsertSkuFeatures",
			Handler:    _FeatureStore_UpsertSkuFeatures_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _FeatureStore_Lock_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _FeatureStore_Release_Handler,
		},
		{
			MethodName: "OrderExists",
			Handler:    _FeatureStore_OrderExists_Handler,
		},
		{
			MethodName: "MarkOrder",
			Handler:    _FeatureStore_MarkOrder_Handler,
		},
		{
			MethodName: "ImpressionExists",
			Handler:    _FeatureStore_ImpressionExists_Handler,
		},
		{
			MethodName: "MarkImpression",
			Handler:    _FeatureStore_MarkImpression_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feature_store.proto",
}
