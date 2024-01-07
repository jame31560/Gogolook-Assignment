package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTaskStatusEnum(t *testing.T) {
	caseList := []struct {
		input    int8
		hasErr   bool
		expected TaskStatusEnum
	}{
		{input: int8(TaskStatusNone), hasErr: true},
		{input: int8(TaskStatusIncomplete), expected: TaskStatusIncomplete},
		{input: int8(TaskStatusEnd), hasErr: true},
		{input: -1, hasErr: true},
	}

	for _, testCase := range caseList {
		result, err := ToTaskStatusEnum(testCase.input)

		if testCase.hasErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.EqualValues(t, testCase.expected, result)
		}
	}
}
