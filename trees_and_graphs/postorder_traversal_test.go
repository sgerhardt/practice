package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostorder(t *testing.T) {
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

	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(tree))
}
