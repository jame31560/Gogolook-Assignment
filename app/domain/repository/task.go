package repository

import aggregate "task/app/domain/model/aggreate"

type TaskRepoInterface interface{
  CreateTask (*aggregate.Task) error
  GetTaskByID (string) (*aggregate.Task, error)
}
