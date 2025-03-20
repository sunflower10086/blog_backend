package poster

import (
	"context"
	"sort"

	pb "sunflower-blog-svc/api/blog/v1"
	"sunflower-blog-svc/app/blog/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PosterService) CreateTags(ctx context.Context, req *pb.CreateTagsReq) (*emptypb.Empty, error) {
	bizTags := make([]*biz.Tag, 0, len(req.Names))
	for _, v := range req.Names {
		bizTags = append(bizTags, &biz.Tag{
			Name: v,
		})
	}

	if err := s.tagUc.CreateTag(ctx, bizTags); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *PosterService) DelTags(ctx context.Context, req *pb.Id) (*emptypb.Empty, error) {
	return nil, s.tagUc.DelTag(ctx, int64(req.Id))
}

func (s *PosterService) ListTags(ctx context.Context, req *emptypb.Empty) (*pb.ListTagsResp, error) {
	tagList, err := s.tagUc.ListTag(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListTagsResp{}
	for _, tag := range tagList {
		resp.Tags = append(resp.Tags, &pb.Tag{
			Id:   int32(tag.Id),
			Name: tag.Name,
		})
	}
	return resp, nil
}

func (s *PosterService) StatTags(ctx context.Context, empty *emptypb.Empty) (*pb.StatTagsResp, error) {
	statData, err := s.tagUc.TagWithCount(ctx)
	if err != nil {
		return nil, err
	}

	keys := make([]int32, 0, len(statData))
	for k := range statData {
		keys = append(keys, int32(k))
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	resp := &pb.StatTagsResp{}
	for _, v := range keys {
		for _, tags := range statData[int(v)] {
			resp.TagStat = append(resp.TagStat, &pb.StatTagsResp_TagStat{
				Id:    int32(tags.Id),
				Name:  tags.Name,
				Count: v,
			})
		}
	}

	return resp, nil
}
