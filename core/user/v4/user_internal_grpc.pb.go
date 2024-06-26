// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v4

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

// UserInternalClient is the client API for UserInternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInternalClient interface {
	GetNames(ctx context.Context, in *GetNamesReq, opts ...grpc.CallOption) (*GetNamesResp, error)
}

type userInternalClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInternalClient(cc grpc.ClientConnInterface) UserInternalClient {
	return &userInternalClient{cc}
}

func (c *userInternalClient) GetNames(ctx context.Context, in *GetNamesReq, opts ...grpc.CallOption) (*GetNamesResp, error) {
	out := new(GetNamesResp)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.UserInternal/GetNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInternalServer is the server API for UserInternal service.
// All implementations must embed UnimplementedUserInternalServer
// for forward compatibility
type UserInternalServer interface {
	GetNames(context.Context, *GetNamesReq) (*GetNamesResp, error)
	mustEmbedUnimplementedUserInternalServer()
}

// UnimplementedUserInternalServer must be embedded to have forward compatible implementations.
type UnimplementedUserInternalServer struct {
}

func (UnimplementedUserInternalServer) GetNames(context.Context, *GetNamesReq) (*GetNamesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNames not implemented")
}
func (UnimplementedUserInternalServer) mustEmbedUnimplementedUserInternalServer() {}

// UnsafeUserInternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInternalServer will
// result in compilation errors.
type UnsafeUserInternalServer interface {
	mustEmbedUnimplementedUserInternalServer()
}

func RegisterUserInternalServer(s grpc.ServiceRegistrar, srv UserInternalServer) {
	s.RegisterService(&UserInternal_ServiceDesc, srv)
}

func _UserInternal_GetNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInternalServer).GetNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.UserInternal/GetNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInternalServer).GetNames(ctx, req.(*GetNamesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInternal_ServiceDesc is the grpc.ServiceDesc for UserInternal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInternal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmux.core.user.v4.UserInternal",
	HandlerType: (*UserInternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNames",
			Handler:    _UserInternal_GetNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v4/user_internal.proto",
}
