package trees_and_graphs

func findTarget(root *TreeNode, k int) bool {
	// this stores k-value so that we can quickly check if a given node adds up to our target k
	values := map[int]struct{}{}

	node := root
	var stack []*TreeNode

	stack = append(stack, node)

	for len(stack) > 0 {
		// pop a node off the stack
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := values[node.Val]; ok {
			return true
		}
		values[k-node.Val] = struct{}{}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return false
}
