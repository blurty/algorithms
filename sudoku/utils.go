package sudoku

// judge if a Sudoku is valid
// The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
func IsValidSudoku(board [][]byte) bool {
	var visit1, visit2, visit3 [9][9]bool

	if len(board) < 9 || len(board[0]) < 9 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0' - 1
				k := i/3*3 + j/3
				if visit1[i][num] || visit2[j][num] || visit3[k][num] {
					return false
				}
				visit1[i][num], visit2[j][num], visit3[k][num] = true, true, true
			}
		}
	}
	return true
}
