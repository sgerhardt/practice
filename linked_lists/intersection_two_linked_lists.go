package linked_lists

// Given the heads of two singly linked-lists headA and headB,
// return the node at which the two lists intersect.
// If the two linked lists have no intersection at all, return null.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeSet := map[*ListNode]struct{}{}

	node := headA
	for node != nil {
		nodeSet[node] = struct{}{}
		node = node.Next
	}

	node = headB
	for node != nil {
		_, exists := nodeSet[node]
		if exists {
			return node
		}
		node = node.Next
	}

	return nil
}
