package models

type Judgement struct {
	Model

	SubmissionID uint64

	TestCase string
	Score    uint
	Status   JudgeStatus `sql:"type:judge_status"`
}
