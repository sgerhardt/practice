package linked_lists

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
// Return the linked list sorted as well.
//
// Example 1:
//
// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
func deleteAllDuplicates(head *ListNode) *ListNode {
	valueToOccurrence := map[int]int{}
	currentNode := head
	// first pass, count the duplicates
	for currentNode != nil {
		valueToOccurrence[currentNode.Val]++
		currentNode = currentNode.Next
	}

	//second pass, skip over duplicates
	currentNode = head
	var previousNode *ListNode
	for currentNode != nil {
		occurrences, _ := valueToOccurrence[currentNode.Val]
		if occurrences > 1 {
			for i := 0; i < occurrences; i++ {
				currentNode = currentNode.Next
			}
			if previousNode == nil {
				// duplicate was at head
				head = currentNode
				continue
			}
			previousNode.Next = currentNode
			continue
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return head
}

// differs from above solution by using another list instead of two passes
func deleteDuplicatesOnePass(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  head.Val + 1, // increment to guarantee the dummy node is not same as the original head. (this handles the first node not being falsely picked up as a duplicate since there's no previous yet...)
		Next: head.Next,
	}

	pre := dummy
	cur := head

	iterator := dummy // the first node in the new unduplicated(result) list

	for cur != nil && cur.Next != nil {
		if cur.Val != pre.Val && cur.Val != cur.Next.Val {
			// we found a unique node, we connect it at the tail of the unduplicated list, and update the iterator node.
			iterator.Next = cur
			iterator = iterator.Next
		}
		pre = cur
		cur = cur.Next
	}

	if pre.Val != cur.Val { // the last node needs to be dealt with independently
		iterator.Next = cur
		iterator = iterator.Next
	}

	// Notice example like: 1 2 2, need to point 1 to nil at last
	iterator.Next = nil

	return dummy.Next
}

//public ListNode deleteDuplicates(ListNode head) {
//    if (head == null) {
//        return null;
//    }
//
//    ListNode dummy = new ListNode(0 == head.val ? 1 : 0);// to guarantee the dummy node is not same as the original head.
//    dummy.next = head;
//
//    ListNode pre = dummy;
//    ListNode cur = head;
//
//    ListNode first = dummy; // the first node in the new unduplicated(result) list.
//
//    while (cur != null && cur.next != null) {
//        if (cur.val != pre.val && cur.val != cur.next.val) { // we found a unique node, we connect it at the tail of the unduplicated list, and update the first node.
//            first.next = cur;
//            first = first.next;
//        }
//        pre = cur;
//        cur = cur.next;
//    }
//
//    if (pre.val != cur.val) {  // the last node needs to be dealt with independently
//        first.next = cur;
//        first = first.next;
//    }
//
//    first.next = null;  // the subsequent list is duplicate.
//
//    return dummy.next;
//}
