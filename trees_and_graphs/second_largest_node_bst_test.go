package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecondLargestNodeBinaryTree(t *testing.T) {
	//simple case
	root := &BinaryNode{
		Val: 5,
		Left: &BinaryNode{
			Val: 3,
			Left: &BinaryNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryNode{
			Val: 8,
			Left: &BinaryNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, int64(8), secondLargestNodeBinaryTree(root).Val)

	// largest tree has subtree
	root = &BinaryNode{
		Val: 5,
		Left: &BinaryNode{
			Val: 3,
			Left: &BinaryNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryNode{
			Val: 8,
			Left: &BinaryNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val: 12,
				Left: &BinaryNode{
					Val: 10,
					Left: &BinaryNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &BinaryNode{
						Val:   11,
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
		},
	}

	assert.Equal(t, int64(11), secondLargestNodeBinaryTree(root).Val)
}
