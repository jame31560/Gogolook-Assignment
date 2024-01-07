package repository

import aggregate "task/app/domain/model/aggreate"

type TaskRepoInterface interface{
  CreateTask (*aggregate.Task) error
  DeleteTask (string) error
  GetTaskByID (string) (*aggregate.Task, error)
  QueryTaskList (string, []int8) ([]*aggregate.Task, error)
  UpdateTaskByID (string, *aggregate.Task) error
}
