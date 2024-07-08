package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordPattern(t *testing.T) {
	assert.True(t, wordPattern("abba", "dog cat cat dog"))
	assert.False(t, wordPattern("abba", "dog cat cat fish"))
	assert.False(t, wordPattern("abba", "dog dog dog dog"))
}
