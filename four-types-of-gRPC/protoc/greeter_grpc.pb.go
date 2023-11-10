// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: protoc/greeter.proto

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

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Sayhello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	// rpc SayHello(HelloRequest) returns (HelloResponse);
	SayHelloServerStreaming(ctx context.Context, in *Namelist, opts ...grpc.CallOption) (GreeterService_SayHelloServerStreamingClient, error)
	SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayHelloClientStreamingClient, error)
	SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayHelloBidirectionalStreamingClient, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Sayhello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreeterService/Sayhello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) SayHelloServerStreaming(ctx context.Context, in *Namelist, opts ...grpc.CallOption) (GreeterService_SayHelloServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[0], "/greet_service.GreeterService/SayHelloServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayHelloServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterService_SayHelloServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterServiceSayHelloServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greeterServiceSayHelloServerStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterServiceClient) SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayHelloClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[1], "/greet_service.GreeterService/SayHelloClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayHelloClientStreamingClient{stream}
	return x, nil
}

type GreeterService_SayHelloClientStreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessageList, error)
	grpc.ClientStream
}

type greeterServiceSayHelloClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greeterServiceSayHelloClientStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterServiceSayHelloClientStreamingClient) CloseAndRecv() (*MessageList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterServiceClient) SayHelloBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreeterService_SayHelloBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[2], "/greet_service.GreeterService/SayHelloBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceSayHelloBidirectionalStreamingClient{stream}
	return x, nil
}

type GreeterService_SayHelloBidirectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterServiceSayHelloBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greeterServiceSayHelloBidirectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterServiceSayHelloBidirectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	Sayhello(context.Context, *NoParam) (*HelloResponse, error)
	// rpc SayHello(HelloRequest) returns (HelloResponse);
	SayHelloServerStreaming(*Namelist, GreeterService_SayHelloServerStreamingServer) error
	SayHelloClientStreaming(GreeterService_SayHelloClientStreamingServer) error
	SayHelloBidirectionalStreaming(GreeterService_SayHelloBidirectionalStreamingServer) error
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) Sayhello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayhello not implemented")
}
func (UnimplementedGreeterServiceServer) SayHelloServerStreaming(*Namelist, GreeterService_SayHelloServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStreaming not implemented")
}
func (UnimplementedGreeterServiceServer) SayHelloClientStreaming(GreeterService_SayHelloClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStreaming not implemented")
}
func (UnimplementedGreeterServiceServer) SayHelloBidirectionalStreaming(GreeterService_SayHelloBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStreaming not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_Sayhello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Sayhello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreeterService/Sayhello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Sayhello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_SayHelloServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Namelist)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServiceServer).SayHelloServerStreaming(m, &greeterServiceSayHelloServerStreamingServer{stream})
}

type GreeterService_SayHelloServerStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterServiceSayHelloServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greeterServiceSayHelloServerStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreeterService_SayHelloClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServiceServer).SayHelloClientStreaming(&greeterServiceSayHelloClientStreamingServer{stream})
}

type GreeterService_SayHelloClientStreamingServer interface {
	SendAndClose(*MessageList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterServiceSayHelloClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greeterServiceSayHelloClientStreamingServer) SendAndClose(m *MessageList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterServiceSayHelloClientStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreeterService_SayHelloBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServiceServer).SayHelloBidirectionalStreaming(&greeterServiceSayHelloBidirectionalStreamingServer{stream})
}

type GreeterService_SayHelloBidirectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterServiceSayHelloBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greeterServiceSayHelloBidirectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterServiceSayHelloBidirectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sayhello",
			Handler:    _GreeterService_Sayhello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStreaming",
			Handler:       _GreeterService_SayHelloServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStreaming",
			Handler:       _GreeterService_SayHelloClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStreaming",
			Handler:       _GreeterService_SayHelloBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protoc/greeter.proto",
}
