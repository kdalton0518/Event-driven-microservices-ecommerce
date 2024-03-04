// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: protos/order_service.proto

package pb

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
	OrderService_GetManyProducts_FullMethodName = "/OrderService/GetManyProducts"
	OrderService_GetProduct_FullMethodName      = "/OrderService/GetProduct"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetManyProducts(ctx context.Context, in *GetManyOrdersRequest, opts ...grpc.CallOption) (*GetManyOrdersResponse, error)
	GetProduct(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetManyProducts(ctx context.Context, in *GetManyOrdersRequest, opts ...grpc.CallOption) (*GetManyOrdersResponse, error) {
	out := new(GetManyOrdersResponse)
	err := c.cc.Invoke(ctx, OrderService_GetManyProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetProduct(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetManyProducts(context.Context, *GetManyOrdersRequest) (*GetManyOrdersResponse, error)
	GetProduct(context.Context, *GetOrderRequest) (*OrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetManyProducts(context.Context, *GetManyOrdersRequest) (*GetManyOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyProducts not implemented")
}
func (UnimplementedOrderServiceServer) GetProduct(context.Context, *GetOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetManyProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetManyProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetManyProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetManyProducts(ctx, req.(*GetManyOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetProduct(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManyProducts",
			Handler:    _OrderService_GetManyProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _OrderService_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order_service.proto",
}