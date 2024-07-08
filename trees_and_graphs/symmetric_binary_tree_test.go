package trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetric(t *testing.T) {
	//tree [1,2,2,3,4,4,3]
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
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryNode{
			Val: 2,
			Left: &BinaryNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.True(t, isSymmetric(root))
}
