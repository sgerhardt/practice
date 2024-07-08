package linked_lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedLists(t *testing.T) {
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

	DeleteNode(b)

	for a != nil {
		fmt.Println(a.value)
		a = a.next
	}
}

func TestCycle(t *testing.T) {
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

	assert.False(t, ContainsCycle(a))
	assert.False(t, ContainsCycleUsingRunners(a))

	c.next = a

	assert.True(t, ContainsCycle(a))
	assert.True(t, ContainsCycleUsingRunners(a))
}
