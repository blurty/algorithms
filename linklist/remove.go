package linklist

// Given a sorted linked list, delete all duplicates
// such that each element appear only once.
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ret := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return ret
}

// Given a sorted linked list, delete all nodes that have
// duplicate numbers, leaving only distinct numbers from
// the original list.
func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	ret := pre
	for head != nil && head.Next != nil {
		flag := false
		for head.Next != nil && head.Val == head.Next.Val {
			flag = true
			head = head.Next
		}
		if flag && head.Next != nil && head.Next.Next != nil && head.Next.Val == head.Next.Next.Val {
			head = head.Next
			continue
		}
		if flag {
			pre.Next = head.Next
		} else {
			pre.Next = head
		}
		pre = pre.Next
		if pre == nil {
			break
		}
		head = pre.Next
	}
	return ret.Next
}
