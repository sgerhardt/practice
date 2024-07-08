package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum(input))

	input2 := []int{-1, 0, 1, 1, 2, -1, -4}
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum(input2))
}
