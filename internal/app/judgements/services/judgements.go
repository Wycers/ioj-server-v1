package services

import (
	"errors"

	"github.com/infinity-oj/server/internal/pkg/crypto"

	"github.com/infinity-oj/server/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	List()
	Create(tp string, properties map[string]string, inputs [][]byte) (*models.Judgement, error)
	Update() error

	PullJudgement(judgementType string) (string, *repositories.JudgementElement)
	PushJudgement(token string, outputs [][]byte) error
}

type DefaultJudgementsService struct {
	logger            *zap.Logger
	Repository        repositories.JudgementsRepository
	FileService       FilesService
	SubmissionService SubmissionsService
	tokenMap          map[string]*repositories.JudgementElement
}

func (d DefaultJudgementsService) List() {
	d.Repository.List()
}

func (d DefaultJudgementsService) Create(tp string, properties map[string]string, inputs [][]byte) (*models.Judgement, error) {
	judgement, err := d.Repository.Create(tp, properties, inputs)

	if err != nil {
		d.logger.Error("error while create judgement", zap.Error(err))
		return nil, err
	}

	if tp == "basic/file" {
		fileSpace, ok := properties["fileSpace"]
		if !ok {
			d.logger.Error("missing fileSpace", zap.String("jid", judgement.JudgementId))
			return judgement, err
		}
		fileName, ok := properties["fileName"]
		if !ok {
			d.logger.Error("missing fileName", zap.String("jid", judgement.JudgementId))
			return judgement, err
		}

		file, err := d.FileService.FetchFile(fileSpace, fileName)
		if err != nil {
			d.logger.Error("error while fetching file", zap.Error(err))
			return judgement, err
		}

		judgement.Outputs = crypto.EasyEncode([][]byte{file})
		err = d.Repository.Update(judgement)
		if err != nil {
			d.logger.Error("error while setting outputs", zap.Error(err))
			return judgement, err
		}

	} else {
		err := d.Repository.PutJudgementInQueue(judgement)
		if err != nil {
			return judgement, err
		}
	}

	return judgement, err
}

func (d DefaultJudgementsService) Update() error {
	panic("implement me")
}

func (d DefaultJudgementsService) PullJudgement(judgementType string) (token string, judgementElement *repositories.JudgementElement) {
	judgementElement = d.Repository.FetchJudgementInQueue(judgementType)
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

func NewJudgementsService(
	logger *zap.Logger,
	Repository repositories.JudgementsRepository,
	filesService FilesService,
	submissionService SubmissionsService,
) JudgementsService {
	return &DefaultJudgementsService{
		logger:            logger.With(zap.String("type", "DefaultJudgementService")),
		Repository:        Repository,
		FileService:       filesService,
		SubmissionService: submissionService,
		tokenMap:          make(map[string]*repositories.JudgementElement),
	}
}
