package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sunflower-blog-svc/internal/biz"
	"sunflower-blog-svc/internal/data/gormgen/model"
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
		Id:       user.ID,
		UserName: user.Username,
		Account:  user.Account,
		Password: user.Password,
	}

	return replyUser, nil
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	userQuery := u.data.DB.User.WithContext(ctx)
	userModel := &model.User{
		Account:     user.Account,
		Description: user.Description,
		Password:    user.Password,
		Username:    user.UserName,
	}

	if err := userQuery.Save(userModel); err != nil {
		return nil, errors.Wrap(err, "创建用户失败")
	}

	return nil, nil
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
