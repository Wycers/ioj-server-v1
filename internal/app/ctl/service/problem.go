package service

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
)

type ProblemService interface {
	CreateProblem(title, locale string) (*models.Page, error)
}

type DefaultProblemService struct {
	problemSrv proto.ProblemsClient
}

func NewProblemService(problemSrv proto.ProblemsClient) ProblemService {
	return &DefaultProblemService{
		problemSrv: problemSrv,
	}
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
