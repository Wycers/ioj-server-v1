package services

import (
	"fmt"
	"log"

	"github.com/Infinity-OJ/Server/internal/app/submissions/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var specialKey = "imf1nlTy0j"

type SubmissionsService interface {
	Create(submitterID uint64, problemID string, userSpace string) (s *models.Submission, err error)
	Judge(submission *models.Submission) error
}

type DefaultSubmissionService struct {
	logger           *zap.Logger
	ProblemService   ProblemsService
	JudgementService JudgementsService
	FileService      FilesService
	Repository       repositories.SubmissionRepository
}

func (d DefaultSubmissionService) Create(submitterID uint64, problemId string, userSpace string) (s *models.Submission, err error) {
	s, err = d.Repository.Create(submitterID, problemId, userSpace)
	if err != nil {
		return nil, err
	}

	d.Judge(s)
	return
}

type Meta struct {
	TestCases []string `yaml:"testcases"`
}

func (d DefaultSubmissionService) Judge(submission *models.Submission) error {
	problem, err := d.ProblemService.Fetch(submission.ProblemID)
	if err != nil {
		return err
	}

	meta, err := d.FileService.FetchMetaFile(problem.PrivateSpace)
	if err != nil {
		return err
	}

	m := Meta{}
	err = yaml.Unmarshal(meta, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for k, v := range m.TestCases {
		fmt.Println(k, v)
		if err := d.JudgementService.Create(problem.PublicSpace, problem.PrivateSpace, submission.UserSpace, v); err != nil {
			return err
		}
	}
	return nil
}

func NewSubmissionService(logger *zap.Logger, ProblemService ProblemsService, Repository repositories.SubmissionRepository) SubmissionsService {
	return &DefaultSubmissionService{
		logger:         logger.With(zap.String("type", "DefaultSubmissionService")),
		ProblemService: ProblemService,
		Repository:     Repository,
	}
}
