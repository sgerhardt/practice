package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, input)

	input2 := []int{2, 0, 2, 1, 1, 0}
	sortColorsDutch(input2)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, input)
}
