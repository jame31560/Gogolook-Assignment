package http_service

import (
	"task/app/interface/http/middle"
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
	middleware middle.Handler,
	taskHttpHandler task_http_handler.TaskHttpHandlerInterface,
) *HttpService {
	app := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	app.GET("/swagger/*any", gin_swagger.WrapHandler(swaggerfiles.Handler))

	apiGroup := app.Group("/api")
	{
		apiGroup.POST("/tasks", middleware.HandleFunc(taskHttpHandler.CreateTask))
    apiGroup.DELETE("/tasks/:task_id", middleware.HandleFunc(taskHttpHandler.DeleteTask))
		apiGroup.GET("/tasks", middleware.HandleFunc(taskHttpHandler.GetTaskList))
		apiGroup.PUT("/tasks/:task_id", middleware.HandleFunc(taskHttpHandler.EditTask))
	}

	return &HttpService{
		App: app,
	}
}

type RouterInterface interface {
	Router(*gin.RouterGroup)
}
