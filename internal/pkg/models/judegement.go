package models

type Judgement struct {
	Model

	JudgementId string
	Type        string
	Status      JudgeStatus `sql:"type:judge_status"`

	Property string

	Inputs  string
	Outputs string
}
