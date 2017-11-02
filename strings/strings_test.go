package strings

import (
	"reflect"
	"sort"
	"testing"
)

type tester struct {
	input []string
	index int
}

func TestStrStr(t *testing.T) {
	tests := []tester{
		tester{[]string{"hello world", "world"}, 6},
		tester{[]string{"", ""}, 0},
		tester{[]string{"a", ""}, 0},
		tester{[]string{"", "a"}, -1},
	}

	for _, v := range tests {
		got := StrStr(v.input[0], v.input[1])
		if got != v.index {
			t.Fatalf("want:%d, got:%d", v.index, got)
		}
	}
}

func TestConcatStrings(t *testing.T) {
	input := []string{"hello", "world", "hi"}
	want := []string{"helloworldhi", "hellohiworld", "worldhellohi", "worldhihello", "hihelloworld", "hiworldhello"}
	got := ConcatStrings(input)
	sort.Strings(want)
	sort.Strings(got)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want:%v, got:%v", want, got)
	}
}

func TestMultiply(t *testing.T) {
	type tester struct {
		x, y, z string
	}
	tests := []tester{
		tester{"123", "456", "56088"},
		tester{"0", "0", "0"},
		tester{"999", "999", "998001"},
		tester{"52", "60", "3120"},
		tester{"9133", "0", "0"},
	}
	for _, v := range tests {
		got := Multiply(v.x, v.y)
		if got != v.z {
			t.Fatalf("want:%s, got:%s", v.z, got)
		}
	}
}

func TestAddBinary(t *testing.T) {
	type tester struct {
		x, y, z string
	}
	tests := []tester{
		tester{"11", "1", "100"},
		tester{"110101", "0101", "111010"},
	}
	for _, v := range tests {
		got := AddBinary(v.x, v.y)
		if got != v.z {
			t.Fatalf("want:%s, got:%s", v.z, got)
		}
	}
}
