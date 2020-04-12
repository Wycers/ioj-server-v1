package models

type Problem struct {
	Model

	Group        int    `gorm:"not null; unique_index:idx2; default: 0"`
	Locale       string `json:"locale"`
	ProblemId    string `json:"pid" gorm:"unique_index:idx2"`
	PublicSpace  string `json:"pub_space"`
	PrivateSpace string `json:"pri_space"`
}
