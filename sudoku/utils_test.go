package sudoku

import (
	"reflect"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	tests := [][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	got := IsValidSudoku(tests)
	if got != true {
		t.Errorf("want:true, got:false")
	}
}

func TestSolveSudoku(t *testing.T) {
	tests := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	want := [][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}
	SolveSudoku(tests)
	if !reflect.DeepEqual(want, tests) {
		t.Errorf("\nwant:%v\ngot:%v", want, tests)
	}
}
