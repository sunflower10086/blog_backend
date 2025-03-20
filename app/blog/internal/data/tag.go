package data

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"

	"sunflower-blog-svc/app/blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ biz.TagRepo = (*tagRepo)(nil)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func (t *tagRepo) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	err := t.data.DB.Transaction(fn)
	if err != nil {
		return errors.Wrap(err, "Transaction 出错")
	}

	return nil
}

func (t *tagRepo) TagWithCount(ctx context.Context) (map[int][]*biz.Tag, error) {
	// 联表查询
	// SELECT tag.id, tag.name, COUNT(DISTINCT post.id) AS post_count
	// FROM tag
	// INNER JOIN post ON post.tags @> jsonb_build_array(tag.id)
	// GROUP BY tag.id, tag.name
	tagMap := make(map[int][]*biz.Tag)
	rows, err := t.data.DB.WithContext(ctx).Model(&model.Tag{}).
		Joins("INNER JOIN post ON post.tags @> jsonb_build_array(tag.id)").
		Select("tag.id, tag.name, COUNT(DISTINCT post.id) AS post_count").
		Group("tag.id, tag.name").Rows()
	if err != nil {
		return nil, errors.Wrap(err, "TagWithCount")
	}
	defer rows.Close()

	for rows.Next() {
		var tag model.Tag
		var postCount int
		err = rows.Scan(&tag.ID, &tag.Name, &postCount)
		if err != nil {
			return nil, errors.Wrap(err, "TagWithCount")
		}
		tagMap[postCount] = append(tagMap[postCount], &biz.Tag{
			Id:   int64(tag.ID),
			Name: tag.Name,
		})
	}

	return tagMap, nil
}

func (t *tagRepo) BatchCreateTag(ctx context.Context, tx *gorm.DB, tags []*biz.Tag) error {
	if tx == nil {
		tx = t.data.DB
	}

	modelData := make([]*model.Tag, 0, len(tags))
	for _, v := range tags {
		modelData = append(modelData, &model.Tag{
			Name: v.Name,
		})
	}

	err := tx.WithContext(ctx).CreateInBatches(modelData, 100).Error
	if err != nil {
		return errors.Wrap(err, "BatchCreateTag")
	}

	return nil
}

func (t *tagRepo) DelTag(ctx context.Context, tx *gorm.DB, id int64) error {
	if tx == nil {
		tx = t.data.DB
	}

	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Tag{}).Error
	if err != nil {
		return errors.Wrap(err, "DelTag")
	}

	return nil
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "model", "data/tag")),
	}
}

func (t *tagRepo) TagList(ctx context.Context) ([]*biz.Tag, error) {
	tags := make([]*model.Tag, 0)
	err := t.data.DB.WithContext(ctx).Find(&tags).Error
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
