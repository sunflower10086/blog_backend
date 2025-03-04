package model

import "gorm.io/gorm"

const TableNameTag = "tag"

type Tag struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (t *Tag) TableName() string {
	return TableNameTag
}
