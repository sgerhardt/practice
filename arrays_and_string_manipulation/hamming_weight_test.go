package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHamingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(0000000000000000000000000000001))
}
