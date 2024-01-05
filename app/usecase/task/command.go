package task

type CreateTaskCmd struct {
	Name string
}

type DeleteTaskCmd struct {
	ID string
}

type EditTaskCmd struct{}

type GetTaskListCmd struct{}
