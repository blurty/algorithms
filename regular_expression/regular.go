package regular_expression

/*func IsMatch(s string, p string) bool {
	if len(p) < 1 {
		return len(s) == 0
	}
	// x* matches empty string or at least one character: x* -> xx*
	// *s is to ensure s is non-empty
	if len(p) > 1 && '*' == p[1] {
		return IsMatch(s, p[2:]) || len(s) > 0 && (p[0] == s[0] || '.' == p[0]) && IsMatch(s[1:], p)
	} else {
		return len(s) > 0 && (p[0] == s[0] || '.' == p[0]) && IsMatch(s[1:], p[1:])
	}
}*/

/**
 * f[i][j]: if s[0..i-1] matches p[0..j-1]
 * if p[j - 1] != '*'
 *      f[i][j] = f[i - 1][j - 1] && s[i - 1] == p[j - 1]
 * if p[j - 1] == '*', denote p[j - 2] with x
 *      f[i][j] is true iff any of the following is true
 *      1) "x*" repeats 0 time and matches empty: f[i][j - 2]
 *      2) "x*" repeats >= 1 times and matches "x*x": s[i - 1] == x && f[i - 1][j]
 * '.' matches any single character
 */
/*
 0true	1false	2	3	4	5	6
 1false
 2false
 3false
*/
func IsMatch(s string, p string) bool {
	row, column := len(s), len(p)
	f := make([][]bool, row+1)
	for i := range f {
		f[i] = make([]bool, column+1)
	}
	f[0][0] = true
	for i := 1; i <= row; i++ {
		f[i][0] = false
	}
	// p[0.., j - 3, j - 2, j - 1] matches empty iff p[j - 1] is '*' and p[0..j - 3] matches empty
	for j := 1; j <= column; j++ {
		f[0][j] = j > 1 && '*' == p[j-1] && f[0][j-2]
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if p[j-1] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i-1] == p[j-1] || '.' == p[j-1])
			} else {
				// p[0] cannot be '*' so no need to check "j > 1" here
				f[i][j] = f[i][j-2] || (s[i-1] == p[j-2] || '.' == p[j-2]) && f[i-1][j]
			}
		}
	}
	return f[row][column]
}
