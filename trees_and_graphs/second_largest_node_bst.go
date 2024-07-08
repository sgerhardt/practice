package trees_and_graphs

func secondLargestNodeBinaryTree(root *BinaryNode) *BinaryNode {

	type nodeWithPredecessor struct {
		node        *BinaryNode
		predecessor *BinaryNode
	}

	nodesWithPredecessor := []nodeWithPredecessor{}
	nodesWithPredecessor = append(nodesWithPredecessor, nodeWithPredecessor{
		node:        root,
		predecessor: nil,
	})

	// Do a depth first search using our queue as a stack
	for len(nodesWithPredecessor) > 0 {
		// pop a node and its bounds from the top of our stack. (Remember, Depth first search uses a stack...)
		nb := nodesWithPredecessor[len(nodesWithPredecessor)-1]

		// check edge cases
		if nb.predecessor == nil && nb.node.Right == nil && nb.node.Left != nil {
			return nb.node.Left
		} else if nb.predecessor == nil && nb.node.Right == nil && nb.node.Left == nil {
			// only a root node...
			return nil
		}

		nodesWithPredecessor = nodesWithPredecessor[:len(nodesWithPredecessor)-1]

		if nb.node.Right != nil {
			// node has more elements on Right
			nodesWithPredecessor = append(nodesWithPredecessor, nodeWithPredecessor{
				node:        nb.node.Right,
				predecessor: nb.node,
			})
		} else if nb.node.Right == nil {
			// we reached the last largest item
			// if a node exists to the Left, check for a subtree, otherwise the parent is second largest...
			if nb.node.Left != nil {
				// The second-largest item is the largest item in the Left sub-tree
				return findLargest(nb.node.Left)
			} else {
				return nb.predecessor
			}
		}

	}

	return nil
}

func findLargest(node *BinaryNode) *BinaryNode {
	current := node
	for current != nil {
		if current.Right == nil {
			return current
		}
		current = current.Right
	}
	return nil
}
