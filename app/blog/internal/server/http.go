package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	v1 "sunflower-blog-svc/api/blog/v1"
	"sunflower-blog-svc/app/blog/internal/conf"
	"sunflower-blog-svc/app/blog/internal/pkg/middlewares"
	"sunflower-blog-svc/app/blog/internal/service"
	"sunflower-blog-svc/pkg/httpencoder"
	"sunflower-blog-svc/pkg/middlewares/validate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	bc *conf.Bootstrap,
	user *service.UserService,
	poster *service.PosterService,
	logger log.Logger,
) *http.Server {
	c := bc.Server
	confJwt := bc.Jwt
	opts := []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(recovery.Recovery(), tracing.Server()).Prefix("/api").Build(),
			middlewares.Jwt(confJwt.GetAccessSecret()),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-Token", "Accept"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		)),
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

	opts = append(opts, http.ResponseEncoder(httpencoder.SuccessEncoder))
	opts = append(opts, http.ErrorEncoder(httpencoder.ErrorEncoder))
	srv := http.NewServer(opts...)

	v1.RegisterPosterHTTPServer(srv, poster)
	v1.RegisterUserHTTPServer(srv, user)
	return srv
}
