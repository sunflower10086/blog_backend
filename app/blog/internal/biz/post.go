package biz

import (
	"context"
	"fmt"

	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Post is a Post view object.
type Post struct {
	Id         int64
	Title      string
	CreatedAt  int64
	UpdatedAt  int64
	Content    string
	Cover      string
	Tags       []int32
	CategoryId int64
}

func (p *Post) String() string {
	return fmt.Sprintf("id: %d, title: %s, content: %s", p.Id, p.Title, p.Content)
}

// PosterRepo is a Greater userRepo.
type PosterRepo interface {
	Save(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, int64) (*Post, error)
	List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*Post, int64, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id int64) error
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
		return nil, 0, errx.Internal(err, "查询文章列表失败")
	}

	return posts, count, nil
}

func (uc *PosterUseCase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	post, err := uc.repo.Create(ctx, post)
	if err != nil {
		err = errx.Internal(err, "创建帖子失败").WithMetadata(map[string]string{
			"post": post.String(),
		})
		return nil, err
	}

	return post, nil
}

func (uc *PosterUseCase) SavePost(ctx context.Context, post *Post) error {
	post, err := uc.repo.Save(ctx, post)
	if err != nil {
		err = errx.Internal(err, "保存帖子失败").WithMetadata(map[string]string{
			"post": post.String(),
		})
		return err
	}
	return nil
}

func (uc *PosterUseCase) DelPost(ctx context.Context, id int64) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		err = errx.Internal(err, "删除帖子失败").WithMetadata(map[string]string{
			"post_id": fmt.Sprintf("%d", id),
		})
		return err
	}
	return nil
}
