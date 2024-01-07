package task_mock

import (
	aggregate "task/app/domain/model/aggreate"
	"task/app/domain/repository"
	"task/app/infra/enum"
	"time"
)

type taskRepoMock struct{}

func NewTaskRepoMock() repository.TaskRepoInterface {
	return &taskRepoMock{}
}

func (repo *taskRepoMock) CreateTask(task *aggregate.Task) error {
	return nil
}

func (repo *taskRepoMock) GetTaskByID(ID string) (*aggregate.Task, error) {
	return &aggregate.Task{
		ID:         ID,
		Status:     enum.TaskStatusIncomplete,
		Name:       "Task",
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}, nil
}

func (repo *taskRepoMock) DeleteTask(ID string) error {
	return nil
}

func (repo *taskRepoMock) QueryTaskList(name string, statusList []int8) ([]*aggregate.Task, error) {
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

	return []*aggregate.Task{
		task1,
		task2,
	}, nil
}

func (repo *taskRepoMock) UpdateTaskByID(ID string, task *aggregate.Task) error {
	return nil
}
