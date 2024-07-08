package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	//assert.Equal(t, expected, mergeIntervals(intervals))
	//
	//intervals2 := [][]int{{1, 4}, {2, 3}}
	//expected2 := [][]int{{1, 4}}
	//assert.Equal(t, expected2, mergeIntervals(intervals2))

	intervals3 := [][]int{{1, 4}, {0, 2}, {3, 5}}
	expected3 := [][]int{{0, 5}}
	assert.Equal(t, expected3, mergeIntervals(intervals3))
}
