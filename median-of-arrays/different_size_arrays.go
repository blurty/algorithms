// Copyright 2017 Blurt. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// A go program to find median of two sorted arrays of
// unequal sizes

package main

import "fmt"

// select max one
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// select min one
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A utility function to find median of two integers
func MO2(a, b int) float64 {
	return (float64(a) + float64(b)) / 2
}

// A utility function to find median of three integers
func MO3(a, b, c int) int {
	return a + b + c - max(a, max(b, c)) - min(a, min(b, c))
}

// A utility function to find median of four integers
func MO4(a, b, c, d int) float64 {
	f := max(a, max(b, max(c, d)))
	g := min(a, min(b, min(c, d)))
	return float64(a+b+c+d-f-g) / 2
}

// Utility function to find median of single array
func medianSingle(a []int, n int) float64 {
	if n == 0 {
		return float64(-1)
	}
	if n%2 == 0 {
		return float64(a[n/2]+a[n/2-1]) / 2
	}
	return float64(a[n/2])
}

// This function assumes that N is smaller than or eq
// This function returns -1 if both arrays are empty
func findMedianUtil(a, b []int, m, n int) float64 {
	if m == 0 {
		return medianSingle(b, n)
	}
	if m == 1 {
		// Case 1: If the larger array also has one element,
		// simply call MO2()
		if n == 1 {
			return MO2(a[0], b[0])
		}

		// Case 2: If the larger array has odd number of elements,
		// then consider the middle 3 elements of larger array and
		// the only element of smaller array. Take few examples
		// like following
		// A[] = {9}, B[] = {5, 8, 10, 20, 30} and
		// A[] = {1}, B[] = {5, 8, 10, 20, 30}
		if n&1 == 1 {
			return MO2(b[n/2], MO3(a[0], b[n/2-1], b[n/2]+1))
		}

		// Case 3: If the larger array has even number of element,
		// then median will be one of the following 3 elements
		// ... The middle two elements of larger array
		// ... The only element of smaller array
		return float64(MO3(a[0], b[n/2-1], b[n/2]))
	}

	if m == 2 {
		// Case 4: If the larger array also has two elements,
		// simply call MO4()
		if n == 2 {
			return MO4(a[0], a[1], b[0], b[1])
		}

		// Case 5: If the larger array has odd number of elements,
		// then median will be one of the following 3 elements
		// 1. Middle element of larger array
		// 2. Max of first element of smaller array and element
		//    just before the middle in bigger array
		// 3. Min of second element of smaller array and element
		//    just after the middle in bigger array

		if n&1 == 1 {
			return float64(MO3(b[n/2], max(a[0], b[n/2-1]), min(a[1], b[n/2+1])))
		}

		// Case 6: If the larger array has even number of elements,
		// then median will be one of the following 4 elements
		// 1) & 2) The middle two elements of larger array
		// 3) Max of first element of smaller array and element
		//    just before the first middle element in bigger array
		// 4. Min of second element of smaller array and element
		//    just after the second middle in bigger array
		return MO4(b[n/2-1], b[n/2], max(a[0], b[n/2-2]), min(a[1], b[n/2+1]))
	}

	idxA := (m - 1) / 2
	idxB := (n - 1) / 2

	// if A[idxA] <= B[idxB], then median must exist in
	//    A[idxA....] and B[....idxB]
	if a[idxA] <= b[idxB] {
		return findMedianUtil(a[idxA:], b, m/2+1, n-idxA)
	}

	// if A[idxA] > B[idxB], then median must exist in
	//    A[...idxA] and B[idxB....]
	return findMedianUtil(a, b[idxA:], m/2+1, n-idxA)
}

// A wrapper function around findMedianUtil(). This function
// makes sure that smaller array is passed as first argument
// to findMedianUtil
func findMedian(a, b []int, m, n int) float64 {
	if m > n {
		return findMedianUtil(b, a, n, m)
	}
	return findMedianUtil(a, b, m, n)
}

// Driver program to test above functions
func main() {
	a := []int{900}
	b := []int{5, 8, 10, 20}

	alen := len(a)
	blen := len(b)

	fmt.Println(findMedian(a, b, alen, blen))
}
