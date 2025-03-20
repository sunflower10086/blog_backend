package service

import (
	"github.com/google/wire"
	"sunflower-blog-svc/app/blog/internal/service/poster"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(poster.NewPosterService, NewUserService)
