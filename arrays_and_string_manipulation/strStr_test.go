package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrstr(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	assert.Equal(t, 2, strStr(haystack, needle))
	haystack = "aaaaa"
	needle = "bba"
	assert.Equal(t, -1, strStr(haystack, needle))
}
