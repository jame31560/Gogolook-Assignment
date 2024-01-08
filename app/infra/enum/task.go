package enum

import (
	"net/http"
	"task/app/pkg/status"
)

type TaskStatusEnum int8

const (
	TaskStatusIncomplete TaskStatusEnum = iota
	TaskStatusCompleted
	TaskStatusEnd
)

func GetAllTaskStatusIntList() []int8 {
	return []int8{
		int8(TaskStatusIncomplete),
		int8(TaskStatusCompleted),
	}
}

func ToTaskStatusEnum(i int8) (TaskStatusEnum, error) {
	err := status.ErrorStatus.WithHttpCode(http.StatusBadRequest).WithMsg("Unknown status enum")
	if i < 0 {
		return 0, err
	}
	if i >= int8(TaskStatusEnd) {
		return 0, err
	}
	return TaskStatusEnum(i), nil
}
