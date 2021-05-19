package algorithm

func partition(arr []int, l, r int, less func(i, j int) bool) int {
	if l >= r {
		return l
	}
	k, i, j := l, l+1, r
	for i <= j {
		if less(i, k) {
			arr[k], arr[i] = arr[i], arr[k]
			k, i = k+1, i+1
			continue
		}

		// 否则的话，需要移动j交换位置
		if !less(j, k) {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return i - 1
}

func quickSort(arr []int, l, r int, less func(i, j int) bool) {
	if r-l < 1 {
		return
	}
	m := partition(arr, l, r, less)
	quickSort(arr, l, m-1, less)
	quickSort(arr, m+1, r, less)
}

func QuickSort(arr []int, less func(i, j int) bool) {
	quickSort(arr, 0, len(arr)-1, less)
}
