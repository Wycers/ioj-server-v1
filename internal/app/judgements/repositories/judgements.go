package repositories

import (
	"container/list"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/Infinity-OJ/Server/internal/pkg/models"

	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Mutex struct {
	sync.Mutex
}

const mutexLocked = 1 << iota

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}

type Judgement struct {
	Mutex  Mutex
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
}

type MysqlJudgementsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	queue  *list.List
}

func (m MysqlJudgementsRepository) Fetch() *Judgement {
	var next *list.Element
	for e := m.queue.Front(); e != nil; e = next {
		next = e.Next()
		judgement, ok := e.Value.(Judgement)
		if !ok {
			m.queue.Remove(e)
		}
		if judgement.Mutex.TryLock() {
			judgement.Status = "idle"
			judgement.Token = random.RandStringRunes(8)
			judgement.Done = make(chan bool)
			return &judgement
		}
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

	m.queue.PushBack(Judgement{
		Mutex:        Mutex{},
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
	}
}
