package trees_and_graphs

// You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion.
// It is guaranteed that the new value does not exist in the original BST.
//
// Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.
// Input: root = [4,2,7,1,3], val = 5
// Output: [4,2,7,1,3,5]
// Explanation: Another accepted tree is:
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	node := root
	var stack []*TreeNode
	stack = append(stack, node)

	for len(stack) > 0 {
		//pop off the stack
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if val < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
			} else {
				stack = append(stack, node.Left)
			}
		} else if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
			} else {
				stack = append(stack, node.Right)
			}
		}
	}
	return root
}
