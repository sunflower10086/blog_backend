package model

import "gorm.io/gorm"

const TableNamePostTag = "post_tag"

type PostTag struct {
	gorm.Model
	PostID int64 `gorm:"column:post_id;type:int8;not null;index" json:"post_id"`
	TagID  int64 `gorm:"column:tag_id;type:int8;not null" json:"tag_id"`
}

func (p *PostTag) TableName() string {
	return TableNamePostTag
}
