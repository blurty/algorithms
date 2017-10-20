# 矩阵问题

### 目录

- [矩阵问题](#矩阵问题)
    - [顺时针旋转矩阵90度](#顺时针旋转矩阵90度)
    - [逆时针旋转矩阵90度](#逆时针旋转矩阵90度)
    - [矩阵以螺旋序排列](#矩阵以螺旋序排列)

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

[算法查看注释](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#10)


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

[算法查看注释](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#28)

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

[查看算法](https://github.com/BlurtHeart/algorithms/tree/master/matrix/matrix.go#40)