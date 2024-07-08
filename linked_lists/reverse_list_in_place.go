package linked_lists

func reverse(headNode *LinkedListNode) *LinkedListNode {
	currentNode := headNode
	var previousNode *LinkedListNode
	var nextNode *LinkedListNode

	for currentNode != nil {
		// copy a pointer to the next element
		// before we overwrite currentNode.next
		nextNode = currentNode.next

		// reverse the 'next' pointer
		currentNode.next = previousNode

		// step forward in the list
		previousNode = currentNode
		currentNode = nextNode
	}

	//return the previous node since current node is nil after exiting loop
	return previousNode
}
