package linklist

import "testing"

func convertListToSlice(head *ListNode) []int {
	array := make([]int, 0)
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	return array
}

func TestMergeTwoLists(t *testing.T) {
	// first case:
	// input: [1, 2, 5], [2, 3, 6]
	// output: [1, 2, 2, 3, 5, 6]
	a := &ListNode{Val: 1}
	a1 := &ListNode{Val: 2}
	a2 := &ListNode{Val: 5}
	a1.Next = a2
	a.Next = a1

	c := &ListNode{Val: 2}
	c1 := &ListNode{Val: 3}
	c2 := &ListNode{Val: 6}
	c1.Next = c2
	c.Next = c1

	b := MergeTwoLists(a, c)
	bs := convertListToSlice(b)
	output := []int{1, 2, 2, 3, 5, 6}
	if len(bs) != len(output) {
		t.Fatalf("want:%v, got:%v", output, bs)
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] != output[i] {
			t.Fatalf("want:%v, got:%v", output, bs)
		}
	}
}

func TestMergeNLists(t *testing.T) {
	a := &ListNode{Val: 1}
	a1 := &ListNode{Val: 2}
	a2 := &ListNode{Val: 5}
	a1.Next = a2
	a.Next = a1

	c := &ListNode{Val: 2}
	c1 := &ListNode{Val: 3}
	c2 := &ListNode{Val: 6}
	c1.Next = c2
	c.Next = c1

	d := &ListNode{Val: 3}
	d1 := &ListNode{Val: 4}
	d2 := &ListNode{Val: 5}
	d1.Next = d2
	d.Next = d1

	m := []*ListNode{a, c, d}

	b := MergeNLists(m)
	bs := convertListToSlice(b)
	output := []int{1, 2, 2, 3, 3, 4, 5, 5, 6}
	if len(bs) != len(output) {
		t.Fatalf("want:%v, got:%v", output, bs)
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] != output[i] {
			t.Fatalf("want:%v, got:%v", output, bs)
		}
	}
}
