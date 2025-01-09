package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	UserName    string
	Password    string
	Description string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	FindByAccount(ctx context.Context, account string) (*User, error)
}

// UserUseCase is a User useCase.
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase new a User useCase.
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) UserInfoByAccount(ctx context.Context, account string) (*User, error) {
	userInfo, err := uc.repo.FindByAccount(ctx, account)
	if err != nil {
		return nil, errors.Wrap(err, "根据账号查找用户信息失败")
	}

	return userInfo, nil
}
