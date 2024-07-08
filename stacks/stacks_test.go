package stacks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	ss := &StackWithMax{}
	ss.push(1)
	ss.push(2)
	ss.push(3)
	assert.Equal(t, 3, ss.pop())
	assert.Equal(t, 2, ss.pop())
	assert.Equal(t, 1, ss.pop())
}

func TestStack_GetMax(t *testing.T) {
	ss := &StackWithMax{}
	ss.push(5)
	ss.push(25)
	ss.push(3)

	assert.Equal(t, 25, ss.GetMax())

	ss.pop()
	assert.Equal(t, 25, ss.GetMax())
	ss.pop()
	assert.Equal(t, 5, ss.GetMax())
	ss.pop()
	assert.Equal(t, 0, ss.GetMax())
}

func TestSimpleStack(t *testing.T) {
	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n])

		stack = stack[:n] // Pop
	}
}
