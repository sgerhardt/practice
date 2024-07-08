package stacks

//Implement a queue with 2 stacks. Your queue should have an enqueue and a dequeue method, and it should be "first in first out" (FIFO).
//
//Optimize for the time cost of m calls on your queue. These can be any mix of enqueue and dequeue calls.
//
//Assume you already have a stack implementation, and it gives O(1) time push and pop.
//
// SOLUTION: Let's call our stacks inStack and outStack.
//
//For enqueue, we simply push the enqueued item onto inStack.

/*
For dequeue on an empty outStack, the oldest item is at the bottom of inStack.
So we dig to the bottom of inStack by pushing each item one-by-one onto outStack until we reach the bottom item, which we return.
After moving everything from inStack to outStack, the item that was enqueued the 2nd longest ago (after the item we just returned) is at the top of outStack,
the item enqueued 3rd longest ago is just below it, etc.
So to dequeue on a non-empty outStack, we simply return the top item from outStack.
*/

type QueueFromStacks struct {
	inStack  StackInt
	outStack StackInt
}

func (q *QueueFromStacks) Enqueue(i int) {
	q.inStack.push(i)
}

func (q *QueueFromStacks) Dequeue() int {
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			q.outStack.push(q.inStack.pop())
		}
		return q.outStack.pop()
	}
	return q.outStack.pop()
}
