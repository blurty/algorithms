package array

import (
	"testing"
	"reflect"
)

func TestRemoveDuplicate(t *testing.T) {
	a := []int{1, 1, 2, 2}
	length := RemoveDuplicates(a)
	want := []int{1, 2}
	if length != 2 || !reflect.DeepEqual(want, a[:length]) {
		t.Fatalf("want:%v, got:%v", want, a[:length])
	}
}
