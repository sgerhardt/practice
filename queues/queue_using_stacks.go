package queues

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		// our item is at bottom of in stack
		for len(this.inStack) > 0 {
			// move all in stack items to out stack
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1:][0])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	popVal := this.outStack[len(this.outStack)-1:][0]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return popVal
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return 0
	}

	if len(this.outStack) == 0 {
		// our item is at bottom of in stack
		return this.inStack[0:1][0]
	}
	return this.outStack[len(this.outStack)-1:][0]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
