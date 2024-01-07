package aggregate

import (
	"task/app/infra/enum"
	"time"
)

type Task struct {
	ID         string
	Name       string
	Status     enum.TaskStatusEnum
	UpdateTime time.Time
	CreateTime time.Time
}
