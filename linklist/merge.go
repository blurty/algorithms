package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{}
	oldStart := start

	for {
		if l1 == nil && l2 == nil {
			break
		} else if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				start.Next = l1
				l1 = l1.Next
			} else {
				start.Next = l2
				l2 = l2.Next
			}
		} else if l1 == nil {
			start.Next = l2
			l2 = l2.Next
		} else {
			start.Next = l1
			l1 = l1.Next
		}
		start = start.Next
	}

	return oldStart.Next
}
