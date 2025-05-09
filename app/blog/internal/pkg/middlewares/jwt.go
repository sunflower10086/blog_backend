package middlewares

import (
	"context"

	v1 "sunflower-blog-svc/api/gen/blog/v1"
	"sunflower-blog-svc/app/blog/internal/pkg/ctxdata"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/errx"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

var noNeedLogin = map[string]struct{}{
	v1.OperationUserLogin:        {},
	v1.OperationPosterListPosts:  {},
	v1.OperationPosterGetPost:    {},
	v1.OperationUserRootUserInfo: {},
}

var NeedLogin = map[string]struct{}{
	v1.OperationPosterCreatePost: {},
	v1.OperationPosterUpdatePost: {},
	v1.OperationUserUserInfo:     {},
}

var ErrUnauthorized = errx.New(codex.CodeNeedLogin, "授权已过期或授权异常,请重新授权")

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		//if _, ok := noNeedLogin[operation]; ok {
		//	return false
		//}
		//return true
		if _, ok := NeedLogin[operation]; ok {
			return true
		}
		return false
	}
}

// 设置用户信息
func setUserInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, ErrUnauthorized
			}
			claimInfo := claim.(jwt2.MapClaims)
			ctx = context.WithValue(ctx, ctxdata.CtxKeyUid, claimInfo[ctxdata.CtxKeyUid])
			return handler(ctx, req)
		}
	}
}

func Jwt(accessKey string) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwt2.Token) (interface{}, error) { return []byte(accessKey), nil },
			jwt.WithSigningMethod(jwt2.SigningMethodHS256),
			jwt.WithClaims(func() jwt2.Claims {
				return jwt2.MapClaims{}
			}),
		),

		setUserInfo(),
	).
		Match(NewWhiteListMatcher()).
		Build()
}
