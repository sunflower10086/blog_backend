package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sunflower-blog-svc/pkg/errx"
)

type Category struct {
	Id   int64
	Name string
}

type CategoryRepo interface {
	CategoryList(ctx context.Context) ([]*Category, error)
}

type CategoryUseCase struct {
	CategoryRepo CategoryRepo
	log          *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	return &CategoryUseCase{CategoryRepo: repo, log: log.NewHelper(logger)}
}

func (uc *CategoryUseCase) ListCategory(ctx context.Context) ([]*Category, error) {
	CategoryList, err := uc.CategoryRepo.CategoryList(ctx)
	if err != nil {
		err = errx.Internal(err, "获取Category列表出错")
		return nil, err
	}

	return CategoryList, nil
}
