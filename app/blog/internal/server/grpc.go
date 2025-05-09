package server

import (
	v1 "sunflower-blog-svc/api/gen/blog/v1"
	"sunflower-blog-svc/app/blog/internal/conf"
	"sunflower-blog-svc/app/blog/internal/service"
	"sunflower-blog-svc/app/blog/internal/service/poster"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, user *service.UserService, poster *poster.PosterService, logger log.Logger) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterPosterServer(srv, poster)
	v1.RegisterUserServer(srv, user)
	return srv
}
