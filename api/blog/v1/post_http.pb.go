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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPosterCreatePost = "/blog.v1.Poster/CreatePost"
const OperationPosterDeletePost = "/blog.v1.Poster/DeletePost"
const OperationPosterGetPost = "/blog.v1.Poster/GetPost"
const OperationPosterListCategory = "/blog.v1.Poster/ListCategory"
const OperationPosterListPosts = "/blog.v1.Poster/ListPosts"
const OperationPosterListTags = "/blog.v1.Poster/ListTags"
const OperationPosterUpdatePost = "/blog.v1.Poster/UpdatePost"

type PosterHTTPServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	DeletePost(context.Context, *DeletePostRequest) (*emptypb.Empty, error)
	// GetPost 获取单个博客详情
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	ListCategory(context.Context, *emptypb.Empty) (*ListCategoryResp, error)
	// ListPosts 获取博客列表
	ListPosts(context.Context, *ListPostsRequest) (*ListPostsResponse, error)
	ListTags(context.Context, *emptypb.Empty) (*ListTagsResp, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
}

func RegisterPosterHTTPServer(s *http.Server, srv PosterHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/post", _Poster_CreatePost0_HTTP_Handler(srv))
	r.PUT("/api/v1/post", _Poster_UpdatePost0_HTTP_Handler(srv))
	r.DELETE("/api/v1/post/{post_id}", _Poster_DeletePost0_HTTP_Handler(srv))
	r.GET("/api/v1/posts", _Poster_ListPosts0_HTTP_Handler(srv))
	r.GET("/api/v1/post/{post_id}", _Poster_GetPost0_HTTP_Handler(srv))
	r.GET("/api/v1/tags", _Poster_ListTags0_HTTP_Handler(srv))
	r.GET("/api/v1/categories", _Poster_ListCategory0_HTTP_Handler(srv))
}

func _Poster_CreatePost0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterCreatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePost(ctx, req.(*CreatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _Poster_UpdatePost0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePostRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterUpdatePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePost(ctx, req.(*UpdatePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Post)
		return ctx.Result(200, reply)
	}
}

func _Poster_DeletePost0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePostRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterDeletePost)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePost(ctx, req.(*DeletePostRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
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

func _Poster_ListTags0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterListTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTags(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTagsResp)
		return ctx.Result(200, reply)
	}
}

func _Poster_ListCategory0_HTTP_Handler(srv PosterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPosterListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCategoryResp)
		return ctx.Result(200, reply)
	}
}

type PosterHTTPClient interface {
	CreatePost(ctx context.Context, req *CreatePostRequest, opts ...http.CallOption) (rsp *Post, err error)
	DeletePost(ctx context.Context, req *DeletePostRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPost(ctx context.Context, req *GetPostRequest, opts ...http.CallOption) (rsp *Post, err error)
	ListCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListCategoryResp, err error)
	ListPosts(ctx context.Context, req *ListPostsRequest, opts ...http.CallOption) (rsp *ListPostsResponse, err error)
	ListTags(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListTagsResp, err error)
	UpdatePost(ctx context.Context, req *UpdatePostRequest, opts ...http.CallOption) (rsp *Post, err error)
}

type PosterHTTPClientImpl struct {
	cc *http.Client
}

func NewPosterHTTPClient(client *http.Client) PosterHTTPClient {
	return &PosterHTTPClientImpl{client}
}

func (c *PosterHTTPClientImpl) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/api/v1/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPosterCreatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PosterHTTPClientImpl) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/v1/post/{post_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPosterDeletePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
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

func (c *PosterHTTPClientImpl) ListCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListCategoryResp, error) {
	var out ListCategoryResp
	pattern := "/api/v1/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPosterListCategory))
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

func (c *PosterHTTPClientImpl) ListTags(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListTagsResp, error) {
	var out ListTagsResp
	pattern := "/api/v1/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPosterListTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PosterHTTPClientImpl) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...http.CallOption) (*Post, error) {
	var out Post
	pattern := "/api/v1/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPosterUpdatePost))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
