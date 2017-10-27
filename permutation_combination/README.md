# 排列组合

### 目录

- [排列组合](#排列组合)
    - [括弧问题](#括弧问题)
    - [一串数字的排列组合](#一串数字的排列组合)
    - [一串数字的排列组合去重后的结果](#一串数字的排列组合去重后的结果)
    - [计算第k个组合](#计算第k个组合)

## 括弧问题

### Longest Valid Parentheses

> Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
For "(()", the longest valid parentheses substring is "()", which has length = 2.
Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/permutation_combination/parentheses.go#L80)

## 一串数字的排列组合

> Given a collection of distinct numbers, return all possible permutations.
  For example,
  [1,2,3] have the following permutations:

    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]

### 查看代码

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/permutation_combination/permutation.go#L47)

## 一串数字的排列组合去重后的结果

> Given a collection of numbers that might contain duplicates, return all possible unique permutations.
  For example,
  [1,1,2] have the following unique permutations:

    [
      [1,1,2],
      [1,2,1],
      [2,1,1]
    ]

### 查看代码

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/permutation_combination/permutation.go#L73)

## 计算第k个组合

> The set [1,2,3,…,n] contains a total of n! unique permutations.
  By listing and labeling all of the permutations in order,
  We get the following sequence (ie, for n = 3):

    "123"
    "132"
    "213"
    "231"
    "312"
    "321"
    
  Given n and k, return the kth permutation sequence.

  Note: Given n will be between 1 and 9 inclusive.

### 查看代码

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/permutation_combination/permutation.go#L133)