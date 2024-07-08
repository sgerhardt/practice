package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	ints := []int{0, 1, 0, 3, 12}
	moveZeroes(ints)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, ints)

	ints2 := []int{0}
	moveZeroes(ints2)
	assert.Equal(t, []int{0}, ints2)

	ints3 := []int{0, 0, 1}
	moveZeroes(ints3)
	assert.Equal(t, []int{1, 0, 0}, ints3)
}
