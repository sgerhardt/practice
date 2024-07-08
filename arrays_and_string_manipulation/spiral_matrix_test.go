package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	expected := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5}}
	assert.Equal(t, expected, generateMatrix(3))
}
