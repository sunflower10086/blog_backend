package validate

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/errx"
)

type validator interface {
	Validate() error
}

// Validator is a validator middleware.
func Validator(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err = v.Validate(); err != nil {
					return nil, errx.New(codex.CodeInternalErr, "validate 未通过").WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}
