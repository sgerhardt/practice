package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	input := []int{1, 2, 2}
	assert.Equal(t, 2, removeDuplicates(input))
	assert.Equal(t, []int{1, 2, 2}, input)
	input2 := []int{0, 0, 0, 0, 0}
	assert.Equal(t, 1, removeDuplicates(input2))
	assert.Equal(t, []int{0, 0, 0, 0, 0}, input2)
}
