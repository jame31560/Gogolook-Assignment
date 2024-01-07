package task

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCheckName(t *testing.T) {
	taskSvc := NewTaskService()

	// 空白名稱
	result := taskSvc.CheckName("")
	assert.False(t, result)

	// 正確名稱
	result = taskSvc.CheckName("Name")
	assert.True(t, result)
}

func TestNewTaskID(t *testing.T) {
	taskSvc := NewTaskService()

	taskID := taskSvc.NewTaskID()
	_, err := uuid.Parse(taskID)
	assert.NoError(t, err)
}
