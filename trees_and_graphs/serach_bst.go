package trees_and_graphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//You are given the root of a binary search tree (BST) and an integer val.
//
//Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.
//Input: root = [4,2,7,1,3], val = 2
//Output: [2,1,3]
//

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if node := searchBST(root.Left, val); node != nil {
		if node.Val == val {
			return node
		}
	}
	if node := searchBST(root.Right, val); node != nil {
		if node.Val == val {
			return node
		}
	}

	return nil

}
