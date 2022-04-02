// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: helloworld.proto

package helloworld

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

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloworldClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Helloworld_StreamClient, error)
	PingPong(ctx context.Context, opts ...grpc.CallOption) (Helloworld_PingPongClient, error)
}

type helloworldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldClient(cc grpc.ClientConnInterface) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/helloworld.Helloworld/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Helloworld_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Helloworld_ServiceDesc.Streams[0], "/helloworld.Helloworld/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloworldStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Helloworld_StreamClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type helloworldStreamClient struct {
	grpc.ClientStream
}

func (x *helloworldStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloworldClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (Helloworld_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &Helloworld_ServiceDesc.Streams[1], "/helloworld.Helloworld/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloworldPingPongClient{stream}
	return x, nil
}

type Helloworld_PingPongClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type helloworldPingPongClient struct {
	grpc.ClientStream
}

func (x *helloworldPingPongClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloworldPingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloworldServer is the server API for Helloworld service.
// All implementations must embed UnimplementedHelloworldServer
// for forward compatibility
type HelloworldServer interface {
	Call(context.Context, *Request) (*Response, error)
	Stream(*StreamingRequest, Helloworld_StreamServer) error
	PingPong(Helloworld_PingPongServer) error
	mustEmbedUnimplementedHelloworldServer()
}

// UnimplementedHelloworldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (UnimplementedHelloworldServer) Call(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedHelloworldServer) Stream(*StreamingRequest, Helloworld_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedHelloworldServer) PingPong(Helloworld_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedHelloworldServer) mustEmbedUnimplementedHelloworldServer() {}

// UnsafeHelloworldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloworldServer will
// result in compilation errors.
type UnsafeHelloworldServer interface {
	mustEmbedUnimplementedHelloworldServer()
}

func RegisterHelloworldServer(s grpc.ServiceRegistrar, srv HelloworldServer) {
	s.RegisterService(&Helloworld_ServiceDesc, srv)
}

func _Helloworld_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Helloworld/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloworldServer).Stream(m, &helloworldStreamServer{stream})
}

type Helloworld_StreamServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type helloworldStreamServer struct {
	grpc.ServerStream
}

func (x *helloworldStreamServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Helloworld_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloworldServer).PingPong(&helloworldPingPongServer{stream})
}

type Helloworld_PingPongServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type helloworldPingPongServer struct {
	grpc.ServerStream
}

func (x *helloworldPingPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloworldPingPongServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Helloworld_ServiceDesc is the grpc.ServiceDesc for Helloworld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Helloworld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Helloworld_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Helloworld_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPong",
			Handler:       _Helloworld_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
