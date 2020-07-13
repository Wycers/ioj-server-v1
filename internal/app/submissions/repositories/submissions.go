package repositories

import (
	"container/list"
	"sync"

	"github.com/infinity-oj/server/internal/pkg/nodeEngine"

	"github.com/google/uuid"

	"github.com/infinity-oj/server/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SubmissionRepository interface {
	FetchSubmissionBySubmissionId(submissionId string) (*models.Submission, error)
	Create(submitterID uint64, problemId, userSpace string) (s *models.Submission, err error)
	Update(s *models.Submission) error

	CreateSubmissionInQueue(s *models.Submission) *SubmissionElement
	FetchSubmissionInQueueById(submissionId string) *SubmissionElement
}

type MysqlSubmissionsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	queue  *list.List
	mutex  *sync.Mutex
}

type SubmissionElement struct {
	id string

	obj    *models.Submission
	graph  *nodeEngine.Graph
	result map[int][]byte
}

type JudgementElement struct {
	Type       string
	Properties map[string]interface{}
	Inputs     [][]byte
}

func (se SubmissionElement) FindUpstreams() (res []*JudgementElement) {

	ids := se.graph.Run()

	for _, block := range ids {
		var inputs [][]byte
		for _, linkId := range block.Inputs {
			if data, ok := se.result[linkId]; ok {
				inputs = append(inputs, data)
			} else {
				panic("something wrong")
			}
		}

		res = append(res, &JudgementElement{
			Type:       block.Type,
			Properties: block.Properties,
			Inputs:     inputs,
		})
	}

	return
}

func (m MysqlSubmissionsRepository) CreateSubmissionInQueue(s *models.Submission) *SubmissionElement {

	graph := nodeEngine.NewGraphByFile("graph.json")

	result := make(map[int][]byte)

	submission := &SubmissionElement{
		obj:    s,
		graph:  graph,
		result: result,
	}

	m.queue.PushBack(submission)

	return submission
}

func (m MysqlSubmissionsRepository) FetchSubmissionInQueueById(submissionId string) *SubmissionElement {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for se := m.queue.Front(); se != nil; se = se.Next() {
		judgement, ok := se.Value.(*SubmissionElement)
		if !ok {
			continue
		}

		if judgement.obj.SubmissionId == submissionId {
			return judgement
		}
	}

	return nil
}

func (m MysqlSubmissionsRepository) FetchSubmissionBySubmissionId(submissionId string) (*models.Submission, error) {
	submission := &models.Submission{}
	err := m.db.Where("submission_id = ?", submissionId).First(&submissionId).Error
	return submission, err
}

func (m MysqlSubmissionsRepository) Create(submitterId uint64, problemId, userSpace string) (s *models.Submission, err error) {
	s = &models.Submission{
		SubmissionId: uuid.New().String(),
		SubmitterId:  submitterId,
		ProblemID:    problemId,
		Judgements:   nil,
		UserSpace:    userSpace,
		Status:       models.Pending,
	}

	if err = m.db.Create(s).Error; err != nil {
		return nil, errors.Wrapf(err,
			" create submission with username: %d, problemID: %s, userSpace: %s",
			submitterId, problemId, userSpace,
		)
	}

	return
}

func (m MysqlSubmissionsRepository) GetUpstreams() (res []*JudgementElement) {

	for se := m.queue.Front(); se != nil; se = se.Next() {
		judgement, ok := se.Value.(*SubmissionElement)

		if !ok {
			continue
		}

		res = append(res, judgement.FindUpstreams()...)
	}

	return
}

func (m MysqlSubmissionsRepository) Update(s *models.Submission) (err error) {
	err = m.db.Save(s).Error
	return
}

func NewMysqlSubmissionsRepository(logger *zap.Logger, db *gorm.DB) SubmissionRepository {
	return &MysqlSubmissionsRepository{
		logger: logger.With(zap.String("type", "SubmissionRepository")),
		db:     db,
		queue:  list.New(),
		mutex:  &sync.Mutex{},
	}
}
