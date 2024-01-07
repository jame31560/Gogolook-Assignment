package task

import (
	"context"
	"net/http"
	"task/app/domain/model/aggregate"
	"task/app/domain/repository"
	task_service "task/app/domain/service/task"
	"task/app/infra/enum"
	"task/app/pkg/status"
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
	if ok := usecase.taskService.CheckName(task.Name); !ok {
		return nil, status.CreateError.WithHttpCode(http.StatusBadRequest).WithMsg("Name is required.")
	}

	task.Status = enum.TaskStatusIncomplete
	task.ID = usecase.taskService.NewTaskID()

	err := usecase.taskRepo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	event := &CreateTaskEvent{
		ID: task.ID,
	}
	return event, nil
}

func (usecase *taskUsecase) DeleteTask(
	ctx context.Context,
	cmd *DeleteTaskCmd,
) (*DeleteTaskEvent, error) {
	err := usecase.taskRepo.DeleteTask(cmd.ID)
	if err != nil {
		return nil, err
	}

	event := &DeleteTaskEvent{
		cmd.ID,
	}
	return event, nil
}

func (usecase *taskUsecase) EditTask(
	ctx context.Context,
	cmd *EditTaskCmd,
) (*EditTaskEvent, error) {
	if ok := usecase.taskService.CheckName(cmd.Name); !ok {
		return nil, status.UpdateError.WithHttpCode(http.StatusBadRequest).WithMsg("Name is required.")
	}
	status, err := enum.ToTaskStatusEnum(cmd.Status)
	if err != nil {
		return nil, err
	}

	task := &aggregate.Task{
		Status: status,
		Name:   cmd.Name,
	}

	err = usecase.taskRepo.UpdateTaskByID(cmd.ID, task)
	if err != nil {
		return nil, err
	}

	event := &EditTaskEvent{
		ID: cmd.ID,
	}
	return event, nil
}

func (usecase *taskUsecase) GetTaskList(
	ctx context.Context,
	cmd *GetTaskListCmd,
) (*GetTaskListEvent, error) {
	event := &GetTaskListEvent{
		TaskList: make([]*TaskDto, 0),
	}

	if cmd.ID != "" {
		task, err := usecase.taskRepo.GetTaskByID(cmd.ID)
		if err != nil {
			return nil, err
		}
		event.TaskList = []*TaskDto{
			{
				ID:     task.ID,
				Name:   task.Name,
				Status: int8(task.Status),
			},
		}
		return event, nil
	}

	taskList, err := usecase.taskRepo.QueryTaskList(cmd.Name, cmd.Status)
	if err != nil {
		return nil, err
	}

	for _, task := range taskList {
		event.TaskList = append(event.TaskList, &TaskDto{
			ID:         task.ID,
			Name:       task.Name,
			Status:     int8(task.Status),
			CreateTime: task.CreateTime,
			UpdateTime: task.UpdateTime,
		})
	}

	return event, nil
}
