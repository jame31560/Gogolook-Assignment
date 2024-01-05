package http_service

import (
	task_http_handler "task/app/interface/http/task"
	"task/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
)

type HttpService struct {
	App *gin.Engine
}

func NewHttpService(
	taskHttpHandler task_http_handler.TaskHttpHandlerInterface,
) *HttpService {
	app := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	app.GET("/swagger/*any", gin_swagger.WrapHandler(swaggerfiles.Handler))

	apiGroup := app.Group("/api")
	{
		apiGroup.POST("/tasks", taskHttpHandler.CreateTask)
		apiGroup.DELETE("/tasks", taskHttpHandler.DeleteTask)
		apiGroup.GET("/tasks", taskHttpHandler.GetTaskList)
		apiGroup.PUT("/tasks", taskHttpHandler.EditTask)
	}

	return &HttpService{
		App: app,
	}
}

type RouterInterface interface {
	Router(*gin.RouterGroup)
}
