package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverlapIntervals(t *testing.T) {
	input := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	assert.Equal(t, 1, eraseOverlapIntervals(input))

	input = [][]int{{1, 2}, {1, 2}, {1, 2}}
	assert.Equal(t, 2, eraseOverlapIntervals(input))

	input = [][]int{{1, 2}, {2, 3}}
	assert.Equal(t, 0, eraseOverlapIntervals(input))

	input = [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	assert.Equal(t, 2, eraseOverlapIntervals(input))

	input = [][]int{{0, 1}, {3, 4}, {1, 2}}
	assert.Equal(t, 0, eraseOverlapIntervals(input))
}
