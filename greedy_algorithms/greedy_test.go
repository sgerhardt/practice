package greedy_algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestProductFromThree(t *testing.T) {
	assert.Equal(t, 120, highestProductFromThreeInts([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 120, highestProductFromThreeInts([]int{6, 2, 5, 4, 3, 1}))
	assert.Equal(t, 120, highestProductFromThreeInts([]int{-10, -10, 1, 3, 2}))
	//
}
