package grpcservers

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/judgements/services"
	"go.uber.org/zap"
)

type JudgementsService struct {
	logger  *zap.Logger
	service services.JudgementsService
}

func (s *JudgementsService) SubmitJudgement(ctx context.Context, req *proto.SubmitJudgementRequest) (res *proto.SubmitJudgementResponse, err error) {
	err = s.service.CreateJudgement(req.GetSubmissionId(), req.GetPublicSpace(), req.GetPrivateSpace(), req.GetUserSpace(), req.GetTestCase())
	if err != nil {
		res = &proto.SubmitJudgementResponse{
			Status: proto.Status_error,
			Score:  0,
		}
		return nil, errors.Wrap(err, "create judgement failed")
	}
	res = &proto.SubmitJudgementResponse{
		Status: proto.Status_success,
		Score:  0,
	}
	return
}

func (s *JudgementsService) FetchFile(ctx context.Context, req *proto.FetchJudgeFileRequest) (res *proto.FetchJudgeFileResponse, err error) {
	fetchRes, err := s.service.FetchFile(req.GetFileSpace(), req.GetFilename())
	if err != nil {
		return
	}
	res = &proto.FetchJudgeFileResponse{
		Status: proto.Status_success,
		File:   fetchRes,
		Sha1:   "",
	}
	return
}

func (s *JudgementsService) FetchHashFile(context.Context, *proto.FetchFileHashRequest) (*proto.FetchFileHashResponse, error) {
	panic("implement me")
}

func (s *JudgementsService) FetchJudgement(context.Context, *proto.FetchJudgementRequest) (*proto.ReturnJudgementRequest, error) {
	panic("implement me")
}

func (s *JudgementsService) ReturnJudgement(context.Context, *proto.ReturnJudgementRequest) (*proto.ReturnJudgementResponse, error) {
	panic("implement me")
}

func NewJudgementsServer(logger *zap.Logger, js services.JudgementsService) (*JudgementsService, error) {
	return &JudgementsService{
		logger:  logger,
		service: js,
	}, nil
}
