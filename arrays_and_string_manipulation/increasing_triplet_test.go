package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	assert.True(t, increasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.True(t, increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	assert.True(t, increasingTriplet([]int{20, 100, 10, 12, 5, 13}))

	assert.False(t, increasingTriplet([]int{5, 4, 3, 2, 1}))
}
