package models

type Profile struct {
	Model

	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
