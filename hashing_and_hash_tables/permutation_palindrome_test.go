package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromePossible(t *testing.T) {
	assert.True(t, palindromePossible("civic"))
	assert.True(t, palindromePossible("ivicc"))

	assert.False(t, palindromePossible("civil"))
	assert.False(t, palindromePossible("livci"))
}
