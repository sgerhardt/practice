package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInOrderTraversal(t *testing.T) {

	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(tree))

	bstree := &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	assert.Equal(t, []int{3, 4, 5}, inorderTraversal(bstree))
}
