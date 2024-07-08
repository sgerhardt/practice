package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicate(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	assert.Equal(t, expected, deleteDuplicates(input))
}
