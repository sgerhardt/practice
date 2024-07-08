package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordCloud(t *testing.T) {
	input := "After beating the eggs, Dana read the next step: Add milk and eggs, then add flour and sugar."
	cloud := wordCloud(input)
	assert.Equal(t, 2, cloud["add"])
	assert.Equal(t, 2, cloud["eggs"])
	assert.Equal(t, 1, cloud["step"])
}
