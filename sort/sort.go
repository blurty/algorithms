package sort

func BubbleSort(data []int) {
	for i:=0; i<len(data)-1; i++ {
		for j:=0; j<len(data)-i-1; j++ {
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
	for i:=1; i<len(data); i++ {
		if data[i] < data[i-1] {
			// x copy as sentry, j act as temp index
			x, j := data[i], i-1
			// move item at i-1 to i first
			data[i] = data[i-1]
			// lookup the insert position
			for j>=0 && x < data[j] {
				data[j+1] = data[j]
				j--
			}
			// insert to the correct position
			data[j+1] = x
		}
	}
}