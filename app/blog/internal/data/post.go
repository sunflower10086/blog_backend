package data

import (
	"context"
	"encoding/json"

	"sunflower-blog-svc/app/blog/internal/data/converter"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
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
	postEnt := &model.Post{
		Title:   post.Title,
		Content: post.Content,
	}
	err := r.data.DB.WithContext(ctx).Create(postEnt).Error
	if err != nil {
		return nil, errors.Wrap(err, "创建 post 失败")
	}

	respPost, err := converter.ConverterBizPost(postEnt)
	if err != nil {
		return nil, errors.Wrap(err, "转换 post 失败")
	}

	return respPost, nil
}

func (r *posterRepo) Update(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	entTags := make([]int64, 0, len(g.Tags))
	for _, tag := range g.Tags {
		entTags = append(entTags, int64(tag))
	}
	entTagsBytes, _ := json.Marshal(entTags)

	postEnt := &model.Post{
		ID:         g.Id,
		Title:      g.Title,
		Content:    g.Content,
		Cover:      g.Cover,
		CategoryId: g.CategoryId,
		Tags:       datatypes.JSON(entTagsBytes),
	}
	err := r.data.DB.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", g.Id).
		Updates(postEnt).Error
	if err != nil {
		return nil, errors.Wrap(err, "save post data field")
	}
	return g, nil
}

func (r *posterRepo) FindByID(ctx context.Context, postId int64) (*biz.Post, error) {
	post := &model.Post{}
	err := r.data.DB.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", postId).
		First(post).Error

	switch {
	case err == nil:
		return converter.ConverterBizPost(post)
	case errors.Is(err, gorm.ErrRecordNotFound):
		r.log.WithContext(ctx).Errorf("post is not found, post_id:{%d}", postId)
		return nil, err
	default:
		return nil, errors.Wrap(err, "查询帖子出错")
	}
}

func (r *posterRepo) List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*biz.Post, int64, error) {
	query1 := r.data.DB.WithContext(ctx).
		Model(&model.Post{}).
		Order("id DESC")

	var total int64
	err := query1.Count(&total).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "查询帖子总数出错")
	}

	postList := make([]*model.Post, 0)
	err = query1.
		Scopes(Paginate(int64(pageNum), int64(pageSize))).
		Find(&postList).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "查询帖子出错")
	}

	resp := make([]*biz.Post, 0, len(postList))
	for _, post := range postList {
		bizPost, err := converter.ConverterBizPost(post)
		if err != nil {
			return nil, 0, errors.Wrap(err, "转换post出错")
		}
		resp = append(resp, bizPost)
	}

	return resp, total, nil
}

func (r *posterRepo) Delete(ctx context.Context, id int64) error {
	err := r.data.DB.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", id).
		Delete(&model.Post{}).Error
	if err != nil {
		return errors.Wrap(err, "删除帖子出错")
	}

	return nil
}

func (r *posterRepo) IncrViews(ctx context.Context, postId int) error {
	err := r.data.DB.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ?", postId).
		Update("views", gorm.Expr("views + ?", 1)).Error
	if err != nil {
		return errors.Wrap(err, "更新帖子浏览量出错")
	}

	return nil
}
