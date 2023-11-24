// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: api/user/v1alpha1/user.proto

package v1alpha1

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

const (
	UserService_SignIn_FullMethodName      = "/user.v1alpha1.UserService/SignIn"
	UserService_SignOut_FullMethodName     = "/user.v1alpha1.UserService/SignOut"
	UserService_GetUserInfo_FullMethodName = "/user.v1alpha1.UserService/GetUserInfo"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	SignIn(ctx context.Context, in *SignIn, opts ...grpc.CallOption) (*User, error)
	SignOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
	GetUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignIn(ctx context.Context, in *SignIn, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, UserService_SignOut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	SignIn(context.Context, *SignIn) (*User, error)
	SignOut(context.Context, *emptypb.Empty) (*Msg, error)
	GetUserInfo(context.Context, *emptypb.Empty) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) SignIn(context.Context, *SignIn) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedUserServiceServer) SignOut(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *emptypb.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignIn(ctx, req.(*SignIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SignOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignOut(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1alpha1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _UserService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _UserService_SignOut_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1alpha1/user.proto",
}