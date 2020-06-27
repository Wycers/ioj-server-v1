package repositories

import (
	"container/list"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/infinity-oj/server/internal/pkg/models"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Judgement struct {
	Status string
	Time   uint64

	Id string

	PublicSpace  string
	PrivateSpace string
	UserSpace    string

	Obj *models.Judgement

	Graph *Graph
	Tasks *list.List

	Result map[int][]byte
}

type Task struct {
	Status      string
	JudgementId string
	TaskId      string
	Type        string
	Properties  map[string]interface{}

	block *Block
}

type JudgementsRepository interface {
	Create(submissionId uint64, publicSpace, privateSpace, userSpace string) error
	FetchTask(string) (*Task, [][]byte)
	ReturnTask(judgementId, taskId string, results [][]byte) error

	Fetch() *Judgement
	List()
}

type MysqlJudgementsRepository struct {
	logger *zap.Logger
	db     *gorm.DB

	mutex *sync.Mutex

	judgementMap map[string]*Judgement // map judgement id to judgement
	taskMap      map[string]*Task      // map task id to task

	judgements *list.List // judgements list
}

func (m MysqlJudgementsRepository) List() {
	fmt.Println("=== START ===")
	for je := m.judgements.Front(); je != nil; je = je.Next() {
		judgement, ok := je.Value.(*Judgement)
		if !ok {
			fmt.Println(je.Value)
			continue
		}

		fmt.Println("=================")
		for te := judgement.Tasks.Front(); te != nil; te = te.Next() {
			task, ok := te.Value.(*Task)
			if !ok {
				fmt.Println(te.Value)
				continue
			}
			fmt.Printf("Judgement Id: %s\nTask Id: %s\nType: %s\nStatus: %s\n", task.JudgementId, task.TaskId, task.Type, task.Status)
		}
	}
	fmt.Println("==== END ====")
}

func (m MysqlJudgementsRepository) Fetch() *Judgement {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return nil
}

func (m MysqlJudgementsRepository) GetTaskInputs(task *Task) [][]byte {
	var inputs [][]byte

	return inputs
}

// FetchTask returns task with specific task type.
func (m MysqlJudgementsRepository) FetchTask(taskType string) (*Task, [][]byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for je := m.judgements.Front(); je != nil; je = je.Next() {
		judgement, ok := je.Value.(*Judgement)
		if !ok {
			fmt.Println(je.Value)
			continue
		}

		if judgement.Status == "pending" {
			judgement.Status = "running"
		}

		ids := judgement.Graph.Run()

		for _, block := range ids {
			fmt.Println("New task!", block.Type)
			judgement.Tasks.PushBack(&Task{
				Status:      "pending",
				JudgementId: judgement.Id,
				TaskId:      uuid.New().String(),
				Type:        block.Type,
				Properties:  nil,

				block: block,
			})
		}

		for te := judgement.Tasks.Front(); te != nil; te = te.Next() {
			task, ok := te.Value.(*Task)
			if !ok {
				fmt.Println(te.Value)
				continue
			}

			if task.Status != "pending" {
				continue
			}
			if task.Type != taskType {
				continue
			}

			task.Status = "working"

			var inputs [][]byte

			block := task.block
			for _, linkId := range block.Inputs {
				if data, ok := judgement.Result[linkId]; ok {
					inputs = append(inputs, data)
				} else {
					panic("something wrong")
				}
			}

			return task, inputs
		}
	}
	return nil, nil
}

// ReturnTask sets task result.
func (m MysqlJudgementsRepository) ReturnTask(judgementId, taskId string, results [][]byte) error {
	judgement, task := m.findJudgementTask(judgementId, taskId)
	if task == nil {
		return errors.New("unknown task: " + taskId)
	}

	task.Status = "done"

	//fmt.Println(task.block.Output)
	//fmt.Println(results)

	if len(task.block.Output) != len(results) {
		return errors.New("output slots mismatch")
	}

	blockId := task.block.Id
	for index, result := range results {
		fmt.Println(blockId, index)
		links := judgement.Graph.findLinkBySourcePort(blockId, index)
		fmt.Println(links, result)
		for _, link := range links {
			fmt.Println(link.Id)
			judgement.Result[link.Id] = result
			fmt.Println("set link", link.Id, "to", result)
		}
	}
	task.block.Done()

	return nil
}

func (m MysqlJudgementsRepository) Create(submissionId uint64, publicSpace, privateSpace, userSpace string) error {
	judgementObj := &models.Judgement{
		SubmissionID: submissionId,
		Score:        0,
		Status:       models.Pending,
	}

	//if err := m.db.Create(&judgementObj).Error; err != nil {
	//	return nil
	//}

	graph := NewGraphByFile("graph.json")

	judgement := &Judgement{
		Id:           uuid.New().String(),
		Status:       "pending",
		PublicSpace:  publicSpace,
		PrivateSpace: privateSpace,
		UserSpace:    userSpace,
		Obj:          judgementObj,

		Graph:  graph,
		Tasks:  list.New(),
		Result: make(map[int][]byte),
	}

	m.judgements.PushBack(judgement)

	return nil
}

func NewMysqlJudgementsRepository(logger *zap.Logger, db *gorm.DB) JudgementsRepository {
	return &MysqlJudgementsRepository{
		logger: logger.With(zap.String("type", "JudgementsRepository")),
		db:     db,
		mutex:  &sync.Mutex{},
		// TODO: Use it for better performance
		//judgementMap: make(map[string]*Judgement),
		//taskMap:      make(map[string]*Task),
		judgements: list.New(),
	}
}

func (m MysqlJudgementsRepository) findJudgement(judgementId string) *Judgement {
	for je := m.judgements.Front(); je != nil; je = je.Next() {
		judgement, ok := je.Value.(*Judgement)
		if !ok {
			continue
		}

		if judgement.Id != judgementId {
			continue
		}

		return judgement
	}
	return nil
}
func (j *Judgement) findTask(taskId string) *Task {
	for te := j.Tasks.Front(); te != nil; te = te.Next() {
		task, ok := te.Value.(*Task)
		if !ok {
			fmt.Println(te.Value)
			continue
		}

		if task.TaskId != taskId {
			continue
		}
		return task
	}
	return nil
}
func (m MysqlJudgementsRepository) findJudgementTask(judgementId, taskId string) (*Judgement, *Task) {
	judgement := m.findJudgement(judgementId)
	if judgement == nil {
		return nil, nil
	}
	task := judgement.findTask(taskId)
	return judgement, task
}
