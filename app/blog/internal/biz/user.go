package biz

import (
	"context"
	"strconv"
	"time"

	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/errors"

	"sunflower-blog-svc/app/blog/internal/conf"
	"sunflower-blog-svc/app/blog/internal/pkg/jwtc"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/helper/encrypt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// WrongUserNameOrPassword is user not found.
var WrongUserNameOrPassword = errx.New(codex.CodeWrongUserNameOrPassword, "数据库找不到用户")

type User struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Account     string `json:"account"`
	Password    string `json:"password"`
	Description string `json:"description"`
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

func (uu *UserUseCase) UserInfoById(ctx context.Context, id int64) (*User, error) {
	metadata := map[string]string{
		"uid": strconv.FormatInt(id, 10),
	}

	userInfo, err := uu.userRepo.FindByID(ctx, id)
	if err != nil {
		err = errx.Internal(err, "根据id查找用户信息失败").WithMetadata(metadata)
		return nil, err
	}

	if userInfo == nil {
		return nil, errx.New(codex.CodeUserNotExist, "用户不存在").WithMetadata(metadata)
	}

	return userInfo, nil
}

func (uu *UserUseCase) UserInfoByAccount(ctx context.Context, account string) (*User, error) {
	userInfo, err := uu.userRepo.FindByAccount(ctx, account)
	if err != nil {
		return nil, errx.Internal(err, "根据账号查找用户信息失败")
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
		return nil, errx.Internal(err, "查询数据库失败").WithMetadata(map[string]string{
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
		return nil, errx.Internal(err, "生成token失败").WithMetadata(map[string]string{
			"account_secret": uu.jwtConf.AccessSecret,
		})
	}

	token = &Token{
		AccessToken: jwtToken,
	}

	return token, nil
}

func (uu *UserUseCase) Register(ctx context.Context, account, password string) error {
	hashPwd := encrypt.PasswordHash(password)
	user := &User{
		Account:     account,
		Password:    hashPwd,
		Description: account,
	}

	_, err := uu.userRepo.Save(ctx, user)
	if err != nil {
		return errx.Internal(err, "注册用户失败")
	}

	return nil
}
