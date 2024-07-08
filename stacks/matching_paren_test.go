package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchingParen(t *testing.T) {
	input := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."
	assert.Equal(t, 79, matchingParen(input, 10))
}
