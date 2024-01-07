package task

import (
	"task/app/domain/model/aggreate"
	"task/app/infra/enum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	repo := &taskRepo{
		taskList: make([]*aggregate.Task, 0),
	}
	task := &aggregate.Task{
		ID:     "7780fc95-7a0a-45ef-a985-21c3e744c1d7",
		Name:   "Task",
		Status: enum.TaskStatusCompleted,
	}

	err := repo.CreateTask(task)
	assert.NoError(t, err)
	assert.Len(t, repo.taskList, 1)
}

func TestGetTaskByID(t *testing.T) {
	ID := "7780fc95-7a0a-45ef-a985-21c3e744c1d7"
	task := &aggregate.Task{
		ID:     ID,
		Name:   "Task",
		Status: enum.TaskStatusCompleted,
	}

	repo := &taskRepo{
		taskList: []*aggregate.Task{
			task,
		},
	}

	// 存在資料
	resultTask, err := repo.GetTaskByID(ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultTask)
	assert.EqualValuesf(t, ID, resultTask.ID, "Expected to get task with ID '%s', but got %s", ID, resultTask.ID)

	// 不存在資料
	resultTask, err = repo.GetTaskByID("123")
	assert.Nil(t, resultTask)
	assert.Error(t, err)
}

func TestDeleteTask(t *testing.T) {
	ID := "7780fc95-7a0a-45ef-a985-21c3e744c1d7"
	task := &aggregate.Task{
		ID:     ID,
		Name:   "Task",
		Status: enum.TaskStatusCompleted,
	}

	repo := &taskRepo{
		taskList: []*aggregate.Task{
			task,
		},
	}

	// 不存在資料
	err := repo.DeleteTask("123")
	assert.Error(t, err)
	assert.Len(t, repo.taskList, 1)

	// 存在資料
	err = repo.DeleteTask(ID)
	assert.NoError(t, err)
	assert.Len(t, repo.taskList, 0)
}

func TestUpdateTaskByID(t *testing.T) {
	ID := "7780fc95-7a0a-45ef-a985-21c3e744c1d7"
	task := &aggregate.Task{
		ID:     ID,
		Name:   "Task",
		Status: enum.TaskStatusCompleted,
	}

	repo := &taskRepo{
		taskList: []*aggregate.Task{
			task,
		},
	}

	// 成功更新
	updatedName := "UpdatedTask"
	updatedStatus := enum.TaskStatusIncomplete
	newTask := &aggregate.Task{
		Name:   updatedName,
		Status: updatedStatus,
	}
	err := repo.UpdateTaskByID(ID, newTask)
	assert.NoError(t, err)
	assert.Len(t, repo.taskList, 1)
	assert.EqualValues(t, repo.taskList[0].ID, ID)
	assert.EqualValues(t, repo.taskList[0].Name, updatedName)
	assert.EqualValues(t, repo.taskList[0].Status, updatedStatus)

	// 不存在資料更新
	err = repo.UpdateTaskByID("123", newTask)
	assert.Error(t, err)
	assert.Len(t, repo.taskList, 1)
	assert.EqualValues(t, repo.taskList[0].ID, ID)
	assert.EqualValues(t, repo.taskList[0].Name, updatedName)
	assert.EqualValues(t, repo.taskList[0].Status, updatedStatus)
}

func TestQueryTaskList(t *testing.T) {
	task1 := &aggregate.Task{
		ID:     "7780fc95-7a0a-45ef-a985-21c3e744c1d7",
		Name:   "Task1",
		Status: enum.TaskStatusIncomplete,
	}

	task2 := &aggregate.Task{
		ID:     "99239a34-7a0a-45ef-a985-21c3e744c1d7",
		Name:   "Task2",
		Status: enum.TaskStatusCompleted,
	}

	repo := &taskRepo{
		taskList: []*aggregate.Task{
			task1,
			task2,
		},
	}

	caseList := []struct {
		nameParam       string
		statusListParam []int8
		length          int
		firstID         string
	}{
		{nameParam: "", statusListParam: []int8{}, length: 0},
		{nameParam: "0", statusListParam: []int8{}, length: 0},
		{nameParam: "1", statusListParam: []int8{}, length: 0},
		{nameParam: "", statusListParam: enum.GetAllTaskStatusIntList(), length: 2},
		{nameParam: "0", statusListParam: enum.GetAllTaskStatusIntList(), length: 0},
		{nameParam: "1", statusListParam: enum.GetAllTaskStatusIntList(), length: 1, firstID: task1.ID},
		{nameParam: "", statusListParam: []int8{int8(enum.TaskStatusCompleted)}, length: 1, firstID: task2.ID},
		{nameParam: "0", statusListParam: []int8{int8(enum.TaskStatusCompleted)}, length: 0},
		{nameParam: "1", statusListParam: []int8{int8(enum.TaskStatusCompleted)}, length: 0},
		{nameParam: "2", statusListParam: []int8{int8(enum.TaskStatusCompleted)}, length: 1, firstID: task2.ID},
	}

	for _, testCase := range caseList {
		result, err := repo.QueryTaskList(testCase.nameParam, testCase.statusListParam)
		assert.NoError(t, err)
		assert.Len(t, result, testCase.length)
		if testCase.length == 0 || testCase.firstID == "" {
			continue
		}
		assert.EqualValues(t, testCase.firstID, result[0].ID)
	}
}
