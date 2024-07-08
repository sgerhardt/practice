package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertBST(t *testing.T) {

	input := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}

	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	assert.Equal(t, expected, insertIntoBST(input, 5))
}
