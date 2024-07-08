package hackerrank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrading(t *testing.T) {
	assert.Equal(t, []int32{75, 67, 40, 33}, gradingStudents([]int32{73, 67, 38, 33}))
}
