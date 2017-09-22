package permutation

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombineString(t *testing.T) {
	got := CombineString("abc", "def")
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	sort.Strings(got)
	sort.Strings(expect)
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("got:%v, want:%v", got, expect)
	}
}
