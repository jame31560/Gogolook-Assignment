package task

import (
	"context"
	"task/app/domain/model/aggregate"
	task_service "task/app/domain/service/task"
	task_mock "task/app/infra/database/mock/task"
	"task/app/infra/enum"
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

func TestGetTaskList(t *testing.T) {
	repoMock := task_mock.NewTaskRepoMock()
	svcMock := task_service.NewTaskServiceMock()
	usecase := &taskUsecase{
		taskRepo:    repoMock,
		taskService: svcMock,
	}

	caseList := []struct {
		name                string
		cmd                 *GetTaskListCmd
		mockGetTaskByID     *aggregate.Task
		mockGetTaskByIDErr  status.Status
		mockQueryTaskList   []*aggregate.Task
		mockQueryErr        status.Status
		expectedTaskListLen int
		hasError            bool
	}{
		{
			name: "Get by ID - Success",
			cmd:  &GetTaskListCmd{ID: "taskID"},
			mockGetTaskByID: &aggregate.Task{
				ID:     "taskID",
				Name:   "Task",
				Status: enum.TaskStatusCompleted,
			},
			expectedTaskListLen: 1,
			hasError:            false,
		},
		{
			name:               "Get by ID - Repository error",
			cmd:                &GetTaskListCmd{ID: "taskID"},
			mockGetTaskByIDErr: status.ErrorStatus,
			hasError:           true,
		},
		{
			name: "Query - Success",
			cmd:  &GetTaskListCmd{Name: "Task", Status: []int8{int8(enum.TaskStatusIncomplete)}},
			mockQueryTaskList: []*aggregate.Task{
				{
					ID:     "taskID1",
					Name:   "TaskName1",
					Status: enum.TaskStatusIncomplete,
				},
				{
					ID:     "taskID2",
					Name:   "TaskName2",
					Status: enum.TaskStatusIncomplete,
				},
			},
			expectedTaskListLen: 2,
			hasError:            false,
		},
		{
			name:                "Query - Repository error",
			cmd:                 &GetTaskListCmd{Name: "TaskName", Status: []int8{int8(enum.TaskStatusIncomplete)}},
			mockQueryErr:        status.ErrorStatus,
			expectedTaskListLen: 0,
			hasError:            true,
		},
	}

	for _, testCase := range caseList {
		t.Run(testCase.name, func(t *testing.T) {
			repoMock.GetTaskByIDFunc = func(ID string) (*aggregate.Task, error) {
				return testCase.mockGetTaskByID, testCase.mockGetTaskByIDErr
			}
			repoMock.QueryTaskListFunc = func(name string, statusList []int8) ([]*aggregate.Task, error) {
				return testCase.mockQueryTaskList, testCase.mockQueryErr
			}

			event, err := usecase.GetTaskList(context.Background(), testCase.cmd)

			if testCase.hasError {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, testCase.expectedTaskListLen, len(event.TaskList))
			assert.NoError(t, err)
		})
	}
}
