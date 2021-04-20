package algorithm

func next(needle string) []int {
	n := len(needle)
	re := make([]int, n)
	if n <= 1 {
		return re
	}
	pos, cnd := 2, 0
	for pos < n {
		if needle[pos-1] == needle[cnd] {
			re[pos], cnd, pos = cnd+1, cnd+1, pos+1
		} else if cnd > 0 {
			cnd = re[cnd]
		} else {
			re[pos], pos = 0, pos+1
		}
	}
	return re
}

func kmp(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	next := next(needle)
	i, j := 0, 0
	for i < n {
		if j >= m {
			break
		}
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
		} else {
			if j > 0 {
				j = next[j]
			} else {
				i, j = i+1, 0
			}
		}
	}
	if j >= m {
		return i - m
	}
	return -1
}
