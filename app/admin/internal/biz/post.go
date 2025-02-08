package biz

import (
	"context"
	"github.com/HiBugEnterprise/gotools/errorx"

	"sunflower-blog-svc/app/admin/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type Post struct {
	Title       string
	Content     string
	CreatedAt   int64
	UpdatedAt   int64
	Tags        []string
	Categories  string
	Cover       string
	Description string
}

// PostRepo is a Greater userRepo.
type PostRepo interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	Save(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, int64) (*Post, error)
	FindByAccount(ctx context.Context, account string) (*Post, error)
}

// PostUseCase is a Post useCase.
type PostUseCase struct {
	postRepo PostRepo
	log      *log.Helper

	jwtConf *conf.Jwt
}

// NewPostUseCase new a Post useCase.
func NewPostUseCase(repo PostRepo, logger log.Logger, jwtConf *conf.Jwt) *PostUseCase {
	return &PostUseCase{postRepo: repo, log: log.NewHelper(logger), jwtConf: jwtConf}
}

func (p *PostUseCase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	postInfo, err := p.postRepo.Create(ctx, post)
	if err != nil {
		err = errorx.Internal(err, "创建 Post 失败").WithMetadata(errorx.Metadata{
			"post": post,
		})
		return nil, err
	}
	return postInfo, nil
}
