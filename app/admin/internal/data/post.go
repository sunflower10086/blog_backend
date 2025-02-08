package data

import (
	"context"
	"github.com/pkg/errors"
	postpb "sunflower-blog-svc/api/blog/v1"
	"sunflower-blog-svc/app/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PostRepo = (*postRepo)(nil)

type postRepo struct {
	postGrpcClient postpb.PosterClient

	log *log.Helper
}

func NewPostRepo(client postpb.PosterClient, logger log.Logger) biz.PostRepo {
	return &postRepo{
		postGrpcClient: client,
		log:            log.NewHelper(log.With(logger, "module", "post-repo")),
	}
}

func (p *postRepo) Create(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	grpcPost := &postpb.CreatePostRequest{
		Post: &postpb.Post{
			BaseInfo: &postpb.PostBaseInfo{
				Title: post.Title,
				// 这里传入 file-key, 图片的存储在 biz 层存到 oss 服务
				Cover:       post.Cover,
				Tags:        post.Tags,
				Description: post.Description,
				Categories:  post.Categories,
			},
			Content: post.Content,
		},
	}
	createPost, err := p.postGrpcClient.CreatePost(ctx, grpcPost)
	if err != nil {
		return nil, errors.Wrap(err, "请求 post grpc create 错误")
	}

	resp := &biz.Post{
		Title:       createPost.BaseInfo.GetTitle(),
		Content:     createPost.GetContent(),
		CreatedAt:   int64(createPost.GetBaseInfo().GetCreatedAt()),
		UpdatedAt:   int64(createPost.GetBaseInfo().GetUpdatedAt()),
		Tags:        createPost.GetBaseInfo().GetTags(),
		Categories:  createPost.GetBaseInfo().GetCategories(),
		Cover:       createPost.GetBaseInfo().GetCover(),
		Description: createPost.GetBaseInfo().GetDescription(),
	}
	return resp, nil
}

func (p *postRepo) Save(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	grpcPost := &postpb.CreatePostRequest{
		Post: &postpb.Post{
			BaseInfo: &postpb.PostBaseInfo{
				Title: post.Title,
				// 这里传入 file-key, 图片的存储在 biz 层存到 oss 服务
				Cover:       post.Cover,
				Tags:        post.Tags,
				Description: post.Description,
				Categories:  post.Categories,
			},
			Content: post.Content,
		},
	}
	createPost, err := p.postGrpcClient.CreatePost(ctx, grpcPost)
	if err != nil {
		return nil, errors.Wrap(err, "请求 post grpc create 错误")
	}
	resp := &biz.Post{
		Title:       createPost.BaseInfo.GetTitle(),
		Content:     createPost.GetContent(),
		CreatedAt:   int64(createPost.GetBaseInfo().GetCreatedAt()),
		UpdatedAt:   int64(createPost.GetBaseInfo().GetUpdatedAt()),
		Tags:        createPost.GetBaseInfo().GetTags(),
		Categories:  createPost.GetBaseInfo().GetCategories(),
		Cover:       createPost.GetBaseInfo().GetCover(),
		Description: createPost.GetBaseInfo().GetDescription(),
	}

	return resp, nil
}

func (p *postRepo) Update(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postRepo) FindByID(ctx context.Context, i int64) (*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postRepo) FindByAccount(ctx context.Context, account string) (*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}
