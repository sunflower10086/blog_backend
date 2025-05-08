package service

import (
	"sunflower-blog-svc/app/blog/internal/service/mq"
	"sunflower-blog-svc/app/blog/internal/service/poster"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	poster.NewPosterService,
	NewUserService,
	mq.NewPageViewConsumer,
	mq.NewLikeActionConsumer,
	mq.NewHandlerRegistry,
	mq.NewSenderFactory,
)
