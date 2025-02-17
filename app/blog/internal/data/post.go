package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"
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
		log:  log.NewHelper(log.With(logger, "module", "data/post")),
	}
}

func (r *posterRepo) Create(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	q := r.data.DB.Post.WithContext(ctx)

	postEnt := &model.Post{
		Title:   post.Title,
		Content: post.Content,
	}
	err := q.Create(postEnt)
	if err != nil {
		return nil, errors.Wrap(err, "创建 post 失败")
	}

	respPost := postEnt.ConverterBizPost()

	return respPost, nil
}

func (r *posterRepo) Save(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	q := r.data.DB.Post.WithContext(ctx)

	postEnt := &model.Post{
		ID:      g.Id,
		Title:   g.Title,
		Content: g.Content,
	}
	if err := q.Save(postEnt); err != nil {
		return nil, errors.Wrap(err, "save post data field")

	}
	return g, nil
}

func (r *posterRepo) Update(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	q := r.data.DB.Post.WithContext(ctx)

	postEnt := &model.Post{
		ID:      g.Id,
		Title:   g.Title,
		Content: g.Content,
	}
	if err := q.Save(postEnt); err != nil {
		return nil, errors.Wrap(err, "save post data field")

	}
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
	q := r.data.DB.Post.WithContext(ctx)
	q = q.Scopes(Paginate(int64(pageNum), int64(pageSize)))

	total, err := q.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, "查询帖子总数出错")
	}

	postList, err := q.Find()
	if err != nil {
		return nil, 0, errors.Wrap(err, "查询帖子出错")
	}

	resp := make([]*biz.Post, 0, len(postList))
	for _, post := range postList {
		resp = append(resp, post.ConverterBizPost())
	}

	return resp, total, nil
}
