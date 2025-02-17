package httpencoder

import (
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/HiBugEnterprise/gotools/httpc"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	stdhttp "net/http"
)

func SuccessEncoder(w http.ResponseWriter, r *http.Request, resp interface{}) error {
	var body httpc.Response
	body.Code = stdhttp.StatusOK
	body.Msg = errorx.CodeSuccess.Msg()
	body.Data = resp
	httpx.OkJsonCtx(r.Context(), w, body)
	return nil
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	// 针对 401 错误特殊处理
	if se.Code == stdhttp.StatusUnauthorized {
		httpc.JwtUnauthorizedResult(w, r, err)
		return
	}
	httpc.RespError(w, r, err)
}
