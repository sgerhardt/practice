package trees_and_graphs

///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//
//

//Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
//
//
//Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,2,3]
//Example 2:
//
//Input: root = []
//Output: []
//Example 3:
//
//Input: root = [1]
//Output: [1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder traversal is used to create a copy of the tree.
// Preorder traversal is also used to get prefix expression on an expression tree. Please see http://en.wikipedia.org/wiki/Polish_notation to know why prefix expressions are useful.
// Example: Preorder traversal for the above-given figure is 1 2 4 5 3.
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	ret := []int{}

	for len(stack) > 0 {
		//pop the stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			ret = append(ret, node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}
	return ret
}
