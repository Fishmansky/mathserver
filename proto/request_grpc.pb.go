// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/request.proto

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

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathServiceClient interface {
	Fibo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	IsPrime(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Fibo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.MathService/Fibo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) IsPrime(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.MathService/IsPrime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
// All implementations should embed UnimplementedMathServiceServer
// for forward compatibility
type MathServiceServer interface {
	Fibo(context.Context, *Request) (*Response, error)
	IsPrime(context.Context, *Request) (*Response, error)
}

// UnimplementedMathServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (UnimplementedMathServiceServer) Fibo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibo not implemented")
}
func (UnimplementedMathServiceServer) IsPrime(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPrime not implemented")
}

// UnsafeMathServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServiceServer will
// result in compilation errors.
type UnsafeMathServiceServer interface {
	mustEmbedUnimplementedMathServiceServer()
}

func RegisterMathServiceServer(s grpc.ServiceRegistrar, srv MathServiceServer) {
	s.RegisterService(&MathService_ServiceDesc, srv)
}

func _MathService_Fibo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Fibo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MathService/Fibo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Fibo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_IsPrime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).IsPrime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MathService/IsPrime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).IsPrime(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MathService_ServiceDesc is the grpc.ServiceDesc for MathService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MathService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibo",
			Handler:    _MathService_Fibo_Handler,
		},
		{
			MethodName: "IsPrime",
			Handler:    _MathService_IsPrime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/request.proto",
}