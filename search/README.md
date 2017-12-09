# 查找

## 二分查找／折半查找

### 基本思想

二分查找的基本思想是将n个元素分成大致相等的两部分，取a[n/2]与x做比较，如果x=a[n/2],则找到x,算法中止；如果x<a[n/2],则只要在数组a的左半部分继续搜索x,如果x>a[n/2],则只要在数组a的右半部搜索x.

### 时间复杂度

    logn

### 代码

[查看代码](https://github.com/blurty/algorithms/blob/master/search/search.go#L4)