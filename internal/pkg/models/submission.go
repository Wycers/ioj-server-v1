package models

type Submission struct {
	Model

	// submitter ID of this submission
	SubmitterId uint64
	Submitter   User

	ProblemID string
	UserSpace string

	Code string

	TimeUsed   uint
	MemoryUsed uint

	Judgements []Judgement

	Score  uint
	Status JudgeStatus `sql:"type:judge_status"`
}
