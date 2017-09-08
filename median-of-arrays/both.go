package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < 1 && len(nums2) < 1 {
		return float64(0)
	}

	var mean int

	odd := (len(nums1) + len(nums2)) % 2
	if odd > 0 {
		mean = (len(nums1) + len(nums2) + 1) / 2
	} else {
		mean = (len(nums1)+len(nums2)+1)/2 + 1
	}

	x, y := 0, 0
	for i, j, flag := 0, 0, 0; flag < mean; flag++ {
		x = y
		if i == len(nums1) {
			y = nums2[j]
			j++
			continue
		}
		if j == len(nums2) {
			y = nums1[i]
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			y = nums1[i]
			i++
		} else {
			y = nums2[j]
			j++
		}
	}
	if odd%2 > 0 {
		return float64(y)
	} else {
		return float64(x+y) / 2
	}
}

// Driver program to test above functions
func main() {
	a := []int{900}
	b := []int{5, 8, 10, 20}

	fmt.Println(findMedianSortedArrays(a, b))
}
