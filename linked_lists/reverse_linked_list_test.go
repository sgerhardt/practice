package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	a := &LinkedListNode{
		value: "A",
	}
	b := &LinkedListNode{
		value: "B",
	}
	c := &LinkedListNode{
		value: "C",
	}
	a.next = b
	b.next = c

	assert.Equal(t, c, reverse(a))
	assert.Equal(t, c.next, b)
	assert.Equal(t, c.next.next, a)
}
