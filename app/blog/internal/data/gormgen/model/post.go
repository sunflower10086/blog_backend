package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"time"

	"sunflower-blog-svc/app/blog/internal/biz"

	"gorm.io/gorm"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;type:int8" json:"id"`
	CreatedAt  *time.Time     `gorm:"column:created_at;type:timestamptz(6)" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:timestamptz(6)" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz(6);index" json:"deleted_at"`
	Title      string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Content    string         `gorm:"column:content;type:text;not null" json:"content"`
	AuthorID   int64          `gorm:"column:author_id;type:int8;not null" json:"author_id"`
	Cover      string         `gorm:"column:cover;type:varchar(255);not null" json:"cover"`
	CategoryId int64          `gorm:"column:category_id;type:int2;not null;default:0" json:"category_id"`
	Tags       datatypes.JSON `gorm:"column:tags;type:jsonb;default:'[]'" json:"tags"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}

func (p *Post) AfterFind(tx *gorm.DB) error {
	// todo: 使用 file-key 去 oss 请求图片
	return nil
}

func (p *Post) ConverterBizPost() (*biz.Post, error) {
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
