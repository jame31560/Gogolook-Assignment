package task

import "task/app/interface/http/middle"

type TaskHttpHandlerInterface interface {
	CreateTask(*middle.Context)
	DeleteTask(*middle.Context)
	GetTaskList(*middle.Context)
	EditTask(*middle.Context)
}
