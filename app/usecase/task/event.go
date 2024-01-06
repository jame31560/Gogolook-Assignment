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

type TaskDto struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int8   `json:"status"`
}

type GetTaskListEvent struct {
	TaskList []*TaskDto `json:"task_list"`
}
