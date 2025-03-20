package converter

import (
	"encoding/json"
	"fmt"
	"sunflower-blog-svc/app/blog/internal/biz"
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"
)

func ConverterBizPost(p *model.Post) (*biz.Post, error) {
	tags := make([]int32, 0)
	err := json.Unmarshal(p.Tags, &tags)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal model.Post.Tags failed: %v", err)
	}
	return &biz.Post{
		Id:         p.ID,
		Title:      p.Title,
		CreatedAt:  p.CreatedAt.Unix(),
		UpdatedAt:  p.UpdatedAt.Unix(),
		Content:    p.Content,
		Tags:       tags,
		Cover:      p.Cover,
		CategoryId: p.CategoryId,
	}, nil
}
