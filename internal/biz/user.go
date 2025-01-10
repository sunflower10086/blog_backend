package biz

import (
	"context"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sunflower-blog-svc/internal/conf"
	"sunflower-blog-svc/internal/pkg/jwtc"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/helper/encrypt"
	"time"
)

var (
	// ErrUserNotFound is user not found.
	WrongUserNameOrPassword = errorx.New("biz user", int(codex.CodeWrongUserNameOrPassword), codex.CodeWrongUserNameOrPassword.Msg()).Show()
)

type User struct {
	Id          int64
	UserName    string
	Password    string
	Description string
}

type Token struct {
	AccessToken string
}

// UserRepo is a Greater userRepo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	FindByAccount(ctx context.Context, account string) (*User, error)
}

// UserUseCase is a User useCase.
type UserUseCase struct {
	userRepo UserRepo
	log      *log.Helper

	jwtConf *conf.Jwt
}

// NewUserUseCase new a User useCase.
func NewUserUseCase(repo UserRepo, logger log.Logger, jwtConf *conf.Jwt) *UserUseCase {
	return &UserUseCase{userRepo: repo, log: log.NewHelper(logger), jwtConf: jwtConf}
}

func (uu *UserUseCase) UserInfoByAccount(ctx context.Context, account string) (*User, error) {
	userInfo, err := uu.userRepo.FindByAccount(ctx, account)
	if err != nil {
		return nil, errors.Wrap(err, "根据账号查找用户信息失败")
	}

	return userInfo, nil
}

func (uu *UserUseCase) Login(ctx context.Context, account, password string) (token *Token, err error) {
	// 查找用户信息
	user, err := uu.userRepo.FindByAccount(ctx, account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, WrongUserNameOrPassword
		}
		return nil, errorx.Internal(err, "根据账号查找用户信息失败").WithMetadata(errorx.Metadata{
			"account": account,
		})
	}

	// 判断密码是否正确
	if !encrypt.PasswordVerify(password, user.Password) {
		return nil, WrongUserNameOrPassword
	}

	// 生成 token
	iat := time.Now().Unix()
	exp := iat + uu.jwtConf.AccessExpire
	payload := jwtc.Payload{
		Uid: user.Id,
		Iat: iat,
		Exp: exp,
	}
	jwtToken, err := jwtc.GenJwtToken(uu.jwtConf.AccessSecret, &payload)
	if err != nil {
		return nil, errorx.Internal(err, "生成token失败").WithMetadata(errorx.Metadata{
			"account_secret": uu.jwtConf.AccessSecret,
		})
	}

	token = &Token{
		AccessToken: jwtToken,
	}

	return token, nil
}
