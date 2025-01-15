package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"sunflower-blog-svc/app/admin/internal/biz"

	pb "sunflower-blog-svc/api/admin/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	userUc *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(logger log.Logger, userUc *biz.UserUseCase) *UserService {
	return &UserService{
		userUc: userUc,
		logger: log.NewHelper(log.With(logger, "service", "User")),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) UserInfo(ctx context.Context, req *emptypb.Empty) (*pb.UserInfoReply, error) {
	return &pb.UserInfoReply{}, nil
}
