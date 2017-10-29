# 链表

### 目录

- [链表](#链表)
    - [合并链表](#合并链表)
    - [交换结点](#交换结点)
    - [按K反转结点](#按K反转结点)
    - [轮换结点](#轮换结点)

## 合并链表

### merge two sorted lists
> Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of
the first two lists.

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/linklist/merge.go#L11)

### merge zero or more sorted lists
> Merge n sorted linked lists and return it as one sorted list.
Analyze and describe its complexity.

*approach*

算法过程：

- 将n个链表两两捉对，并且合并每组组合
- 以上操作之后，n个链表组合成了n/2个链表，重复以上步骤随后变为n/4, n/8 ...
- 重复以上步骤直到获得最后一个链表。

我们在每次捉对组合的时候需要遍历N个节点（N为所有链表的节点总数），然后重复这个过程log2(n)次

如图所示：

![](https://github.com/BlurtHeart/markdownphotos/blob/master/algorithms/algorithm-merge_n_sorted_lists.png)

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/linklist/merge.go#L46)

## 交换结点

### Swap Nodes in Pairs
> Given a linked list, swap every two adjacent nodes and return its head.
  For example:
  Given 1->2->3->4, you should return the list as 2->1->4->3.
  Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/linklist/merge.go#L75)

## 按K反转结点
> Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
  k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
  You may not alter the values in the nodes, only nodes itself may be changed.
  Only constant memory is allowed.
  For example:
  Given this linked list: 1->2->3->4->5
  For k = 2, you should return: 2->1->4->3->5
  For k = 3, you should return: 3->2->1->4->5

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/linklist/merge.go#L90)

## 轮换结点
> Given a list, rotate the list to the right by k places, where k is non-negative.
  For example:
  Given 1->2->3->4->5->NULL and k = 2,
  return 4->5->1->2->3->NULL.

[查看代码](https://github.com/BlurtHeart/algorithms/blob/master/linklist/merge.go#L113)