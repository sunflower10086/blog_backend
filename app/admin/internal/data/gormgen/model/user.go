package model

import "time"

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;type:int8" json:"id"`
	Username    string    `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Account     string    `gorm:"column:account;type:varchar(255);not null" json:"account"`
	Password    string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Description string    `gorm:"column:description;type:text;not null" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamptz(6)" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamptz(6)" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:timestamptz(6);index" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
