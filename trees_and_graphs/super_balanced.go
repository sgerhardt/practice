package trees_and_graphs

import "fmt"

// A tree is "superbalanced" (made-up proeprty) if the difference between the depths of any two leaf nodes is no greater than one.
// Your first thought might be to write a recursive function, thinking, "the tree is balanced if the Left subtree is balanced and the Right subtree is balanced." This kind of approach works well for some other tree problems.
//
// But this isn't quite true. Counterexample: suppose that from the root of our tree:
//
// The Left subtree only has leaves at depths 10 and 11.
// The Right subtree only has leaves at depths 11 and 12.
// Both subtrees are balanced, but from the root we will have leaves at 3 different depths.
//
// We could instead have our recursive function get the list of distinct leaf depths for each subtree. That could work fine. But let's come up with an iterative solution instead. It's usually better to use an iterative solution instead of a recursive one because it avoids stack overflow. ↴
// Depth-first ↴ and breadth-first ↴ are common ways to traverse a tree. Which one should we use here?
//
// The worst-case time and space costs of both are the same—you could make a case for either.
//
// But one characteristic of our algorithm is that it could short-circuit and return False as soon as it finds two leaves with depths more than 1 apart. So maybe we should use a traversal that will hit leaves as quickly as possible...
// Depth-first traversal will generally hit leaves before breadth-first, so let's go with that. How could we write a depth-first walk that also keeps track of our depth?
func checkSuperbalance(node *BinaryNode) int {
	if node == nil {
		return -1
	} else {
		/* compute the depth of each subtree */
		lDepth := checkSuperbalance(node.Left)
		rDepth := checkSuperbalance(node.Right)

		/* use the larger one */
		if lDepth-rDepth > 1 || rDepth-lDepth > 1 {
			fmt.Printf("not superbalanced!\n")
			return 0
		}
		if lDepth > rDepth {
			return lDepth + 1
		} else {
			return rDepth + 1
		}
	}
}

func cakeIsSuperBalance(root *BinaryNode) bool {
	if root == nil {
		return true
	}

	//short circuit as soon as we find more than 2
	depths := map[int]struct{}{}

	// treat this list as a stack that will store tuples of of (node, depth)
	nodeDepths := []nodeDepth{}
	nodeDepths = append(nodeDepths, nodeDepth{
		node:  root,
		depth: 0,
	})

	for len(nodeDepths) > 0 {
		// pop a node and its depth from the top of our stack. (Remember, Depth first search uses a stack...)
		nd := nodeDepths[len(nodeDepths)-1]
		nodeDepths = nodeDepths[:len(nodeDepths)-1]
		if nd.node.Left == nil && nd.node.Right == nil {
			// we found a leaf

			// We only care if it's a new depth
			if _, found := depths[nd.depth]; !found {
				depths[nd.depth] = struct{}{}
				//                # Two ways we might now have an unbalanced tree:
				//                #   1) more than 2 different leaf depths
				//                #   2) 2 leaf depths that are more than 1 apart
				if len(depths) > 2 || len(depths) == 2 && depthDiffGreaterThan1(depths) {
					return false
				}

			}
		} else {
			// this isn't a leaf, keep stepping down
			if nd.node.Left != nil {
				nodeDepths = append(nodeDepths, nodeDepth{nd.node.Left, nd.depth + 1})
			}
			if nd.node.Right != nil {
				nodeDepths = append(nodeDepths, nodeDepth{nd.node.Right, nd.depth + 1})
			}
		}
	}
	return true
}

func depthDiffGreaterThan1(depths map[int]struct{}) bool {
	//assuming len 2
	depthVal := make([]int, len(depths))

	count := 0
	for depth, _ := range depths {
		depthVal[count] = depth
		count++
	}

	diff := depthVal[0] - depthVal[1]
	if diff > 1 || diff < -1 {
		return true
	}
	return false
}
