package trees_and_graphs

func postorderTraversal(root *TreeNode) []int {
	var nodeValues []int
	if root == nil {
		return nodeValues
	}

	var nodes []*TreeNode

	// traverse down right of tree, adding node values to the _front_ of nodeValues.
	// If we get to a nil right node, pop off previous node from stack and continue traverse
	temp := root
	for len(nodes) > 0 || temp != nil {
		if temp != nil {
			nodes = append(nodes, temp)
			// add value to front
			nodeValues = append([]int{temp.Val}, nodeValues...)
			temp = temp.Right
		} else {
			// temp is nil, pop a node
			temp = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			// we have traversed the right side, not traverse the left...
			temp = temp.Left
		}
	}
	return nodeValues
}
