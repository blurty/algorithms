// Copyright 2017 Blurt. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// A Simple Merge based O(n) solution to find median of
// two sorted arrays

package main

import "fmt"

// This function returns median of ar1[] and ar2[].
// Assumptions in this function:
// Both ar1[] and ar2[] are sorted arrays
// Both have n elements
func getMedian_method_1(ar1, ar2 []int, n int) float64 {
	var i, j int
	m1, m2 := -1, -1

	// Since there are 2n elements, median will be average
	// of elements at index n-1 and n in the array obtained after
	// merging ar1 and ar2
	for count := 0; count <= n; count++ {
		// Store the prev median
		m1 = m2
		// Below is to handle case where all elements of ar1[] are
		// smaller or bigger than smallest(or first) element of ar2[]
		if i == n {
			m2 = ar2[0]
			break
		} else if j == n {
			m2 = ar1[0]
			break
		}

		if ar1[i] < ar2[j] {
			m2 = ar1[i]
			i++
		} else {
			m2 = ar2[j]
			j++
		}
	}
	return (float64(m1) + float64(m2)) / 2
}

func main() {
	ar1 := []int{1, 12, 15, 26, 38}
	ar2 := []int{2, 13, 17, 30, 45}

	n := len(ar1)
	fmt.Println(getMedian_method_1(ar1, ar2, n))
}
