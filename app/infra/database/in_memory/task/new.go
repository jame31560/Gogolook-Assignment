package task

import "task/app/domain/repository"

type taskRepo struct {}

func NewTaskRepo() repository.TaskRepoInterface {
	return &taskRepo{}
}
