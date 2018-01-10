package sort

func BubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func QuickSort(data []int) {
	if len(data) > 0 {
		privotLoc := partition(data, 0, len(data)-1)
		QuickSort(data[:privotLoc])
		QuickSort(data[privotLoc+1:])
	}
}

// cut data into two parts in quick sort
func partition(data []int, low, high int) int {
	privotKey := data[low]
	for low < high {
		for low < high && data[high] >= privotKey {
			high--
		}
		data[low], data[high] = data[high], data[low]
		for low < high && data[low] <= privotKey {
			low++
		}
		data[low], data[high] = data[high], data[low]
	}
	return low
}

// insert sort
func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			// x copy as sentry, j act as temp index
			x, j := data[i], i-1
			// move item at i-1 to i first
			data[i] = data[i-1]
			// lookup the insert position
			for j >= 0 && x < data[j] {
				data[j+1] = data[j]
				j--
			}
			// insert to the correct position
			data[j+1] = x
		}
	}
}

// shell insert sort
// similar to InsertSort, two functions can be combined
func shellInsertSort(data []int, n, dk int) {
	for i := dk; i < n; i++ {
		if data[i] < data[i-dk] {
			j := i - dk
			x := data[i]
			data[i] = data[i-dk]
			for j >= 0 && x < data[j] {
				data[j+dk] = data[j]
				j -= dk
			}
			data[j+dk] = x
		}
	}
}

// shell sort
func ShellSort(data []int) {
	length := len(data)
	dk := length / 2
	for dk >= 1 {
		shellInsertSort(data, length, dk)
		dk /= 2
	}
}

// select sort
func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		key := i + findMinKeyIndex(nums[i:])
		if key != i {
			nums[i], nums[key] = nums[key], nums[i]
		}
	}
}

// find the index of min key
func findMinKeyIndex(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[index] {
			index = i
		}
	}
	return index
}

// an improvement to select sort
func DoubleSelectSort(nums []int) {
	length := len(nums)
	for i := 1; i <= length/2; i++ {
		min, max := i-1, i-1
		for j := i; j <= length-i; j++ {
			if nums[j] > nums[max] {
				max = j
				continue
			}
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i-1], nums[min] = nums[min], nums[i-1]
		// avoid collision between max exchange and min exchange
		if max == i-1 {
			max = min
		}
		nums[length-i], nums[max] = nums[max], nums[length-i]
	}
}

// heap sort
func HeapSort(nums []int) {
	buildingHeap(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
}

// building the heap from slice
func buildingHeap(nums []int) {
	length := len(nums)
	for i := (length - 1) / 2; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}
}

func adjustHeap(nums []int, s, length int) {
	// child here is left child, child++ is right child
	child := 2*s + 1
	for child < length {
		if child+1 < length && nums[child] < nums[child+1] {
			child++
		}
		if nums[s] < nums[child] {
			nums[s], nums[child] = nums[child], nums[s]
			s = child
			child = 2*s + 1
		} else {
			break
		}
	}
}

// merge sort
func MergeSort(nums []int) {
	length := len(nums)
	ret := make([]int, length)
	privot := 1
	for privot < length {
		s := privot
		privot *= 2
		i := 0
		for i+privot < length {
			merge(nums, ret, i, i+s-1, i+privot-1)
			i += privot
		}
		if i+s < length {
			merge(nums, ret, i, i+s-1, length-1)
		}
		// exchange nums and ret to assure nums always merge to ret
		nums, ret = ret, nums
	}
}

// merge r[i...m] and r[m+1...n] to rf[i...n]
func merge(r, rf []int, i, m, n int) {
	j, k := m+1, i
	for ; i <= m && j <= n; k++ {
		if r[j] < r[i] {
			rf[k] = r[j]
			j++
		} else {
			rf[k] = r[i]
			i++
		}
	}
	for ; i <= m; i++ {
		rf[k] = r[i]
		k++
	}
	for ; j <= n; j++ {
		rf[k] = r[j]
		k++
	}
}

// another implement for quick sort
func QuickSort2(data []int) {
	if len(data) < 2 {
		return
	}
	head, tail := 0, len(data)-1
	i := 1
	for head < tail {
		if data[i] < data[head] {
			data[i], data[head] = data[head], data[i]
			i++
			head++
		} else {
			data[i], data[tail] = data[tail], data[i]
			tail--
		}
	}
	QuickSort2(data[:head])
	QuickSort2(data[head+1:])
}

// improvement of quick sort
func QuickSort3(data []int) {
	if len(data) < 2 {
		return
	} else if len(data) == 2 {
		if data[0] > data[1] {
			data[1], data[0] = data[0], data[1]
		}
		return
	}
	head, tail := 0, len(data)-1
	mid := (data[0] + data[tail] + data[(tail+1)/2]) / 3
	for head <= tail {
		if data[head] <= mid {
			head++
		} else {
			data[head], data[tail] = data[tail], data[head]
			tail--
		}
	}
	QuickSort3(data[:head])
	QuickSort3(data[head:])
}
