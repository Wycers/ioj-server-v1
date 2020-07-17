package service

import (
	"context"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/pkg/errors"
)

type SubmissionService interface {
	Create(submitterId uint64, problemId, userSpace string) error
	DispatchJudgement(submissionId string) (judgementId string, err error)
}

type DefaultSubmissionService struct {
	submissionSrv proto.SubmissionsClient
}

func (d DefaultSubmissionService) Create(submitterId uint64, problemId, userSpace string) error {
	req := &proto.CreateSubmissionRequest{
		SubmitterId: submitterId,
		ProblemId:   problemId,
		UserSpace:   userSpace,
	}

	rsp, err := d.submissionSrv.CreateSubmission(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "create submission error")
	}

	if rsp.GetStatus() == proto.Status_error {
		return errors.New("Failed!")
	}

	return nil
}

func (d DefaultSubmissionService) DispatchJudgement(submissionId string) (string, error) {

	req := &proto.DispatchJudgeRequest{
		SubmissionId: submissionId,
	}

	rsp, err := d.submissionSrv.DispatchJudge(context.TODO(), req)
	if err != nil {
		return "", errors.Wrap(err, "dispatch submission error")
	}

	return rsp.JudgementId, nil
}

func NewSubmissionService(client proto.SubmissionsClient) SubmissionService {
	return &DefaultSubmissionService{
		submissionSrv: client,
	}
}
