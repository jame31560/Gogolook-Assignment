package task

type CreateTaskEvent struct {
	ID string `json:"id"`
}

type DeleteTaskEvent struct {
	ID string `json:"id"`
}

type EditTaskEvent struct {
	ID string `json:"id"`
}

type GetTaskListEvent struct{}
