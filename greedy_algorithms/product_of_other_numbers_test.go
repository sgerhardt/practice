package greedy_algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfOtherNumbers(t *testing.T) {
	ints := []int{1, 7, 3, 4}
	assert.Equal(t, []int{84, 12, 28, 21}, getProductsOfAllIntsExceptAtIndexBruteForce(ints))
	assert.Equal(t, []int{84, 12, 28, 21}, getProductsOfAllIntsExceptAtIndex(ints))

	ints = []int{0, 7, 3, 4}
	assert.Equal(t, []int{84, 0, 0, 0}, getProductsOfAllIntsExceptAtIndexBruteForce(ints))
	assert.Equal(t, []int{84, 0, 0, 0}, getProductsOfAllIntsExceptAtIndex(ints))
}
