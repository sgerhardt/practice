package linked_lists

type LinkedListNode struct {
	value string
	next  *LinkedListNode
}

// DeleteNode deletes a node from a singly-linked list, given only a variable pointing to that node.
// be careful—there are some potential problems with this implementation:
// First, it doesn't work for deleting the last node in the list.
// Second, this technique can cause some unexpected side-effects. For example, let's say we call:
//
//	a = LinkedListNode(3)
//
// b = LinkedListNode(8)
// c = LinkedListNode(2)
//
// a.next = b
// b.next = c
//
// delete_node(b)
//
// Python 3.6
// There are two potential side-effects:
//
// Any references to the input node have now effectively been reassigned to its next node.
// In our example, we "deleted" the node assigned to the variable b, but in actuality we just gave it a new value (2) and a new next!
// If we had another pointer to b somewhere else in our code and we were assuming it still had its old value (8), that could cause bugs.
// If there are pointers to the input node's original next node, those pointers now point to a "dangling" node
// (a node that's no longer reachable by walking down our list). In our example above, c is now dangling.
// If we changed c, we'd never encounter that new value by walking down our list from the head to the tail.
func DeleteNode(node *LinkedListNode) {
	//node = node.next
	node.value = node.next.value
	node.next = nil
}

// ContainsCycle returns whether a linked list contains a cycle.
// A cycle occurs when a node’s next points back to a previous node in the list.
// The linked list is no longer linear with a beginning and end—instead, it cycles through a loop of nodes.
// Runtime is O(n)
// Space is O(n)
func ContainsCycle(node *LinkedListNode) bool {
	nodes := map[*LinkedListNode]struct{}{}
	for node != nil {
		// see if node is in set
		if _, ok := nodes[node]; ok {
			// must be a cycle...
			return true
		} else {
			nodes[node] = struct{}{}
		}
		node = node.next
	}
	return false
}

// ContainsCycleUsingRunners returns whether a linked list contains a cycle.
// A cycle occurs when a node’s next points back to a previous node in the list.
// The linked list is no longer linear with a beginning and end—instead, it cycles through a loop of nodes.
// Runtime is O(n)
// Space is O(1)
func ContainsCycleUsingRunners(node *LinkedListNode) bool {
	slowRunner := node
	fastRunner := node

	for fastRunner != nil && fastRunner.next != nil {
		slowRunner = slowRunner.next
		fastRunner = fastRunner.next.next

		if fastRunner == slowRunner {
			return true
		}
	}
	return false
}

func returnStartOfCycleList(head *LinkedListNode) *LinkedListNode {
	node := head
	nodes := map[*LinkedListNode]struct{}{}
	for node != nil {
		// see if node is in set
		if _, ok := nodes[node]; ok {
			// must be a cycle...
			return node
		} else {
			nodes[node] = struct{}{}
		}
		node = node.next
	}
	return nil
}
