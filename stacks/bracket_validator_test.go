package stacks

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStackString_BracketsValid(t *testing.T) {
	//"{ [ ] ( ) }" should return true
	//"{ [ ( ] ) }" should return false
	//"{ [ }" should return false
	assert.True(t, bracketsValid(strings.Split("{ [ ] ( ) }", ``)))
	assert.False(t, bracketsValid(strings.Split("{ [ ( ] ) }", ``)))
	assert.False(t, bracketsValid(strings.Split("{ [ }", ``)))
	assert.False(t, bracketsValid(strings.Split("{ [ ( ] ) }", ``)))
}
