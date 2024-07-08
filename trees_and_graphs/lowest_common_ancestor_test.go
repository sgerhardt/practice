package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCA(t *testing.T) {

	input := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	}
	assert.Equal(t, input, lowestCommonAncestor(input, input.Left, input.Right))
}
