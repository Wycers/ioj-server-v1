package models

type Judgement struct {
	Model

	TestCase string

	Score  uint
	Status JudgeStatus `sql:"type:judge_status"`
}
