package quickSort

func QuickSort(arr []int, lo, hi int) []int {
	if hi-lo < 2 {
		return nil
	}
	mi := partition(arr, lo, hi)
	QuickSort(arr, lo, mi)
	QuickSort(arr, mi+1, hi)
	return arr
}

func partition(arr []int, lo, hi int) int {
	hi--
	pivot := arr[lo]
	for lo < hi {
		for lo < hi && pivot <= arr[hi] {
			hi--
		}
		arr[lo] = arr[hi]
		for lo < hi && pivot >= arr[lo] {
			lo++
		}
		arr[hi] = arr[lo]
	}
	arr[lo] = pivot
	return lo
}
