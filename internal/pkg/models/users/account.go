package users

import "time"

type Account struct {
	ID       uint64 `gorm:"PRIMARY_KEY"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Username string
	Hash     string
	Salt     string
}
