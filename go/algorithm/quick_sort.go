package algorithm

func partition(arr []int, l, r int, less func(i, j int) bool) int {
	swap := func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }
	if l >= r {
		return l
	}
	k, i, j := l, l+1, r
	for {
		for i < r && less(i, k) {
			i++
		}
		for j > l && !less(j, k) {
			j--
		}

		if i >= j {
			break
		}
		swap(i, j)
		i, j = i+1, j-1
	}
	swap(j, k)
	return j
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
