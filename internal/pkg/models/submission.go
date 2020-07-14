package models

type Submission struct {
	Model

	SubmissionId string
	// submitter ID of this submission
	SubmitterId uint64
	Submitter   User

	ProblemID string
	UserSpace string

	Code string

	TimeUsed   uint
	MemoryUsed uint

	Score  uint
	Status JudgeStatus `sql:"type:judge_status"`
}
