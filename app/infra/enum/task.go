package enum

type TaskStatusEnum int8

const (
	TaskStatusIncomplete TaskStatusEnum = iota + 1
	TaskStatusCompleted
)

func GetAllTaskStatusIntList() []int8 {
	return []int8{
		int8(TaskStatusIncomplete),
		int8(TaskStatusCompleted),
	}
}
