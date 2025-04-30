package data

import (
	"gorm.io/gorm"
)

var RecordNotFound = gorm.ErrRecordNotFound

func Paginate(pageNum, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
