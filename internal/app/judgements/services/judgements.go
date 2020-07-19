package services

import (
	"container/list"
	"errors"
	"sync"
	"time"

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

	ReportJudgement(element *repositories.JudgementElement)
}

type DefaultJudgementsService struct {
	logger            *zap.Logger
	Repository        repositories.JudgementsRepository
	FileService       FilesService
	SubmissionService SubmissionsService
	tokenMap          map[string]*repositories.JudgementElement

	waitList *list.List
	mutex    *sync.Mutex
}

func (d DefaultJudgementsService) ReportJudgement(element *repositories.JudgementElement) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.waitList.PushBack(element)
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

		judgementElement, err := d.Repository.WrapJudgement(judgement)
		if err != nil {
			return judgement, err
		}
		d.ReportJudgement(judgementElement)
		go func() {
			for {
				err = d.SubmissionService.ReportJudgement(judgement.JudgementId, [][]byte{file})
				if err == nil {
					break
				}
				time.Sleep(time.Second)
			}
		}()
	} else {
		judgementElement, err := d.Repository.WrapJudgement(judgement)
		if err != nil {
			return judgement, err
		}
		d.Repository.PutJudgementInQueue(judgementElement)
	}

	return judgement, err
}

func (d DefaultJudgementsService) Update() error {
	panic("implement me")
}

func (d DefaultJudgementsService) PullJudgement(judgementType string) (token string, element *repositories.JudgementElement) {
	d.logger.Info("pull judgement", zap.String("type", judgementType))
	element = d.Repository.FetchJudgementInQueue(judgementType)
	if element != nil {
		d.logger.Info("pull judgement", zap.String("judgement id", element.JudgementId))
		// TODO: use jwt
		token = uuid.New().String()
		d.tokenMap[token] = element
	} else {
		d.logger.Info("pull judgement: nothing")
	}
	return
}

func (d DefaultJudgementsService) PushJudgement(token string, outputs [][]byte) error {
	element, ok := d.tokenMap[token]

	if !ok {
		return errors.New("invalid token")
	}

	d.logger.Info("push judgement",
		zap.String("type", element.Type))
	err := d.Repository.ReturnJudgementInQueue(element, outputs)
	if err != nil {
		d.logger.Error("Push judgement", zap.Error(err))
		return err
	}

	delete(d.tokenMap, token)
	err = d.SubmissionService.ReportJudgement(element.JudgementId, outputs)
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

		waitList: list.New(),
		mutex:    &sync.Mutex{},
	}
}
