package repository

import "task/app/domain/model/aggregate"

type TaskRepoInterface interface {
	CreateTask(*aggregate.Task) error
	DeleteTask(string) error
	GetTaskByID(string) (*aggregate.Task, error)
	QueryTaskList(string, []int8) ([]*aggregate.Task, error)
	UpdateTaskByID(string, *aggregate.Task) error
}
