package leetcode

var (
	cnt_equal    = 0
	cnt_no_equal = 0
)

func partition_equal(arr []int, l, r int) int {
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]

	}
	if l >= r {
		return l
	}
	k, i, j := l, l+1, r
	for {
		for i < r && arr[i] < arr[k] {
			i++
		}
		for j > l && arr[j] >= arr[k] {
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

func quickSort_equal(arr []int, l, r int) {
	if r-l < 1 {
		return
	}
	m := partition_equal(arr, l, r)
	quickSort_equal(arr, l, m-1)
	quickSort_equal(arr, m+1, r)
}

func QuickSort_equal(arr []int) {
	quickSort_equal(arr, 0, len(arr)-1)
}

func partition_no_equal(arr []int, l, r int) int {
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]

	}
	if l >= r {
		return l
	}
	k, i, j := l, l+1, r
	for {
		for i < r && arr[i] < arr[k] {
			i++
		}
		for j > l && arr[j] > arr[k] {
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

func quickSort_no_equal(arr []int, l, r int) {
	if r-l < 1 {
		return
	}
	m := partition_no_equal(arr, l, r)
	quickSort_no_equal(arr, l, m-1)
	quickSort_no_equal(arr, m+1, r)
}

func QuickSort_no_equal(arr []int) {
	quickSort_no_equal(arr, 0, len(arr)-1)
}
