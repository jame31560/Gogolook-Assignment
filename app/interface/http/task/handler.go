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

// @Summary	Create a task
// @Schemes
// @Description	Create a task.
// @Param	data body task_usecase.CreateTaskCmd true "Create Task"
// @Tags Task
// @Accept json
// @Produce json
// @Success	201	{object} task_usecase.CreateTaskEvent
// @Router /tasks [post]
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
		return
	}

	ctx.Response(status.CreateSuccess, event)
}

// @Summary	Delete a task
// @Schemes
// @Description	Delete a task.
// @Param task_id path string true "Task ID"
// @Tags Task
// @Accept json
// @Produce json
// @Success	200 {object} task_usecase.DeleteTaskEvent
// @Router /tasks/{task_id} [delete]
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
		return
	}

	ctx.Response(status.GeneralSuccess, event)
}

// @Summary	Get task List
// @Schemes
// @Description	Get tasks.
// @Description	Theree has two mod.
// @Description	- ID: if use ID, other field will be ignore, and only full match ID's task will be response.
// @Description	- Search: if ID is empty, will search by other field.
// @Param id query string false "id" example(string)
// @Param name query string false "name" example(string)
// @Param status query []int false "status enums" Enums(1, 2)
// @Tags Task
// @Accept json
// @Produce json
// @Success	200	{object} task_usecase.GetTaskListEvent
// @Router /tasks [get]
func (s *taskHttpHandler) GetTaskList(ctx *middle.Context) {
	cmd := &task_usecase.GetTaskListCmd{}
	if err := ctx.BindQuery(cmd); err != nil {
		ctx.ErrorRes(status.QueryError.WithHttpCode(http.StatusBadRequest).WithMsg("Request format incorect"))
		return
	}

	event, err := s.taskUsecase.GetTaskList(ctx, cmd)
	if err != nil {
		ctx.ErrorRes(err)
		return
	}

	ctx.Response(status.GeneralSuccess, event)
}

// @Summary	Edit task List
// @Schemes
// @Description	Edit tasks.
// @Param task_id path string true "Task ID"
// @Param	data body task_usecase.EditTaskCmd true "Edit Task"
// @Tags Task
// @Accept json
// @Produce json
// @Success	200	{object} task_usecase.EditTaskEvent
// @Router /tasks/{task_id} [put]
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
		return
	}

	ctx.Response(status.GeneralSuccess, event)
}
