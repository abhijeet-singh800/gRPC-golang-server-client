// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: calc.proto

package calc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CalcService_Add_FullMethodName       = "/calc.CalcService/Add"
	CalcService_Sub_FullMethodName       = "/calc.CalcService/Sub"
	CalcService_Mul_FullMethodName       = "/calc.CalcService/Mul"
	CalcService_Div_FullMethodName       = "/calc.CalcService/Div"
	CalcService_SumUpTo_FullMethodName   = "/calc.CalcService/SumUpTo"
	CalcService_CountUpTo_FullMethodName = "/calc.CalcService/CountUpTo"
)

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Add(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error)
	Sub(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error)
	Mul(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error)
	Div(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error)
	SumUpTo(ctx context.Context, in *NumList, opts ...grpc.CallOption) (*One, error)
	CountUpTo(ctx context.Context, in *One, opts ...grpc.CallOption) (*NumList, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Add(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(One)
	err := c.cc.Invoke(ctx, CalcService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Sub(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(One)
	err := c.cc.Invoke(ctx, CalcService_Sub_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Mul(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(One)
	err := c.cc.Invoke(ctx, CalcService_Mul_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) Div(ctx context.Context, in *Two, opts ...grpc.CallOption) (*One, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(One)
	err := c.cc.Invoke(ctx, CalcService_Div_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) SumUpTo(ctx context.Context, in *NumList, opts ...grpc.CallOption) (*One, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(One)
	err := c.cc.Invoke(ctx, CalcService_SumUpTo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) CountUpTo(ctx context.Context, in *One, opts ...grpc.CallOption) (*NumList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NumList)
	err := c.cc.Invoke(ctx, CalcService_CountUpTo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility.
type CalcServiceServer interface {
	Add(context.Context, *Two) (*One, error)
	Sub(context.Context, *Two) (*One, error)
	Mul(context.Context, *Two) (*One, error)
	Div(context.Context, *Two) (*One, error)
	SumUpTo(context.Context, *NumList) (*One, error)
	CountUpTo(context.Context, *One) (*NumList, error)
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalcServiceServer struct{}

func (UnimplementedCalcServiceServer) Add(context.Context, *Two) (*One, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalcServiceServer) Sub(context.Context, *Two) (*One, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (UnimplementedCalcServiceServer) Mul(context.Context, *Two) (*One, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (UnimplementedCalcServiceServer) Div(context.Context, *Two) (*One, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (UnimplementedCalcServiceServer) SumUpTo(context.Context, *NumList) (*One, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumUpTo not implemented")
}
func (UnimplementedCalcServiceServer) CountUpTo(context.Context, *One) (*NumList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUpTo not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}
func (UnimplementedCalcServiceServer) testEmbeddedByValue()                     {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Two)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Add(ctx, req.(*Two))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Two)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_Sub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sub(ctx, req.(*Two))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Two)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_Mul_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Mul(ctx, req.(*Two))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Two)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_Div_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Div(ctx, req.(*Two))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_SumUpTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).SumUpTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_SumUpTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).SumUpTo(ctx, req.(*NumList))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_CountUpTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(One)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).CountUpTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcService_CountUpTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).CountUpTo(ctx, req.(*One))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalcService_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _CalcService_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _CalcService_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _CalcService_Div_Handler,
		},
		{
			MethodName: "SumUpTo",
			Handler:    _CalcService_SumUpTo_Handler,
		},
		{
			MethodName: "CountUpTo",
			Handler:    _CalcService_CountUpTo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}
