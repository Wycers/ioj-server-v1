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
	submissionId := request.SubmissionId

	s.logger.Info("dispatch judgement rpc", zap.String("submission id", submissionId))
	err := s.service.DispatchJudgement(submissionId)
	if err != nil {
		s.logger.Error("error:", zap.Error(err))
		return nil, err
	}
	return &proto.DispatchJudgeResponse{}, nil
}

func (s *SubmissionService) ReturnJudgement(ctx context.Context, request *proto.ReturnJudgementRequest) (*proto.ReturnJudgementResponse, error) {
	judgementId := request.JudgementId
	outputs := request.Outputs

	err := s.service.ReturnJudgement(judgementId, outputs)
	if err != nil {
		s.logger.Error("return judgement error", zap.Error(err))
		return nil, err
	}
	return &proto.ReturnJudgementResponse{}, nil
}

func NewSubmissionsServer(logger *zap.Logger, ps services.SubmissionsService) (*SubmissionService, error) {
	return &SubmissionService{
		logger:  logger,
		service: ps,
	}, nil
}
