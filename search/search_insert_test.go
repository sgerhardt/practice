package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchBinary(t *testing.T) {

	assert.Equal(t, 0, searchInsert([]int{1, 3, 5}, 1))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 1, searchInsert([]int{1, 3}, 2))

	assert.Equal(t, 1, searchInsert([]int{1, 3, 5}, 2))

	assert.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
	assert.Equal(t, 1, searchInsert([]int{1}, 2))
	assert.Equal(t, 3, searchInsert([]int{3, 5, 7, 9, 10}, 8))
}
