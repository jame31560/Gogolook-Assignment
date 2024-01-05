package task

type CreateTaskEvent struct {
	ID string
}

type DeleteTaskEvent struct{
  ID string
}

type EditTaskEvent struct{
  ID string
}

type GetTaskListEvent struct{}
