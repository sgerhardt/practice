package math

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleArithmetic(t *testing.T) {
	val, err := simpleArithmetic("2 + 2")
	require.NoError(t, err)
	assert.Equal(t, 4, val)

	val, err = simpleArithmetic("2 + 2 - 3")
	require.NoError(t, err)
	assert.Equal(t, 1, val)
}
