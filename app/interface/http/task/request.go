package task

type CreateTaskReq struct {
	Name string `json:"name"`
}

type EditTaskReq struct {
	ID     string
	Name   string `json:"name"`
	Status int8   `json:"status"`
}

type DeleteTaskReq struct {
	ID string
}

type GetTaskListReq struct {
	ID     string
	Name   string
	Status []int8
}
