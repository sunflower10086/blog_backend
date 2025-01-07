package service

import (
	"github.com/google/wire"
	"sunflower-blog-svc/internal/service/poster"
	"sunflower-blog-svc/internal/service/user"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(poster.NewPosterService, user.NewUserService)
