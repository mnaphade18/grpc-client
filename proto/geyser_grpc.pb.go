// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: geyser.proto

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

// GeyserClient is the client API for Geyser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeyserClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Geyser_SubscribeClient, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	GetLatestBlockhash(ctx context.Context, in *GetLatestBlockhashRequest, opts ...grpc.CallOption) (*GetLatestBlockhashResponse, error)
	GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error)
	GetSlot(ctx context.Context, in *GetSlotRequest, opts ...grpc.CallOption) (*GetSlotResponse, error)
	IsBlockhashValid(ctx context.Context, in *IsBlockhashValidRequest, opts ...grpc.CallOption) (*IsBlockhashValidResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type geyserClient struct {
	cc grpc.ClientConnInterface
}

func NewGeyserClient(cc grpc.ClientConnInterface) GeyserClient {
	return &geyserClient{cc}
}

func (c *geyserClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Geyser_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Geyser_ServiceDesc.Streams[0], "/geyser.Geyser/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &geyserSubscribeClient{stream}
	return x, nil
}

type Geyser_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeUpdate, error)
	grpc.ClientStream
}

type geyserSubscribeClient struct {
	grpc.ClientStream
}

func (x *geyserSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *geyserSubscribeClient) Recv() (*SubscribeUpdate, error) {
	m := new(SubscribeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *geyserClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geyserClient) GetLatestBlockhash(ctx context.Context, in *GetLatestBlockhashRequest, opts ...grpc.CallOption) (*GetLatestBlockhashResponse, error) {
	out := new(GetLatestBlockhashResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/GetLatestBlockhash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geyserClient) GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error) {
	out := new(GetBlockHeightResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/GetBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geyserClient) GetSlot(ctx context.Context, in *GetSlotRequest, opts ...grpc.CallOption) (*GetSlotResponse, error) {
	out := new(GetSlotResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/GetSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geyserClient) IsBlockhashValid(ctx context.Context, in *IsBlockhashValidRequest, opts ...grpc.CallOption) (*IsBlockhashValidResponse, error) {
	out := new(IsBlockhashValidResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/IsBlockhashValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geyserClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/geyser.Geyser/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeyserServer is the server API for Geyser service.
// All implementations must embed UnimplementedGeyserServer
// for forward compatibility
type GeyserServer interface {
	Subscribe(Geyser_SubscribeServer) error
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	GetLatestBlockhash(context.Context, *GetLatestBlockhashRequest) (*GetLatestBlockhashResponse, error)
	GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error)
	GetSlot(context.Context, *GetSlotRequest) (*GetSlotResponse, error)
	IsBlockhashValid(context.Context, *IsBlockhashValidRequest) (*IsBlockhashValidResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedGeyserServer()
}

// UnimplementedGeyserServer must be embedded to have forward compatible implementations.
type UnimplementedGeyserServer struct {
}

func (UnimplementedGeyserServer) Subscribe(Geyser_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGeyserServer) Ping(context.Context, *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGeyserServer) GetLatestBlockhash(context.Context, *GetLatestBlockhashRequest) (*GetLatestBlockhashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlockhash not implemented")
}
func (UnimplementedGeyserServer) GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (UnimplementedGeyserServer) GetSlot(context.Context, *GetSlotRequest) (*GetSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlot not implemented")
}
func (UnimplementedGeyserServer) IsBlockhashValid(context.Context, *IsBlockhashValidRequest) (*IsBlockhashValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlockhashValid not implemented")
}
func (UnimplementedGeyserServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedGeyserServer) mustEmbedUnimplementedGeyserServer() {}

// UnsafeGeyserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeyserServer will
// result in compilation errors.
type UnsafeGeyserServer interface {
	mustEmbedUnimplementedGeyserServer()
}

func RegisterGeyserServer(s grpc.ServiceRegistrar, srv GeyserServer) {
	s.RegisterService(&Geyser_ServiceDesc, srv)
}

func _Geyser_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GeyserServer).Subscribe(&geyserSubscribeServer{stream})
}

type Geyser_SubscribeServer interface {
	Send(*SubscribeUpdate) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type geyserSubscribeServer struct {
	grpc.ServerStream
}

func (x *geyserSubscribeServer) Send(m *SubscribeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *geyserSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Geyser_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geyser_GetLatestBlockhash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockhashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).GetLatestBlockhash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/GetLatestBlockhash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).GetLatestBlockhash(ctx, req.(*GetLatestBlockhashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geyser_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/GetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).GetBlockHeight(ctx, req.(*GetBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geyser_GetSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).GetSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/GetSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).GetSlot(ctx, req.(*GetSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geyser_IsBlockhashValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBlockhashValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).IsBlockhashValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/IsBlockhashValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).IsBlockhashValid(ctx, req.(*IsBlockhashValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geyser_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeyserServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geyser.Geyser/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeyserServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Geyser_ServiceDesc is the grpc.ServiceDesc for Geyser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geyser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geyser.Geyser",
	HandlerType: (*GeyserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Geyser_Ping_Handler,
		},
		{
			MethodName: "GetLatestBlockhash",
			Handler:    _Geyser_GetLatestBlockhash_Handler,
		},
		{
			MethodName: "GetBlockHeight",
			Handler:    _Geyser_GetBlockHeight_Handler,
		},
		{
			MethodName: "GetSlot",
			Handler:    _Geyser_GetSlot_Handler,
		},
		{
			MethodName: "IsBlockhashValid",
			Handler:    _Geyser_IsBlockhashValid_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Geyser_GetVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Geyser_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "geyser.proto",
}
