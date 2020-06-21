package repositories

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/infinity-oj/server/internal/pkg/models"

	"github.com/infinity-oj/server/internal/pkg/utils/random"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Judgement struct {
	Status string
	Token  string
	Time   uint64
	Done   chan bool

	PublicSpace  string
	PrivateSpace string
	UserSpace    string
	TestCase     string
	Obj          *models.Judgement
}

type JudgementsRepository interface {
	Create(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error
	Fetch() *Judgement
	List()
}

type MysqlJudgementsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	queue  *list.List
	mutex  *sync.Mutex
}

func (m MysqlJudgementsRepository) List() {
	for e := m.queue.Front(); e != nil; e = e.Next() {
		judgement, ok := e.Value.(Judgement)
		if !ok {
			fmt.Println("...")
		}
		fmt.Printf("Address: %x\nTestCase: %s\nStatus: %s\n", e.Value, judgement.TestCase, judgement.Status)
	}
}

func (m MysqlJudgementsRepository) Fetch() *Judgement {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var next *list.Element
	for e := m.queue.Front(); e != nil; e = next {
		next = e.Next()
		judgement, ok := e.Value.(*Judgement)
		if !ok || judgement.Status == "done" {
			m.queue.Remove(e)
			continue
		}

		judgement.Status = "idle"
		judgement.Token = random.RandStringRunes(8)
		judgement.Done = make(chan bool)
		return judgement
	}
	return nil
}

func (m MysqlJudgementsRepository) Create(submissionId uint64, publicSpace, privateSpace, userSpace, testCase string) error {
	judgement := &models.Judgement{
		SubmissionID: submissionId,
		TestCase:     testCase,
		Score:        0,
		Status:       models.Pending,
	}

	if err := m.db.Create(&judgement).Error; err != nil {
		return nil
	}

	m.queue.PushBack(&Judgement{
		Status:       "",
		Token:        random.RandStringRunes(8),
		Done:         make(chan bool),
		PublicSpace:  publicSpace,
		PrivateSpace: privateSpace,
		UserSpace:    userSpace,
		TestCase:     testCase,
		Obj:          judgement,
	})
	return nil
}

func NewMysqlJudgementsRepository(logger *zap.Logger, db *gorm.DB) JudgementsRepository {
	return &MysqlJudgementsRepository{
		logger: logger.With(zap.String("type", "JudgementsRepository")),
		db:     db,
		queue:  list.New(),
		mutex:  &sync.Mutex{},
	}
}
