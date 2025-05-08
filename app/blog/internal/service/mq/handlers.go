package mq

import (
	"sunflower-blog-svc/app/blog/internal/pkg/constants"
	pkgmq "sunflower-blog-svc/pkg/mq"
)

type HandlerRegistry struct {
	handlers map[string]pkgmq.Handler
}

func NewHandlerRegistry(
	pageViewConsumer *PageViewConsumer,
	likeActionConsumer *LikeActionConsumer,
) *HandlerRegistry {
	handlers := map[string]pkgmq.Handler{
		constants.MQConsumerPageViewName:   pageViewConsumer.Handle,
		constants.MQConsumerLikeActionName: likeActionConsumer.Handle,
	}
	return &HandlerRegistry{
		handlers: handlers,
	}
}

func (r *HandlerRegistry) Get(name string) (pkgmq.Handler, bool) {
	h, ok := r.handlers[name]
	return h, ok
}
