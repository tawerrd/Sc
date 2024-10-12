// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/provider/plugin.proto

package providerpb

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
	Oauth2Plugin_Init_FullMethodName        = "/proto.Oauth2Plugin/Init"
	Oauth2Plugin_Provider_FullMethodName    = "/proto.Oauth2Plugin/Provider"
	Oauth2Plugin_NewAuthURL_FullMethodName  = "/proto.Oauth2Plugin/NewAuthURL"
	Oauth2Plugin_GetUserInfo_FullMethodName = "/proto.Oauth2Plugin/GetUserInfo"
)

// Oauth2PluginClient is the client API for Oauth2Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Oauth2PluginClient interface {
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*Enpty, error)
	Provider(ctx context.Context, in *Enpty, opts ...grpc.CallOption) (*ProviderResp, error)
	NewAuthURL(ctx context.Context, in *NewAuthURLReq, opts ...grpc.CallOption) (*NewAuthURLResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
}

type oauth2PluginClient struct {
	cc grpc.ClientConnInterface
}

func NewOauth2PluginClient(cc grpc.ClientConnInterface) Oauth2PluginClient {
	return &oauth2PluginClient{cc}
}

func (c *oauth2PluginClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*Enpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Enpty)
	err := c.cc.Invoke(ctx, Oauth2Plugin_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) Provider(ctx context.Context, in *Enpty, opts ...grpc.CallOption) (*ProviderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProviderResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_Provider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) NewAuthURL(ctx context.Context, in *NewAuthURLReq, opts ...grpc.CallOption) (*NewAuthURLResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewAuthURLResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_NewAuthURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2PluginClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, Oauth2Plugin_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Oauth2PluginServer is the server API for Oauth2Plugin service.
// All implementations must embed UnimplementedOauth2PluginServer
// for forward compatibility.
type Oauth2PluginServer interface {
	Init(context.Context, *InitReq) (*Enpty, error)
	Provider(context.Context, *Enpty) (*ProviderResp, error)
	NewAuthURL(context.Context, *NewAuthURLReq) (*NewAuthURLResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	mustEmbedUnimplementedOauth2PluginServer()
}

// UnimplementedOauth2PluginServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOauth2PluginServer struct{}

func (UnimplementedOauth2PluginServer) Init(context.Context, *InitReq) (*Enpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedOauth2PluginServer) Provider(context.Context, *Enpty) (*ProviderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provider not implemented")
}
func (UnimplementedOauth2PluginServer) NewAuthURL(context.Context, *NewAuthURLReq) (*NewAuthURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAuthURL not implemented")
}
func (UnimplementedOauth2PluginServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedOauth2PluginServer) mustEmbedUnimplementedOauth2PluginServer() {}
func (UnimplementedOauth2PluginServer) testEmbeddedByValue()                      {}

// UnsafeOauth2PluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Oauth2PluginServer will
// result in compilation errors.
type UnsafeOauth2PluginServer interface {
	mustEmbedUnimplementedOauth2PluginServer()
}

func RegisterOauth2PluginServer(s grpc.ServiceRegistrar, srv Oauth2PluginServer) {
	// If the following call pancis, it indicates UnimplementedOauth2PluginServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Oauth2Plugin_ServiceDesc, srv)
}

func _Oauth2Plugin_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_Provider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).Provider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_Provider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).Provider(ctx, req.(*Enpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_NewAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuthURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).NewAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_NewAuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).NewAuthURL(ctx, req.(*NewAuthURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Plugin_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2PluginServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Plugin_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2PluginServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth2Plugin_ServiceDesc is the grpc.ServiceDesc for Oauth2Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth2Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Oauth2Plugin",
	HandlerType: (*Oauth2PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Oauth2Plugin_Init_Handler,
		},
		{
			MethodName: "Provider",
			Handler:    _Oauth2Plugin_Provider_Handler,
		},
		{
			MethodName: "NewAuthURL",
			Handler:    _Oauth2Plugin_NewAuthURL_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Oauth2Plugin_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/provider/plugin.proto",
}
