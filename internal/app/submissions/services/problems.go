package services

import (
	"context"
	"fmt"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/models"
	"github.com/pkg/errors"
)

type ProblemsService interface {
	Fetch(ProblemID string) (*models.Problem, error)
}

type DefaultProblemService struct {
	problemSrv proto.ProblemsClient
}

func (s *DefaultProblemService) Fetch(ProblemID string) (*models.Problem, error) {
	// get problem
	req := &proto.FetchProblemRequest{
		ProblemId: ProblemID,
	}
	res, err := s.problemSrv.FetchProblem(context.TODO(), req)
	if err != nil {
		return nil, errors.Wrap(err, "fetch problem error")
	}
	problem := &models.Problem{
		Group:        int(res.GetProblem().GetGroup()),
		Locale:       res.GetProblem().GetLocale(),
		ProblemId:    res.GetProblem().GetProblemId(),
		PublicSpace:  res.GetProblem().GetPublicSpace(),
		PrivateSpace: res.GetProblem().GetPrivateSpace(),
	}
	return problem, nil
}

func (s *DefaultProblemService) CreateProblem(title, locale string) (*models.Page, error) {
	// get detail
	req := &proto.CreateProblemRequest{
		Title:  title,
		Locale: locale,
	}

	pd, err := s.problemSrv.CreateProblem(context.TODO(), req)
	if err != nil {
		return nil, errors.Wrap(err, "create problem error")
	}

	fmt.Println(pd.GetStatus().String())
	return nil, nil
}

func NewProblemService(problemSrv proto.ProblemsClient) ProblemsService {
	return &DefaultProblemService{
		problemSrv: problemSrv,
	}
}
