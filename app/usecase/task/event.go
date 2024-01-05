package task

type CreateTaskEvent struct {
	ID string
}

type DeleteTaskEvent struct{
  ID string
}

type EditTaskEvent struct{}

type GetTaskListEvent struct{}
