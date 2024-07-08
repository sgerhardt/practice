package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	input1 := []int{1, 2, 2, 1}
	input2 := []int{2, 2}
	assert.Equal(t, []int{2, 2}, intersect(input1, input2))

	input1 = []int{4, 9, 5}
	input2 = []int{9, 4, 9, 8, 4}
	assert.Contains(t, [][]int{{4, 9}, []int{9, 4}}, intersect(input1, input2))

	input1 = []int{2, 1}
	input2 = []int{1, 1}
	assert.Equal(t, []int{1}, intersect(input1, input2))

	input1 = []int{2, 1}
	input2 = []int{1, 1}
	assert.Equal(t, []int{1}, intersect(input1, input2))

	input1 = []int{2, 2}
	input2 = []int{2, 2}
	assert.Equal(t, []int{2, 2}, intersect(input1, input2))
}
