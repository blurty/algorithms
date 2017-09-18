package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes
// of the first two lists.
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

// Merge n sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
func MergeNLists(lists []*ListNode) *ListNode {
	amount := len(lists)
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = MergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	if amount > 0 {
		return lists[0]
	} else {
		return nil
	}
}

// swap every two adjacent nodes and return its head.
// Given 1->2->3->4, return the list as 2->1->4->3.
// only swap the value in each node
//func SwapPairs(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	ptr := head
//	flag := 0
//	for ptr.Next != nil {
//		if flag%2 == 0 {
//			ptr.Next.Val, ptr.Val = ptr.Val, ptr.Next.Val
//		}
//		ptr = ptr.Next
//		flag++
//	}
//	return head
//}
// avoid modify the values in the list, only nodes itself changed
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	nextHead := newHead.Next
	newHead.Next = head
	head.Next = SwapPairs(nextHead)
	return newHead
}
