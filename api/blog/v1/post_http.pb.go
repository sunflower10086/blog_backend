// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             v4.25.3
// source: blog/v1/post.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPosterGetPost = "/blog.v1.Poster/GetPost"
const OperationPosterListPosts = "/blog.v1.Poster/ListPosts"

type PosterHTTPServer interface {
	// GetPost 获取单个博客详情
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	// ListPosts 获取博客列表
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
}

func RegisterPosterHTTPServer(s *http.Server, srv PosterHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/posts", _Poster_ListPosts0_HTTP_Handler(srv))
	r.GET("/api/v1/post/{post_id}", _Poster_GetPost0_HTTP_Handler(srv))
}

func _Poster_ListPosts0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPostsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterListPosts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPosts(ctx, req.(*ListPostsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPostsResponse)
		return ctx.Result(200, reply)
	}
}

func _Poster_GetPost0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterGetPost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPost(ctx, req.(*GetPostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

type PosterHTTPClient interface {
	GetPost(ctx context.Context, req *GetPostRequest, opts ...http.CallOption) (rsp *Post, err error)
	ListPosts(ctx context.Context, req *ListPostsRequest, opts ...http.CallOption) (rsp *ListPostsResponse, err error)
}

type PosterHTTPClientImpl struct {
	cc *http.Client
}

func NewPosterHTTPClient(client *http.Client) PosterHTTPClient {
	return &PosterHTTPClientImpl{client}
}

func (c *PosterHTTPClientImpl) GetPost(ctx context.Context, in *GetPostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/api/v1/post/{post_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPosterGetPost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PosterHTTPClientImpl) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...http.CallOption) (*ListPostsResponse, error) {
	var out ListPostsResponse
	pattern := "/api/v1/posts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPosterListPosts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
