package trees_and_graphs

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)

	return root
}
