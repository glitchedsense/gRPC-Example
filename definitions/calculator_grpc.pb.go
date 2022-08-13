// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: definitions/calculator.proto

package __

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

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Calculate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Calculate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/Calculator/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Calculate(context.Context, *Input) (*Output, error)
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Calculate(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Calculate(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calculator_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definitions/calculator.proto",
}
