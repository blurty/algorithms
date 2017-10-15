# 现实问题

### 目录

- [现实问题](#现实问题)
    - [蓄水](#trapping-rain-water)

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

[传送门](https://github.com/BlurtHeart/algorithms/tree/master/reality/reality.go#L7)