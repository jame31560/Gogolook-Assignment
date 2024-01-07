package task

import (
	task_service "task/app/domain/service/task"
	task_mock "task/app/infra/database/mock/task"
)

func newUsecaseForTest() TaskUsecaseInterface {
	repoMock := task_mock.NewTaskRepoMock()
	taskSvc := task_service.NewTaskService()
	usecase := NewTaskUsecase(repoMock, taskSvc)
	return usecase
}
