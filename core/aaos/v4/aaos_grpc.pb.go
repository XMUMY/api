// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v4

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AAOSClient is the client API for AAOS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AAOSClient interface {
	GetInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Info, error)
	GetTimetable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Timetable, error)
	GetCourses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Courses, error)
	GetExams(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Exams, error)
	GetTranscript(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Transcript, error)
}

type aAOSClient struct {
	cc grpc.ClientConnInterface
}

func NewAAOSClient(cc grpc.ClientConnInterface) AAOSClient {
	return &aAOSClient{cc}
}

func (c *aAOSClient) GetInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/xmux.core.aaos.v4.AAOS/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aAOSClient) GetTimetable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Timetable, error) {
	out := new(Timetable)
	err := c.cc.Invoke(ctx, "/xmux.core.aaos.v4.AAOS/GetTimetable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aAOSClient) GetCourses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/xmux.core.aaos.v4.AAOS/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aAOSClient) GetExams(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Exams, error) {
	out := new(Exams)
	err := c.cc.Invoke(ctx, "/xmux.core.aaos.v4.AAOS/GetExams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aAOSClient) GetTranscript(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Transcript, error) {
	out := new(Transcript)
	err := c.cc.Invoke(ctx, "/xmux.core.aaos.v4.AAOS/GetTranscript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AAOSServer is the server API for AAOS service.
// All implementations must embed UnimplementedAAOSServer
// for forward compatibility
type AAOSServer interface {
	GetInfo(context.Context, *emptypb.Empty) (*Info, error)
	GetTimetable(context.Context, *emptypb.Empty) (*Timetable, error)
	GetCourses(context.Context, *emptypb.Empty) (*Courses, error)
	GetExams(context.Context, *emptypb.Empty) (*Exams, error)
	GetTranscript(context.Context, *emptypb.Empty) (*Transcript, error)
	mustEmbedUnimplementedAAOSServer()
}

// UnimplementedAAOSServer must be embedded to have forward compatible implementations.
type UnimplementedAAOSServer struct {
}

func (UnimplementedAAOSServer) GetInfo(context.Context, *emptypb.Empty) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedAAOSServer) GetTimetable(context.Context, *emptypb.Empty) (*Timetable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimetable not implemented")
}
func (UnimplementedAAOSServer) GetCourses(context.Context, *emptypb.Empty) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedAAOSServer) GetExams(context.Context, *emptypb.Empty) (*Exams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExams not implemented")
}
func (UnimplementedAAOSServer) GetTranscript(context.Context, *emptypb.Empty) (*Transcript, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranscript not implemented")
}
func (UnimplementedAAOSServer) mustEmbedUnimplementedAAOSServer() {}

// UnsafeAAOSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AAOSServer will
// result in compilation errors.
type UnsafeAAOSServer interface {
	mustEmbedUnimplementedAAOSServer()
}

func RegisterAAOSServer(s grpc.ServiceRegistrar, srv AAOSServer) {
	s.RegisterService(&AAOS_ServiceDesc, srv)
}

func _AAOS_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAOSServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.aaos.v4.AAOS/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAOSServer).GetInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AAOS_GetTimetable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAOSServer).GetTimetable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.aaos.v4.AAOS/GetTimetable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAOSServer).GetTimetable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AAOS_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAOSServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.aaos.v4.AAOS/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAOSServer).GetCourses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AAOS_GetExams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAOSServer).GetExams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.aaos.v4.AAOS/GetExams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAOSServer).GetExams(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AAOS_GetTranscript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAOSServer).GetTranscript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.aaos.v4.AAOS/GetTranscript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAOSServer).GetTranscript(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AAOS_ServiceDesc is the grpc.ServiceDesc for AAOS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AAOS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmux.core.aaos.v4.AAOS",
	HandlerType: (*AAOSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _AAOS_GetInfo_Handler,
		},
		{
			MethodName: "GetTimetable",
			Handler:    _AAOS_GetTimetable_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _AAOS_GetCourses_Handler,
		},
		{
			MethodName: "GetExams",
			Handler:    _AAOS_GetExams_Handler,
		},
		{
			MethodName: "GetTranscript",
			Handler:    _AAOS_GetTranscript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aaos/v4/aaos.proto",
}
