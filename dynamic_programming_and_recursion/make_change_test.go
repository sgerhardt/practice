package dynamic_programming_and_recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePossibilitiesTopDown(t *testing.T) {
	assert.Equal(t, 4, changePossibilitiesTopDownRecursive(4, []int{1, 2, 3}, 0))

	assert.Equal(t, 4, changePossibilitiesBottomUpIterative(4, []int{1, 2, 3}))

	assert.Equal(t, 3, changePossibilitiesBottomUpIterative(5, []int{1, 2}))

	assert.Equal(t, 5, changePossibilitiesBottomUpIterative(5, []int{1, 2, 3}))
}
