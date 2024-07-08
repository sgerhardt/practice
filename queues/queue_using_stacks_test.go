package queues

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {

	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	assert.Equal(t, 1, myQueue.Peek())
	assert.Equal(t, 1, myQueue.Pop())
	assert.False(t, myQueue.Empty())

}
