// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: bing.proto

package bing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Bing_Work_FullMethodName = "/bing.Bing/Work"
)

// BingClient is the client API for Bing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BingClient interface {
	Work(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type bingClient struct {
	cc grpc.ClientConnInterface
}

func NewBingClient(cc grpc.ClientConnInterface) BingClient {
	return &bingClient{cc}
}

func (c *bingClient) Work(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Bing_Work_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BingServer is the server API for Bing service.
// All implementations must embed UnimplementedBingServer
// for forward compatibility
type BingServer interface {
	Work(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedBingServer()
}

// UnimplementedBingServer must be embedded to have forward compatible implementations.
type UnimplementedBingServer struct {
}

func (UnimplementedBingServer) Work(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedBingServer) mustEmbedUnimplementedBingServer() {}

// UnsafeBingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BingServer will
// result in compilation errors.
type UnsafeBingServer interface {
	mustEmbedUnimplementedBingServer()
}

func RegisterBingServer(s grpc.ServiceRegistrar, srv BingServer) {
	s.RegisterService(&Bing_ServiceDesc, srv)
}

func _Bing_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BingServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bing_Work_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BingServer).Work(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Bing_ServiceDesc is the grpc.ServiceDesc for Bing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bing.Bing",
	HandlerType: (*BingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Work",
			Handler:    _Bing_Work_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bing.proto",
}