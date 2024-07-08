package trees_and_graphs

import "math"

// Check that a binary tree is a valid binary search tree.
// A binary search tree is a binary tree where the nodes are ordered in a specific way. For every node:
// -The nodes to the Left are smaller than the current node.
// -The nodes to the Right are larger than the current node.
func isValidBinarySearchTree(root *BinaryNode) bool {

	type nodeWithBounds struct {
		node       *BinaryNode
		lowerBound int64
		upperBound int64
	}

	nodesWithBounds := []nodeWithBounds{}
	nodesWithBounds = append(nodesWithBounds, nodeWithBounds{
		node:       root,
		lowerBound: math.MinInt,
		upperBound: math.MaxInt,
	})

	// Do a depth first search using our slice as a stack
	for len(nodesWithBounds) > 0 {
		// pop a node and its bounds from the top of our stack. (Remember, Depth first search uses a stack...)
		nb := nodesWithBounds[len(nodesWithBounds)-1]
		println(nb.node.Val)
		if nb.node.Val < nb.lowerBound || nb.node.Val > nb.upperBound {
			return false
		}
		nodesWithBounds = nodesWithBounds[:len(nodesWithBounds)-1] // as part of the pop, remove the top value from the stack...

		// this isn't a leaf, keep stepping down
		if nb.node.Left != nil {
			nodesWithBounds = append(nodesWithBounds, nodeWithBounds{
				node:       nb.node.Left,
				lowerBound: nb.lowerBound,
				upperBound: nb.node.Val,
			})
		}
		if nb.node.Right != nil {
			nodesWithBounds = append(nodesWithBounds, nodeWithBounds{
				node:       nb.node.Right,
				lowerBound: nb.node.Val,
				upperBound: nb.upperBound,
			})
		}

	}

	return true
}
