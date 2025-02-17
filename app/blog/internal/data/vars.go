package data

import (
	"gorm.io/gen"
)

func Paginate(pageNum, pageSize int64) func(db gen.Dao) gen.Dao {
	return func(db gen.Dao) gen.Dao {
		if pageNum <= 0 {
			pageNum = 1
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
