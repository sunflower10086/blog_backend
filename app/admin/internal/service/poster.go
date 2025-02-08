package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"sunflower-blog-svc/app/admin/internal/biz"

	pb "sunflower-blog-svc/api/admin/v1"
)

type PosterService struct {
	pb.UnimplementedPosterServer

	pUc    *biz.PostUseCase
	logger *log.Helper
}

func NewPosterService(pUc *biz.PostUseCase, logger log.Logger) *PosterService {
	return &PosterService{
		pUc:    pUc,
		logger: log.NewHelper(log.With(logger, "service", "Post")),
	}
}

func (s *PosterService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post := &biz.Post{
		Title:      req.Post.BaseInfo.Title,
		Content:    req.Post.Content,
		Cover:      req.Post.BaseInfo.Cover,
		Tags:       req.Post.BaseInfo.Tags,
		Categories: req.Post.BaseInfo.Categories,
	}
	postInfo, err := s.pUc.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	resp := &pb.Post{
		BaseInfo: &pb.PostBaseInfo{
			Title:       postInfo.Title,
			Cover:       postInfo.Cover,
			Tags:        postInfo.Tags,
			Categories:  postInfo.Categories,
			Description: postInfo.Description,
		},
		Content: postInfo.Content,
	}

	return resp, nil
}
func (s *PosterService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
}
func (s *PosterService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PosterService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
}
func (s *PosterService) ListPost(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	return &pb.ListPostsResponse{}, nil
}
