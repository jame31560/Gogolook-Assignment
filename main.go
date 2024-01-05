package main

import (
	task_service "task/app/domain/service/task"
	task_repo "task/app/infra/database/in_memory/task"
	http_service "task/app/interface/http"
	"task/app/interface/http/middle"
	task_http_handler "task/app/interface/http/task"
	task_usecase "task/app/usecase/task"
)

func main() {
	taskRepo := task_repo.NewTaskRepo()

  taskService := task_service.NewTaskService()

	taskUsecase := task_usecase.NewTaskUsecase(taskRepo, taskService)

	taskHttpHandler := task_http_handler.NewTaskHttpHandler(taskUsecase)

  middleware := middle.NewMiddleware()

	app := http_service.NewHttpService(middleware, taskHttpHandler)

	app.App.Run(":8080")
}
