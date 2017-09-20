package integer

import "testing"

func TestDivide(t *testing.T) {
	tests := [][3]int{
		[3]int{-1, 1, -1},
		[3]int{1, -1, -1},
		[3]int{MinInt, -1, MaxInt},
		[3]int{5, 2, 2},
		[3]int{3, 1, 3},
		[3]int{1, 3, 0},
		[3]int{0, 1, 0},
	}
	for _, v := range tests {
		got := Divide(v[0], v[1])
		if got != v[2] {
			t.Fatalf("want:%d, got:%d", v[2], got)
		}
	}
}
