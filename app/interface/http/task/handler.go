package task

import (
	"task/app/interface/http/middle"
	task_usecase "task/app/usecase/task"
)

type taskHttpHandler struct {
	taskUsecase task_usecase.TaskUsecaseInterface
}

func NewTaskHttpHandler(
	taskUsecase task_usecase.TaskUsecaseInterface,
) TaskHttpHandlerInterface {
	return &taskHttpHandler{
		taskUsecase: taskUsecase,
	}
}

func (s *taskHttpHandler) CreateTask(ctx *middle.Context) {
	req := &CreateTaskReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {

		return
	}
}

func (s *taskHttpHandler) DeleteTask(ctx *middle.Context) {
}

func (s *taskHttpHandler) GetTaskList(ctx *middle.Context) {
}

func (s *taskHttpHandler) EditTask(ctx *middle.Context) {
}
