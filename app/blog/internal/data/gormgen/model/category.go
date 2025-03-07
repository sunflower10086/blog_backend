package model

import "gorm.io/gorm"

const TableNameCategory = "category"

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (t *Category) TableName() string {
	return TableNameCategory
}
