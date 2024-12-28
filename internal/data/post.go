package data

import (
	"context"

	"sunflower-blog-svc/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PosterRepo = (*posterRepo)(nil)

type posterRepo struct {
	data *Data
	log  *log.Helper
}

// NewPosterRepo .
func NewPosterRepo(data *Data, logger log.Logger) biz.PosterRepo {
	return &posterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *posterRepo) Save(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	return g, nil
}

func (r *posterRepo) Update(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	return g, nil
}

func (r *posterRepo) FindByID(context.Context, int64) (*biz.Post, error) {
	return nil, nil
}

func (r *posterRepo) ListByHello(context.Context, string) ([]*biz.Post, error) {
	return nil, nil
}

func (r *posterRepo) ListAll(context.Context) ([]*biz.Post, error) {
	return nil, nil
}

func (r *posterRepo) List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*biz.Post, int64, error) {
	return []*biz.Post{
		{
			Title:   "test",
			Content: "test_content",
		},
	}, 1, nil
}
