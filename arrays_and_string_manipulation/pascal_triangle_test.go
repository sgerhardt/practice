package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, generate(1))
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5))
}

func TestPascalTriangleAtIndex(t *testing.T) {
	assert.Equal(t, []int{1}, getRow(0))
	assert.Equal(t, []int{1, 2, 1}, getRow(2))
}
