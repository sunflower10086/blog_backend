package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sunflower-blog-svc/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "sunflower-blog-svc/api/blog/v1"
)

type Service struct {
	pb.UnimplementedUserServer

	userUc *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(logger log.Logger, userUc *biz.UserUseCase) *Service {
	return &Service{
		userUc: userUc,
		logger: log.NewHelper(log.With(logger, "service", "User")),
	}
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *Service) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *Service) UserInfo(ctx context.Context, req *emptypb.Empty) (*pb.UserInfoReply, error) {
	return &pb.UserInfoReply{}, nil
}
