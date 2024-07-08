package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrderTraverse(t *testing.T) {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	expected := [][]int{{3}, {9, 20}, {15, 7}}
	assert.Equal(t, expected, levelOrder(input))

	input = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	expected = [][]int{{1}, {2, 3}, {4, 5}}
	assert.Equal(t, expected, levelOrder(input))

}
