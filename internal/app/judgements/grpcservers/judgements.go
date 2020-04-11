package grpcservers

import (
	"context"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/judgements/services"
	"go.uber.org/zap"
)

type JudgementsService struct {
	logger      *zap.Logger
	service     services.JudgementsService
	fileService proto.FilesClient
}

func (s *JudgementsService) SubmitJudgement(context.Context, *proto.SubmitJudgementRequest) (*proto.SubmitJudgementResponse, error) {
	panic("implement me")
}

func (s *JudgementsService) FetchFile(ctx context.Context, req *proto.FetchJudgeFileRequest) (res *proto.FetchJudgeFileResponse, err error) {
	fetchReq := proto.FetchFileRequest{
		SpaceName: req.FileSpace,
		FilePath:  req.Filename,
	}
	fetchRes, err := s.fileService.FetchFile(ctx, &fetchReq)
	if err != nil {
		return
	}
	res = &proto.FetchJudgeFileResponse{
		Status: proto.Status_success,
		File:   fetchRes.File.Data,
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

func NewJudgementsServer(logger *zap.Logger, ps services.JudgementsService, fs proto.FilesClient) (*JudgementsService, error) {
	return &JudgementsService{
		logger:      logger,
		service:     ps,
		fileService: fs,
	}, nil
}
