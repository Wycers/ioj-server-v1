package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	CreateJudgement(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error
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
	Map         map[string]*repositories.Judgement
	FileService FilesService
}

type TaskInputs struct {
	Arguments map[string]interface{}
	Slots     [][]byte
}

func (d DefaultJudgementsService) FetchJudgementTask(taskType string) (*repositories.Task, [][]byte) {
	task, inputs := d.Repository.FetchTask(taskType)
	if task == nil {
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

func (d DefaultJudgementsService) CreateJudgement(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error {
	err := d.Repository.Create(submissionId, publicSpace, privateSpace, userSpace)
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

func (d DefaultJudgementsService) FetchJudgement() *repositories.Judgement {
	judgement := d.Repository.Fetch()

	token := uuid.New().String()
	if judgement != nil {
		d.Map[token] = judgement
	}

	return judgement
}

func (d DefaultJudgementsService) FinishJudgement(token string, score uint64, msg string) error {
	judgement, ok := d.Map[token]
	fmt.Println(token)
	fmt.Printf("%x\n", &judgement)
	if ok {
		if judgement.Status == "idle" {
			judgement.Status = "done"
		}
	} else {
		return errors.New("qwq")
	}
	return nil
}

func NewJudgementsService(logger *zap.Logger, Repository repositories.JudgementsRepository, filesService FilesService) JudgementsService {
	return &DefaultJudgementsService{
		logger:      logger.With(zap.String("type", "DefaultJudgementService")),
		Repository:  Repository,
		Map:         make(map[string]*repositories.Judgement),
		FileService: filesService,
	}
}
