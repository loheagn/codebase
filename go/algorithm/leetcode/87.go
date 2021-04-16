package leetcode

import (
	"reflect"
)

type Status int

const (
	False Status = iota - 1
	NotSet
	True
)

var mem [][]map[int]Status

func compare(s1, s2 string, i, j, n int) bool {
	if mem[i][j][n] != NotSet {
		return mem[i][j][n] == True
	}
	if s1[i:i+n] == s2[j:j+n] {
		mem[i][j][n] = True
		return true
	}

	hash := func(s string, base int) map[byte]int {
		re := make(map[byte]int)
		for i := base; i < base+n; i++ {
			re[s[i]]++
		}
		return re
	}
	if !reflect.DeepEqual(hash(s1, i), hash(s2, j)) {
		mem[i][j][n] = False
		return false
	}

	for k := 1; k < n; k++ {
		if compare(s1, s2, i+k, j+k, n-k) && compare(s1, s2, i, j, k) {
			mem[i][j][n] = True
			return true
		}

		if compare(s1, s2, i+k, j, n-k) && compare(s1, s2, i, j+n-k, k) {
			mem[i][j][n] = True
			return true
		}
	}

	mem[i][j][n] = False
	return false
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	if n <= 0 {
		return true
	}

	// init
	mem = make([][]map[int]Status, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]map[int]Status, n)
		for j := 0; j < n; j++ {
			mem[i][j] = make(map[int]Status)
		}
	}

	return compare(s1, s2, 0, 0, n)
}
