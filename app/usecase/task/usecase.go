package task

import (
	"context"
	aggregate "task/app/domain/model/aggreate"
	"task/app/domain/repository"
	task_service "task/app/domain/service/task"
	"task/app/infra/enum"
)

type taskUsecase struct {
	taskRepo    repository.TaskRepoInterface
	taskService task_service.TaskServiceInterface
}

func NewTaskUsecase(
	taskRepo repository.TaskRepoInterface,
	taskService task_service.TaskServiceInterface,
) TaskUsecaseInterface {
	return &taskUsecase{
		taskRepo:    taskRepo,
		taskService: taskService,
	}
}

func (usecase *taskUsecase) CreateTask(
	ctx context.Context,
	cmd *CreateTaskCmd,
) (*CreateTaskEvent, error) {
	task := new(aggregate.Task)

	task.Name = cmd.Name
	task.Status = enum.TaskStatusIncomplete
  task.ID = usecase.taskService.NewTaskID()

	err := usecase.taskRepo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	event := new(CreateTaskEvent)
	return event, nil
}

func (usecase *taskUsecase) DeleteTask(
	ctx context.Context,
	cmd *DeleteTaskCmd,
) (*DeleteTaskEvent, error) {
	event := new(DeleteTaskEvent)
	return event, nil
}

func (usecase *taskUsecase) EditTask(
	ctx context.Context,
	cmd *EditTaskCmd,
) (*EditTaskEvent, error) {
	event := new(EditTaskEvent)
	return event, nil
}

func (usecase *taskUsecase) GetTaskList(
	ctx context.Context,
	cmd *GetTaskListCmd,
) (*GetTaskListEvent, error) {
	event := new(GetTaskListEvent)
	return event, nil
}
