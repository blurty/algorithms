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

type noder struct {
	items []byte
}

// add v to visit's three relevant items
// visit is an array, so give its address to initVisit to change it
func addVisit(visit *[3][9]noder, i, j int, v byte) {
	(*visit)[0][i].items = append((*visit)[0][i].items, v)
	(*visit)[1][j].items = append((*visit)[1][j].items, v)
	(*visit)[2][i/3*3+j/3].items = append((*visit)[2][i/3*3+j/3].items, v)
}

// delete the final item of visit's three relevant items
// visit is an array, so give its address to initVisit to change it
func deleteVisit(visit *[3][9]noder, i, j int) {
	(*visit)[0][i].items = (*visit)[0][i].items[:len((*visit)[0][i].items)-1]
	(*visit)[1][j].items = (*visit)[1][j].items[:len((*visit)[1][j].items)-1]
	(*visit)[2][i/3*3+j/3].items = (*visit)[2][i/3*3+j/3].items[:len((*visit)[2][i/3*3+j/3].items)-1]
}

// init visit
// visit is an array, so give its address to initVisit to change it
func initVisit(visit *[3][9]noder, board [][]byte) {
	// apply space for noder
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			(*visit)[i][j].items = make([]byte, 0, 9)
		}
	}
	// init visit by walk board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				addVisit(visit, i, j, board[i][j])
			}
		}
	}
}

// board has the unique solution
func SolveSudoku(board [][]byte) {
	var visit [3][9]noder
	initVisit(&visit, board)
	sudokuRecursion(board, &visit, 0)
}

// find next step, if find, return true. Otherwise, return false
// board is a slice, so no need to give its address
func sudokuRecursion(board [][]byte, visit *[3][9]noder, x int) bool {
	var k byte
	for i := x; i < 81; i++ {
		// m: the row of i
		// n: the column of i
		m, n := i/9, i%9
		if board[m][n] == '.' {
			for k = '1'; k <= '9'; k++ {
				if !isInByteSlice(visit[0][m].items, k) && !isInByteSlice(visit[1][n].items, k) && !isInByteSlice(visit[2][m/3*3+n/3].items, k) {
					board[m][n] = k
					addVisit(visit, m, n, k)
					result := sudokuRecursion(board, visit, i+1)
					if !result {
						board[m][n] = '.'
						deleteVisit(visit, m, n)
					} else {
						return true
					}
				}
			}
			return false
		}
	}
	return true
}

// check whether b in bArr
func isInByteSlice(bArr []byte, b byte) bool {
	for i := 0; i < len(bArr); i++ {
		if bArr[i] == b {
			return true
		}
	}
	return false
}
