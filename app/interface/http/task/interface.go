package task

import "github.com/gin-gonic/gin"

type TaskHttpHandlerInterface interface {
  CreateTask(*gin.Context)
  DeleteTask(*gin.Context)
  GetTaskList(*gin.Context)
  EditTask(*gin.Context)
}
