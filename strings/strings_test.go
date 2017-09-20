package strings

import (
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
