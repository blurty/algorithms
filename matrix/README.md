# 矩阵问题

### 目录

- [矩阵问题](#矩阵问题)
    - [顺时针旋转矩阵90度](#顺时针旋转矩阵90度)
    - [逆时针旋转矩阵90度](#逆时针旋转矩阵90度)
    - [矩阵以螺旋序排列](#矩阵以螺旋序排列)
    - [以螺旋序生成矩阵](#以螺旋序生成矩阵)
    - [查找有序矩阵中的元素](#查找有序矩阵中的元素)
    - [矩阵置零](#矩阵置零)

## 顺时针旋转矩阵90度

You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

**Note:**

You have to rotate the image in-place, which means you have to modify
the input 2D matrix directly. DO NOT allocate another 2D matrix and do
the rotation.

*Example*

    Given input matrix =
    [
      [1,2,3],
      [4,5,6],
      [7,8,9]
    ],

    rotate the input matrix in-place such that it becomes:
    [
      [7,4,1],
      [8,5,2],
      [9,6,3]
    ]

### 算法以及代码

[算法查看注释](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#L10)


## 逆时针旋转矩阵90度

You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (anti clockwise).

**Note:**

You have to rotate the image in-place, which means you have to modify
the input 2D matrix directly. DO NOT allocate another 2D matrix and do
the rotation.

*Example*

    Given input matrix =
    [
      [1,2,3],
      [4,5,6],
      [7,8,9]
    ],

    rotate the input matrix in-place such that it becomes:
    [
      [3,6,9],
      [2,5,8],
      [1,4,7]
    ]

### 算法以及代码

[算法查看注释](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#L28)

## 矩阵以螺旋序排列

将一个m x n的矩阵以螺旋序排列并返回。

Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

### 举例

**输入：**

    [
     [ 1, 2, 3 ],
     [ 4, 5, 6 ],
     [ 7, 8, 9 ]
    ]

**输出：**

    [1,2,3,6,9,8,7,4,5]

### 算法以及代码

 I traverse right and increment rowBegin, then traverse down and decrement colEnd,
 then I traverse left and decrement rowEnd, and finally I traverse up and increment
 colBegin.

 when traverse left or up it has to check whether the row or col still exists to
 prevent duplicates.

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#L40)

## 以螺旋序生成矩阵

给一个数字n，返回一个n x n的矩阵，矩阵以螺旋序生成。

### 举例

**输入：**

    3

**输出：**

    [
        [ 1, 2, 3 ],
        [ 8, 9, 4 ],
        [ 7, 6, 5 ]
    ]

### 算法以及代码

本题与[矩阵以螺旋序排列](#矩阵以螺旋序排列)有异曲同工之妙。所以解法思路也一样。

从1到n*n依次找到在矩阵中对应的下标。很容易看出来矩阵下标的增长共有4个方向，我们以[rowInc, colInc]
分别来表示一维和二维下标的增量。可知四个取值分别为：

    [0, 1], [1, 0], [0, -1], [-1, 0]

我们以rowBegin, rowEnd, colBegin, colEnd分别来表示下标增长的四个边界。每次下标越过边界时，需要重置
边界值和下标值。分析可得如下四种情况：

1. row > rowEnd，此时向右增长到边界，需要向下增长，同时重置下标值。colEnd--;row, col = rowEnd, colEnd; rowInc, colInc = 0, -1。
2. 同上分析，分别可得出其他三种情况的取值。具体见代码。

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#L80)

## 查找有序矩阵中的元素

在一个m * n的二维矩阵中查找一个元素是否存在。矩阵是有序的，有以下两个特性：

1. 矩阵中的每一行都是从左到右有序的。
2. 矩阵中的每一行的第一个元素都比上一行的最后一个元素大。

### 分析

我们可以把这个二维矩阵看作一个有序的数组。那么对于一个有序数组中查找一个元素再简单不过了，我们使用二分查找法。

### 代码实现

[传送门](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#L123)

## 矩阵置零

一个m x n的矩阵，如果某个元素为0，则把该元素所在的整行和整列都置0。

### O(mn) Space

    func SetZeroes(matrix [][]int) {
        var rows, cols int
    	if rows = len(matrix); rows == 0 {
    		return
    	}
    	if cols = len(matrix[0]); cols == 0 {
    		return
    	}

        flagArr := make([][]bool, rows)
	    for i := 0; i < rows; i++ {
	    	flagArr[i] = make([]bool, cols)
	    }
	    for i := 0; i < rows; i++ {
	    	for j := 0; j < cols; j++ {
	    		if matrix[i][j] == 0 {
	    			for m := 0; m < cols; m++ {
	    				flagArr[i][m] = true
	    			}
	    			for n := 0; n < rows; n++ {
	    				flagArr[n][j] = true
	    			}
	    		}
	    	}
	    }
	    for i := 0; i < rows; i++ {
	    	for j := 0; j < cols; j++ {
	    		if flagArr[i][j] {
	    			matrix[i][j] = 0
	    		}
	    	}
	    }
	}

### O(m+n) Space

    func SetZeroes(matrix [][]int) {
    	var rows, cols int
    	if rows = len(matrix); rows == 0 {
    		return
    	}
    	if cols = len(matrix[0]); cols == 0 {
    		return
    	}

    	flagRows := make([]bool, rows)
    	flagCols := make([]bool, cols)
    	for i := 0; i < rows; i++ {
    		for j := 0; j < cols; j++ {
    			if matrix[i][j] == 0 {
    				flagRows[i] = true
    				flagCols[j] = true
    			}
    		}
    	}
    	for i := 0; i < rows; i++ {
    		for j := 0; j < cols; j++ {
    			if flagRows[i] || flagCols[j] {
    				matrix[i][j] = 0
    			}
    		}
    	}
    }

### O(1)

    func SetZeroes(matrix [][]int) {
    	var rows, cols int
    	if rows = len(matrix); rows == 0 {
    		return
    	}
    	if cols = len(matrix[0]); cols == 0 {
    		return
    	}

    	col0 := 1
    	for i := 0; i < rows; i++ {
    		if matrix[i][0] == 0 {
    			col0 = 0
    		}
    		for j := 1; j < cols; j++ {
    			if matrix[i][j] == 0 {
    				matrix[i][0], matrix[0][j] = 0, 0
    			}
    		}
    	}

    	for i := rows - 1; i >= 0; i-- {
    		for j := cols - 1; j > 0; j-- {
    			if matrix[i][0] == 0 || matrix[0][j] == 0 {
    				matrix[i][j] = 0
    			}
    		}
    		if col0 == 0 {
    			matrix[i][0] = 0
    		}
    	}
    }