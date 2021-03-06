// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlClient interface {
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	SetupTest(ctx context.Context, in *SetupTestRequest, opts ...grpc.CallOption) (*SetupTestResponse, error)
	ExecuteTest(ctx context.Context, in *ExecuteTestRequest, opts ...grpc.CallOption) (*ExecuteTestResponse, error)
	TeardownTest(ctx context.Context, in *TeardownTestRequest, opts ...grpc.CallOption) (*TeardownTestResponse, error)
}

type controlClient struct {
	cc grpc.ClientConnInterface
}

func NewControlClient(cc grpc.ClientConnInterface) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/grpc.Control/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) SetupTest(ctx context.Context, in *SetupTestRequest, opts ...grpc.CallOption) (*SetupTestResponse, error) {
	out := new(SetupTestResponse)
	err := c.cc.Invoke(ctx, "/grpc.Control/SetupTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) ExecuteTest(ctx context.Context, in *ExecuteTestRequest, opts ...grpc.CallOption) (*ExecuteTestResponse, error) {
	out := new(ExecuteTestResponse)
	err := c.cc.Invoke(ctx, "/grpc.Control/ExecuteTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) TeardownTest(ctx context.Context, in *TeardownTestRequest, opts ...grpc.CallOption) (*TeardownTestResponse, error) {
	out := new(TeardownTestResponse)
	err := c.cc.Invoke(ctx, "/grpc.Control/TeardownTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
// All implementations must embed UnimplementedControlServer
// for forward compatibility
type ControlServer interface {
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	SetupTest(context.Context, *SetupTestRequest) (*SetupTestResponse, error)
	ExecuteTest(context.Context, *ExecuteTestRequest) (*ExecuteTestResponse, error)
	TeardownTest(context.Context, *TeardownTestRequest) (*TeardownTestResponse, error)
	mustEmbedUnimplementedControlServer()
}

// UnimplementedControlServer must be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (UnimplementedControlServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedControlServer) SetupTest(context.Context, *SetupTestRequest) (*SetupTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupTest not implemented")
}
func (UnimplementedControlServer) ExecuteTest(context.Context, *ExecuteTestRequest) (*ExecuteTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteTest not implemented")
}
func (UnimplementedControlServer) TeardownTest(context.Context, *TeardownTestRequest) (*TeardownTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeardownTest not implemented")
}
func (UnimplementedControlServer) mustEmbedUnimplementedControlServer() {}

// UnsafeControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServer will
// result in compilation errors.
type UnsafeControlServer interface {
	mustEmbedUnimplementedControlServer()
}

func RegisterControlServer(s grpc.ServiceRegistrar, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Control/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_SetupTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).SetupTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Control/SetupTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).SetupTest(ctx, req.(*SetupTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_ExecuteTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).ExecuteTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Control/ExecuteTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).ExecuteTest(ctx, req.(*ExecuteTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_TeardownTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).TeardownTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Control/TeardownTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).TeardownTest(ctx, req.(*TeardownTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _Control_Shutdown_Handler,
		},
		{
			MethodName: "SetupTest",
			Handler:    _Control_SetupTest_Handler,
		},
		{
			MethodName: "ExecuteTest",
			Handler:    _Control_ExecuteTest_Handler,
		},
		{
			MethodName: "TeardownTest",
			Handler:    _Control_TeardownTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
