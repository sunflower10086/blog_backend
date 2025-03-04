package httpencoder

import (
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
	"sunflower-blog-svc/pkg/codex"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SuccessEncoder(w http.ResponseWriter, r *http.Request, resp interface{}) error {
	var body Response
	body.Code = stdhttp.StatusOK
	body.Msg = errorx.CodeSuccess.Msg()
	body.Data = resp

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(body)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(stdhttp.StatusOK)
	_, err = w.Write(data)
	return err
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 返回码均是200
	w.WriteHeader(stdhttp.StatusOK)

	se := errors.FromError(err)
	body := Response{Code: int(codex.CodeInternalErr), Msg: codex.CodeInternalErr.Msg()}

	if se.Code != int32(codex.CodeInternalErr) {
		body.Code = int(se.GetCode())
		body.Msg = se.GetMessage()
	}

	data, err := codec.Marshal(body)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	return
}
