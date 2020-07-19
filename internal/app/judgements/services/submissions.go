package services

import (
	"context"

	proto "github.com/infinity-oj/api/protobuf-spec"
)

type SubmissionsService interface {
	ReportJudgement(judgementId string, outputs [][]byte) error
}

type DefaultSubmissionsService struct {
	submissionSrv proto.SubmissionsClient
}

func (d DefaultSubmissionsService) ReportJudgement(judgementId string, outputs [][]byte) error {
	request := &proto.ReturnJudgementRequest{
		JudgementId: judgementId,
		Outputs:     outputs,
	}

	_, err := d.submissionSrv.ReturnJudgement(context.TODO(), request)
	if err != nil {
		return err
	}

	return nil
}

func NewSubmissionsService(client proto.SubmissionsClient) SubmissionsService {
	return &DefaultSubmissionsService{
		submissionSrv: client,
	}
}
