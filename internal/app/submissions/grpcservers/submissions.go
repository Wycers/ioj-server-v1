package grpcservers

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/submissions/services"
	"go.uber.org/zap"
)

type SubmissionService struct {
	logger  *zap.Logger
	service services.SubmissionsService
}

func (s *SubmissionService) mustEmbedUnimplementedSubmissionsServer() {
	panic("implement me")
}

func (s *SubmissionService) CreateSubmission(ctx context.Context, req *proto.CreateSubmissionRequest) (res *proto.CreateSubmissionResponse, err error) {
	if _, err := s.service.Create(req.GetSubmitterId(), req.GetProblemId(), req.GetUserSpace()); err != nil {
		return nil, errors.Wrapf(err, "Create submission failed")
	} else {
		res = &proto.CreateSubmissionResponse{
			Status: proto.Status_success,
		}
	}
	return
}

func (s *SubmissionService) DispatchJudge(ctx context.Context, request *proto.DispatchJudgeRequest) (*proto.DispatchJudgeResponse, error) {
	panic("implement me")
}

func (s *SubmissionService) ReturnJudgement(ctx context.Context, request *proto.ReturnJudgementRequest) (*proto.ReturnJudgementResponse, error) {
	panic("implement me")
}

func NewSubmissionsServer(logger *zap.Logger, ps services.SubmissionsService) (*SubmissionService, error) {
	return &SubmissionService{
		logger:  logger,
		service: ps,
	}, nil
}
