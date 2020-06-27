package grpcservers

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"go.uber.org/zap"
)

type JudgementsService struct {
	logger  *zap.Logger
	service services.JudgementsService
}

func (s *JudgementsService) FetchJudgementTask(ctx context.Context, request *proto.FetchJudgementTaskRequest) (*proto.FetchJudgementTaskResponse, error) {
	taskType := request.GetType()

	task, _ := s.service.FetchJudgementTask(taskType)

	if task == nil {
		return nil, nil
	}

	rsp := &proto.FetchJudgementTaskResponse{
		Token:     "",
		Arguments: nil,
		Slots:     nil,
	}

	return rsp, nil

}

func (s *JudgementsService) ReturnJudgementTask(ctx context.Context, request *proto.ReturnJudgementTaskRequest) (*proto.ReturnJudgementTaskResponse, error) {
	panic("implement me")
}

func (s *JudgementsService) mustEmbedUnimplementedJudgementsServer() {
	panic("implement me")
}

func (s *JudgementsService) ListJudgements(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error) {
	s.service.List()
	return &proto.ListResponse{}, nil
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
	judgement := s.service.FetchJudgementByToken(req.GetToken())
	if judgement != nil {
		space := ""
		switch strings.ToLower(req.GetSpace()) {
		case "public":
			space = judgement.PublicSpace
			break
		case "private":
			space = judgement.PrivateSpace
			break
		case "user":
			space = judgement.UserSpace
			break
		default:
			return nil, errors.New("invalid space type")
		}

		fetchRes, err := s.service.FetchFile(space, req.GetFilename())

		if err != nil {
			return nil, err
		}
		res = &proto.FetchJudgeFileResponse{
			Status: proto.Status_success,
			File:   fetchRes,
			Sha1:   "",
		}
		return res, nil
	} else {
		panic("No such judgement")
	}
}

func (s *JudgementsService) FetchHashFile(context.Context, *proto.FetchFileHashRequest) (*proto.FetchFileHashResponse, error) {
	panic("implement me")
}

func (s *JudgementsService) FetchJudgement(ctx context.Context, req *proto.FetchJudgementRequest) (res *proto.FetchJudgementResponse, err error) {
	judgement := s.service.FetchJudgement()

	if judgement == nil {
		res = &proto.FetchJudgementResponse{}
	} else {
		res = &proto.FetchJudgementResponse{
			Token:            judgement.Token,
			TestCase:         judgement.TestCase,
			TimeLimit:        0,
			MemoryLimit:      0,
			FileIoInputName:  "",
			FileIoOutputName: "",
		}
	}

	return
}

func (s *JudgementsService) ReturnJudgement(ctx context.Context, req *proto.ReturnJudgementRequest) (res *proto.ReturnJudgementResponse, err error) {
	err = s.service.FinishJudgement(req.GetToken(), req.GetScore(), req.GetMsg())

	if err != nil {
		res = &proto.ReturnJudgementResponse{
			Status: proto.Status_error,
		}
	} else {
		res = &proto.ReturnJudgementResponse{
			Status: proto.Status_success,
		}
	}
	return
}

func NewJudgementsServer(logger *zap.Logger, js services.JudgementsService) (*JudgementsService, error) {
	return &JudgementsService{
		logger:  logger,
		service: js,
	}, nil
}
