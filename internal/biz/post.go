package biz

import (
	"context"
	"github.com/HiBugEnterprise/gotools/errorx"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Post is a Post view object.
type Post struct {
	Title   string
	Content string
}

// PosterRepo is a Greater repo.
type PosterRepo interface {
	Save(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, int64) (*Post, error)
	ListByHello(context.Context, string) ([]*Post, error)
	ListAll(context.Context) ([]*Post, error)
	List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*Post, int64, error)
}

// PosterUseCase is a Post useCase.
type PosterUseCase struct {
	repo PosterRepo
	log  *log.Helper
}

// NewPosterUseCase new a Post useCase.
func NewPosterUseCase(repo PosterRepo, logger log.Logger) *PosterUseCase {
	return &PosterUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PosterUseCase) Posts(ctx context.Context, pageNum, pageSize int, tags []string, categories string) ([]*Post, int64, error) {
	posts, count, err := uc.repo.List(ctx, pageNum, pageSize, tags, categories)
	if err != nil {
		return nil, 0, errorx.Internal(err, "查询文章列表失败")
	}

	return posts, count, nil
}
