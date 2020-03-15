package models

type Profile struct {
	Model

	UserId uint64 `json:"uid"`
	Locale string `json:"locale" gorm:"default: 'en'"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
