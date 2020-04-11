package service

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/pkg/errors"
)

type SubmissionService interface {
	Create(submitterId uint64, problemId, userSpace string) error
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

	pd, err := d.submissionSrv.CreateSubmission(context.TODO(), req)
	if err != nil {
		return errors.Wrap(err, "create submission error")
	}

	if pd.GetStatus() == proto.Status_error {
		return errors.New("Failed!")
	}
	return nil
}

func NewSubmissionService(client proto.SubmissionsClient) SubmissionService {
	return &DefaultSubmissionService{
		submissionSrv: client,
	}
}
