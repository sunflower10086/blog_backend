package biz

import (
	"context"

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

// NewUserUseCase new a User useCase.
func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{TagRepo: repo, log: log.NewHelper(logger)}
}
