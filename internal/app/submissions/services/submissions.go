package services

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/infinity-oj/server/internal/app/submissions/repositories"
	"github.com/infinity-oj/server/internal/pkg/models"
)

var specialKey = "imf1nlTy0j"

type SubmissionsService interface {
	Create(submitterID uint64, problemID string, userSpace string) (s *models.Submission, err error)

	DeliverJudgement(element *repositories.Process) error
	DispatchJudgement(submissionId string) error
	ReturnJudgement(judgementId string, outputs [][]byte) error
}

type DefaultSubmissionService struct {
	logger           *zap.Logger
	problemService   ProblemsService
	JudgementService JudgementsService
	fileService      FilesService
	Repository       repositories.SubmissionRepository

	processMap map[string]string
	idMap      map[string]int
}

func (d DefaultSubmissionService) Create(submitterID uint64, problemId string, userSpace string) (s *models.Submission, err error) {
	s, err = d.Repository.Create(submitterID, problemId, userSpace)
	if err != nil {
		return nil, err
	}
	return
}

func (d DefaultSubmissionService) DispatchJudgement(submissionId string) error {
	submission, err := d.Repository.FetchSubmissionBySubmissionId(submissionId)
	d.logger.Info("dispatch judgement")
	if err != nil {
		d.logger.Error("error:", zap.Error(err), zap.String("submissionId", submissionId))
		return err
	}

	if submission == nil {
		d.logger.Error("unknown submission")
		return errors.New("unknown submission")
	}

	submissionElement := d.Repository.CreateProcess(submission)

	err = d.DeliverJudgement(submissionElement)
	if err != nil {
		d.logger.Error("dispatch judgement error", zap.Error(err))
		return err
	}
	return nil
}

func (d DefaultSubmissionService) DeliverJudgement(element *repositories.Process) error {

	upstreams := element.FindUpstreams()

	for _, upstream := range upstreams {
		upstreamType := upstream.Type

		fmt.Println(upstream.Properties)

		d.logger.Info("create judgement",
			zap.String("process Id", element.ProcessId),
			zap.Int("block Id", upstream.Id),
			zap.String("judgement type", upstreamType),
		)

		judgementId, err := d.JudgementService.Create(
			context.TODO(),
			upstreamType,
			upstream.Properties,
			upstream.Inputs,
		)

		if err != nil {
			d.logger.Error("create judgement error", zap.Error(err))
			return err
		}
		d.logger.Info("create judgement success", zap.String("judgement id", judgementId))

		d.idMap[judgementId] = upstream.Id
		d.processMap[judgementId] = element.ProcessId
	}

	return nil

}

func (d DefaultSubmissionService) ReturnJudgement(judgementId string, outputs [][]byte) error {
	d.logger.Info("return judgement",
		zap.String("judgement id", judgementId),
	)

	blockId, ok := d.idMap[judgementId]
	if !ok {
		err := errors.New("unknown judgement id")
		d.logger.Error("unknown judgement id", zap.String("judgement id", judgementId))
		return err
	}
	processId, ok := d.processMap[judgementId]
	if !ok {
		err := errors.New("unknown judgement id")
		d.logger.Error("unknown judgement id", zap.String("judgement id", judgementId))
		return err
	}

	submissionElement := d.Repository.FetchProcess(processId)
	if submissionElement == nil {
		err := errors.New("internal error: unknown submission")
		d.logger.Error("return judgement failed", zap.Error(err))
		return err
	}
	err := submissionElement.SetOutputs(blockId, outputs)
	if err != nil {
		d.logger.Error("return judgement failed", zap.Error(err))
	}

	err = d.DeliverJudgement(submissionElement)
	return err
}

func NewSubmissionService(
	logger *zap.Logger,
	ProblemService ProblemsService,
	Repository repositories.SubmissionRepository,
	FileService FilesService,
	JudgementService JudgementsService,
) SubmissionsService {
	return &DefaultSubmissionService{
		logger:           logger.With(zap.String("type", "DefaultSubmissionService")),
		problemService:   ProblemService,
		fileService:      FileService,
		JudgementService: JudgementService,
		Repository:       Repository,
		processMap:       make(map[string]string),
		idMap:            make(map[string]int),
	}
}
