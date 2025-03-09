package service

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
	return &pb.Post{
		BaseInfo: &pb.PostBaseInfo{
			Title:      "test",
			Id:         1,
			CreatedAt:  int32(time.Now().Unix()),
			UpdatedAt:  int32(time.Now().Unix()),
			Tags:       []int32{1, 2},
			CategoryId: 1,
			Cover:      "https://hibug.bj.bcebos.com/study-analysis/105217275132383232/105217275132383232_175173842891706368_1_0_1735895898.png",
		},
		Content: "test_content",
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

	err := s.postUc.SavePost(ctx, bizPost)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PosterService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	err := s.postUc.DelPost(ctx, int64(req.PostId))
	return nil, err
}

func (s *PosterService) ListTags(ctx context.Context, req *emptypb.Empty) (*pb.ListTagsResp, error) {
	tagList, err := s.tagUc.ListTag(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListTagsResp{}
	for _, tag := range tagList {
		resp.Tags = append(resp.Tags, &pb.ListTagsResp_Tag{
			Id:   int32(tag.Id),
			Name: tag.Name,
		})
	}
	return resp, nil
}
func (s *PosterService) ListCategory(ctx context.Context, req *emptypb.Empty) (*pb.ListCategoryResp, error) {
	categoryList, err := s.categoryUc.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListCategoryResp{}
	for _, category := range categoryList {
		resp.Categories = append(resp.Categories, &pb.ListCategoryResp_Category{
			Id:   int32(category.Id),
			Name: category.Name,
		})
	}
	return resp, nil
}
