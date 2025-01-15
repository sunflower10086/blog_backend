//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"sunflower-blog-svc/app/admin/internal/biz"
	"sunflower-blog-svc/app/admin/internal/conf"
	"sunflower-blog-svc/app/admin/internal/data"
	"sunflower-blog-svc/app/admin/internal/server"
	"sunflower-blog-svc/app/admin/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Bootstrap, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
