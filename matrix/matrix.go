package matrix

/*
 * clockwise rotate
 * first reverse up to down, then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
 */
func Rotate(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/*
 * anticlockwise rotate
 * first swap the symmetry, then reverse left to right
 * 1 2 3     1 4 7     3 6 9
 * 4 5 6  => 2 5 8  => 2 5 8
 * 7 8 9     3 6 9     1 4 7
 */
func AntiRotate(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
}

// return all elements of the matrix in spiral order
func SpiralOrder(matrix [][]int) []int {
	var ret []int
	if len(matrix) == 0 {
		return ret
	}
	rowBegin, rowEnd := 0, len(matrix)-1
	columnBegin, columnEnd := 0, len(matrix[0])-1
	ret = make([]int, 0, len(matrix)*len(matrix[0]))

	for rowBegin <= rowEnd && columnBegin <= columnEnd {
		for i := columnBegin; i <= columnEnd; i++ {
			ret = append(ret, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			ret = append(ret, matrix[i][columnEnd])
		}
		columnEnd--

		// if-condition is to avoid 1 x 3 matrix
		if rowBegin <= rowEnd {
			for i := columnEnd; i >= columnBegin; i-- {
				ret = append(ret, matrix[rowEnd][i])
			}
		}
		rowEnd--

		// if-condition is to avoid 3 x 1 matrix
		if columnBegin <= columnEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				ret = append(ret, matrix[i][columnBegin])
			}
		}
		columnBegin++
	}
	return ret
}
