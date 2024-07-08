package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	//assert.Equal(t, expected, addTwoNumbers(l1, l2))
	assert.Equal(t, expected, addTwoNumbersOptimal(l1, l2))
}
