package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRansom(t *testing.T) {
	assert.False(t, canConstruct("a", "b"))
	assert.False(t, canConstruct("aa", "ab"))
	assert.True(t, canConstruct("aa", "aab"))
}
