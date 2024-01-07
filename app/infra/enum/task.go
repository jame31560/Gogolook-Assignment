package enum

import (
	"net/http"
	"task/app/pkg/status"
)

type TaskStatusEnum int8

const (
	TaskStatusNone TaskStatusEnum = iota
	TaskStatusIncomplete
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
	if i <= int8(TaskStatusNone) {
		return TaskStatusNone, err
	}
	if i >= int8(TaskStatusEnd) {
		return TaskStatusNone, err
	}
	return TaskStatusEnum(i), nil
}
