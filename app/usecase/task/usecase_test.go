package task

import (
	"context"
	aggregate "task/app/domain/model/aggreate"
	task_service "task/app/domain/service/task"
	task_mock "task/app/infra/database/mock/task"
	"task/app/pkg/status"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	repoMock := task_mock.NewTaskRepoMock()
	svcMock := task_service.NewTaskServiceMock()
	usecase := &taskUsecase{
		taskRepo:    repoMock,
		taskService: svcMock,
	}

	caseList := []struct {
		name                 string
		cmd                  *CreateTaskCmd
		mockCheckName        bool
		mockNewUuid          string
		mockCreateRepoHasErr bool
		expectedEvent        *CreateTaskEvent
		hasError             bool
	}{
		{
			name:          "Success",
			cmd:           &CreateTaskCmd{Name: "Task"},
			mockCheckName: true,
			mockNewUuid:   "taskID",
			expectedEvent: &CreateTaskEvent{
				ID: "taskID",
			},
			hasError: false,
		},
		{
			name:          "Empty name",
			cmd:           &CreateTaskCmd{Name: ""},
			mockCheckName: false,
			expectedEvent: nil,
			hasError:      true,
		},
		{
			name:                 "Repository error",
			cmd:                  &CreateTaskCmd{Name: "Task"},
			mockCheckName:        true,
			mockCreateRepoHasErr: true,
			expectedEvent:        nil,
			hasError:             true,
		},
	}

	for _, testCase := range caseList {
		t.Run(testCase.name, func(t *testing.T) {
			svcMock.NewTaskIDFunc = func() string {
				return testCase.mockNewUuid
			}
			svcMock.CheckNameFunc = func(name string) bool {
				return testCase.mockCheckName
			}
			repoMock.CreateTaskFunc = func(task *aggregate.Task) error {
				if testCase.mockCreateRepoHasErr {
					return status.ErrorStatus
				}
				return nil
			}

			event, err := usecase.CreateTask(context.Background(), testCase.cmd)

			assert.EqualValues(t, testCase.expectedEvent, event)
			if testCase.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	repoMock := task_mock.NewTaskRepoMock()
	svcMock := task_service.NewTaskServiceMock()
	usecase := &taskUsecase{
		taskRepo:    repoMock,
		taskService: svcMock,
	}

	caseList := []struct {
		name                 string
		cmd                  *DeleteTaskCmd
		mockDeleteRepoHasErr bool
		expectedEvent        *DeleteTaskEvent
		hasError             bool
	}{
		{
			name: "Success",
			cmd:  &DeleteTaskCmd{ID: "taskID"},
			expectedEvent: &DeleteTaskEvent{
				ID: "taskID",
			},
		},
		{
			name:                 "Repository error",
			cmd:                  &DeleteTaskCmd{ID: "taskID"},
			mockDeleteRepoHasErr: true,
			hasError:             true,
		},
	}

	for _, testCase := range caseList {
		t.Run(testCase.name, func(t *testing.T) {
			repoMock.DeleteTaskFunc = func(ID string) error {
				if testCase.mockDeleteRepoHasErr {
					return status.ErrorStatus
				}
				return nil
			}

			event, err := usecase.DeleteTask(context.Background(), testCase.cmd)

			assert.EqualValues(t, testCase.expectedEvent, event)
			if testCase.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestEditTask(t *testing.T) {
	repoMock := task_mock.NewTaskRepoMock()
	svcMock := task_service.NewTaskServiceMock()
	usecase := &taskUsecase{
		taskRepo:    repoMock,
		taskService: svcMock,
	}

	caseList := []struct {
		name                 string
		cmd                  *EditTaskCmd
		mockCheckName        bool
		mockUpdateRepoHasErr bool
		expectedEvent        *EditTaskEvent
		hasError             bool
	}{
		{
			name:          "Success",
			cmd:           &EditTaskCmd{ID: "taskID", Name: "NewTask", Status: 1},
			mockCheckName: true,
			expectedEvent: &EditTaskEvent{ID: "taskID"},
			hasError:      false,
		},
		{
			name:          "Empty name",
			cmd:           &EditTaskCmd{ID: "taskID", Name: "", Status: 1},
			mockCheckName: false,
			expectedEvent: nil,
			hasError:      true,
		},
		{
			name:          "TaskStatus incorrect",
			cmd:           &EditTaskCmd{ID: "taskID", Name: "NewTask", Status: 99},
			mockCheckName: true,
			expectedEvent: nil,
			hasError:      true,
		},
		{
			name:                 "Repository update error",
			cmd:                  &EditTaskCmd{ID: "taskID", Name: "NewTask", Status: 1},
			mockCheckName:        true,
			mockUpdateRepoHasErr: true,
			expectedEvent:        nil,
			hasError:             true,
		},
	}

	for _, testCase := range caseList {
		t.Run(testCase.name, func(t *testing.T) {
			svcMock.CheckNameFunc = func(name string) bool {
				return testCase.mockCheckName
			}
			repoMock.UpdateTaskByIDFunc = func(ID string, task *aggregate.Task) error {
				if testCase.mockUpdateRepoHasErr {
					return status.ErrorStatus
				}
				return nil
			}

			event, err := usecase.EditTask(context.Background(), testCase.cmd)

			assert.EqualValues(t, testCase.expectedEvent, event)
			if testCase.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
