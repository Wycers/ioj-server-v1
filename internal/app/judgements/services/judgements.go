package services

import (
	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/pkg/errors"

	"github.com/Infinity-OJ/Server/internal/app/judgements/repositories"
	"go.uber.org/zap"
)

type JudgementsService interface {
	FetchJudgement() *repositories.Judgement
	FinishJudgement(token string, score uint64) error
}

type DefaultJudgementsService struct {
	logger      *zap.Logger
	Repository  repositories.JudgementsRepository
	mp          map[string]*repositories.Judgement
	fileService proto.FilesClient
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

func NewJudgementsService(logger *zap.Logger, Repository repositories.JudgementsRepository, fileSerivce proto.FilesClient) JudgementsService {
	return &DefaultJudgementsService{
		logger:      logger.With(zap.String("type", "DefaultJudgementservice")),
		Repository:  Repository,
		mp:          make(map[string]*repositories.Judgement),
		fileService: fileSerivce,
	}
}
