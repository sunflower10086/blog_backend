package data

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.TagRepo = (*tagRepo)(nil)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "model", "data/tag")),
	}
}

func (t *tagRepo) TagList(ctx context.Context) ([]*biz.Tag, error) {
	tagQuery := t.data.DB.Tag
	tags, err := tagQuery.WithContext(ctx).Find()
	if err != nil {
		return nil, errors.Wrap(err, "ListTag")
	}
	res := make([]*biz.Tag, 0)
	for _, tag := range tags {
		res = append(res, &biz.Tag{
			Id:   int64(tag.ID),
			Name: tag.Name,
		})
	}
	return res, nil
}
