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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// Login with basic authorization.
	// Will try to register during first login.
	// Returns error if information not adequate to create user.
	Login(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoginResp, error)
	// Register and login with provided information.
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*LoginResp, error)
	// GetProfile of specified user.
	GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*Profile, error)
	// RefreshDevice push key and last seen.
	RefreshDevice(ctx context.Context, in *RefreshDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetDevices of authenticated user.
	GetDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDevicesResp, error)
	// RevokeDevices will revoke firebase refresh token and delete all devices associated with user.
	RevokeDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RefreshDevice(ctx context.Context, in *RefreshDeviceReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/RefreshDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDevicesResp, error) {
	out := new(GetDevicesResp)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/GetDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RevokeDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/xmux.core.user.v4.User/RevokeDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// Login with basic authorization.
	// Will try to register during first login.
	// Returns error if information not adequate to create user.
	Login(context.Context, *emptypb.Empty) (*LoginResp, error)
	// Register and login with provided information.
	Register(context.Context, *RegisterReq) (*LoginResp, error)
	// GetProfile of specified user.
	GetProfile(context.Context, *GetProfileReq) (*Profile, error)
	// RefreshDevice push key and last seen.
	RefreshDevice(context.Context, *RefreshDeviceReq) (*emptypb.Empty, error)
	// GetDevices of authenticated user.
	GetDevices(context.Context, *emptypb.Empty) (*GetDevicesResp, error)
	// RevokeDevices will revoke firebase refresh token and delete all devices associated with user.
	RevokeDevices(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *emptypb.Empty) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetProfile(context.Context, *GetProfileReq) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServer) RefreshDevice(context.Context, *RefreshDeviceReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshDevice not implemented")
}
func (UnimplementedUserServer) GetDevices(context.Context, *emptypb.Empty) (*GetDevicesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedUserServer) RevokeDevices(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeDevices not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetProfile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RefreshDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RefreshDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/RefreshDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RefreshDevice(ctx, req.(*RefreshDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetDevices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RevokeDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RevokeDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.core.user.v4.User/RevokeDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RevokeDevices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmux.core.user.v4.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _User_GetProfile_Handler,
		},
		{
			MethodName: "RefreshDevice",
			Handler:    _User_RefreshDevice_Handler,
		},
		{
			MethodName: "GetDevices",
			Handler:    _User_GetDevices_Handler,
		},
		{
			MethodName: "RevokeDevices",
			Handler:    _User_RevokeDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v4/user.proto",
}