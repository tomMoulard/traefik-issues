// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/whoami.proto

package whoamigrpc

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

// WhoamiClient is the client API for Whoami service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhoamiClient interface {
	// Sends a greeting
	AskWhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error)
}

type whoamiClient struct {
	cc grpc.ClientConnInterface
}

func NewWhoamiClient(cc grpc.ClientConnInterface) WhoamiClient {
	return &whoamiClient{cc}
}

func (c *whoamiClient) AskWhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error) {
	out := new(WhoAmIReply)
	err := c.cc.Invoke(ctx, "/whoami.Whoami/AskWhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhoamiServer is the server API for Whoami service.
// All implementations must embed UnimplementedWhoamiServer
// for forward compatibility
type WhoamiServer interface {
	// Sends a greeting
	AskWhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error)
	mustEmbedUnimplementedWhoamiServer()
}

// UnimplementedWhoamiServer must be embedded to have forward compatible implementations.
type UnimplementedWhoamiServer struct {
}

func (UnimplementedWhoamiServer) AskWhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskWhoAmI not implemented")
}
func (UnimplementedWhoamiServer) mustEmbedUnimplementedWhoamiServer() {}

// UnsafeWhoamiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhoamiServer will
// result in compilation errors.
type UnsafeWhoamiServer interface {
	mustEmbedUnimplementedWhoamiServer()
}

func RegisterWhoamiServer(s grpc.ServiceRegistrar, srv WhoamiServer) {
	s.RegisterService(&Whoami_ServiceDesc, srv)
}

func _Whoami_AskWhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhoamiServer).AskWhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/whoami.Whoami/AskWhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhoamiServer).AskWhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Whoami_ServiceDesc is the grpc.ServiceDesc for Whoami service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Whoami_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "whoami.Whoami",
	HandlerType: (*WhoamiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskWhoAmI",
			Handler:    _Whoami_AskWhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/whoami.proto",
}
