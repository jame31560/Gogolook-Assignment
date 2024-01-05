package task

import "github.com/google/uuid"

type taskService struct{}

func NewTaskService() TaskServiceInterface {
	return &taskService{}

}

func (svc *taskService) NewTaskID() string {
  return uuid.New().String()
}
