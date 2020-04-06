package grpcservers

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/app/problems/services"
	"go.uber.org/zap"
)

type ProblemService struct {
	logger  *zap.Logger
	service services.ProblemsService
}

func (s *ProblemService) CreateProblem(ctx context.Context, req *proto.CreateProblemRequest) (res *proto.CreateProblemResponse, err error) {
	if _, err := s.service.Create(req.Title, req.Locale); err != nil {
		return nil, errors.Wrapf(err, "CreateUser user failed")
	} else {
		res = &proto.CreateProblemResponse{
			Status: proto.Status_success,
		}
	}
	return
}

func (s *ProblemService) FetchProblem(context.Context, *proto.FetechProblemRequest) (*proto.FetchProblemResponse, error) {
	panic("implement me")
}

func NewUsersServer(logger *zap.Logger, ps services.ProblemsService) (*ProblemService, error) {
	return &ProblemService{
		logger:  logger,
		service: ps,
	}, nil
}
