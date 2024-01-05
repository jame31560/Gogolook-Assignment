package task

import (
	"net/http"
	aggregate "task/app/domain/model/aggreate"
	"task/app/domain/repository"
	"task/app/pkg/status"
	"time"
)

type taskRepo struct {
	taskList []*aggregate.Task
}

func NewTaskRepo() repository.TaskRepoInterface {
	return &taskRepo{
		taskList: make([]*aggregate.Task, 0),
	}
}

func (repo *taskRepo) CreateTask(task *aggregate.Task) error {
	now := time.Now()
	task.CreateTime = now
	task.UpdateTime = now
	repo.taskList = append(repo.taskList, task)
	return nil
}

func (repo *taskRepo) GetTaskByID(ID string) (*aggregate.Task, error) {
	for _, task := range repo.taskList {
		if task.ID == ID {
			return task, nil
		}
	}
	return nil, status.QueryError.WithHttpCode(http.StatusNotFound).WithMsg("Task not found")
}

func (repo *taskRepo) DeleteTask(ID string) error {
	for idx, task := range repo.taskList {
		if task.ID == ID {
			repo.taskList = append(repo.taskList[:idx], repo.taskList[idx+1:]...)
			return nil
		}
	}
	return status.DeleteError.WithHttpCode(http.StatusNotFound).WithMsg("Task not found")
}
