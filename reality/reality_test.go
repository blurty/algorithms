package reality

import "testing"

func TestTrap(t *testing.T) {
	type tester struct {
		input  []int
		output int
	}
	tests := []tester{
		tester{
			input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			output: 6,
		},
		tester{
			input:  []int{0, 0, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			output: 5,
		},
	}
	for _, v := range tests {
		got := Trap(v.input)
		if got != v.output {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestClimbStairs(t *testing.T) {
	type tester struct {
		input, output int
	}
	tests := []tester{
		tester{2, 2},
		tester{3, 3},
	}
	for _, v := range tests {
		got := ClimbStairs(v.input)
		if got != v.output {
			t.Errorf("want:%v, got:%v", v.output, got)
		}
	}
}

func TestIsWordExists(t *testing.T) {
	type tester struct {
		board  [][]byte
		word   string
		result bool
	}
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	tests := []tester{
		tester{
			board:  board,
			word:   "ABCCED",
			result: true,
		},
		tester{
			board:  board,
			word:   "SEE",
			result: true,
		},
		tester{
			board:  board,
			word:   "ABCB",
			result: false,
		},
		tester{
			board: [][]byte{
				[]byte{'a', 'b'},
				[]byte{'c', 'd'},
			},
			word:   "acdb",
			result: true,
		},
	}
	for _, v := range tests {
		got := IsWordExists(v.board, v.word)
		if got != v.result {
			t.Errorf("word:%v, want:%v, got:%v", v.word, v.result, got)
		}
	}
}
