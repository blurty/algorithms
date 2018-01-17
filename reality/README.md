# 现实问题

### 目录

- [现实问题](#现实问题)
    - [蓄水](#trapping-rain-water)
    - [爬楼梯](#climbing-stairs)
    - [查找单词](#search-word)
    - [计算最大矩形面积](#Largest-Rectangle-Area)

## Trapping Rain Water

*蓄水*

> Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.

### 图示

        []{0, 1, 0, 2, 0, 1}
    y
    ^
    |
    |            __
    |    __  ~~ |  | ~~  __
    |__ |__| __ |__| __ |__| __ > x

    ~下的面积为下雨后能够蓄水的量，可以简单的算出来是2。

### 解法

        instead of calculating area by height*width, we can think it in a cumulative way. In other words,
    sum water amount of each bin(width=1).
        Search from left to right and maintain a max height of left and right separately, which is like a
    one-side wall of partial container. Fix the higher one and flow water from the lower part. For example, if current height of left is lower, we fill water in the left bin. Until left meets right, we filled the whole container.

### 查看代码

[传送门](https://github.com/blurty/algorithms/tree/master/reality/reality.go#L7)

## Climbing Stairs

*爬楼梯*

### 问题描述

假设你在爬楼梯，每次只能爬一个或两个台阶。问有多少种方式能爬n个台阶。

### 示例

示例一：

    Input: 2
    Output:  2
    Explanation:  There are two ways to climb to the top.

    1. 1 step + 1 step
    2. 2 steps

示例二：

    Input: 3
    Output:  3
    Explanation:  There are three ways to climb to the top.

    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step

### 算法描述

假设我们要爬n个台阶，以[n]表示爬n个台阶的方式总数，相应的，以[n-1]和[n-2]分别表示爬n-1和n-2个台阶的总数。
分析一下，从n-1处到达n处，一次性爬一个台阶，而从n-2处到达n处，一次性爬两个台阶。所以，他们之间的关系是
`[n] = [n-1] + [n-2]`。在这个过程中不会存在重复或者缺少的情况。因为我们在最后一步把他们区分开来了。

现在让我们来分析一下整个算法的过程。

1. 当n为0时，[n] = 0
2. 当n为1时，我们只有一个方法，[n] = 1
3. 当n为2时，我们可以通过一次性爬两个台阶或者一次一个台阶的方式，[n] = 2
4. 当n大于2时，[n] = [n-1] + [n-2]

通过以上分析可知，这就是一个起点不同的伪斐波那契数列。

### 查看代码

[传送门](https://github.com/blurty/algorithms/tree/master/reality/reality.go#L34)

## Search Word

*在二维数组中查找单词*

### 问题描述

给一个二维数组以及一个单词，查找单词是否存在。要求单词必须由相邻单元格的字母组成，所谓相邻指
单元格的四邻域。没个单元格不可重复使用。

### 示例

示例一：

二维数组为：

    [
        ['A','B','C','E'],
        ['S','F','C','S'],
        ['A','D','E','E']
    ]

- word = "ABCCED", -> returns true,
- word = "SEE", -> returns true,
- word = "ABCB", -> returns false.

### 查看代码

[传送门](https://github.com/blurty/algorithms/tree/master/reality/reality.go#L59)

## Largest Rectangle Area

*寻找最大的矩形面积*

### 问题描述

给一个高度数组组成的一系列相邻矩形，计算这一系列相邻矩形所能组成的最大矩形面积。

### 示例

For example,

Given heights = [2,1,5,6,2,3],

return 10.

### 查看代码

[传送门](https://github.com/blurty/algorithms/tree/master/reality/reality.go#L140)