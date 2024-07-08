package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromeSubstring(t *testing.T) {

	assert.Equal(t, longestPalindromeSubstring("babad"), "bab")
}
