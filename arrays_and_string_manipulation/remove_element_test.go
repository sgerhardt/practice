package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestRemoveElement(t *testing.T) {
	elements := []int{3, 2, 2, 3}
	val := removeElement(elements, 3)
	assert.Equal(t, 2, val)

	elements2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := removeElement(elements2, 2)
	assert.Equal(t, 5, val2)
}
