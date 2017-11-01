# 路径问题

### 目录

- [路径问题](#路径问题)
    - [在矩阵中寻找路径总数](#在矩阵中寻找路径总数)
    - [在有障碍的矩阵中寻找路径总数](#在有障碍的矩阵中寻找路径总数)
    - [计算最小路径总值](#计算最小路径总值)

## 在矩阵中寻找路径总数

有一个m*n的棋盘（矩阵），一个骑士在棋盘的左上角，他要走到棋盘的右下角，一次只能走一步，每次只能向右走一格或者向下走一格。

计算一共有多少条路径可以走。

### 算法

算法分析类似[计算最小路径总值](#计算最小路径总值)

### 时间复杂度

    O(m*n)

### 空间复杂度

    O(min(m, n))

### 代码

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/maze/maze.go#L7)

## 在有障碍的矩阵中寻找路径总数

与[在矩阵中寻找路径总数](#在矩阵中寻找路径总数)类似，只是在矩阵中加入障碍物，
障碍物用1来表示，如下所示：

    [
        [0,0,0],
        [0,1,0],
        [0,0,0],
    ]

以上矩阵的唯一路径总数是2。

### 算法

算法分析类似[计算最小路径总值](#计算最小路径总值)

### 代码

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/maze/maze.go#L38)

## 计算最小路径总值

有一个m x n的矩阵，矩阵中的每一个元素代表此处的路径权值，计算从左上角到右下角（只能向右或下走）的最小路径总值。

例如：

    [
        [1,3,1],
        [1,5,1],
        [4,2,1],
    ]

路径1→3→1→1→1是最小的，最小路径总值为7。

### 算法

1. 这是一个典型的DP问题。假设到达位置`(i,j)`最小的路径总值是`S[i][j]`，它的取值是`S[i][j] = min(S[i - 1][j], S[i][j - 1]) + grid[i][j]`
2. 当然，还是要处理一些边界情况。边界情况有两个，当位置在首行时，`S[i - 1][j]`不存在；当位置在首列时，`S[i][j - 1]`不存在。
3. 假设矩阵是`[1, 1, 1, 1]`，那么到达每一个位置的最小路径总值是一个累积的过程，结果是`[1, 2, 3, 4]`。

根据以上分析，可以写如下未经优化的代码：

    func MinPathSum(grid [][]int) int {
    	height, width := len(grid), len(grid[0])
    	vec := make([][]int, height)
    	for i := 0; i < height; i++ {
    		vec[i] = make([]int, width)
    	}
    	vec[0][0] = grid[0][0]
    	for i := 1; i < height; i++ {
    		vec[i][0] = vec[i-1][0] + grid[i][0]
    	}
    	for i := 1; i < width; i++ {
    		vec[0][i] = vec[0][i-1] + grid[0][i]
    	}

    	for i := 1; i < height; i++ {
    		for j := 1; j < width; j++ {
    			if vec[i-1][j] < vec[i][j-1] {
    				vec[i][j] = vec[i-1][j] + grid[i][j]
    			} else {
    				vec[i][j] = vec[i][j-1] + grid[i][j]
    			}
    		}
    	}
    	return vec[height-1][width-1]
    }

可以看到，每次当我们更新`sum[i][j]`时，我们仅仅需要`sum[i - 1][j]`以及`sum[i][j - 1]`的值。
所以我们不需要维护整个m*n的矩阵。仅仅维护两列就足够了。据此可以优化以上代码：

    func MinPathSum(grid [][]int) int {
    	height, width := len(grid), len(grid[0])
        vec1 := make([]int, width)
        vec2 := make([]int, width)
        vec1[0] = grid[0][0]
        for i := 1; i < width; i++ {
        	vec1[i] = vec1[i-1] + grid[0][i]
        }
        for i := 1; i < height; i++ {
        	vec2[0] = vec1[0] + grid[i][0]
        	for j := 1; j < width; j++ {
        		if vec2[j-1] < vec1[j] {
        			vec2[j] = vec2[j-1] + grid[i][j]
        		} else {
        			vec2[j] = vec1[j] + grid[i][j]
        		}
        	}
        	vec1, vec2 = vec2, vec1
        }
        return vec1[width-1]
    }

观察以上代码，可以看出维护vec是为了重新设置vec[i]的值，所以我们仅仅需要维护一个vector。现在的空间复杂度被进一步降低，代码如下：

    func MinPathSum(grid [][]int) int {
    	if len(grid) == 0 || len(grid[0]) == 0 {
    		return 0
    	}
    	width := len(grid[0])
    	vec := make([]int, width)
    	vec[0] = grid[0][0]
    	for i := 1; i < width; i++ {
    		vec[i] = vec[i-1] + grid[0][i]
    	}
    	for i := 1; i < len(grid); i++ {
    		vec[0] += grid[i][0]
    		for j := 1; j < width; j++ {
    			if vec[j-1] < vec[j] {
    				vec[j] = vec[j-1] + grid[i][j]
    			} else {
    				vec[j] += grid[i][j]
    			}
    		}
    	}
    	return vec[width-1]
    }

### 代码

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/maze/maze.go#L60)