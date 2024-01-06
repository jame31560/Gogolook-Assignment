package task

type CreateTaskCmd struct {
	Name string `json:"name"`
}

type DeleteTaskCmd struct {
	ID string
}

type EditTaskCmd struct {
  ID     string `json:"-"`
	Name   string `json:"name"`
	Status int8   `json:"status"`
}

type GetTaskListCmd struct {
	ID     string
	Name   string
	Status []int8
}
