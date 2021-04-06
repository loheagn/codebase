package algorithm

func binarySearch(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := l + (r-l)/2
		if a[m] > x {
			r = m - 1
		} else if a[m] <= x {
			l = m + 1
		}
	}
	return r
}
