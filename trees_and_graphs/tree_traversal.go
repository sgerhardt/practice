package trees_and_graphs

func traverseTree(root *BinaryNode) {
	if root == nil {
		return
	}

	var nodes []*BinaryNode
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		//pop a node
		node := nodes[len(nodes)-1]
		println(node.Val)
		nodes = nodes[:len(nodes)-1]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
}

//	                    1
//	                     \
//	                     2
//						   /
//	                   3
//
// Gives 1,3,2 (note this is not "in order" for non-bst trees...)
func inOrderTraverseTree(root *BinaryNode) []int64 {
	var nodeValues []int64
	if root == nil {
		return nodeValues
	}

	var nodes []*BinaryNode
	temp := root
	for len(nodes) > 0 || temp != nil {
		if temp != nil {
			nodes = append(nodes, temp)
			temp = temp.Left
		} else {
			//pop a node
			temp = nodes[len(nodes)-1]
			println(temp.Val)
			nodeValues = append(nodeValues, temp.Val)
			nodes = nodes[:len(nodes)-1]
			temp = temp.Right
		}
	}
	return nodeValues
}
