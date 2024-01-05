package task

import "github.com/google/uuid"

type taskService struct{}

func NewTaskService() TaskServiceInterface {
	return &taskService{}

}

func (svc *taskService) CheckName(name string) bool {
  if name == "" {
    return false
  }
  return true
}

func (svc *taskService) NewTaskID() string {
  return uuid.New().String()
}
