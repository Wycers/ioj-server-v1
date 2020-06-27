package repositories

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFile = flag.String("f", "configs/judgements.yml", "set config file which viper will loading.")

func TestCreateJudgement(t *testing.T) {
	flag.Parse()

	sto, err := CreateJudgementsRepository(*configFile)
	if err != nil {
		t.Fatalf("create judgement Repository error,%+v", err)
	}
	t.Run("BeforeCreate", func(t *testing.T) {
		task, inputs := sto.FetchTask("*")
		assert.Nil(t, task)
		assert.Nil(t, inputs)
	})

	t.Run("Create", func(t *testing.T) {
		err := sto.Create(0, "test_pub", "test_pri", "test_usr")
		assert.NoError(t, err)
	})

	t.Run("AfterCreate", func(t *testing.T) {
		task, inputs := sto.FetchTask("*")
		assert.Nil(t, task)
		assert.Nil(t, inputs)
	})

	t.Run("AfterCreate", func(t *testing.T) {
		task1, inputs1 := sto.FetchTask("basic/file")
		assert.NotNil(t, task1)
		assert.Equal(t, 0, len(inputs1))
		assert.Equal(t, "basic/file", task1.Type)

		task2, inputs2 := sto.FetchTask("basic/file")
		assert.NotNil(t, task2)
		assert.Equal(t, 0, len(inputs2))
		assert.Equal(t, "basic/file", task2.Type)

		fmt.Println(task1.TaskId)
		fmt.Println(task2.TaskId)

		assert.False(t, task1.TaskId == task2.TaskId && task1.JudgementId == task2.JudgementId)

		var err error
		err = sto.ReturnTask(task1.JudgementId, task1.TaskId, [][]byte{
			[]byte("1 2 3"),
		})
		assert.Nil(t, err)
		err = sto.ReturnTask(task2.JudgementId, task2.TaskId, [][]byte{
			[]byte("1 2 3"),
		})
		assert.Nil(t, err)

		for {
			task, inputs := sto.FetchTask("basic/file")

			if task == nil {
				break
			}
			assert.Equal(t, 0, len(inputs))

			err = sto.ReturnTask(task.JudgementId, task.TaskId, [][]byte{
				[]byte("1 2 3"),
			})

			assert.Nil(t, err)
		}

		{

			task, inputs := sto.FetchTask("builder/Clang")

			err = sto.ReturnTask(task.JudgementId, task.TaskId, [][]byte{
				[]byte("1 2 3"),
			})

			assert.Equal(t, [][]byte{[]byte("1 2 3")}, inputs)
			assert.Nil(t, err)
		}

		{
			task, inputs := sto.FetchTask("basic/file")

			assert.Nil(t, inputs)
			assert.Nil(t, task)
		}

		sto.List()
	})
}
