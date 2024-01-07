package task

type taskServiceMock struct {
	CheckNameFunc func(name string) bool
	NewTaskIDFunc func() string
}

func NewTaskServiceMock() *taskServiceMock {
	return &taskServiceMock{}
}

func (svc *taskServiceMock) CheckName(name string) bool {
	if svc.CheckNameFunc != nil {
		return svc.CheckNameFunc(name)
	}
	return true
}

func (svc *taskServiceMock) NewTaskID() string {
	if svc.NewTaskIDFunc != nil {
		return svc.NewTaskIDFunc()
	}
	return "7780fc95-7a0a-45ef-a985-21c3e744c1d7"
}
