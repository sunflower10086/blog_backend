package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;type:int8" json:"id"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamptz(6)" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamptz(6)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz(6)" json:"deleted_at"`
	Title     string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	AuthorID  int64          `gorm:"column:author_id;type:int8;not null" json:"author_id"`
	Cover     string         `gorm:"column:cover;type:varchar(255);not null" json:"cover"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}