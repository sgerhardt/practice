package trees_and_graphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Searches through the left subtree, then the root, then the right subtree. This is a DFS\
//Input: root = [1,null,2,3]
//
//Output: [1,2,3]
//In the case of binary search trees (BST), Inorder traversal gives nodes in non-decreasing order.
//To get nodes of BST in non-increasing order, a variation of Inorder traversal where Inorder traversal s reversed can be used.
//Example: In order traversal for the above-given figure is 4 2 5 1 3.
func inorderTraversal(root *TreeNode) []int {
	var nodeValues []int
	if root == nil {
		return nodeValues
	}

	var nodes []*TreeNode

	temp := root
	for len(nodes) > 0 || temp != nil {
		if temp != nil {
			nodes = append(nodes, temp)
			temp = temp.Left
		} else {
			//pop a node
			temp = nodes[len(nodes)-1]
			nodeValues = append(nodeValues, temp.Val)
			nodes = nodes[:len(nodes)-1]
			temp = temp.Right
		}
	}
	return nodeValues
}
