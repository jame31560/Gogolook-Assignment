package enum

type TaskStatusEnum int8

const (
	TaskStatusIncomplete TaskStatusEnum = iota + 1
  TaskStatusCompleted
)
