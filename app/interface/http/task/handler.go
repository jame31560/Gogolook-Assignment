package task

import (
	task_usecase "task/app/usecase/task"

	"github.com/gin-gonic/gin"
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

func (s *taskHttpHandler) CreateTask(ctx *gin.Context) {
}

func (s *taskHttpHandler) DeleteTask(ctx *gin.Context) {
}

func (s *taskHttpHandler) GetTaskList(ctx *gin.Context) {
}

func (s *taskHttpHandler) EditTask(ctx *gin.Context) {
}
