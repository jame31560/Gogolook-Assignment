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
	cmd := &task_usecase.CreateTaskCmd{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		errStatus := status.CreateError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect")
		ctx.ErrorRes(errStatus)
		return
	}

	event, err := s.taskUsecase.CreateTask(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	ctx.Response(status.CreateSuccess, event)
}

func (s *taskHttpHandler) DeleteTask(ctx *middle.Context) {
	ID, ok := ctx.Params.Get("task_id")
	if !ok || ID == "" {
		errStatus := status.DeleteError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect")
		ctx.ErrorRes(errStatus)
		return
	}

	cmd := &task_usecase.DeleteTaskCmd{
		ID: ID,
	}

	event, err := s.taskUsecase.DeleteTask(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	ctx.Response(status.GeneralSuccess, event)
}

func (s *taskHttpHandler) GetTaskList(ctx *middle.Context) {
	cmd := &task_usecase.GetTaskListCmd{}
	if err := ctx.BindQuery(cmd); err != nil {
		ctx.ErrorRes(status.QueryError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect"))
		return
	}

	event, err := s.taskUsecase.GetTaskList(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	ctx.Response(status.GeneralSuccess, event)
}

func (s *taskHttpHandler) EditTask(ctx *middle.Context) {
	badRequestStatus := status.UpdateError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect")

	cmd := &task_usecase.EditTaskCmd{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.ErrorRes(badRequestStatus)
		return
	}

	ID, ok := ctx.Params.Get("task_id")
	if !ok || ID == "" {
		ctx.ErrorRes(badRequestStatus)
		return
	}
	cmd.ID = ID

	event, err := s.taskUsecase.EditTask(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
	}

	ctx.Response(status.GeneralSuccess, event)
}
