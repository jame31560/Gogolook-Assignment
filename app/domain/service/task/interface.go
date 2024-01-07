package task

type TaskServiceInterface interface {
	NewTaskID() string
	CheckName(string) bool
}
