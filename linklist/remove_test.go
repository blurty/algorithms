package linklist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteDuplicates2(t *testing.T) {
	type tester struct {
		intput, output []int
	}
	tests := []tester{
		tester{
			intput: nil,
			output: nil,
		},
		tester{
			intput: []int{1, 1, 1, 2, 2, 3},
			output: []int{3},
		},
		tester{
			intput: []int{1, 2},
			output: []int{1, 2},
		},
		tester{
			intput: []int{1, 1},
			output: nil,
		},
		tester{
			intput: []int{1, 2, 2, 3, 3, 4},
			output: []int{1, 4},
		},
	}
	for _, v := range tests {
		intputLink := createLinkList(v.intput)
		got := DeleteDuplicates2(intputLink)
		gotSlice := Link2Slice(got)
		if !reflect.DeepEqual(gotSlice, v.output) {
			t.Errorf("want linklist:%v, got linklist:%v", v.output, gotSlice)
		}
	}
}

func createLinkList(data []int) *ListNode {
	head := new(ListNode)
	privot := head
	for _, d := range data {
		node := &ListNode{d, nil}
		head.Next = node
		head = head.Next
	}
	return privot.Next
}

func Link2Slice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

func printLink(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
