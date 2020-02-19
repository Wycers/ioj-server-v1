package models

import "time"

type Account struct {
	ID        uint64 `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Username  string
	Hash      string
	Salt      string
}
