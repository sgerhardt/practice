package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdersFirstComeFirstServe(t *testing.T) {
	takeOut := []int{1, 3, 5}
	dineIn := []int{2, 4, 6}
	served := []int{1, 2, 4, 6, 5, 3}
	isFCFS := ordersFirstComeFirstServe(takeOut, dineIn, served)
	assert.False(t, isFCFS)
	isFCFS = ordersFirstComeFirstServe2(takeOut, dineIn, served)
	assert.False(t, isFCFS)

	takeOut = []int{17, 8, 24}
	dineIn = []int{12, 19, 2}
	served = []int{17, 8, 12, 19, 24, 2}
	isFCFS = ordersFirstComeFirstServe(takeOut, dineIn, served)
	assert.True(t, isFCFS)
	isFCFS = ordersFirstComeFirstServe2(takeOut, dineIn, served)
	assert.True(t, isFCFS)
}
