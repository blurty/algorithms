## 对有序数组去重

看题目：

> Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this in place with constant memory.
For example,
Given input array nums = [1,1,2],
Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

简单翻译一下，输入一个有序数组，将数组去重，使得每个元素在数组中只会出现一次。输出是去重后的数组长度。

限制是不能使用额外的数组。即使数组在返回的长度之后还有元素也没有关系。

*举个例子：*

    输入是[]int{1, 1, 2}，输出应该是2。去重后的数组为[]int{1, 2}，即使你去重完之后是[]int{1, 2, 1}也是可以的。

*算法：*

1. 首先数组中的第一个元素肯定是不需要去重的，所以我们从数组的第二个元素开始处理。
2. 初始化flag为0，表示去重的个数。
3. 如果位于i的元素的值等于前一个去重之后的数，则flag加1。反之，则将该元素往前移flag位。

举个例子，假设输入数组为：

    1   1   2   2

i = 1时，元素等于1，flag=0，此时它与位于i-1-flag=0处的元素1相等，表示它需要去重，我们将flag增加为1

i = 2时，元素等于2，flag=1，此时它与位于i-1-flag=1处的元素1不相等，表示它不需要去重，我们需要将此元素往前移flag位。移位之后数组就变成了：

    1   2   2   2

i = 3时，元素等于2，flag=1，此时它与位于i-1-flag=1处的元素2相等，表示它需要去重，我们将flag增加为2

于是，我们返回的数组长度就是原数组长度4减去去重的个数2 = 2.

[查看代码](https://github.com/BlurtHeart/algorithms/tree/master/array/duplicate.go)

**为什么要拿每个元素跟i-1-flag处的元素作比较呢？**

答：因为i-1-flag是数组中的上一个有效数字（去重之后的数组），又因为数组是有序的，所以我们只需要跟上一个有效数字比较就可以。


## 查找第一个缺少的正整数

> 给定一个乱序的整型数组array，查找第一个缺少的正整数。

### 示例

1. array:[1,2,0], 缺少正整数3
2. array:[3,4,-1,1], 缺少正整数2

### 要求

算法的时间复杂度应该是O(n)，并且不占用额外的空间。

### 查看代码

[传送门](https://github.com/BlurtHeart/algorithms/tree/master/array/find.go#4)

## 从数组首元素跳到尾元素

> Given an array of non-negative integers, you are initially positioned at the first index of the array.
  Each element in the array represents your maximum jump length at that position.
  Your goal is to reach the last index in the minimum number of jumps.
  For example:
  Given array A = [2,3,1,1,4]
  The minimum number of jumps to reach the last index is 2. (Jump 1 step from index 0 to 1, then 3 steps to the last index.)
  *Note:*
  You can assume that you can always reach the last index.

### 时间复杂度

    O(n)

### 查看代码

[传送门](https://github.com/BlurtHeart/algorithms/tree/master/array/find.go#24)

## 判断数组是否可达

> Given an array of non-negative integers, you are initially positioned at the first index of the array.
  Each element in the array represents your maximum jump length at that position.
  Determine if you are able to reach the last index.
  For example:
  A = [2,3,1,1,4], return true.
  A = [3,2,1,0,4], return false.

### 时间复杂度

    O(n)

### 查看代码

[传送门](https://github.com/BlurtHeart/algorithms/tree/master/array/find.go#41)
