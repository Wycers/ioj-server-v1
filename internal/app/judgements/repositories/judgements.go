package repositories

import (
	"container/list"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/Infinity-OJ/Server/internal/pkg/utils/random"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Mutex struct {
	sync.Mutex
}
type Judgement struct {
	Mutex  Mutex
	Status string
	Token  string
	Time   uint64
	Done   chan bool
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}

type JudgementsRepository interface {
	Create() error
	Fetch() *Judgement
}

type MysqlJudgementsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	queue  *list.List
}

const mutexLocked = 1 << iota

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

func (m MysqlJudgementsRepository) Create() error {
	m.queue.PushBack(Judgement{})
	panic("implement me")
}

func NewMysqlJudgementsRepository(logger *zap.Logger, db *gorm.DB) JudgementsRepository {
	return &MysqlJudgementsRepository{
		logger: logger.With(zap.String("type", "JudgementsRepository")),
		db:     db,
		queue:  list.New(),
	}
}
