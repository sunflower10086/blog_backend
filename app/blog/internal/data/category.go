package data

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/biz"

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
	CategoryQuery := t.data.DB.Category
	Categories, err := CategoryQuery.WithContext(ctx).Find()
	if err != nil {
		return nil, errors.Wrap(err, "ListCategory")
	}

	res := make([]*biz.Category, 0)
	for _, Category := range Categories {
		res = append(res, &biz.Category{
			Id:   int64(Category.ID),
			Name: Category.Name,
		})
	}
	return res, nil
}
