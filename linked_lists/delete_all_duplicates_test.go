package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAllDuplicates(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}

	assert.Equal(t, expected, deleteAllDuplicates(input))

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	assert.Equal(t, expected, deleteDuplicatesOnePass(input))

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	expected = &ListNode{
		Val: 2,
	}

	assert.Equal(t, expected, deleteAllDuplicates(input))
}
