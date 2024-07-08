package trees_and_graphs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTree(t *testing.T) {
	root := &BinaryNode{
		Val: 1,
		Left: &BinaryNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &BinaryNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	root.insert(5)
	root.insert(7)

	//assert.Equal(t, int64(5), root.Right.Right.Val)

	traverse(root)

	fmt.Println(computeMaxDepth(root))

	checkSuperbalance(root)

	println(cakeIsSuperBalance(root))
}

func TestSearch(t *testing.T) {
	//preorder
	root := &BinaryNode{
		Val: 1,
		Left: &BinaryNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &BinaryNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	assert.False(t, isValidBinarySearchTree(root))

	root = &BinaryNode{
		Val: 1,
		Left: &BinaryNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &BinaryNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	assert.False(t, isValidBinarySearchTree(root))

	// checking node against parent is insufficient
	root = &BinaryNode{
		Val: 50,
		Left: &BinaryNode{
			Val: 30,
			Left: &BinaryNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   60,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryNode{
			Val: 80,
			Left: &BinaryNode{
				Val:   70,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   90,
				Left:  nil,
				Right: nil,
			},
		},
	}
	assert.False(t, isValidBinarySearchTree(root))
}

func TestLevelOrderTraverseTree(t *testing.T) {
	root := &BinaryNode{
		Val: 1,
		Left: &BinaryNode{
			Val: 2,
			Left: &BinaryNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryNode{
			Val: 3,
			Left: &BinaryNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	traverseTree(root)
	//bstRoot := &BinaryNode{
	//	Val: 2,
	//	Right: &BinaryNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Left: &BinaryNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//traverseTree(bstRoot)
	//assert.Equal(t, []int64{1, 2, 3}, traverseTree(root))
}

func TestInOrderTraverseTree(t *testing.T) {
	root := &BinaryNode{
		Val:  1,
		Left: nil,
		Right: &BinaryNode{
			Val: 2,
			Left: &BinaryNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	assert.Equal(t, []int64{1, 3, 2}, inOrderTraverseTree(root))

	bstRoot := &BinaryNode{
		Val: 2,
		Right: &BinaryNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Left: &BinaryNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	// NOTE: For a BST, in order gives the nodes in non-decreaseing order
	assert.Equal(t, []int64{1, 2, 3}, inOrderTraverseTree(bstRoot))
}
