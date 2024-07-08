package memoization

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 0, fib(0, nil))
	assert.Equal(t, 1, fib(1, nil))
	assert.Equal(t, 1, fib(2, nil))
	assert.Equal(t, 2, fib(3, nil))
	assert.Equal(t, 3, fib(4, nil))

	assert.Equal(t, 6765, fib(20, nil))
	assert.Equal(t, 75025, fib(25, nil))
	assert.Equal(t, 832040, fib(30, nil))
	assert.Equal(t, 102334155, fib(40, nil))

	// Without memoization, you're gonna be waiting a looooooooong time!!!
	assert.Equal(t, 12586269025, fib(50, nil))
	assert.Equal(t, 2111485077978050, fib(75, nil))
}
