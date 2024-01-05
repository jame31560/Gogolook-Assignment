package task

type CreateTaskCmd struct {
	Name string
}

type DeleteTaskCmd struct {
	ID string
}

type EditTaskCmd struct {
	ID     string
	Name   string
	Status int8
}

type GetTaskListCmd struct{}
