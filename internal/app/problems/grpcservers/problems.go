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

func (s *ProblemService) FetchProblem(ctx context.Context, req *proto.FetchProblemRequest) (res *proto.FetchProblemResponse, err error) {
	if p, err := s.service.FetchProblem(req.GetProblemId()); err != nil {
		return nil, errors.Wrapf(err, "Fetch problem failed")
	} else {
		res = &proto.FetchProblemResponse{
			Problem: &proto.Problem{
				Group:        int32(p.Group),
				Locale:       p.Locale,
				ProblemId:    p.ProblemId,
				PublicSpace:  p.PublicSpace,
				PrivateSpace: p.PrivateSpace,
			},
		}
	}
	return
}

func (s *ProblemService) CreateProblem(ctx context.Context, req *proto.CreateProblemRequest) (res *proto.CreateProblemResponse, err error) {
	if _, err := s.service.CreateProblem(req.Title, req.Locale); err != nil {
		return nil, errors.Wrapf(err, "Create problem failed")
	} else {
		res = &proto.CreateProblemResponse{
			Status: proto.Status_success,
		}
	}
	return
}

func NewUsersServer(logger *zap.Logger, ps services.ProblemsService) (*ProblemService, error) {
	return &ProblemService{
		logger:  logger,
		service: ps,
	}, nil
}
