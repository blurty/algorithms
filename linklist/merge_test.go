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

func TestSwapPairs(t *testing.T) {
	a := &ListNode{Val: 1}
	a1 := &ListNode{Val: 2}
	a2 := &ListNode{Val: 3}
	a3 := &ListNode{Val: 4}
	a2.Next = a3
	a1.Next = a2
	a.Next = a1

	b := SwapPairs(a)
	if b.Val != 2 {
		t.Fatalf("first node, want:2, got:%d", b.Val)
	}
	if b.Next == nil {
		t.Fatalf("second node, want:1, got:nil")
	} else if b.Next.Val != 1 {
		t.Fatalf("second node, want:1, got:%d", b.Next.Val)
	}
	if b.Next.Next == nil {
		t.Fatalf("third node, want:4, got:nil")
	} else if b.Next.Next.Val != 4 {
		t.Fatalf("third node, want:4, got:%d", b.Next.Next.Val)
	}
	if b.Next.Next.Next == nil {
		t.Fatalf("forth node, want:3, got:nil")
	} else if b.Next.Next.Next.Val != 3 {
		t.Fatalf("third node, want:3, got:%d", b.Next.Next.Next.Val)
	}
	if b.Next.Next.Next.Next != nil {
		t.Fatalf("no fifth node, got:%d", b.Next.Next.Next.Next.Val)
	}
}

func TestReverseKGroup(t *testing.T) {
	a := &ListNode{Val: 1}
	a1 := &ListNode{Val: 2}
	a2 := &ListNode{Val: 3}
	a3 := &ListNode{Val: 4}
	a4 := &ListNode{Val: 5}
	a3.Next = a4
	a2.Next = a3
	a1.Next = a2
	a.Next = a1

	b := ReverseKGroup(a, 3) // 3, 2, 1, 4, 5
	if b.Val != 3 {
		t.Fatalf("first node, want:3, got:%d", b.Val)
	}
	if b.Next == nil {
		t.Fatalf("second node, want:2, got:nil")
	} else if b.Next.Val != 2 {
		t.Fatalf("second node, want:2, got:%d", b.Next.Val)
	}
	if b.Next.Next == nil {
		t.Fatalf("third node, want:1, got:nil")
	} else if b.Next.Next.Val != 1 {
		t.Fatalf("third node, want:1, got:%d", b.Next.Next.Val)
	}
	if b.Next.Next.Next == nil {
		t.Fatalf("forth node, want:4, got:nil")
	} else if b.Next.Next.Next.Val != 4 {
		t.Fatalf("third node, want:4, got:%d", b.Next.Next.Next.Val)
	}
	if b.Next.Next.Next.Next == nil {
		t.Fatalf("forth node, want:5, got:nil")
	} else if b.Next.Next.Next.Next.Val != 5 {
		t.Fatalf("third node, want:5, got:%d", b.Next.Next.Next.Next.Val)
	}
	if b.Next.Next.Next.Next.Next != nil {
		t.Fatalf("no fifth node, got:%d", b.Next.Next.Next.Next.Next.Val)
	}
}
