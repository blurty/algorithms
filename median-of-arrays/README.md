## Question
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

##### 前提条件
两个数组已排序

### Contents
- [Same Size Arrays](#same-size-arrays) - Median of Two-Same-Size-Sorted-Arrays
- [Different Size Arrays](#different-size-arrays) - Median of Two-Different-Size-Arrays
- [Both](#both) - combine both same size and different size

### Same Size Arrays

#### Example:
    nums1 = [1, 2]
    nums2 = [3, 4]

    The median is (2 + 3)/2 = 2.5

##### 中位数
>Median: In probability theory and statistics, a median is described
as the number separating the higher half of a sample, a population,
or a probability distribution, from the lower half.

>The median of a finite list of numbers can be found by arranging all
the numbers from lowest value to highest value and picking the middle one.

#### Method

##### 1. Method 1 (Simply count while Merging)
Use merge procedure of merge sort. Keep track of count while comparing elements of two arrays. If count becomes n(For 2n elements), we have reached the median. Take the average of the elements at indexes n-1 and n in the merged array. See the below implementation.

##### 2. Method 2 (By comparing the medians of two arrays)
This method works by first getting medians of the two sorted arrays and then comparing them.

### Different Size Arrays

#### Example:

    nums1 = [1, 3]
    nums2 = [2]

    The median is 2.0


### Both


##### 参考文献
1. [GeeksforGeeks](http://www.geeksforgeeks.org/median-of-two-sorted-arrays/) - median-of-two-sorted-arrays
2. [GeeksforGeeks](http://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/) - median-of-two-sorted-arrays-of-different-sizes