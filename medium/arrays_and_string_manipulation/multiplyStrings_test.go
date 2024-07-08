package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "5535", multiply("123", "45"))
}
