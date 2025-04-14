package poster

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "sunflower-blog-svc/api/gen/blog/v1"
)

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
