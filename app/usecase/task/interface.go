package task

import "context"

type TaskUsecaseInterface interface {
	CreateTask(ctx context.Context, cmd *CreateTaskCmd) (*CreateTaskEvent, error)

	DeleteTask(ctx context.Context, cmd *DeleteTaskCmd) (*DeleteTaskEvent, error)

	EditTask(ctx context.Context, cmd *EditTaskCmd) (*EditTaskEvent, error)

	GetTaskList(ctx context.Context, cmd *GetTaskListCmd) (*GetTaskListEvent, error)
}
