package addnums

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddNums(l1 *ListNode, l2 *ListNode) *ListNode {
	out := ListNode{
		Next: nil,
		Val:  l1.Val + l2.Val,
	}

	cur := out

	i := 1
	carry := 0
	for i < 2 {
		sum := carry

		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		next := ListNode{
			Next: nil,
			Val:  sum,
		}

		cur.Next = &next
		cur = next
	}

	return &out
}
