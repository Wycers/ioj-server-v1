package services

import (
	"errors"

	"github.com/infinity-oj/server/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	List()
	Create(string, string, [][]byte) (*models.Judgement, error)
	Update() error

	PullJudgement(judgementType string) (string, *repositories.JudgementElement)
	PushJudgement(token string, outputs [][]byte) error
}

type DefaultJudgementsService struct {
	logger      *zap.Logger
	Repository  repositories.JudgementsRepository
	FileService FilesService
	tokenMap    map[string]*repositories.JudgementElement
}

func (d DefaultJudgementsService) List() {
	panic("implement me")
}

func (d DefaultJudgementsService) Create(tp string, properties string, inputs [][]byte) (*models.Judgement, error) {
	judgement, err := d.Repository.Create(tp, properties, inputs)
	return judgement, err
}

func (d DefaultJudgementsService) Update() error {
	panic("implement me")
}

func (d DefaultJudgementsService) PullJudgement(judgementType string) (token string, judgementElement *repositories.JudgementElement) {
	judgementElement = d.Repository.FetchJudgementInQueueBy(judgementType)
	if judgementElement != nil {
		// TODO: use jwt
		token = uuid.New().String()
		d.tokenMap[token] = judgementElement
	}
	return "", nil
}

func (d DefaultJudgementsService) PushJudgement(token string, outputs [][]byte) error {
	judgementElement, ok := d.tokenMap[token]

	if !ok {
		return errors.New("invalid token")
	}

	err := d.Repository.ReturnJudgementInQueue(judgementElement, outputs)
	if err != nil {
		delete(d.tokenMap, token)
	}

	return err
}

func NewJudgementsService(logger *zap.Logger, Repository repositories.JudgementsRepository, filesService FilesService) JudgementsService {
	return &DefaultJudgementsService{
		logger:      logger.With(zap.String("type", "DefaultJudgementService")),
		Repository:  Repository,
		FileService: filesService,
		tokenMap:    make(map[string]*repositories.JudgementElement),
	}
}
