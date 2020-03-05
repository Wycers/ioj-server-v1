package models

type Account struct {
	Model

	Group    int    `gorm:"not null; unique_index:idx1; default: 0"`
	Username string `gorm:"not null; unique_index:idx1"`
	Hash     string `gorm:"not null;"`
	Salt     string `gorm:"not null;"`
}
