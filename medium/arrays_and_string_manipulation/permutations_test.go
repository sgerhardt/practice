package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	expected := [][]int{{0, 1}, {1, 0}}
	//assert.Equal(t, expected, permute([]int{5, 4, 6, 2}))
	assert.Equal(t, expected, permute([]int{0, 1}))

	//expected = [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	//assert.Equal(t, expected, permute([]int{1, 2, 3}))

}
