package biz

import (
	"context"
	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/log"
)

type Tag struct {
	Id   int64
	Name string
}

type TagRepo interface {
	TagList(ctx context.Context) ([]*Tag, error)
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
