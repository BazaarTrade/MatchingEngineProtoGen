// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/matchingEngine.proto

package pbM

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MatchingEngine_PlaceOrder_FullMethodName              = "/pbM.MatchingEngine/PlaceOrder"
	MatchingEngine_CancelOrder_FullMethodName             = "/pbM.MatchingEngine/CancelOrder"
	MatchingEngine_GetCurrentOrders_FullMethodName        = "/pbM.MatchingEngine/GetCurrentOrders"
	MatchingEngine_GetOrders_FullMethodName               = "/pbM.MatchingEngine/GetOrders"
	MatchingEngine_CreateOrderBook_FullMethodName         = "/pbM.MatchingEngine/CreateOrderBook"
	MatchingEngine_DeleteOrderBook_FullMethodName         = "/pbM.MatchingEngine/DeleteOrderBook"
	MatchingEngine_StreamTrades_FullMethodName            = "/pbM.MatchingEngine/StreamTrades"
	MatchingEngine_StreamOrderBookSnapshot_FullMethodName = "/pbM.MatchingEngine/StreamOrderBookSnapshot"
)

// MatchingEngineClient is the client API for MatchingEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchingEngineClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*Orders, error)
	CancelOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*Order, error)
	GetCurrentOrders(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Orders, error)
	GetOrders(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Orders, error)
	CreateOrderBook(ctx context.Context, in *OrderBookSymbol, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteOrderBook(ctx context.Context, in *OrderBookSymbol, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StreamTrades(ctx context.Context, in *Pair, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Trades], error)
	StreamOrderBookSnapshot(ctx context.Context, in *Pair, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderBookSnapshot], error)
}

type matchingEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchingEngineClient(cc grpc.ClientConnInterface) MatchingEngineClient {
	return &matchingEngineClient{cc}
}

func (c *matchingEngineClient) PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*Orders, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Orders)
	err := c.cc.Invoke(ctx, MatchingEngine_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) CancelOrder(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, MatchingEngine_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) GetCurrentOrders(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Orders, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Orders)
	err := c.cc.Invoke(ctx, MatchingEngine_GetCurrentOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) GetOrders(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Orders, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Orders)
	err := c.cc.Invoke(ctx, MatchingEngine_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) CreateOrderBook(ctx context.Context, in *OrderBookSymbol, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MatchingEngine_CreateOrderBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) DeleteOrderBook(ctx context.Context, in *OrderBookSymbol, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MatchingEngine_DeleteOrderBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingEngineClient) StreamTrades(ctx context.Context, in *Pair, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Trades], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MatchingEngine_ServiceDesc.Streams[0], MatchingEngine_StreamTrades_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Pair, Trades]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MatchingEngine_StreamTradesClient = grpc.ServerStreamingClient[Trades]

func (c *matchingEngineClient) StreamOrderBookSnapshot(ctx context.Context, in *Pair, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderBookSnapshot], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MatchingEngine_ServiceDesc.Streams[1], MatchingEngine_StreamOrderBookSnapshot_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Pair, OrderBookSnapshot]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MatchingEngine_StreamOrderBookSnapshotClient = grpc.ServerStreamingClient[OrderBookSnapshot]

// MatchingEngineServer is the server API for MatchingEngine service.
// All implementations must embed UnimplementedMatchingEngineServer
// for forward compatibility.
type MatchingEngineServer interface {
	PlaceOrder(context.Context, *PlaceOrderReq) (*Orders, error)
	CancelOrder(context.Context, *OrderID) (*Order, error)
	GetCurrentOrders(context.Context, *UserID) (*Orders, error)
	GetOrders(context.Context, *UserID) (*Orders, error)
	CreateOrderBook(context.Context, *OrderBookSymbol) (*emptypb.Empty, error)
	DeleteOrderBook(context.Context, *OrderBookSymbol) (*emptypb.Empty, error)
	StreamTrades(*Pair, grpc.ServerStreamingServer[Trades]) error
	StreamOrderBookSnapshot(*Pair, grpc.ServerStreamingServer[OrderBookSnapshot]) error
	mustEmbedUnimplementedMatchingEngineServer()
}

// UnimplementedMatchingEngineServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMatchingEngineServer struct{}

func (UnimplementedMatchingEngineServer) PlaceOrder(context.Context, *PlaceOrderReq) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedMatchingEngineServer) CancelOrder(context.Context, *OrderID) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedMatchingEngineServer) GetCurrentOrders(context.Context, *UserID) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentOrders not implemented")
}
func (UnimplementedMatchingEngineServer) GetOrders(context.Context, *UserID) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedMatchingEngineServer) CreateOrderBook(context.Context, *OrderBookSymbol) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderBook not implemented")
}
func (UnimplementedMatchingEngineServer) DeleteOrderBook(context.Context, *OrderBookSymbol) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderBook not implemented")
}
func (UnimplementedMatchingEngineServer) StreamTrades(*Pair, grpc.ServerStreamingServer[Trades]) error {
	return status.Errorf(codes.Unimplemented, "method StreamTrades not implemented")
}
func (UnimplementedMatchingEngineServer) StreamOrderBookSnapshot(*Pair, grpc.ServerStreamingServer[OrderBookSnapshot]) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderBookSnapshot not implemented")
}
func (UnimplementedMatchingEngineServer) mustEmbedUnimplementedMatchingEngineServer() {}
func (UnimplementedMatchingEngineServer) testEmbeddedByValue()                        {}

// UnsafeMatchingEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchingEngineServer will
// result in compilation errors.
type UnsafeMatchingEngineServer interface {
	mustEmbedUnimplementedMatchingEngineServer()
}

func RegisterMatchingEngineServer(s grpc.ServiceRegistrar, srv MatchingEngineServer) {
	// If the following call pancis, it indicates UnimplementedMatchingEngineServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MatchingEngine_ServiceDesc, srv)
}

func _MatchingEngine_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).PlaceOrder(ctx, req.(*PlaceOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).CancelOrder(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_GetCurrentOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).GetCurrentOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_GetCurrentOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).GetCurrentOrders(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).GetOrders(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_CreateOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderBookSymbol)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).CreateOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_CreateOrderBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).CreateOrderBook(ctx, req.(*OrderBookSymbol))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_DeleteOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderBookSymbol)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingEngineServer).DeleteOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchingEngine_DeleteOrderBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingEngineServer).DeleteOrderBook(ctx, req.(*OrderBookSymbol))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingEngine_StreamTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pair)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchingEngineServer).StreamTrades(m, &grpc.GenericServerStream[Pair, Trades]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MatchingEngine_StreamTradesServer = grpc.ServerStreamingServer[Trades]

func _MatchingEngine_StreamOrderBookSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pair)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchingEngineServer).StreamOrderBookSnapshot(m, &grpc.GenericServerStream[Pair, OrderBookSnapshot]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MatchingEngine_StreamOrderBookSnapshotServer = grpc.ServerStreamingServer[OrderBookSnapshot]

// MatchingEngine_ServiceDesc is the grpc.ServiceDesc for MatchingEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchingEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbM.MatchingEngine",
	HandlerType: (*MatchingEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _MatchingEngine_PlaceOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _MatchingEngine_CancelOrder_Handler,
		},
		{
			MethodName: "GetCurrentOrders",
			Handler:    _MatchingEngine_GetCurrentOrders_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _MatchingEngine_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrderBook",
			Handler:    _MatchingEngine_CreateOrderBook_Handler,
		},
		{
			MethodName: "DeleteOrderBook",
			Handler:    _MatchingEngine_DeleteOrderBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTrades",
			Handler:       _MatchingEngine_StreamTrades_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrderBookSnapshot",
			Handler:       _MatchingEngine_StreamOrderBookSnapshot_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/matchingEngine.proto",
}
