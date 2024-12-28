package server

import (
	v1 "sunflower-blog-svc/api/blog/v1"
	"sunflower-blog-svc/internal/conf"
	"sunflower-blog-svc/internal/service"
	"sunflower-blog-svc/pkg/middlewares/validate"

	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/HiBugEnterprise/gotools/httpc"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	stdhttp "net/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, poster *service.PosterService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(recovery.Recovery(), tracing.Server()).Prefix("/api").Build(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts = append(opts, http.ResponseEncoder(successEncoder))
	opts = append(opts, http.ErrorEncoder(errorEncoder))
	srv := http.NewServer(opts...)

	v1.RegisterPosterHTTPServer(srv, poster)
	return srv
}

func successEncoder(w http.ResponseWriter, r *http.Request, resp interface{}) error {
	var body httpc.Response
	body.Code = stdhttp.StatusOK
	body.Msg = errorx.CodeSuccess.Msg()
	body.Data = resp
	httpx.OkJsonCtx(r.Context(), w, body)
	return nil
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	httpc.RespError(w, r, err)
}
