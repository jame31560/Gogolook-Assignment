package task

import "time"

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
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     int8      `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

type GetTaskListEvent struct {
	TaskList []*TaskDto `json:"task_list"`
}
