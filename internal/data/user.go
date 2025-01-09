package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sunflower-blog-svc/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) FindByAccount(ctx context.Context, account string) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
