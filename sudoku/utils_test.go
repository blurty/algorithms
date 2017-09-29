package sudoku

import "testing"

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
