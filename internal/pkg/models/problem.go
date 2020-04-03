package models

type Problem struct {
	Model

	Title       string `json:"title" gorm: "index: title; not null"` // title
	TimeLimit   uint64 `json:"time_limit"`                           // in ms
	MemoryLimit uint64 `json:"memory_limit"`                         // in MB

	Description  string `json:"description"`
	InputFormat  string `json:"input_format"`
	OutputFormat string `json:"output_format"`
	Example      string `json:"example"`
	HintAndLimit string `json:"hint_and_limit"`

	FileIoInputName  string `json:"file_io_input_name"`
	FileIoOutputName string `json:"file_io_output_name"`
}
