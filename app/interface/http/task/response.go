package task

type CreateTaskRes struct {
	ID string `json:"id"`
}

type EditTaskRes struct {
	ID string `json:"id"`
}

type DeleteTaskRes struct {
	ID string `json:"id"`
}

type TaskRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      int8   `json:"status"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type GetTaskListRes struct {
	TaskList []*TaskRes `json:"task_list"`
}
