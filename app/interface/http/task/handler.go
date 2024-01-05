package task

import (
	"net/http"
	"task/app/interface/http/middle"
	"task/app/pkg/status"
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
		errStatus := status.CreateError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect")
		ctx.ErrorRes(errStatus)
		return
	}

	cmd := &task_usecase.CreateTaskCmd{
		Name: req.Name,
	}

	event, err := s.taskUsecase.CreateTask(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	res := &CreateTaskRes{
		ID: event.ID,
	}

	ctx.Response(status.CreateSuccess, res)
}

func (s *taskHttpHandler) DeleteTask(ctx *middle.Context) {
	ID, ok := ctx.Params.Get("task_id")
	if !ok || ID == "" {
		errStatus := status.DeleteError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect")
		ctx.ErrorRes(errStatus)
		return
	}

	req := &DeleteTaskReq{
		ID: ID,
	}

	cmd := &task_usecase.DeleteTaskCmd{
		ID: req.ID,
	}

	event, err := s.taskUsecase.DeleteTask(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	res := &DeleteTaskRes{
		ID: event.ID,
	}

	ctx.Response(status.CreateSuccess, res)
}

func (s *taskHttpHandler) GetTaskList(ctx *middle.Context) {
}

func (s *taskHttpHandler) EditTask(ctx *middle.Context) {
}
