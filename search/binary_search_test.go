package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 0, search([]int{1, 2, 3}, 1))
	assert.Equal(t, 1, search([]int{1, 2, 3}, 2))
	assert.Equal(t, 2, search([]int{1, 2, 3}, 3))
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
	assert.Equal(t, 1, search([]int{-1, 0, 3, 5, 9, 12}, 0))
}
