package services

import (
	"github.com/Infinity-OJ/Server/internal/app/problems/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"go.uber.org/zap"
)

var specialKey = "imf1nlTy0j"

type ProblemsService interface {
	Create(title string) (p *models.Problem, err error)
}

type DefaultProblemService struct {
	logger     *zap.Logger
	Repository repositories.ProblemRepository
}

func (s DefaultProblemService) Create(title string) (p *models.Problem, err error) {
	if p, err = s.Repository.Create(title); err != nil {
		return nil, err
	}
	return
}

func NewProblemService(logger *zap.Logger, Repository repositories.ProblemRepository) ProblemsService {
	return &DefaultProblemService{
		logger:     logger.With(zap.String("type", "DefaultProblemService")),
		Repository: Repository,
	}
}
