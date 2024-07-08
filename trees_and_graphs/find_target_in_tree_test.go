package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTargetInTree(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	assert.True(t, findTarget(tree, 11))
}
