package linked_lists

//You are given the heads of two sorted linked lists list1 and list2.
//
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
//Return the head of the merged linked list.
//
//
//
//Example 1:
//
//
//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//Example 2:
//
//Input: list1 = [], list2 = []
//Output: []
//Example 3:
//
//Input: list1 = [], list2 = [0]
//Output: [0]
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	} else if list2 == nil && list1 != nil {
		return list1
	} else if list2 == nil && list1 == nil {
		return nil
	}

	var mergedList *ListNode
	var headOfMergedList *ListNode

	var previousMergeNode *ListNode
	headNodeSet := false
	for list1 != nil || list2 != nil {

		if list1 == nil {
			mergedList = &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
		} else if list2 == nil {
			mergedList = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			mergedList = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			list1 = list1.Next
		} else {
			mergedList = &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
		}

		if !headNodeSet {
			headOfMergedList = mergedList
			headNodeSet = true
		}

		if previousMergeNode != nil {
			previousMergeNode.Next = mergedList
		}
		previousMergeNode = mergedList
	}

	return headOfMergedList
}
