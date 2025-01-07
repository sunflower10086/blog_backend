package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	UserName    string
	Description string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
	List(ctx context.Context, pageNum int, pageSize int, tags []string, categories string) ([]*User, int64, error)
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
