package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	postpb "sunflower-blog-svc/api/blog/v1"

	pb "sunflower-blog-svc/api/admin/v1"
)

type PosterService struct {
	pb.UnimplementedPosterServer

	postGrpcCli postpb.PosterClient
}

func NewPosterService(postCli postpb.PosterClient) *PosterService {
	return &PosterService{
		postGrpcCli: postCli,
	}
}

func (s *PosterService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
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
