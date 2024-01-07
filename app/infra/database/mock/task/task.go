package task_mock

import (
	aggregate "task/app/domain/model/aggreate"
	"task/app/infra/enum"
	"time"
)

type taskRepoMock struct {
	CreateTaskFunc     func(task *aggregate.Task) error
	GetTaskByIDFunc    func(ID string) (*aggregate.Task, error)
	DeleteTaskFunc     func(ID string) error
	QueryTaskListFunc  func(name string, statusList []int8) ([]*aggregate.Task, error)
	UpdateTaskByIDFunc func(ID string, task *aggregate.Task) error
}

func NewTaskRepoMock() *taskRepoMock {
	return &taskRepoMock{}
}

func (repo *taskRepoMock) CreateTask(task *aggregate.Task) error {
	if repo.CreateTaskFunc != nil {
		return repo.CreateTaskFunc(task)
	}
	return nil
}

func (repo *taskRepoMock) GetTaskByID(ID string) (*aggregate.Task, error) {
	if repo.GetTaskByIDFunc != nil {
		return repo.GetTaskByIDFunc(ID)
	}
	return &aggregate.Task{
		ID:         ID,
		Status:     enum.TaskStatusIncomplete,
		Name:       "Task",
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}, nil
}

func (repo *taskRepoMock) DeleteTask(ID string) error {
	if repo.DeleteTaskFunc != nil {
		return repo.DeleteTaskFunc(ID)
	}
	return nil
}

func (repo *taskRepoMock) QueryTaskList(name string, statusList []int8) ([]*aggregate.Task, error) {
	if repo.QueryTaskListFunc != nil {
		return repo.QueryTaskListFunc(name, statusList)
	}
	return []*aggregate.Task{}, nil
}

func (repo *taskRepoMock) UpdateTaskByID(ID string, task *aggregate.Task) error {
	if repo.UpdateTaskByIDFunc != nil {
		return repo.UpdateTaskByIDFunc(ID, task)
	}
	return nil
}
