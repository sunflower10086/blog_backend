package errx

import (
	"github.com/go-kratos/kratos/v2/errors"
	"sunflower-blog-svc/pkg/codex"
)

func Internal(err error, reason string) *errors.Error {
	return errors.New(int(codex.CodeInternalErr), reason, codex.CodeInternalErr.Msg()).WithCause(err)
}

func New(code codex.ResCode, reason string) *errors.Error {
	return errors.New(int(code), reason, code.Msg())
}

func BadRequest(err error, reason string) *errors.Error {
	return errors.New(int(codex.CodeInvalidParams), reason, err.Error())
}
