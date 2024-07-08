package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidAnagram(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
	assert.False(t, isAnagram("rat", "car"))
	assert.False(t, isAnagram("ab", "a"))
}
