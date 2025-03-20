package biz

import (
	"context"

	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Tag struct {
	Id   int64
	Name string
}

type TagRepo interface {
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error

	TagList(ctx context.Context) ([]*Tag, error)
	TagWithCount(ctx context.Context) (map[int][]*Tag, error)

	BatchCreateTag(ctx context.Context, tx *gorm.DB, tag []*Tag) error
	DelTag(ctx context.Context, tx *gorm.DB, id int64) error
}

type TagUseCase struct {
	TagRepo TagRepo
	log     *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{TagRepo: repo, log: log.NewHelper(logger)}
}

func (uc *TagUseCase) ListTag(ctx context.Context) ([]*Tag, error) {
	tagList, err := uc.TagRepo.TagList(ctx)
	if err != nil {
		err = errx.Internal(err, "获取tag列表出错")
		return nil, err
	}

	return tagList, nil
}

func (uc *TagUseCase) CreateTag(ctx context.Context, tag []*Tag) error {
	err := uc.TagRepo.BatchCreateTag(ctx, nil, tag)
	if err != nil {
		err = errx.Internal(err, "创建tag出错")
		return err
	}

	return nil
}

func (uc *TagUseCase) DelTag(ctx context.Context, id int64) error {
	err := uc.TagRepo.DelTag(ctx, nil, id)
	if err != nil {
		err = errx.Internal(err, "删除tag出错")
		return err
	}
	return nil
}

// TagWithCount
// 返回值 map key 为 tag 对应的文章数量，value 为 tag 列表
func (uc *TagUseCase) TagWithCount(ctx context.Context) (map[int][]*Tag, error) {
	tagMap, err := uc.TagRepo.TagWithCount(ctx)
	if err != nil {
		err = errx.Internal(err, "获取tag列表出错")
		return nil, err
	}
	return tagMap, nil
}
