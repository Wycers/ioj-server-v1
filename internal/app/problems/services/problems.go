package services

import (
	"github.com/Infinity-OJ/Server/internal/app/problems/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProblemsService interface {
	CreateProblem(title, locale string) (p *models.Page, err error)
	FetchProblem(problemId string) (p *models.Problem, err error)
}

type DefaultProblemService struct {
	logger      *zap.Logger
	Repository  repositories.ProblemRepository
	FileService FileService
}

func (s DefaultProblemService) FetchProblem(problemId string) (p *models.Problem, err error) {
	p, err = s.Repository.FindProblemById(problemId)
	return
}

func (s DefaultProblemService) CreateProblem(title, locale string) (p *models.Page, err error) {
	publicSpace, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "generate public space failed")
	}
	privateSpace, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "generate private space failed")
	}

	if err := s.FileService.CreateFileSpace(publicSpace.String()); err != nil {
		return nil, err
	}
	if err := s.FileService.CreateFileSpace(privateSpace.String()); err != nil {
		return nil, err
	}

	if p, err = s.Repository.CreateProblem(title, locale, publicSpace.String(), privateSpace.String()); err != nil {
		return nil, err
	}
	return
}

func NewProblemService(logger *zap.Logger, Repository repositories.ProblemRepository, fileService FileService) ProblemsService {
	return &DefaultProblemService{
		logger:      logger.With(zap.String("type", "DefaultProblemService")),
		Repository:  Repository,
		FileService: fileService,
	}
}
