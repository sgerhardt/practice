package linked_lists

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
//
// Example 1:
//
// Input: head = [1,1,2]
// Output: [1,2]
// Example 2:
//
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head

	for currentNode != nil {
		for currentNode.Next != nil && currentNode.Val == currentNode.Next.Val {
			// duplicate found, assign next value to the value after it
			currentNode.Next = currentNode.Next.Next
		}
		currentNode = currentNode.Next
	}
	return head
}
