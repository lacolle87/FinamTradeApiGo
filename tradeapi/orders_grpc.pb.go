// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: grpc/tradeapi/v1/orders.proto

package tradeapi

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

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersClient interface {
	// Creates new order.
	// Order placement in OrderBook takes some time due to processing speed,
	// that is why this method returns transaction_id, which can be used
	// to find corresponding order in GetOrdersRequest or in OrderEvent message
	// of Events service (EventResponse.event.order).
	// Создать новую заявку.
	// На обработку нового поручения по размещению заявки в биржевой стакан
	// требуется некоторое время, поэтому этот метод возвращает структуру с
	// transaction_id, которая может быть использована для поиска соответствующей
	// заявки через GetOrdersRequest или в сообщении OrderEvent от сервиса событий
	// (EventResponse.event.order).
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResult, error)
	// Cancels order.
	// Отменяет заявку.
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResult, error)
	// Returns the list of orders.
	// Возвращает список заявок.
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResult, error)
}

type ordersClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersClient(cc grpc.ClientConnInterface) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResult, error) {
	out := new(NewOrderResult)
	err := c.cc.Invoke(ctx, "/grpc.tradeapi.v1.Orders/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResult, error) {
	out := new(CancelOrderResult)
	err := c.cc.Invoke(ctx, "/grpc.tradeapi.v1.Orders/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResult, error) {
	out := new(GetOrdersResult)
	err := c.cc.Invoke(ctx, "/grpc.tradeapi.v1.Orders/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
// All implementations must embed UnimplementedOrdersServer
// for forward compatibility
type OrdersServer interface {
	// Creates new order.
	// Order placement in OrderBook takes some time due to processing speed,
	// that is why this method returns transaction_id, which can be used
	// to find corresponding order in GetOrdersRequest or in OrderEvent message
	// of Events service (EventResponse.event.order).
	// Создать новую заявку.
	// На обработку нового поручения по размещению заявки в биржевой стакан
	// требуется некоторое время, поэтому этот метод возвращает структуру с
	// transaction_id, которая может быть использована для поиска соответствующей
	// заявки через GetOrdersRequest или в сообщении OrderEvent от сервиса событий
	// (EventResponse.event.order).
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResult, error)
	// Cancels order.
	// Отменяет заявку.
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResult, error)
	// Returns the list of orders.
	// Возвращает список заявок.
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResult, error)
	mustEmbedUnimplementedOrdersServer()
}

// UnimplementedOrdersServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServer struct {
}

func (UnimplementedOrdersServer) NewOrder(context.Context, *NewOrderRequest) (*NewOrderResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedOrdersServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrdersServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrdersServer) mustEmbedUnimplementedOrdersServer() {}

// UnsafeOrdersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServer will
// result in compilation errors.
type UnsafeOrdersServer interface {
	mustEmbedUnimplementedOrdersServer()
}

func RegisterOrdersServer(s grpc.ServiceRegistrar, srv OrdersServer) {
	s.RegisterService(&Orders_ServiceDesc, srv)
}

func _Orders_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.tradeapi.v1.Orders/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.tradeapi.v1.Orders/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.tradeapi.v1.Orders/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orders_ServiceDesc is the grpc.ServiceDesc for Orders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.tradeapi.v1.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _Orders_NewOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Orders_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Orders_GetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/tradeapi/v1/orders.proto",
}
