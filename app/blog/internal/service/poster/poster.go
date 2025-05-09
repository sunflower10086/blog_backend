package poster

import (
	"context"
	"fmt"
	"time"

	"sunflower-blog-svc/app/blog/internal/service/mq"
	"sunflower-blog-svc/pkg/codex"
	"sunflower-blog-svc/pkg/errx"
	pkgmq "sunflower-blog-svc/pkg/mq"

	"google.golang.org/protobuf/types/known/emptypb"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/pkg/constants"

	"github.com/go-kratos/kratos/v2/log"

	pb "sunflower-blog-svc/api/gen/blog/v1"
)

type PosterService struct {
	pb.UnimplementedPosterServer

	senderFactory *mq.SenderFactory

	postUc     *biz.PosterUseCase
	tagUc      *biz.TagUseCase
	categoryUc *biz.CategoryUseCase
	logger     *log.Helper
}

func NewPosterService(uc *biz.PosterUseCase, tagUc *biz.TagUseCase, categoryUc *biz.CategoryUseCase, senderFactory *mq.SenderFactory, logger log.Logger) *PosterService {
	return &PosterService{
		postUc:        uc,
		tagUc:         tagUc,
		categoryUc:    categoryUc,
		senderFactory: senderFactory,
		logger:        log.NewHelper(log.With(logger, "service", "Post")),
	}
}

func (s *PosterService) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, total, err := s.postUc.Posts(ctx, int(req.Page), int(req.PageSize), []string{}, "")
	if err != nil {
		return nil, err
	}

	postBaseInfoList := make([]*pb.Post, 0, len(posts))
	for _, post := range posts {
		postBaseInfoList = append(postBaseInfoList, &pb.Post{
			Title:     post.Title,
			Id:        int32(post.Id),
			CreatedAt: int32(post.CreatedAt),
			UpdatedAt: int32(post.UpdatedAt),
			Tags:      post.Tags,
			Cover:     post.Cover,
			Views:     int32(post.Views),
		})
	}

	return &pb.ListPostsResponse{
		Posts: postBaseInfoList,
		Total: int32(total),
	}, nil
}

func (s *PosterService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.postUc.GetPostInfo(ctx, int64(req.PostId))
	if err != nil {
		return nil, err
	}

	go func() {
		// 异步发送消息
		sender, err := s.senderFactory.GetSender(constants.MQProducerPageViewName)
		if err != nil {
			s.logger.Errorf("failed to get sender: %v", err)
			return
		}
		inErr := sender.
			Send(context.Background(), pkgmq.NewMessage("postID", []byte(fmt.Sprint(req.PostId))))
		if inErr != nil {
			s.logger.Errorf("failed to send message to kafka: %v", inErr)
		}
	}()

	return &pb.Post{
		Title:      post.Title,
		Id:         int32(post.Id),
		CreatedAt:  int32(post.CreatedAt),
		UpdatedAt:  int32(post.UpdatedAt),
		Tags:       post.Tags,
		CategoryId: int32(post.CategoryId),
		Cover:      post.Cover,
		Views:      int32(post.Views),
		Content:    post.Content,
	}, nil
}

func (s *PosterService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	if err := s.validateTagsAndCategoryIdExist(ctx, req.Tags, req.CategoryId); err != nil {
		return nil, err
	}

	bizPost := &biz.Post{
		Title:      req.Title,
		Content:    req.Content,
		Cover:      req.Content,
		Tags:       req.Tags,
		CategoryId: int64(req.CategoryId),
	}
	post, err := s.postUc.CreatePost(ctx, bizPost)
	if err != nil {
		return nil, err
	}

	resp := &pb.Post{
		Title:      post.Title,
		Id:         int32(post.Id),
		CreatedAt:  int32(time.Now().Unix()),
		UpdatedAt:  int32(time.Now().Unix()),
		Tags:       post.Tags,
		CategoryId: int32(post.CategoryId),
		Cover:      post.Cover,
	}
	return resp, nil
}

func (s *PosterService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	if err := s.validateTagsAndCategoryIdExist(ctx, req.Post.Tags, req.Post.CategoryId); err != nil {
		return nil, err
	}

	bizPost := &biz.Post{
		Id:         int64(req.Post.Id),
		Title:      req.Post.Title,
		Content:    req.Post.Content,
		Cover:      req.Post.Cover,
		Tags:       req.Post.Tags,
		CategoryId: int64(req.Post.CategoryId),
	}

	err := s.postUc.UpdatePost(ctx, bizPost)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PosterService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	err := s.postUc.DelPost(ctx, int64(req.PostId))
	return nil, err
}

func (s *PosterService) validateTagsAndCategoryIdExist(ctx context.Context, tags []int32, categoryId int32) error {
	// 校验 tags 是否合理
	if len(tags) != 0 {
		exist, err := s.tagUc.TagsIsExist(ctx, tags)
		if err != nil {
			return err
		}
		if !exist {
			return errx.New(codex.CodeInvalidTags, "tagsId 不存在").WithMetadata(map[string]string{
				"tags": fmt.Sprintf("%v", tags),
			})
		}
	}

	if categoryId != 0 {
		exist, err := s.categoryUc.CategoryIsExist(ctx, categoryId)
		if err != nil {
			return err
		}
		if !exist {
			return errx.New(codex.CodeInvalidCategoryId, "categoryId 不存在").WithMetadata(map[string]string{
				"categoryId": fmt.Sprintf("%d", categoryId),
			})
		}
	}

	return nil
}
