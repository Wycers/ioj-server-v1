package services

import (
	"github.com/pkg/errors"

	"github.com/Infinity-OJ/Server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	CreateJudgement(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error
	FetchJudgement() *repositories.Judgement
	FinishJudgement(token string, score uint64) error
	FetchFile(fileSpace, fileName string) ([]byte, error)
}

type DefaultJudgementsService struct {
	logger      *zap.Logger
	Repository  repositories.JudgementsRepository
	mp          map[string]*repositories.Judgement
	FileService FilesService
}

func (d DefaultJudgementsService) FetchFile(fileSpace, fileName string) ([]byte, error) {
	return d.FileService.FetchFile(fileSpace, fileName)
}

func (d DefaultJudgementsService) CreateJudgement(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error {
	err := d.Repository.Create(submissionId, publicSpace, privateSpace, userSpace, testCase)
	if err != nil {
		return err
	}
	return nil
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
	d.mp[judgement.Token] = judgement
	return judgement
	//go func() {
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(judgement.Time)*time.Millisecond)
	//	defer cancel()
	//	dosomething(ctx, judgement)
	//	judgement.Done <- true
	//}()
	//select {
	//case res := <-done:
	//	fmt.Println(time.Now(), res, id)
	//case <-time.After(time.Duration(10) * time.Second):
	//	fmt.Println(time.Now(), "timeout ", 10)
	//}
}

func (d DefaultJudgementsService) FinishJudgement(token string, score uint64) error {
	judgement, ok := d.mp[token]
	if ok {
		if judgement.Status == "idle" {
			judgement.Done <- true
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
		mp:          make(map[string]*repositories.Judgement),
		FileService: filesService,
	}
}
