package data

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.CategoryRepo = (*categoryRepo)(nil)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "model", "data/Category")),
	}
}

func (t *categoryRepo) CategoryList(ctx context.Context) ([]*biz.Category, error) {
	categories := make([]*model.Category, 0)
	err := t.data.DB.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, errors.Wrap(err, "ListCategory")
	}

	res := make([]*biz.Category, 0)
	for _, category := range categories {
		res = append(res, &biz.Category{
			Id:   int64(category.ID),
			Name: category.Name,
		})
	}
	return res, nil
}
