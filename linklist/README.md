## LinkList - 链表

### contents

- [合并链表](merge)

## Merge

### merge two sorted lists
> Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of
the first two lists.

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