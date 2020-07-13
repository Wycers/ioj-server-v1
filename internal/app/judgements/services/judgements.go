package services

import (
	"errors"

	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	Create(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error
	FetchJudgement() *repositories.Judgement
	FetchJudgementByToken(token string) *repositories.Judgement
	FinishJudgement(token string, score uint64, msg string) error
	FetchFile(fileSpace, fileName string) ([]byte, error)
	List()

	FetchJudgementTask(string) (*repositories.Task, [][]byte)
}

type DefaultJudgementsService struct {
	logger      *zap.Logger
	Repository  repositories.JudgementsRepository
	FileService FilesService
	tokenMap    map[string]*repositories.JudgementElement
}

type TaskInputs struct {
	Arguments map[string]interface{}
	Slots     [][]byte
}

func (d DefaultJudgementsService) FetchJudgement(taskType string) (*repositories.Task, [][]byte) {
	judgement := d.Repository.FetchJudgementInQueue(taskType)
	if judgement == nil {
		return nil, nil
	}
	return task, inputs
}

func (d DefaultJudgementsService) FetchJudgementByToken(token string) *repositories.Judgement {
	judgement, ok := d.Map[token]
	if ok {
		return judgement
	} else {
		return nil
	}
}

func (d DefaultJudgementsService) FetchFile(fileSpace, fileName string) ([]byte, error) {
	return d.FileService.FetchFile(fileSpace, fileName)
}

func (d DefaultJudgementsService) Create() error {
	judgement, err := d.Repository.Create()
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultJudgementsService) List() {
	d.Repository.List()
}

//func dosomething(ctx context.Context, judgement *repositories.Judgement) {
//	judgement.Mutex.Lock()
//	defer judgement.Mutex.Unlock()
//	select {
//	case <-ctx.Done():
//		fmt.Println(time.Now(), "op timeout", val)
//		return
//	default:
//		//....
//		time.Sleep(time.Duration(2) * time.Second)
//		id = val
//	}
//
//}

func (d DefaultJudgementsService) FetchJudgement() (*repositories.Judgement, error) {
	judgement := d.Repository.Fetch()

	return judgement
}

func (d DefaultJudgementsService) FinishJudgement(token string, outputs [][]byte) error {
	element, ok := d.tokenMap[token]
	if !ok {
		return errors.New("unknown judgement")
	}

	err := d.Repository.ReturnJudgementInQueue(element, outputs)
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
