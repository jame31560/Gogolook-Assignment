package task

import (
	"context"
	"task/app/domain/repository"
)

type taskUsecase struct{
  taskRepo repository.TaskRepoInterface
}

func NewTaskUsecase(
  taskRepo repository.TaskRepoInterface,
) TaskUsecaseInterface {
	return &taskUsecase{
    taskRepo: taskRepo,
  }
}

func (usecase *taskUsecase) CreateTask(
	ctx context.Context,
	cmd *CreateTaskCmd,
) (*CreateTaskEvent, error) {
	return nil, nil
}


func (usecase *taskUsecase) DeleteTask(
	ctx context.Context,
	cmd *DeleteTaskCmd,
) (*DeleteTaskEvent, error) {
	return nil, nil
}


func (usecase *taskUsecase) EditTask(
	ctx context.Context,
	cmd *EditTaskCmd,
) (*EditTaskEvent, error) {
	return nil, nil
}


func (usecase *taskUsecase) GetTaskList(
	ctx context.Context,
	cmd *GetTaskListCmd,
) (*GetTaskListEvent, error) {
	return nil, nil
}
