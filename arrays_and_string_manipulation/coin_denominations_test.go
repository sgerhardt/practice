package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDenoms(t *testing.T) {
	assert.Equal(t, []int{25, 1, 1}, denominations([]int{25, 3, 1}, 27))
	assert.Equal(t, []int{8, 1, 1, 1}, denominations([]int{1, 5, 6, 8}, 11))
}
