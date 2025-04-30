package biz

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Post is a Post view object.
type Post struct {
	Id         int64   `json:"id"`
	Title      string  `json:"title"`
	CreatedAt  int64   `json:"created_at"`
	UpdatedAt  int64   `json:"updated_at"`
	Content    string  `json:"content"`
	Cover      string  `json:"cover"`
	Tags       []int32 `json:"tags"`
	CategoryId int64   `json:"category_id"`
	Views      int64   `json:"views"`
}

func (p *Post) String() string {
	return fmt.Sprintf("id: %d, title: %s, content: %s", p.Id, p.Title, p.Content)
}

// PosterRepo is a Greater userRepo.
type PosterRepo interface {
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, int64) (*Post, error)
	List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*Post, int64, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id int64) error

	IncrViews(ctx context.Context, id int) error
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

func (uc *PosterUseCase) GetPostInfo(ctx context.Context, id int64) (*Post, error) {
	post, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.New(codex.CodePostNotExist, codex.CodePostNotExist.Msg()).WithMetadata(map[string]string{
				"post_id": fmt.Sprintf("%d", id),
			})
		}
		return nil, errx.Internal(err, "查询帖子失败").WithMetadata(map[string]string{
			"post_id": fmt.Sprintf("%d", id),
		})
	}

	return post, nil
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

func (uc *PosterUseCase) UpdatePost(ctx context.Context, post *Post) error {
	post, err := uc.repo.Update(ctx, post)
	if err != nil {
		err = errx.Internal(err, "保存帖子失败")
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

func (uc *PosterUseCase) IncrViews(ctx context.Context, postId int) error {
	err := uc.repo.IncrViews(ctx, postId)
	if err != nil {
		err = errx.Internal(err, "增加帖子浏览量失败").WithMetadata(map[string]string{
			"post_id": fmt.Sprintf("%d", postId),
		})
		return err
	}

	return nil
}
