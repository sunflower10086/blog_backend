package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"sunflower-blog-svc/pkg/errx"
)

type Category struct {
	Id   int64
	Name string
}

type CategoryRepo interface {
	CategoryList(ctx context.Context) ([]*Category, error)

	CategoryIsExist(ctx context.Context, categoryId int32) (bool, error)
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

func (uc *CategoryUseCase) CategoryIsExist(ctx context.Context, categoryId int32) (bool, error) {
	exist, err := uc.CategoryRepo.CategoryIsExist(ctx, categoryId)
	if err != nil {
		err = errx.Internal(err, "count Category 出错").WithMetadata(map[string]string{
			"categoryId": fmt.Sprintf("%v", categoryId),
		})
		return false, err
	}

	return exist, nil
}
