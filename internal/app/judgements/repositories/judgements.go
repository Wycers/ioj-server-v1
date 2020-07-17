package repositories

import (
	"container/list"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/infinity-oj/server/internal/pkg/crypto"

	"github.com/google/uuid"

	"github.com/infinity-oj/server/internal/pkg/models"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type JudgementElement struct {
	Idle        bool
	JudgementId string
	Type        string
	Properties  map[string]string

	Inputs  [][]byte
	Outputs [][]byte

	obj *models.Judgement
}

type JudgementsRepository interface {
	List()
	Fetch() *models.Judgement
	Create(tp string, properties map[string]string, inputs [][]byte) (*models.Judgement, error)
	Update(judgement *models.Judgement) error

	PutJudgementInQueue(judgement *models.Judgement) error
	FetchJudgementInQueue(tp string) *JudgementElement
	ReturnJudgementInQueue(element *JudgementElement, outputs [][]byte) error
}

type MysqlJudgementsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	mutex  *sync.Mutex
	queue  *list.List // judgements list
}

func (m MysqlJudgementsRepository) Create(tp string, properties map[string]string, inputs [][]byte) (*models.Judgement, error) {
	propertiesJson, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}
	propertiesStr := string(propertiesJson)

	judgementId := uuid.New().String()
	judgement := &models.Judgement{
		JudgementId: judgementId,
		Type:        tp,
		Status:      models.Pending,
		Property:    propertiesStr,
		Inputs:      crypto.EasyEncode(inputs),
		Outputs:     "",
	}

	err = m.db.Save(&judgement).Error

	if err != nil {
		return nil, err
	}
	return judgement, nil
}

func (m MysqlJudgementsRepository) Update(judgement *models.Judgement) error {
	err := m.db.Save(&judgement).Error
	return err
}

func (m MysqlJudgementsRepository) PutJudgementInQueue(judgement *models.Judgement) error {

	judgementId := judgement.JudgementId
	tp := judgement.Type

	var properties map[string]string
	propertiesJson := judgement.Property
	if propertiesJson != "" {
		if err := json.Unmarshal([]byte(propertiesJson), &properties); err != nil {
			return err
		}
	}
	inputs, err := crypto.EasyDecode(judgement.Inputs)
	if err != nil {
		return err
	}

	judgementInQueue := &JudgementElement{
		Idle: false,

		JudgementId: judgementId,
		Type:        tp,
		Properties:  properties,

		Inputs:  inputs,
		Outputs: nil,

		obj: judgement,
	}

	m.queue.PushBack(judgementInQueue)

	return nil
}

func (m MysqlJudgementsRepository) List() {
	fmt.Println("=== START ===")

	for te := m.queue.Front(); te != nil; te = te.Next() {
		judgementElement, ok := te.Value.(*JudgementElement)

		if !ok {
			fmt.Println(te.Value)
			continue
		}

		fmt.Printf("id: %s\ntype: %s\n idle: %v\n, inputs: %+v\n, outputs: %+v\n\n",
			judgementElement.JudgementId,
			judgementElement.Type,
			judgementElement.Idle,
			judgementElement.Inputs,
			judgementElement.Properties,
		)
	}

	fmt.Println("==== END ====")
}

func (m MysqlJudgementsRepository) Fetch() *models.Judgement {
	return nil
}

// FetchJudgementInQueue returns task with specific task type.
func (m MysqlJudgementsRepository) FetchJudgementInQueue(taskType string) *JudgementElement {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for te := m.queue.Front(); te != nil; te = te.Next() {
		judgementElement, ok := te.Value.(*JudgementElement)

		if !ok {
			fmt.Println(te.Value)
			continue
		}

		if !judgementElement.Idle {
			continue
		}
		if judgementElement.Type != taskType {
			continue
		}

		judgementElement.Idle = true

		return judgementElement
	}

	return nil
}

func (m MysqlJudgementsRepository) ReturnJudgementInQueue(element *JudgementElement, outputs [][]byte) error {
	judgement := element.obj

	judgement.Outputs = crypto.EasyEncode(outputs)
	element.Outputs = outputs

	if element.Idle {
		for te := m.queue.Front(); te != nil; te = te.Next() {
			je, ok := te.Value.(*JudgementElement)

			if !ok {
				continue
			}

			if je == element {
				m.queue.Remove(te)
				break
			}
		}
	} else {

	}

	err := m.Update(judgement)

	return err
}

func NewMysqlJudgementsRepository(logger *zap.Logger, db *gorm.DB) JudgementsRepository {
	return &MysqlJudgementsRepository{
		logger: logger.With(zap.String("type", "JudgementsRepository")),
		db:     db,
		mutex:  &sync.Mutex{},
		queue:  list.New(),
	}
}
