// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: api/admin/v1/post.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Poster_CreatePost_FullMethodName = "/api.admin.v1.Poster/CreatePost"
	Poster_UpdatePost_FullMethodName = "/api.admin.v1.Poster/UpdatePost"
	Poster_DeletePost_FullMethodName = "/api.admin.v1.Poster/DeletePost"
	Poster_GetPost_FullMethodName    = "/api.admin.v1.Poster/GetPost"
	Poster_ListPost_FullMethodName   = "/api.admin.v1.Poster/ListPost"
)

// PosterClient is the client API for Poster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosterClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	ListPost(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error)
}

type posterClient struct {
	cc grpc.ClientConnInterface
}

func NewPosterClient(cc grpc.ClientConnInterface) PosterClient {
	return &posterClient{cc}
}

func (c *posterClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, Poster_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, Poster_UpdatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Poster_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, Poster_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterClient) ListPost(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*ListPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPostsResponse)
	err := c.cc.Invoke(ctx, Poster_ListPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosterServer is the server API for Poster service.
// All implementations must embed UnimplementedPosterServer
// for forward compatibility.
type PosterServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
	DeletePost(context.Context, *DeletePostRequest) (*emptypb.Empty, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	ListPost(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	mustEmbedUnimplementedPosterServer()
}

// UnimplementedPosterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPosterServer struct{}

func (UnimplementedPosterServer) CreatePost(context.Context, *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPosterServer) UpdatePost(context.Context, *UpdatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPosterServer) DeletePost(context.Context, *DeletePostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPosterServer) GetPost(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPosterServer) ListPost(context.Context, *ListPostsRequest) (*ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedPosterServer) mustEmbedUnimplementedPosterServer() {}
func (UnimplementedPosterServer) testEmbeddedByValue()                {}

// UnsafePosterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosterServer will
// result in compilation errors.
type UnsafePosterServer interface {
	mustEmbedUnimplementedPosterServer()
}

func RegisterPosterServer(s grpc.ServiceRegistrar, srv PosterServer) {
	// If the following call pancis, it indicates UnimplementedPosterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Poster_ServiceDesc, srv)
}

func _Poster_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poster_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poster_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poster_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poster_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poster_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poster_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poster_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poster_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosterServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poster_ListPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosterServer).ListPost(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Poster_ServiceDesc is the grpc.ServiceDesc for Poster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Poster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.Poster",
	HandlerType: (*PosterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Poster_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Poster_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Poster_DeletePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Poster_GetPost_Handler,
		},
		{
			MethodName: "ListPost",
			Handler:    _Poster_ListPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/v1/post.proto",
}
