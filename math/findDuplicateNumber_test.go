package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 2, findDuplicateWithMap([]int{1, 2, 3, 4, 2}))
	assert.Equal(t, 2, findDuplicateWithMath([]int{1, 2, 3, 4, 2}))

	// see triangluar series
	assert.Equal(t, 3, findDuplicateWithMap([]int{1, 2, 3, 3, 4, 5, 6}))
	assert.Equal(t, 3, findDuplicateWithMath([]int{1, 2, 3, 3, 4, 5, 6}))

}
