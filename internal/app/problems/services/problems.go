package services

import (
	"fmt"

	"github.com/Infinity-OJ/Server/internal/app/problems/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"
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
	panic("implement me")
}

func (s DefaultProblemService) CreateProblem(title, locale string) (p *models.Page, err error) {
	publicSpace := random.RandStringRunes(32)
	privateSpace := random.RandStringRunes(32)
	fmt.Println("QwQ")
	fmt.Println(&s.FileService)
	if err := s.FileService.CreateFileSpace(publicSpace); err != nil {
		return nil, err
	}
	if err := s.FileService.CreateFileSpace(privateSpace); err != nil {
		return nil, err
	}

	if p, err = s.Repository.CreateProblem(title, locale, publicSpace, privateSpace); err != nil {
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
