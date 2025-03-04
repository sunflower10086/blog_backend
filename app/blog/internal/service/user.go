package service

import (
	"context"

	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/pkg/ctxdata"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "sunflower-blog-svc/api/blog/v1"
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
	// 1.查询用户信息
	token, err := s.userUc.Login(ctx, req.Account, req.Password)
	if err != nil {
		return nil, err
	}

	loginReply := &pb.LoginReply{
		Token: token.AccessToken,
	}

	return loginReply, nil
}

func (s *UserService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserService) UserInfo(ctx context.Context, req *emptypb.Empty) (*pb.UserInfoReply, error) {
	uid, err := ctxdata.GetUid(ctx)
	if err != nil {
		err = errors.New(500, "获取用户信息失败", "获取用户信息失败")
		return nil, err
	}

	userInfo, err := s.userUc.UserInfoById(ctx, uid)
	if err != nil {
		return nil, err
	}

	userInfoReply := &pb.UserInfoReply{
		Username: userInfo.UserName,
		Email:    userInfo.Account,
		Avatar:   "",
	}
	return userInfoReply, nil
}

func (s *UserService) RootUserInfo(ctx context.Context, req *emptypb.Empty) (*pb.UserInfoReply, error) {
	rootUid := int64(1)

	userInfo, err := s.userUc.UserInfoById(ctx, rootUid)
	if err != nil {
		return nil, err
	}

	userInfoReply := &pb.UserInfoReply{
		Username: userInfo.UserName,
		Email:    userInfo.Account,
		Avatar:   "",
	}
	return userInfoReply, nil
}
