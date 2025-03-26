package poster

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"sunflower-blog-svc/app/blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	pb "sunflower-blog-svc/api/blog/v1"
)

type PosterService struct {
	pb.UnimplementedPosterServer

	postUc     *biz.PosterUseCase
	tagUc      *biz.TagUseCase
	categoryUc *biz.CategoryUseCase
	logger     *log.Helper
}

func NewPosterService(uc *biz.PosterUseCase, tagUc *biz.TagUseCase, categoryUc *biz.CategoryUseCase, logger log.Logger) *PosterService {
	return &PosterService{
		postUc:     uc,
		tagUc:      tagUc,
		categoryUc: categoryUc,
		logger:     log.NewHelper(log.With(logger, "service", "Post")),
	}
}

func (s *PosterService) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, total, err := s.postUc.Posts(ctx, int(req.Page), int(req.PageSize), []string{}, "")
	if err != nil {
		return nil, err
	}

	postBaseInfoList := make([]*pb.PostBaseInfo, 0, len(posts))
	for _, post := range posts {
		postBaseInfoList = append(postBaseInfoList, &pb.PostBaseInfo{
			Title:     post.Title,
			Id:        int32(post.Id),
			CreatedAt: int32(post.CreatedAt),
			UpdatedAt: int32(post.UpdatedAt),
			Tags:      post.Tags,
			Cover:     post.Cover,
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
		if inErr := s.postUc.IncrViews(ctx, int(req.PostId)); inErr != nil {
			log.Errorf("增加帖子浏览量失败: %v", inErr)
		}
	}()

	return &pb.Post{
		BaseInfo: &pb.PostBaseInfo{
			Title:      post.Title,
			Id:         int32(post.Id),
			CreatedAt:  int32(post.CreatedAt),
			UpdatedAt:  int32(post.UpdatedAt),
			Tags:       post.Tags,
			CategoryId: int32(post.CategoryId),
			Cover:      post.Cover,
		},
		Content: post.Content,
	}, nil
}

func (s *PosterService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
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
		BaseInfo: &pb.PostBaseInfo{
			Title:      post.Title,
			Id:         int32(post.Id),
			CreatedAt:  int32(time.Now().Unix()),
			UpdatedAt:  int32(time.Now().Unix()),
			Tags:       post.Tags,
			CategoryId: int32(post.CategoryId),
			Cover:      post.Cover,
		},
	}
	return resp, nil
}

func (s *PosterService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	bizPost := &biz.Post{
		Id:         int64(req.Post.BaseInfo.Id),
		Title:      req.Post.BaseInfo.Title,
		Content:    req.Post.Content,
		Cover:      req.Post.BaseInfo.Cover,
		Tags:       req.Post.BaseInfo.Tags,
		CategoryId: int64(req.Post.BaseInfo.CategoryId),
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
