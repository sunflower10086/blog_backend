package data

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
		ID:          user.Id,
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
	// TODO implement me
	panic("implement me")
}

func (u *userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	userQuery := u.data.DB.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(i)).First()

	switch err {
	case nil:
		resp := &biz.User{
			Id:       user.ID,
			UserName: user.Username,
			Account:  user.Account,
			Password: user.Password,
		}
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, errors.Wrap(err, "根据用户ID查询用户信息失败")
	}
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
