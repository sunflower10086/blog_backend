package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"sunflower-blog-svc/internal/biz"

	pb "sunflower-blog-svc/api/blog/v1"
)

type PosterService struct {
	pb.UnimplementedPosterServer

	uc     *biz.PosterUseCase
	logger *log.Helper
}

func NewPosterService(uc *biz.PosterUseCase, logger log.Logger) *PosterService {
	return &PosterService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "service", "Post")),
	}
}

func (s *PosterService) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, total, err := s.uc.Posts(ctx, int(req.Page), int(req.PageSize), []string{}, "")
	if err != nil {
		return nil, nil
	}

	postBaseInfoList := make([]*pb.PostBaseInfo, 0, len(posts))
	for _, post := range posts {
		postBaseInfoList = append(postBaseInfoList, &pb.PostBaseInfo{
			Title: post.Title,
			Id:    strconv.Itoa(1),
		})
	}

	return &pb.ListPostsResponse{
		Posts: postBaseInfoList,
		Total: int32(total),
	}, nil
}

func (s *PosterService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
}
