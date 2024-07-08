package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrRoot(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 6, mySqrt(36))
}
