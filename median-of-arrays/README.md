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

Time Complexity: O(n)

Use merge procedure of merge sort. Keep track of count while comparing elements of two arrays. If count becomes n(For 2n elements), we have reached the median. Take the average of the elements at indexes n-1 and n in the merged array. See the below implementation.

##### 2. Method 2 (By comparing the medians of two arrays)

Time Complexity: O(logn)

This method works by first getting medians of the two sorted arrays and then comparing them.

*Let ar1 and ar2 be the input arrays.*

###### Algorithm:

    1) Calculate the medians m1 and m2 of the input arrays ar1[] 
       and ar2[] respectively.
    2) If m1 and m2 both are equal then we are done.
       return m1 (or m2)
    3) If m1 is greater than m2, then median is present in one 
       of the below two subarrays.
       a)  From first element of ar1 to m1 (ar1[0...|_n/2_|])
       b)  From m2 to last element of ar2  (ar2[|_n/2_|...n-1])
    4) If m2 is greater than m1, then median is present in one    
       of the below two subarrays.
       a)  From m1 to last element of ar1  (ar1[|_n/2_|...n-1])
       b)  From first element of ar2 to m2 (ar2[0...|_n/2_|])
    5) Repeat the above process until size of both the subarrays 
       becomes 2.
    6) If size of the two arrays is 2 then use below formula to get 
       the median.
        Median = (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1]))/2

###### Example:

    ar1[] = {1, 12, 15, 26, 38}
    ar2[] = {2, 13, 17, 30, 45}
    
For above two arrays m1 = 15 and m2 = 17

For the above ar1[] and ar2[], m1 is smaller than m2. So median is present in one of the following two subarrays.

    [15, 26, 38] and [2, 13, 17]

Let us repeat the process for above two subarrays:

    m1 = 26 m2 = 13.
    
m1 is greater than m2. So the subarrays become

    [15, 26] and [13, 17]
    Now size is 2, so median = (max(ar1[0], ar2[0]) + min(ar1[1], ar2[1]))/2
                       = (max(15, 13) + min(26, 17))/2 
                       = (15 + 17)/2
                       = 16
    
    
### Different Size Arrays

#### Example:

    nums1 = [1, 3]
    nums2 = [2]

    The median is 2.0


### Both


##### 参考文献
1. [GeeksforGeeks](http://www.geeksforgeeks.org/median-of-two-sorted-arrays/) - median-of-two-sorted-arrays
2. [GeeksforGeeks](http://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/) - median-of-two-sorted-arrays-of-different-sizes