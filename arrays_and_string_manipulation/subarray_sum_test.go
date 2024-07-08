package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubarraySums(t *testing.T) {
	//assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
	//assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3))
	//assert.Equal(t, 4, subarraySum([]int{1, 2, 1, 2, 1}, 3))
	assert.Equal(t, 3, subarraySum([]int{1, -1, 0}, 0))
}
