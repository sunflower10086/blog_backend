package mq

import (
	"context"
	"strconv"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/pkg/errx"
	pkgmq "sunflower-blog-svc/pkg/mq"
)

type PageViewConsumer struct {
	postUc *biz.PosterUseCase
}

func NewPageViewConsumer(postUc *biz.PosterUseCase) *PageViewConsumer {
	h := &PageViewConsumer{
		postUc: postUc,
	}
	return h
}

func (h *PageViewConsumer) Handle(ctx context.Context, event pkgmq.Event) error {
	postIdStr := string(event.Value())
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		err = errx.Internal(err, "消息队列接收 postID 转 int 失败")
		return err
	}
	if err = h.postUc.IncrViews(ctx, postId); err != nil {
		err = errx.Internal(err, "增加 post 访问量失败").WithMetadata(map[string]string{
			"post_id": postIdStr,
		})
		return err
	}
	return nil
}

type LikeActionConsumer struct {
}

func NewLikeActionConsumer() *LikeActionConsumer {
	h := &LikeActionConsumer{}
	return h
}

func (h *LikeActionConsumer) Handle(ctx context.Context, event pkgmq.Event) error {
	return nil
}
