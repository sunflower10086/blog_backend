package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"time"

	"sunflower-blog-svc/app/blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

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
			Title:     post.Title,
			Id:        strconv.Itoa(1),
			CreatedAt: int32(time.Now().Unix()),
			UpdatedAt: int32(time.Now().Unix()),
			Tags:      []string{"test", "后端"},
			Cover:     "https://hibug.bj.bcebos.com/study-analysis/105217275132383232/105217275132383232_175173842891706368_1_0_1735895898.png?authorization=bce-auth-v1%2FALTAK7asxOyRDVbLWoO1R5U3HM%2F2025-01-03T09%3A18%3A18Z%2F604800%2Fhost%2Fa3965cf84867ad8c3a1c7c483c17ebd97156d54674a5c4b76040ade4a229b539",
		})
	}

	return &pb.ListPostsResponse{
		Posts: postBaseInfoList,
		Total: int32(total),
	}, nil
}

func (s *PosterService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	return &pb.Post{
		BaseInfo: &pb.PostBaseInfo{
			Title:      "test",
			Id:         "1",
			CreatedAt:  int32(time.Now().Unix()),
			UpdatedAt:  int32(time.Now().Unix()),
			Tags:       []string{"test", "后端"},
			Categories: "test",
			Cover:      "https://hibug.bj.bcebos.com/study-analysis/105217275132383232/105217275132383232_175173842891706368_1_0_1735895898.png",
		},
		Content: "test_content",
	}, nil
}

func (s *PosterService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (s *PosterService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (s *PosterService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
