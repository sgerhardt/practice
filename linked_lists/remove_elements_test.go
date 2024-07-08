package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	assert.Equal(t, expected, removeElements(list, 2))
	//assert.Nil(t, removeElements(nil, 1))

	//list3 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}
	//assert.Equal(t, &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}, removeElements(list3, 1))

	//list4 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val:  1,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//assert.Equal(t, &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: nil,
	//	},
	//}, removeElements(list4, 1))
	//
	//list2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: nil,
	//	},
	//}
	//assert.Nil(t, removeElements(list2, 1))

}
