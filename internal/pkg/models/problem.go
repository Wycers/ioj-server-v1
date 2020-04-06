package models

type Problem struct {
	Model

	Group     int    `gorm:"not null; unique_index:idx1; default: 0"`
	Locale    string `json:"locale"`
	ProblemID string `json:"pid" gorm:"unique_index:idx2"`
	FileSpace string `json:"file_space"`
}
