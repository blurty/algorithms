package reality

// Given n non-negative integers representing an elevation map
// where the width of each bar is 1, compute how much water it
// is able to trap after raining.
// Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
func Trap(height []int) int {
	left, right := 0, len(height)-1 // left and right index of height
	ret := 0                        // total capacity
	maxLeft, maxRight := 0, 0       // max value of left side and right side
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				ret += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				ret += maxRight - height[right]
			}
			right--
		}
	}
	return ret
}

// Each time climb 1 or 2 steps.
// given steps of stairs
// return total distinct ways to climb to the top
func ClimbStairs(n int) int {
	// base cases
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	oneStepBefore, twoStepsBefore := 2, 1
	allWays := 0
	for i := 2; i < n; i++ {
		allWays = oneStepBefore + twoStepsBefore
		twoStepsBefore, oneStepBefore = oneStepBefore, allWays
	}
	return allWays
}

// Given a 2D board and a word, find if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent
// cell, where "adjacent" cells are those horizontally or vertically
// neighboring. The same letter cell may not be used more than once.
func IsWordExists(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if board == nil {
		return false
	}

	visit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visit[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if len(word) == 1 {
					return true
				} else {
					visit[i][j] = true
					if recursiveSearch(board, visit, word[1:], i, j) {
						return true
					}
					visit[i][j] = false
				}
			}
		}
	}
	return false
}

// recursive search each byte of word in board
func recursiveSearch(board [][]byte, visit [][]bool, word string, x, y int) bool {
	if word == "" {
		return true
	}
	if x > 0 && !visit[x-1][y] && board[x-1][y] == word[0] {
		if len(word) == 1 {
			return true
		}
		visit[x-1][y] = true
		if recursiveSearch(board, visit, word[1:], x-1, y) {
			return true
		}
		visit[x-1][y] = false
	}
	if x+1 < len(board) && !visit[x+1][y] && board[x+1][y] == word[0] {
		if len(word) == 1 {
			return true
		}
		visit[x+1][y] = true
		if recursiveSearch(board, visit, word[1:], x+1, y) {
			return true
		}
		visit[x+1][y] = false
	}
	if y > 0 && !visit[x][y-1] && board[x][y-1] == word[0] {
		if len(word) == 1 {
			return true
		}
		visit[x][y-1] = true
		if recursiveSearch(board, visit, word[1:], x, y-1) {
			return true
		}
		visit[x][y-1] = false
	}
	if y+1 < len(board[0]) && !visit[x][y+1] && board[x][y+1] == word[0] {
		if len(word) == 1 {
			return true
		}
		visit[x][y+1] = true
		if recursiveSearch(board, visit, word[1:], x, y+1) {
			return true
		}
		visit[x][y+1] = false
	}
	return false
}

// Given n non-negative integers representing the histogram's bar height
// where the width of each bar is 1, find the area of largest rectangle
// in the histogram.
func LargestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	heights = append(heights, 0) // append 0 to assure calculation down
	var max int
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			startIdx := -1
			if len(stack) > 0 {
				startIdx = stack[len(stack)-1]
			}
			area := height * (i - startIdx - 1)
			if area > max {
				max = area
			}
		}
		stack = append(stack, i)
	}
	return max
}
