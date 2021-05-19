package algorithm

func Kth(arr []int, k int, less func(i, j int) bool) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := partition(arr, l, r, less)
		if m == k {
			return arr[m]
		}
		if m > k {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return 0
}
