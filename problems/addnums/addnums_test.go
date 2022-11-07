package addnums

import (
	"testing"
)

func TestAddNums(t *testing.T) {
	l1 := ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  3,
			},
			Val: 7,
		},
		Val: 5,
	}

	l2 := ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  4,
			},
			Val: 7,
		},
		Val: 1,
	}

	result := AddNums(&l1, &l2)

	var expVals = [3]int{6, 4, 8}

	for i, ch := range expVals {
		if result.Val != ch {
			t.Errorf("num is %d at %d expected %d", result.Val, i, ch)
		}

		if result.Next != nil {
			result = result.Next
		} else {
			break
		}
	}
}
