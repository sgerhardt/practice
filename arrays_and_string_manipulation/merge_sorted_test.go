package arrays_and_string_manipulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSorted(t *testing.T) {
	merged := mergeSortedNlogN([]int{3, 6, 8}, []int{1, 5, 9})
	expected := []int{1, 3, 5, 6, 8, 9}
	assert.Equal(t, expected, merged)

	merged = mergeSortedNlogN([]int{3, 4, 6, 10, 11, 15}, []int{1, 5, 8, 12, 14, 19})
	expected = []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19}
	assert.Equal(t, expected, merged)

	//merged = mergeSortedN([]int{1, 5, 7}, []int{2, 4, 6})
	//assert.Equal(t, []int{1, 2, 4, 5, 6, 7}, merged)

	merged = mergeSortedN([]int{3, 4, 6, 10, 11, 15}, []int{1, 5, 8, 12, 14, 19})
	assert.Equal(t, []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19}, merged)

}
