package greedy_algorithms

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMaxProfit(t *testing.T) {
	stockPrices := []int{10, 7, 5, 8, 11, 9}
	p, err := getMaxProfit(stockPrices)
	require.NoError(t, err)
	assert.Equal(t, 6, p)

	// price goes down all day
	stockPrices = []int{12, 7, 5, 2, 1}
	p, err = getMaxProfit(stockPrices)
	require.NoError(t, err)
	assert.Equal(t, -1, p)
}
