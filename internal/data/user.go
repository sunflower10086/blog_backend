package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sunflower-blog-svc/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) FindByAccount(ctx context.Context, account string) (*biz.User, error) {
	userQuery := u.data.DB.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.Account.Eq(account)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, err
	case err != nil:
		return nil, err
	}

	replyUser := &biz.User{
		Id: user.ID,
	}

	return replyUser, nil
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
