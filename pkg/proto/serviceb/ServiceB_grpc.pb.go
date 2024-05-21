// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: ServiceB.proto

package serviceb

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

// ServiceBClient is the client API for ServiceB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceBClient interface {
	Server(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type serviceBClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceBClient(cc grpc.ClientConnInterface) ServiceBClient {
	return &serviceBClient{cc}
}

func (c *serviceBClient) Server(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ServiceB.ServiceB/Server", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceBServer is the server API for ServiceB service.
// All implementations should embed UnimplementedServiceBServer
// for forward compatibility
type ServiceBServer interface {
	Server(context.Context, *Request) (*Response, error)
}

// UnimplementedServiceBServer should be embedded to have forward compatible implementations.
type UnimplementedServiceBServer struct {
}

func (UnimplementedServiceBServer) Server(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Server not implemented")
}

// UnsafeServiceBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceBServer will
// result in compilation errors.
type UnsafeServiceBServer interface {
	mustEmbedUnimplementedServiceBServer()
}

func RegisterServiceBServer(s grpc.ServiceRegistrar, srv ServiceBServer) {
	s.RegisterService(&ServiceB_ServiceDesc, srv)
}

func _ServiceB_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceBServer).Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceB.ServiceB/Server",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceBServer).Server(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceB_ServiceDesc is the grpc.ServiceDesc for ServiceB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceB.ServiceB",
	HandlerType: (*ServiceBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Server",
			Handler:    _ServiceB_Server_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ServiceB.proto",
}