package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueFromStacks_Enqueue(t *testing.T) {
	q := QueueFromStacks{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 1, q.Dequeue())

	q.Enqueue(4)
	assert.Equal(t, 2, q.Dequeue())
	q.Enqueue(5)
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
}
