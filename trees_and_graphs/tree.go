package trees_and_graphs

type BinaryTree struct {
	root *BinaryNode
}

type BinaryNode struct {
	Val   int64
	Left  *BinaryNode
	Right *BinaryNode
}

type nodeDepth struct {
	node  *BinaryNode
	depth int
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{Val: data, Left: nil, Right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.Val {
		if n.Left == nil {
			n.Left = &BinaryNode{Val: data, Left: nil, Right: nil}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Val: data, Left: nil, Right: nil}
		} else {
			n.Right.insert(data)
		}
	}
}

func traverse(node *BinaryNode) {
	if node == nil {
		return
	}
	// believe this is preoder traversal...
	println(node.Val)
	traverse(node.Left)
	traverse(node.Right)
	return
}

/* computeMaxDepth computes the "maxDepth" of a tree -- the number of
   nodes along the longest path from the root node
   down to the farthest leaf node.*/
// Note that this is simply a traverse with a compare afterwards and returning an int.
// Base case is 0 so if the leafs are nil, we don't add anything
func computeMaxDepth(node *BinaryNode) int {
	if node == nil {
		return 0
	} else {
		/* compute the depth of each subtree */
		lDepth := computeMaxDepth(node.Left)
		rDepth := computeMaxDepth(node.Right)

		/* use the larger one */
		if lDepth > rDepth {
			return lDepth + 1
		} else {
			return rDepth + 1
		}
	}
}
